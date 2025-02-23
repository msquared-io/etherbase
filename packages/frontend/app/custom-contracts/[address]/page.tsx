"use client"

import { useEffect, useState } from "react"
import { useEtherbase, useEtherbaseEvents } from "@msquared/etherbase-client"
import type { EtherbaseEvent } from "@msquared/etherbase-client"
import { Button } from "@/components/ui/button"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { useRouter } from "next/navigation"
import type { Address, AbiEvent, AbiParameter, AbiFunction } from "viem"
import { AnimatedList } from "@/components/ui/animated-list"
import { Separator } from "@/components/ui/separator"
import { parseAbiItem } from "viem"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { useWebThree } from "@msquared/etherbase-client"

interface ExtendedEvent extends EtherbaseEvent {
  index: number
}

interface EventSchema {
  name: string
  args: {
    name: string
    type: string
    indexed: boolean
  }[]
}

interface WriteFunction {
  name: string
  args: {
    name: string
    type: string
  }[]
  payable: boolean
}

export default function CustomContractPage({
  params,
}: {
  params: { address: string }
}) {
  const router = useRouter()
  const { customContracts, fetchCustomContracts, deleteCustomContract } =
    useEtherbase()
  const [isDeleting, setIsDeleting] = useState(false)
  const [connectedAddress, setConnectedAddress] = useState<string>()
  const [latestEvents, setLatestEvents] = useState<
    Record<string, ExtendedEvent[]>
  >({})
  const [eventSchemas, setEventSchemas] = useState<EventSchema[]>([])
  const [writeFunctions, setWriteFunctions] = useState<WriteFunction[]>([])
  const [functionInputs, setFunctionInputs] = useState<
    Record<string, string[]>
  >({})
  const [ethValue, setEthValue] = useState<Record<string, string>>({})
  const [isExecuting, setIsExecuting] = useState<Record<string, boolean>>({})
  const { publicClient, getWalletClient } = useWebThree()

  useEffect(() => {
    fetchCustomContracts()
  }, [fetchCustomContracts])

  useEffect(() => {
    // Get the connected account from MetaMask
    async function getConnectedAccount() {
      if (!window.ethereum) return
      const accounts = (await window.ethereum.request({
        method: "eth_requestAccounts",
      })) as string[]
      setConnectedAddress(accounts[0]?.toLowerCase())
    }
    getConnectedAccount()
  }, [])

  const contract = customContracts.find(
    (c) => c.contractAddress.toLowerCase() === params.address.toLowerCase(),
  )

  useEffect(() => {
    if (!contract) return

    // Parse ABI to get event schemas and write functions
    try {
      const abi = JSON.parse(contract.contractABI) as (AbiEvent | AbiFunction)[]

      // Get events
      const eventItems = abi.filter(
        (item): item is AbiEvent => item.type === "event",
      )
      const schemas = eventItems.map((event) => ({
        name: event.name,
        args: event.inputs.map((input) => ({
          name: input.name ?? "",
          type: input.type,
          indexed: !!input.indexed,
        })),
      }))
      setEventSchemas(schemas)

      // Get write functions
      const writeFunctionItems = abi.filter(
        (item): item is AbiFunction =>
          item.type === "function" &&
          (item.stateMutability === "nonpayable" ||
            item.stateMutability === "payable"),
      )
      const functions = writeFunctionItems.map((func) => ({
        name: func.name,
        args: func.inputs.map((input) => ({
          name: input.name ?? "",
          type: input.type,
        })),
        payable: func.stateMutability === "payable",
      }))
      setWriteFunctions(functions)

      // Initialize function inputs
      const initialInputs: Record<string, string[]> = {}
      const initialExecuting: Record<string, boolean> = {}
      const initialEthValue: Record<string, string> = {}
      for (const func of functions) {
        initialInputs[func.name] = Array(func.args.length).fill("")
        initialExecuting[func.name] = false
        initialEthValue[func.name] = ""
      }
      setFunctionInputs(initialInputs)
      setIsExecuting(initialExecuting)
      setEthValue(initialEthValue)
    } catch (error) {
      console.error("Error parsing ABI:", error)
    }
  }, [contract])

  const handleNewEvent = (event: EtherbaseEvent) => {
    setLatestEvents((prev) => {
      const updated = { ...prev }
      const currentEvents = updated[event.name] || []
      const extendedEvent: ExtendedEvent = {
        ...event,
        index: currentEvents.length + 1,
      }
      updated[event.name] = [...currentEvents, extendedEvent]
      return updated
    })
  }

  // Subscribe to contract events
  useEtherbaseEvents({
    contractAddress: params.address as Address,
    events: eventSchemas.map((schema) => ({ name: schema.name })),
    onEvent: handleNewEvent,
  })

  const handleExecuteFunction = async (func: WriteFunction) => {
    if (!window.ethereum || !contract || !publicClient) return

    try {
      setIsExecuting((prev) => ({ ...prev, [func.name]: true }))

      // Get wallet client
      const walletClient = await getWalletClient()
      if (!walletClient) throw new Error("Failed to get wallet client")

      // Parse the inputs based on their types
      const args = functionInputs[func.name].map((input, idx) => {
        const argType = func.args[idx].type
        if (argType.includes("int")) {
          return BigInt(input)
        }
        if (argType === "bool") {
          return input.toLowerCase() === "true"
        }
        return input
      })

      // Execute the transaction
      const hash = await walletClient.writeContract({
        address: contract.contractAddress as Address,
        abi: JSON.parse(contract.contractABI),
        functionName: func.name,
        args,
        value: func.payable ? BigInt(ethValue[func.name] || "0") : undefined,
        chain: walletClient.chain,
        account: walletClient.account?.address ?? null,
      })

      // Wait for transaction receipt
      await publicClient.waitForTransactionReceipt({ hash })

      alert("Transaction successful!")
    } catch (error) {
      console.error("Error executing function:", error)
      alert(`Error executing function: ${(error as Error).message}`)
    } finally {
      setIsExecuting((prev) => ({ ...prev, [func.name]: false }))
    }
  }

  if (!contract) {
    return (
      <div className="flex flex-col items-center justify-center p-8">
        <h1 className="text-2xl font-bold mb-4">Contract Not Found</h1>
        <p className="text-gray-600 mb-4">
          The contract you're looking for doesn't exist or has been deleted.
        </p>
        <Button onClick={() => router.push("/custom-contracts")}>
          Back to Custom Contracts
        </Button>
      </div>
    )
  }

  const handleDelete = async () => {
    try {
      setIsDeleting(true)
      await deleteCustomContract(contract.contractAddress as Address)
      router.push("/custom-contracts")
    } catch (error) {
      console.error("Error deleting contract:", error)
      alert("Failed to delete contract")
    } finally {
      setIsDeleting(false)
    }
  }

  const canDelete = connectedAddress === contract.addedBy.toLowerCase()

  return (
    <div className="space-y-6">
      <div className="flex justify-between items-center">
        <h1 className="text-2xl font-bold">Custom Contract Details</h1>
        <Button
          onClick={() => router.push("/custom-contracts")}
          variant="outline"
        >
          Back to List
        </Button>
      </div>

      <Card>
        <CardHeader>
          <CardTitle>Contract Information</CardTitle>
        </CardHeader>
        <CardContent className="space-y-4">
          <div>
            <h3 className="font-semibold mb-1">Contract Address</h3>
            <p className="font-mono text-sm">{contract.contractAddress}</p>
          </div>
          <div>
            <h3 className="font-semibold mb-1">Added By</h3>
            <p className="font-mono text-sm">{contract.addedBy}</p>
          </div>
          <div>
            <h3 className="font-semibold mb-1">Contract ABI</h3>
            <pre className="bg-gray-900 p-4 rounded-lg overflow-x-auto text-sm">
              {JSON.stringify(JSON.parse(contract.contractABI), null, 2)}
            </pre>
          </div>
        </CardContent>
      </Card>

      <Separator className="my-4" />

      <div className="space-y-4">
        {eventSchemas.map((schema) => {
          const eventsForType = latestEvents[schema.name] || []
          return (
            <Card key={schema.name}>
              <CardHeader>
                <CardTitle className="text-lg font-semibold -mb-2">
                  {schema.name}
                </CardTitle>
              </CardHeader>
              <CardContent className="flex flex-row gap-5">
                <div className="flex flex-col w-1/6">
                  <h6 className="text-muted-foreground font-semibold">
                    Parameters
                  </h6>
                  {schema.args?.map((arg, idx) => (
                    <div
                      key={`${schema.name}-${arg.name}-${idx}`}
                      className="text-sm"
                    >
                      {arg.name} ({arg.type})
                      {arg.indexed && (
                        <span className="text-xs ml-1 text-muted-foreground">
                          (indexed)
                        </span>
                      )}
                    </div>
                  ))}
                </div>
                <div>
                  <h6 className="text-muted-foreground font-semibold">
                    Events
                  </h6>
                  <div className="flex gap-3 overflow-hidden">
                    <AnimatedList delay={0.2}>
                      {eventsForType.map((evt) => (
                        <Card
                          className="min-w-[350px] max-w-[350px] border border-accent/50 bg-background shadow-sm mt-2"
                          key={evt.index}
                        >
                          <CardHeader>
                            <CardTitle className="text-s font-medium">
                              {`#${evt.index}`}
                            </CardTitle>
                          </CardHeader>
                          <CardContent className="space-y-0.5">
                            <p className="text-xs text-muted-foreground">
                              Block: {evt.block}
                            </p>
                            {Object.entries(evt.args).map(([key, value]) => (
                              <p
                                key={`${evt.index}-${key}`}
                                className="text-xs text-foreground"
                              >
                                {key}: {String(value)}
                              </p>
                            ))}
                          </CardContent>
                        </Card>
                      ))}
                    </AnimatedList>
                    {eventsForType.length === 0 && (
                      <p className="text-sm text-muted-foreground">
                        No events received yet
                      </p>
                    )}
                  </div>
                </div>
              </CardContent>
            </Card>
          )
        })}
      </div>

      <Separator className="my-4" />

      <div className="space-y-4">
        <h2 className="text-xl font-bold">Write Functions</h2>
        {writeFunctions.map((func) => (
          <Card key={func.name}>
            <CardHeader>
              <CardTitle className="text-lg font-semibold">
                {func.name}
              </CardTitle>
            </CardHeader>
            <CardContent className="space-y-4">
              {func.args.map((arg, idx) => (
                <div
                  key={`${func.name}-${arg.name}-${idx}`}
                  className="space-y-2"
                >
                  <Label>
                    {arg.name} ({arg.type})
                  </Label>
                  <Input
                    value={functionInputs[func.name]?.[idx] || ""}
                    onChange={(e) => {
                      const newInputs = { ...functionInputs }
                      newInputs[func.name][idx] = e.target.value
                      setFunctionInputs(newInputs)
                    }}
                    placeholder={`Enter ${arg.type}`}
                  />
                </div>
              ))}
              {func.payable && (
                <div className="space-y-2">
                  <Label>ETH Value (in wei)</Label>
                  <Input
                    value={ethValue[func.name] || ""}
                    onChange={(e) => {
                      setEthValue((prev) => ({
                        ...prev,
                        [func.name]: e.target.value,
                      }))
                    }}
                    placeholder="Enter ETH value in wei"
                  />
                </div>
              )}
              <Button
                onClick={() => handleExecuteFunction(func)}
                disabled={isExecuting[func.name]}
              >
                {isExecuting[func.name] ? "Executing..." : "Execute"}
              </Button>
            </CardContent>
          </Card>
        ))}
      </div>

      {canDelete && (
        <div className="flex justify-end">
          <Button
            onClick={handleDelete}
            disabled={isDeleting}
            variant="destructive"
          >
            {isDeleting ? "Deleting..." : "Delete Contract"}
          </Button>
        </div>
      )}
    </div>
  )
}
