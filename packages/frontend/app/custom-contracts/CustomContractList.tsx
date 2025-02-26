"use client"

import { useRouter } from "next/navigation"
import { useEtherbase, useWebThree } from "@msquared/etherbase-client"
import { useEffect, useState } from "react"
import AddCustomContractButton from "@/components/AddCustomContractButton"
import { Search } from "lucide-react"

export default function CustomContractList() {
  const router = useRouter()
  const { customContracts } = useEtherbase()
  const { getWalletClient } = useWebThree()
  const [searchTerm, setSearchTerm] = useState("")
  const [currentWallet, setCurrentWallet] = useState<string | null>(null)

  useEffect(() => {
    const getCurrentWallet = async () => {
      const walletClient = await getWalletClient()
      if (walletClient?.account) {
        setCurrentWallet(walletClient.account.address.toLowerCase())
      }
    }
    getCurrentWallet()
  }, [getWalletClient])

  // Filter contracts based on search term
  const filteredContracts = customContracts.filter(contract => 
    searchTerm === "" || 
    contract.contractAddress.toLowerCase().includes(searchTerm.toLowerCase())
  )

  // Group contracts by owner
  const groupedContracts = filteredContracts.reduce((groups, contract) => {
    const ownerAddress = contract.addedBy.toLowerCase()
    if (!groups[ownerAddress]) {
      groups[ownerAddress] = []
    }
    groups[ownerAddress].push(contract)
    return groups
  }, {} as Record<string, typeof customContracts>)

  // Sort owners to put current user first
  const sortedOwners = Object.keys(groupedContracts).sort((a, b) => {
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
          <h1 className="text-2xl font-bold">Custom Contracts</h1>
          <AddCustomContractButton />
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
          Showing {filteredContracts.length} of {customContracts.length} contracts
        </p>
        
        {/* Grouped contracts */}
        {filteredContracts.length > 0 ? (
          <div className="space-y-8">
            {sortedOwners.map(ownerAddress => (
              <div key={ownerAddress} className="space-y-4">
                <h2 className="text-xl font-semibold border-b border-gray-700 pb-2">
                  {ownerAddress === currentWallet 
                    ? `Your Custom Contracts (${formatAddress(ownerAddress)})` 
                    : `Contracts by ${formatAddress(ownerAddress)}`}
                </h2>
                <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
                  {groupedContracts[ownerAddress].map(contract => (
                    <div
                      key={contract.contractAddress}
                      onClick={() => router.push(`/custom-contracts/${contract.contractAddress}`)}
                      onKeyDown={(e) => {
                        if (e.key === "Enter") {
                          router.push(`/custom-contracts/${contract.contractAddress}`)
                        }
                      }}
                      tabIndex={0}
                      className="cursor-pointer rounded-lg border border-gray-700 bg-gray-800/20 p-4 hover:shadow-lg transition-all hover:border-gray-500 focus:outline-none focus:ring-2 focus:ring-gray-500"
                    >
                      <h3 className="text-lg font-semibold mb-2 truncate" title={contract.contractAddress}>
                        {formatAddress(contract.contractAddress)}
                      </h3>
                    </div>
                  ))}
                </div>
              </div>
            ))}
          </div>
        ) : (
          <div className="text-center py-8 text-gray-500">
            No contracts found matching your search
          </div>
        )}
      </div>
    </div>
  )
}
