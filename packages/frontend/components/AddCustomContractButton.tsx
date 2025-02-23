"use client"

import { useCallback, useState, type ChangeEvent } from "react"
import { Button } from "@/components/ui/button"
import { useEtherbase } from "@msquared/etherbase-client"
import { useRouter } from "next/navigation"
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog"
import { Input } from "@/components/ui/input"
import { Textarea } from "@/components/ui/textarea"
import { getAddress } from "viem"

export default function AddCustomContractButton() {
  const [isAdding, setIsAdding] = useState(false)
  const [isOpen, setIsOpen] = useState(false)
  const [contractAddress, setContractAddress] = useState("")
  const [contractABI, setContractABI] = useState("")
  const { addCustomContract } = useEtherbase()
  const router = useRouter()

  const isValidAddress = contractAddress && getAddress(contractAddress)
  const isValidABI =
    contractABI &&
    (() => {
      try {
        JSON.parse(contractABI)
        return true
      } catch (error) {
        console.error("Error parsing ABI:", error)
        return false
      }
    })()

  const handleAddContract = useCallback(async () => {
    if (!isValidAddress || !isValidABI) return

    try {
      setIsAdding(true)
      const checksumAddress = getAddress(contractAddress)
      const success = await addCustomContract(checksumAddress, contractABI)
      if (!success) {
        throw new Error("Failed to add custom contract")
      }
      console.log("Added custom contract", success)
      setIsOpen(false)
      router.push(`/custom-contracts/${checksumAddress}`)
    } catch (error) {
      console.error("Error adding custom contract:", error)
    } finally {
      setIsAdding(false)
    }
  }, [
    router,
    addCustomContract,
    contractAddress,
    contractABI,
    isValidAddress,
    isValidABI,
  ])

  return (
    <Dialog open={isOpen} onOpenChange={setIsOpen}>
      <DialogTrigger asChild>
        <Button className="bg-white text-black">Add Custom Contract</Button>
      </DialogTrigger>
      <DialogContent className="sm:max-w-[425px]">
        <DialogHeader>
          <DialogTitle>Add Custom Contract</DialogTitle>
        </DialogHeader>
        <div className="grid gap-4 py-4">
          <div className="grid gap-2">
            <label htmlFor="address">Contract Address</label>
            <Input
              id="address"
              value={contractAddress}
              onChange={(e: ChangeEvent<HTMLInputElement>) =>
                setContractAddress(e.target.value)
              }
              placeholder="0x..."
              className={
                !contractAddress || isValidAddress ? "" : "border-red-500"
              }
            />
            {contractAddress && !isValidAddress && (
              <p className="text-sm text-red-500">
                Please enter a valid EVM contract address
              </p>
            )}
          </div>
          <div className="grid gap-2">
            <label htmlFor="abi">Contract ABI</label>
            <Textarea
              id="abi"
              value={contractABI}
              onChange={(e) => setContractABI(e.target.value.trim())}
              placeholder="[{...}]"
              className={!contractABI || isValidABI ? "" : "border-red-500"}
            />
            {contractABI && !isValidABI && (
              <p className="text-sm text-red-500">
                Please enter a valid JSON ABI
              </p>
            )}
          </div>
          <Button
            onClick={handleAddContract}
            disabled={isAdding || !isValidAddress || !isValidABI}
            className="bg-white text-black"
          >
            {isAdding ? "Adding..." : "Add Contract"}
          </Button>
        </div>
      </DialogContent>
    </Dialog>
  )
}
