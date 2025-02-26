"use client"

import { useCallback, useEffect, useMemo, useState } from "react"
import { useEtherbase, useEtherbaseEvents } from "@msquared/etherbase-client"
import type { EtherbaseEvent } from "@msquared/etherbase-client"
import { Button } from "@/components/ui/button"
import { Card, CardContent, CardHeader, CardTitle, CardDescription } from "@/components/ui/card"
import { useRouter } from "next/navigation"
import type { Address, AbiEvent, AbiParameter, AbiFunction } from "viem"
import { AnimatedList } from "@/components/ui/animated-list"
import { Separator } from "@/components/ui/separator"
import { parseAbiItem } from "viem"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { useWebThree } from "@msquared/etherbase-client"
import { ScrollArea } from "@/components/ui/scroll-area"
import { Badge } from "@/components/ui/badge"
import { 
  CodeIcon, 
  ArrowLeftIcon, 
  ChevronDownIcon, 
  ChevronUpIcon,
  FunctionSquareIcon,
  ActivityIcon,
  TrashIcon
} from "lucide-react"
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs"

interface ExtendedEvent extends EtherbaseEvent {
  index: number
  timestamp?: number
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

// Format timestamp to readable time - moved outside component
const formatTime = (timestamp?: number) => {
  if (!timestamp) return "";
  return new Date(timestamp).toLocaleTimeString([], {
    hour: "2-digit",
    minute: "2-digit",
    second: "2-digit",
  });
};

export default function CustomContractPage({
  params,
}: {
  params: { address: string }
}) {
  const router = useRouter()
  const { customContracts, fetchCustomContracts, deleteCustomContract } =
    useEtherbase()
  const { publicClient, getWalletClient } = useWebThree()
  
  // State hooks - all defined at the top of the component
  const [isDeleting, setIsDeleting] = useState(false)
  const [connectedAddress, setConnectedAddress] = useState<string>()
  const [latestEvents, setLatestEvents] = useState<Record<string, ExtendedEvent[]>>({})
  const [eventSchemas, setEventSchemas] = useState<EventSchema[]>([])
  const [writeFunctions, setWriteFunctions] = useState<WriteFunction[]>([])
  const [functionInputs, setFunctionInputs] = useState<Record<string, string[]>>({})
  const [ethValue, setEthValue] = useState<Record<string, string>>({})
  const [isExecuting, setIsExecuting] = useState<Record<string, boolean>>({})
  const [showAbi, setShowAbi] = useState(false)
  const [activeTab, setActiveTab] = useState<string>("events")
  const [searchTerm, setSearchTerm] = useState("")

  // Derived state using useMemo
  const contract = useMemo(() => 
    customContracts.find(
      (c) => c.contractAddress.toLowerCase() === params.address.toLowerCase()
    ), 
    [customContracts, params.address]
  )
  
  const canDelete = useMemo(() => 
    connectedAddress === contract?.addedBy.toLowerCase(),
    [connectedAddress, contract]
  )
  
  const allEvents = useMemo(() => {
    const events: ExtendedEvent[] = []
    for (const eventList of Object.values(latestEvents)) {
      events.push(...eventList)
    }
    // Sort by timestamp (newest first)
    return events.sort((a, b) => (b.timestamp || 0) - (a.timestamp || 0))
  }, [latestEvents])

  // Event handlers
  const handleNewEvent = useCallback((event: EtherbaseEvent) => {
    setLatestEvents((prev) => {
      const updated = { ...prev }
      const currentEvents = updated[event.name] || []
      const extendedEvent: ExtendedEvent = {
        ...event,
        index: currentEvents.length + 1,
        timestamp: Date.now(),
      }
      updated[event.name] = [extendedEvent, ...currentEvents]
      return updated
    })
  }, [])

  const handleExecuteFunction = useCallback(async (func: WriteFunction) => {
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
  }, [contract, publicClient, getWalletClient, functionInputs, ethValue])

  const handleDelete = useCallback(async () => {
    if (!contract) return
    
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
  }, [contract, deleteCustomContract, router])

  // Effects
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

  // Subscribe to contract events
  useEtherbaseEvents({
    contractAddress: params.address as Address,
    events: useMemo(
      () => eventSchemas.map((schema) => ({ name: schema.name })),
      [eventSchemas],
    ),
    onEvent: handleNewEvent,
  })

  // Early return for contract not found
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

  // Main render
  return (
    <div className="p-4">
      <div className="flex items-center justify-between mb-4">
        <h2 className="scroll-m-20 text-2xl font-semibold tracking-tight">
          Custom Contract: {contract.contractAddress}
        </h2>
        <Button
          onClick={() => router.push("/custom-contracts")}
          variant="outline"
          size="sm"
        >
          <ArrowLeftIcon className="h-4 w-4 mr-2" />
          Back to List
        </Button>
      </div>

      <Card className="mb-6">
        <CardHeader>
          <div className="flex justify-between items-center">
            <CardTitle className="text-lg font-semibold">
              Contract Information
            </CardTitle>
          </div>
          <CardDescription>
            View and interact with your custom contract
          </CardDescription>
        </CardHeader>
        <CardContent className="space-y-4">
          <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div>
              <Label className="text-sm text-muted-foreground">Contract Address</Label>
              <div className="font-mono text-sm bg-accent/5 p-2 rounded-md border border-accent/10 mt-1">
                {contract.contractAddress}
              </div>
            </div>
            <div>
              <Label className="text-sm text-muted-foreground">Added By</Label>
              <div className="font-mono text-sm bg-accent/5 p-2 rounded-md border border-accent/10 mt-1">
                {contract.addedBy}
              </div>
            </div>
          </div>
          
          <div>
            <div 
              className="flex items-center justify-between cursor-pointer p-2 rounded-md hover:bg-accent/5 transition-colors"
              onClick={() => setShowAbi(!showAbi)}
            >
              <Label className="text-sm font-medium cursor-pointer">Contract ABI</Label>
              <Button variant="ghost" size="sm">
                {showAbi ? <ChevronUpIcon className="h-4 w-4" /> : <ChevronDownIcon className="h-4 w-4" />}
              </Button>
            </div>
            
            {showAbi && (
              <ScrollArea className="h-[300px] border rounded-md mt-2">
                <pre className="p-4 text-xs font-mono">
                  {JSON.stringify(JSON.parse(contract.contractABI), null, 2)}
                </pre>
              </ScrollArea>
            )}
          </div>
        </CardContent>
      </Card>

      <Tabs defaultValue="events" onValueChange={setActiveTab} value={activeTab} className="space-y-4">
        <div className="flex justify-between items-center mb-4">
          <TabsList>
            <TabsTrigger value="events">
              <ActivityIcon className="h-4 w-4 mr-2" />
              Events
            </TabsTrigger>
            <TabsTrigger value="functions">
              <FunctionSquareIcon className="h-4 w-4 mr-2" />
              Write Functions
            </TabsTrigger>
          </TabsList>
          
          {activeTab === "events" && (
            <div className="flex items-center gap-2">
              <Input
                placeholder="Filter events..."
                className="w-48 h-8"
                value={searchTerm}
                onChange={(e) => setSearchTerm(e.target.value)}
              />
            </div>
          )}
        </div>

        <TabsContent value="events">
          <Card>
            <CardHeader>
              <div className="flex justify-between items-center">
                <CardTitle className="text-lg font-semibold">
                  All Events
                  <Badge variant="outline" className="ml-2">
                    {allEvents.length} events
                  </Badge>
                </CardTitle>
              </div>
              <CardDescription>
                View all events emitted by this contract
              </CardDescription>
            </CardHeader>
            <CardContent>
              {allEvents.length > 0 ? (
                <ScrollArea className="h-[400px] border rounded-md">
                  <div className="p-2">
                    <table className="w-full">
                      <thead className="sticky top-0 bg-background z-10">
                        <tr className="border-b">
                          <th className="text-left p-2 text-xs font-medium text-muted-foreground">
                            Event
                          </th>
                          <th className="text-left p-2 text-xs font-medium text-muted-foreground">
                            Time
                          </th>
                          <th className="text-left p-2 text-xs font-medium text-muted-foreground">
                            Block
                          </th>
                          <th className="text-left p-2 text-xs font-medium text-muted-foreground">
                            Arguments
                          </th>
                        </tr>
                      </thead>
                      <tbody>
                        {allEvents
                          .filter(
                            (evt) =>
                              !searchTerm ||
                              JSON.stringify(evt)
                                .toLowerCase()
                                .includes(searchTerm.toLowerCase()),
                          )
                          .map((evt) => (
                          <tr
                            key={`${evt.name}-${evt.index}`}
                            className="border-b border-accent/10 hover:bg-accent/5"
                          >
                            <td className="p-2 text-xs">
                              <Badge
                                variant="secondary"
                                className="font-mono"
                              >
                                {evt.name}
                              </Badge>
                            </td>
                            <td className="p-2 text-xs">
                              {formatTime(evt.timestamp)}
                            </td>
                            <td className="p-2 text-xs">{evt.block}</td>
                            <td className="p-2 text-xs">
                              <div className="flex flex-col gap-1">
                                {Object.entries(evt.args).map(
                                  ([key, value]) => (
                                    <div
                                      key={`${evt.name}-${evt.index}-${key}`}
                                    >
                                      <span className="font-medium">
                                        {key}:
                                      </span>{" "}
                                      <span className="font-mono">
                                        {String(value)}
                                      </span>
                                    </div>
                                  ),
                                )}
                              </div>
                            </td>
                          </tr>
                        ))}
                      </tbody>
                    </table>
                  </div>
                </ScrollArea>
              ) : (
                <div className="flex items-center justify-center h-[400px] border rounded-md bg-accent/5">
                  <p className="text-sm text-muted-foreground">
                    {searchTerm
                      ? "No matching events found"
                      : "No events received yet"}
                  </p>
                </div>
              )}
            </CardContent>
          </Card>

          <div className="mt-4 space-y-4">
            {eventSchemas.map((schema) => {
              const eventsForType = latestEvents[schema.name] || []
              const filteredEvents = searchTerm
                ? eventsForType.filter((evt: ExtendedEvent) =>
                    JSON.stringify(evt)
                      .toLowerCase()
                      .includes(searchTerm.toLowerCase()),
                  )
                : eventsForType
              
              return (
                <Card key={schema.name}>
                  <CardHeader>
                    <div className="flex justify-between items-center">
                      <CardTitle className="text-lg font-semibold">
                        {schema.name}
                        <Badge variant="outline" className="ml-2">
                          {eventsForType.length} events
                        </Badge>
                      </CardTitle>
                    </div>
                    <CardDescription>
                      View events of type {schema.name}
                    </CardDescription>
                  </CardHeader>
                  <CardContent>
                    <div className="grid grid-cols-1 md:grid-cols-5 gap-5">
                      <div className="md:col-span-1">
                        <h6 className="text-sm font-medium text-muted-foreground mb-2">
                          Parameters
                        </h6>
                        <div className="border rounded-md p-3 bg-accent/5">
                          {schema.args && schema.args.length > 0 ? (
                            schema.args.map((arg, idx) => (
                              <div
                                key={`${schema.name}-${arg.name}-${idx}`}
                                className="text-sm py-1 border-b last:border-0 border-accent/10"
                              >
                                <span className="font-medium">{arg.name}</span>{" "}
                                <span className="text-xs text-muted-foreground">
                                  ({arg.type})
                                  {arg.indexed && " (indexed)"}
                                </span>
                              </div>
                            ))
                          ) : (
                            <div className="text-sm text-muted-foreground">
                              No parameters
                            </div>
                          )}
                        </div>
                      </div>
                      <div className="md:col-span-4">
                        <h6 className="text-sm font-medium text-muted-foreground mb-2">
                          Events
                        </h6>

                        {filteredEvents.length > 0 ? (
                          <ScrollArea className="h-[400px] border rounded-md">
                            <div className="p-2">
                              <table className="w-full">
                                <thead className="sticky top-0 bg-background z-10">
                                  <tr className="border-b">
                                    <th className="text-left p-2 text-xs font-medium text-muted-foreground">
                                      #
                                    </th>
                                    <th className="text-left p-2 text-xs font-medium text-muted-foreground">
                                      Time
                                    </th>
                                    <th className="text-left p-2 text-xs font-medium text-muted-foreground">
                                      Block
                                    </th>
                                    {schema.args?.map((arg) => (
                                      <th
                                        key={arg.name}
                                        className="text-left p-2 text-xs font-medium text-muted-foreground"
                                      >
                                        {arg.name}
                                      </th>
                                    ))}
                                  </tr>
                                </thead>
                                <tbody>
                                  {filteredEvents.map((evt) => (
                                    <tr
                                      key={evt.index}
                                      className="border-b border-accent/10 hover:bg-accent/5"
                                    >
                                      <td className="p-2 text-xs">{evt.index}</td>
                                      <td className="p-2 text-xs">
                                        {formatTime(evt.timestamp)}
                                      </td>
                                      <td className="p-2 text-xs">{evt.block}</td>
                                      {schema.args?.map((arg) => (
                                        <td
                                          key={`${evt.index}-${arg.name}`}
                                          className="p-2 text-xs font-mono"
                                        >
                                          {String(evt.args[arg.name] ?? "")}
                                        </td>
                                      ))}
                                    </tr>
                                  ))}
                                </tbody>
                              </table>
                            </div>
                          </ScrollArea>
                        ) : (
                          <div className="flex items-center justify-center h-[400px] border rounded-md bg-accent/5">
                            <p className="text-sm text-muted-foreground">
                              {searchTerm
                                ? "No matching events found"
                                : "No events received yet"}
                            </p>
                          </div>
                        )}
                      </div>
                    </div>
                  </CardContent>
                </Card>
              )
            })}
          </div>
        </TabsContent>

        <TabsContent value="functions">
          <Card>
            <CardHeader>
              <div className="flex justify-between items-center">
                <CardTitle className="text-lg font-semibold">
                  Write Functions
                  <Badge variant="outline" className="ml-2">
                    {writeFunctions.length} functions
                  </Badge>
                </CardTitle>
              </div>
              <CardDescription>
                Execute contract functions that modify state
              </CardDescription>
            </CardHeader>
            <CardContent>
              {writeFunctions.length > 0 ? (
                <div className="space-y-6">
                  {writeFunctions.map((func) => (
                    <div key={func.name} className="border rounded-md p-4 bg-accent/5">
                      <h3 className="text-md font-semibold mb-3 flex items-center">
                        <FunctionSquareIcon className="h-4 w-4 mr-2 text-muted-foreground" />
                        {func.name}
                        {func.payable && (
                          <Badge variant="secondary" className="ml-2 text-xs">
                            Payable
                          </Badge>
                        )}
                      </h3>
                      
                      <div className="space-y-4">
                        {func.args.map((arg, idx) => (
                          <div
                            key={`${func.name}-${arg.name}-${idx}`}
                            className="space-y-2"
                          >
                            <Label className="text-sm">
                              {arg.name} <span className="text-xs text-muted-foreground">({arg.type})</span>
                            </Label>
                            <Input
                              value={functionInputs[func.name]?.[idx] || ""}
                              onChange={(e) => {
                                const newInputs = { ...functionInputs }
                                newInputs[func.name][idx] = e.target.value
                                setFunctionInputs(newInputs)
                              }}
                              placeholder={`Enter ${arg.type}`}
                              className="font-mono text-sm"
                            />
                          </div>
                        ))}
                        
                        {func.payable && (
                          <div className="space-y-2">
                            <Label className="text-sm">ETH Value (in wei)</Label>
                            <Input
                              value={ethValue[func.name] || ""}
                              onChange={(e) => {
                                setEthValue((prev) => ({
                                  ...prev,
                                  [func.name]: e.target.value,
                                }))
                              }}
                              placeholder="Enter ETH value in wei"
                              className="font-mono text-sm"
                            />
                          </div>
                        )}
                        
                        <Button
                          onClick={() => handleExecuteFunction(func)}
                          disabled={isExecuting[func.name]}
                          className="mt-2"
                        >
                          {isExecuting[func.name] ? "Executing..." : "Execute"}
                        </Button>
                      </div>
                    </div>
                  ))}
                </div>
              ) : (
                <div className="flex items-center justify-center h-[200px] border rounded-md bg-accent/5">
                  <p className="text-sm text-muted-foreground">
                    No write functions found in contract ABI
                  </p>
                </div>
              )}
            </CardContent>
          </Card>
        </TabsContent>
      </Tabs>

      {canDelete && (
        <div className="flex justify-end mt-6">
          <Button
            onClick={handleDelete}
            disabled={isDeleting}
            variant="destructive"
            size="sm"
          >
            <TrashIcon className="h-4 w-4 mr-2" />
            {isDeleting ? "Deleting..." : "Delete Contract"}
          </Button>
        </div>
      )}
    </div>
  )
}
