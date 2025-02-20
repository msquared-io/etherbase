"use client"

import { useCallback, useState } from "react"
import type { AbiEvent, AbiParameter, Address } from "viem"
import {
  encodeAbiParameters,
  keccak256,
  pad,
  parseAbi,
  stringToBytes,
  toHex,
} from "viem"
import { useEtherbaseContext } from "../EtherbaseProvider"
import { EtherbaseSourceAbi } from "../abi/EtherbaseSource"
import { getConfig } from "../config"
import { fetchEventDefinitionsData } from "../etherbaseSource"
import { WebSocketManager } from "./WebSocketManager"
import type { EtherstoreState } from "./types/state"
import useWebThree from "./useWebThree"
import { Argument } from "./types/events"

export type UseEtherbaseSourceProps = Readonly<{
  sourceAddress: Address
}>

export default function useEtherbaseSource({
  sourceAddress,
}: UseEtherbaseSourceProps) {
  useEtherbaseContext()
  const { httpReaderUrl, wsReaderUrl, wsWriterUrl, useBackend, privateKey } = getConfig()
  const { publicClient, getWalletClient } = useWebThree()

  const getWalletClientInternal = useCallback(async () => {
    const walletClient = await getWalletClient()

    if (!walletClient) {
      throw new Error("Wallet client not found")
    }

    return walletClient
  }, [getWalletClient])

  const executeWriteBrowser = useCallback(
    async (
      writeFunctionName: string,
      args: unknown[],
    ) => {
      if (!publicClient) {
        throw new Error("Public client not found")
      }
      const walletClient = await getWalletClientInternal()

      // check the function name is valid
      if (
        !EtherbaseSourceAbi.find(
          (f: any) => f.type === "function" && f.name === writeFunctionName,
        )
      ) {
        throw new Error(`Invalid function name: ${writeFunctionName}`)
      }

      const hash = await walletClient.writeContract({
        address: sourceAddress,
        abi: EtherbaseSourceAbi,
        // @ts-ignore
        functionName: writeFunctionName,
        // @ts-ignore
        args,
        // currently gas estimation is broken in somnia so hardcoding a high value
        gas: 200000n,
      })
      const receipt = await publicClient.waitForTransactionReceipt({ hash })
      console.log("Execute write receipt:", receipt)
    },
    [sourceAddress, publicClient, getWalletClientInternal],
  )

  type EventDefinition = { name: string; args: Argument[] }

  const getEventAbi = useCallback((event: EventDefinition) => {
    let abi = `event ${event.name}(`
    for (const param of event.args) {
      abi += `${param.argType}`
      if (param.isIndexed) {
        abi += " indexed"
      }
      abi += ` ${param.name}`
      abi += ", "
    }
    if (event.args.length > 0) {
      abi = abi.slice(0, -2)
    }
    abi += ")"
    return abi
  }, [])

  const getEventTopic = useCallback((event: EventDefinition) => {
    return keccak256(
      stringToBytes(
        `${event.name}(${event.args.map((arg) => arg.argType).join(",")})`,
      ),
    )
  }, [])

  const parseEvent = useCallback(
    (event: EventDefinition) => {
      const abi = getEventAbi(event)
      const eventTopic = getEventTopic(event)
      const eventAbi = parseAbi([abi]) as AbiEvent[]
      const eventName = eventAbi[0].name
      const args = eventAbi[0].inputs.map(
        (input: AbiParameter & { indexed?: boolean }) => ({
          name: input.name ?? "",
          argType: input.type,
          isIndexed: !!input?.indexed,
        }),
      )
      return { eventName, eventTopic, args }
    },
    [getEventAbi, getEventTopic],
  )

  const getEventId = useCallback(
    ({
      eventTopic,
      numIndexedArgs,
    }: { eventTopic: string; numIndexedArgs: number }): string => {
      return `${eventTopic}:${numIndexedArgs}`
    },
    [],
  )

  const registerEventBrowser = useCallback(
    async (event: { name: string; args: Argument[] }) => {
      const { eventName, eventTopic, args } = parseEvent(event)

      const id = getEventId({
        eventTopic,
        numIndexedArgs: args.filter((arg: Argument) => arg.isIndexed).length,
      })

      console.log("registering event", eventName, id, eventTopic, args)
      await executeWriteBrowser(
        "registerEventSchema",
        [eventName, id, eventTopic, args],
      )
    },
    [executeWriteBrowser, getEventId, parseEvent],
  )

  const encodeEvent = useCallback(
    async (eventName: string, args: Record<string, unknown>) => {
      if (!publicClient) {
        throw new Error("Public client not found")
      }
      // fetch schema from the contract
      const eventSchema = await publicClient.readContract({
        address: sourceAddress,
        abi: EtherbaseSourceAbi,
        functionName: "getEventSchemaFromName",
        args: [eventName],
      })
      if (!eventSchema) {
        throw new Error(`Event schema not found for ${eventName}`)
      }

      const topics: string[] = []
      const dataTypes: string[] = []
      const dataValues: unknown[] = []

      // Validate number of parameters
      // Ensure parameterValues is an object
      if (typeof args !== "object" || Array.isArray(args)) {
        throw new Error("Parameter values must be an object")
      }

      // Convert parameter values object to array
      const parameterValues = eventSchema.args.map(
        (arg: Argument) => args[arg.name],
      )
      if (parameterValues.length !== eventSchema.args.length) {
        throw new Error("Event arguments/values length mismatch")
      }

      // Process each parameter
      eventSchema.args.forEach((arg: Argument, index: number) => {
        const value = parameterValues[index]

        if (arg.isIndexed) {
          if (arg.argType === "string") {
            topics.push(keccak256(stringToBytes(value as string)))
          } else if (arg.argType.startsWith("bytes")) {
            topics.push(keccak256(value as `0x${string}`))
          } else if (arg.argType === "address") {
            topics.push(pad(value as Address, { size: 32 }))
          } else if (arg.argType === "uint256" || arg.argType === "int256") {
            topics.push(pad(toHex(value as bigint), { size: 32 }))
          } else if (arg.argType === "bool") {
            topics.push(pad(value ? "0x01" : "0x00", { size: 32 }))
          } else {
            topics.push(pad(toHex(value as bigint), { size: 32 }))
          }
        } else {
          dataTypes.push(arg.argType)
          dataValues.push(value)
        }
      })

      const data = encodeAbiParameters(
        dataTypes.map((type) => ({ type })),
        dataValues,
      )

      return { parameterTopics: topics, parameterData: data }
    },
    [publicClient, sourceAddress],
  )

  const emitEventBrowser = useCallback(
    async ({ name, args }: { name: string; args: Record<string, unknown> }) => {
      const { parameterTopics, parameterData } = await encodeEvent(name, args)

      await executeWriteBrowser("emitEvent", [
        name,
        parameterTopics,
        parameterData,
      ])
    },
    [encodeEvent, executeWriteBrowser],
  )

  const emitEventBackend = useCallback(
    async ({ name, args }: { name: string; args: Record<string, unknown> }) => {
      if (!privateKey) {
        throw new Error("No private key configured")
      }

      // Use WebSocket if available, fallback to HTTP
      if (wsWriterUrl) {
        WebSocketManager.get().emitEvent({
          sourceAddress,
          name,
          args,
        })
      } else if (httpReaderUrl) {
        const response = await fetch(`${httpReaderUrl}/emit-event`, {
          method: "POST",
          body: JSON.stringify({ sourceAddress, name, args }),
          headers: {
            "x-private-key": privateKey,
          },
        })
        const data = await response.json()
        console.log("Emit event response:", data)
      } else {
        throw new Error("No backend URL configured")
      }
    },
    [privateKey, httpReaderUrl, wsWriterUrl, sourceAddress],
  )

  const emitEvent = useCallback(
    async ({ name, args }: { name: string; args: Record<string, unknown> }) => {
      if (useBackend) {
        console.log("emitting event backend", name, args)
        await emitEventBackend({ name, args })
      } else {
        console.log("emitting event browser", name, args)
        await emitEventBrowser({ name, args })
      }
    },
    [useBackend, emitEventBackend, emitEventBrowser],
  )

  enum DataType {
    NONE = 0,
    STRING = 1,
    BOOL = 2,
    UINT256 = 3,
    INT256 = 4,
    BYTES = 5,
  }

  type BatchSetValue = {
    segments: string[]
    dataType: DataType
    data: string
  }

  const setValueBrowser = useCallback(
    async (state: EtherstoreState) => {
      const batchSetValues: BatchSetValue[] = []
      const processState = (
        obj: EtherstoreState,
        path: string[] = [],
      ): void => {
        for (const [key, value] of Object.entries(obj)) {
          const currentPath = path.length === 0 ? [key] : [...path, key]
          console.log("currentPath", currentPath)

          if (value === null || value === undefined) {
            batchSetValues.push({
              segments: currentPath,
              dataType: DataType.NONE,
              data: "0x",
            })
          } else if (typeof value === "object") {
            processState(value, currentPath)
          } else if (typeof value === "string") {
            batchSetValues.push({
              segments: currentPath,
              dataType: DataType.STRING,
              data: toHex(value),
            })
          } else if (typeof value === "boolean") {
            batchSetValues.push({
              segments: currentPath,
              dataType: DataType.BOOL,
              data: value ? "0x01" : "0x00",
            })
          } else if (typeof value === "number") {
            if (Number.isInteger(value)) {
              const hexValue = toHex(BigInt(value), { size: 32 })
              batchSetValues.push({
                segments: currentPath,
                dataType: value >= 0 ? DataType.UINT256 : DataType.INT256,
                data: hexValue,
              })
            } else {
              throw new Error("Only integer numbers are supported")
            }
          } else {
            throw new Error(`Unsupported value type: ${typeof value}`)
          }
        }
      }

      console.log("state", state)
      processState(state)
      console.log("batchSetValues", batchSetValues)
      await executeWriteBrowser("setValues", [batchSetValues])
    },
    [executeWriteBrowser],
  )

  const setValueBackend = useCallback(
    async (state: EtherstoreState) => {
      if (!privateKey) return

      // Use WebSocket if available, fallback to HTTP
      if (wsWriterUrl) {
        WebSocketManager.get().setValue({
          contractAddress: sourceAddress,
          state,
        })
      } else if (httpReaderUrl) {
        const response = await fetch(`${httpReaderUrl}/set-value`, {
          method: "POST",
          body: JSON.stringify({ state }),
          headers: {
            "x-private-key": privateKey,
          },
        })
        const data = await response.json()
        console.log("Set value response:", data)
      } else {
        throw new Error("No backend URL configured")
      }
    },
    [privateKey, httpReaderUrl, wsWriterUrl, sourceAddress],
  )

  const setValue = useCallback(
    async (state: EtherstoreState) => {
      if (useBackend) {
        await setValueBackend(state)
      } else {
        await setValueBrowser(state)
      }
    },
    [useBackend, setValueBackend, setValueBrowser],
  )

  const [eventDefinitions, setEventDefinitions] = useState<EventDefinition[]>(
    [],
  )

  const fetchEventDefinitions = useCallback(async () => {
    const data = await fetchEventDefinitionsData(sourceAddress)
    setEventDefinitions(data)
  }, [sourceAddress])

  return {
    registerEvent: registerEventBrowser,
    emitEvent,
    setValue,
    eventDefinitions,
    fetchEventDefinitions,
  }
}
