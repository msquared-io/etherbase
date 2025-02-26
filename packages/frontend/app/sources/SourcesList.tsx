"use client"

import { useRouter } from "next/navigation"
import { useEtherbase, useWebThree } from "@msquared/etherbase-client"
import { useEffect, useState } from "react"
import CreateSourceButton from "@/components/CreateSourceButton"
import { Search } from "lucide-react"

export default function SourcesList() {
  const router = useRouter()
  const { sources, fetchSources } = useEtherbase()
  const { getWalletClient } = useWebThree()
  const [searchTerm, setSearchTerm] = useState("")
  const [currentWallet, setCurrentWallet] = useState<string | null>(null)

  useEffect(() => {
    fetchSources()
  }, [fetchSources])

  useEffect(() => {
    const getCurrentWallet = async () => {
      const walletClient = await getWalletClient()
      if (walletClient?.account) {
        setCurrentWallet(walletClient.account.address.toLowerCase())
      }
    }
    getCurrentWallet()
  }, [getWalletClient])

  // Filter sources based on search term
  const filteredSources = sources.filter(source => 
    searchTerm === "" || 
    source.sourceAddress.toLowerCase().includes(searchTerm.toLowerCase())
  )

  // Group sources by owner
  const groupedSources = filteredSources.reduce((groups, source) => {
    const ownerAddress = source.owner.toLowerCase()
    if (!groups[ownerAddress]) {
      groups[ownerAddress] = []
    }
    groups[ownerAddress].push(source)
    return groups
  }, {} as Record<string, typeof sources>)

  // Sort owners to put current user first
  const sortedOwners = Object.keys(groupedSources).sort((a, b) => {
    if (a === currentWallet) return -1
    if (b === currentWallet) return 1
    return a.localeCompare(b)
  })

  // Function to format address for display
  const formatAddress = (address: string) => {
    return `${address.substring(0, 6)}...${address.substring(address.length - 4)}`
  }

  return (
    <div className="w-full max-w-6xl mx-auto">
      <div className="mb-6">
        <div className="flex justify-between items-center mb-4">
          <h1 className="text-2xl font-bold">Sources</h1>
          <CreateSourceButton fetchSources={fetchSources} />
        </div>
        
        {/* Search filter */}
        <div className="mb-6">
          <div className="relative">
            <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
              <Search className="h-5 w-5 text-gray-400" />
            </div>
            <input
              type="text"
              placeholder="Search by address"
              className="pl-10 w-full p-2 bg-gray-800/20 border border-gray-700 rounded-md focus:outline-none focus:ring-2 focus:ring-gray-500 text-white placeholder-gray-400"
              value={searchTerm}
              onChange={(e) => setSearchTerm(e.target.value)}
            />
          </div>
        </div>
        
        {/* Results count */}
        <p className="text-sm text-gray-500 mb-4">
          Showing {filteredSources.length} of {sources.length} sources
        </p>
        
        {/* Grouped sources */}
        {filteredSources.length > 0 ? (
          <div className="space-y-8">
            {sortedOwners.map(ownerAddress => (
              <div key={ownerAddress} className="space-y-4">
                <h2 className="text-xl font-semibold border-b border-gray-700 pb-2">
                  {ownerAddress === currentWallet 
                    ? `Your Sources (${formatAddress(ownerAddress)})` 
                    : `Sources by ${formatAddress(ownerAddress)}`}
                </h2>
                <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
                  {groupedSources[ownerAddress].map(source => (
                    <div
                      key={source.sourceAddress}
                      onClick={() => router.push(`/sources/${source.sourceAddress}`)}
                      onKeyDown={(e) => {
                        if (e.key === "Enter") {
                          router.push(`/sources/${source.sourceAddress}`)
                        }
                      }}
                      tabIndex={0}
                      className="cursor-pointer rounded-lg border border-gray-700 bg-gray-800/20 p-4 hover:shadow-lg transition-all hover:border-gray-500 focus:outline-none focus:ring-2 focus:ring-gray-500"
                    >
                      <h3 className="text-lg font-semibold mb-2 truncate" title={source.sourceAddress}>
                        {formatAddress(source.sourceAddress)}
                      </h3>
                    </div>
                  ))}
                </div>
              </div>
            ))}
          </div>
        ) : (
          <div className="text-center py-8 text-gray-500">
            No sources found matching your search
          </div>
        )}
      </div>
    </div>
  )
}
