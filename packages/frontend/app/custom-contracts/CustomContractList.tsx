"use client"

import { useRouter } from "next/navigation"
import { useEtherbase, useWebThree } from "@msquared/etherbase-client"
import { useEffect, useState } from "react"
import AddCustomContractButton from "@/components/AddCustomContractButton"

export default function CustomContractList() {
  const router = useRouter()
  const { customContracts, fetchCustomContracts, addCustomContract } =
    useEtherbase()
  const { getWalletClient } = useWebThree()
  const [showOnlyMine, setShowOnlyMine] = useState(false)
  const [currentWallet, setCurrentWallet] = useState<string | null>(null)

  useEffect(() => {
    fetchCustomContracts()
  }, [fetchCustomContracts])

  useEffect(() => {
    const getCurrentWallet = async () => {
      const walletClient = await getWalletClient()
      if (walletClient?.account) {
        setCurrentWallet(walletClient.account.address.toLowerCase())
      }
    }
    getCurrentWallet()
  }, [getWalletClient])

  const filteredContracts =
    showOnlyMine && currentWallet
      ? customContracts.filter(
          (contract) => contract.addedBy.toLowerCase() === currentWallet,
        )
      : customContracts

  return (
    <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-4 flex-col">
      <div className="mb-3 flex gap-4 flex-col items-start">
        <div className="flex gap-4 items-center">
          <h1 className="text-xl font-bold">Custom Contracts</h1>
          <AddCustomContractButton />
        </div>
        <div className="flex items-center gap-2 mb-4">
          <label className="relative inline-flex items-center cursor-pointer">
            <input
              type="checkbox"
              className="sr-only peer"
              checked={showOnlyMine}
              onChange={(e) => setShowOnlyMine(e.target.checked)}
            />
            <div className="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600" />
            <span className="ml-3 text-sm font-medium text-white">
              Show only my contracts
            </span>
          </label>
        </div>
        {filteredContracts.map((customContract) => (
          <div
            key={customContract.contractAddress}
            onClick={() =>
              router.push(`/custom-contracts/${customContract.contractAddress}`)
            }
            onKeyDown={(e) => {
              if (e.key === "Enter") {
                router.push(
                  `/custom-contracts/${customContract.contractAddress}`,
                )
              }
            }}
            className="cursor-pointer rounded-lg border border-gray-200 p-4 hover:shadow-lg transition"
          >
            <h2 className="text-lg font-semibold truncate">
              {customContract.contractAddress}
            </h2>
            <p className="mt-1 text-xs text-gray-600">
              Added By: {customContract.addedBy}
            </p>
            <p className="mt-1 text-sm text-gray-600">Click to view</p>
          </div>
        ))}
      </div>
    </div>
  )
}
