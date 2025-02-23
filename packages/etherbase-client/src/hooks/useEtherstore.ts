"use client"

import { useEffect, useMemo, useRef, useState } from "react"
import type { Address } from "viem"
import { useEtherbaseContext } from "../EtherbaseProvider"
import { getConfig } from "../config"
import { WebSocketManager } from "./WebSocketManager"
import type {
  EtherstoreHookReturn,
  EtherstoreState,
  EtherstoreValue,
  UseEtherstoreHookProps,
} from "./types/state"
import type { WebSocketStateUpdateInner } from "./types/websocket"
import useEtherbaseSource from "./useEtherbaseSource"

function deepMerge(
  target: EtherstoreState,
  source: EtherstoreState,
): EtherstoreState {
  if (!target) {
    return source as EtherstoreState
  }
  if (!source) {
    return target || {}
  }
  const output = { ...target }
  for (const key in source) {
    console.log("key", key, source[key], target[key])
    if (source[key] === null) {
      delete output[key]
    } else if (source[key] instanceof Object && key in target) {
      output[key] = deepMerge(
        target[key] as EtherstoreState,
        source[key] as EtherstoreState,
      )
    } else {
      output[key] = source[key] as EtherstoreValue
    }
  }
  return output
}

export default function useEtherstore({
  contractAddress,
  path,
  options = {},
  onStateChange,
}: UseEtherstoreHookProps): EtherstoreHookReturn {
  useEtherbaseContext()

  const [state, setState] = useState<EtherstoreState>({})
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)
  const keyRef = useRef<string>(Math.random().toString(36).substring(2, 9))

  if (!contractAddress) {
    throw new Error("Contract address must be provided in path")
  }

  // biome-ignore lint/correctness/useExhaustiveDependencies: need a stable reference
  const statePath = useMemo(() => path, [JSON.stringify(path)])
  // biome-ignore lint/correctness/useExhaustiveDependencies: need a stable reference
  const optionsMemo = useMemo(() => options, [JSON.stringify(options)])

  const { useBackend } = getConfig()
  const { setValue } = useEtherbaseSource({ sourceAddress: contractAddress })

  // state is the new state to be set. todo implement merge false
  const update = async (state: EtherstoreState, merge = true) => {
    // todo check if the contract is a source
    // if (!isSource(contractAddress)) {
    //   throw new Error(`Contract ${contractAddress} is not registered as an source`);
    // }

    if (!useBackend) {
      await setValue(state)
      return
    }

    console.log("Updating state", contractAddress, state)

    try {
      await WebSocketManager.get().setValue({
        contractAddress,
        state,
      })
    } catch (err) {
      setError(err instanceof Error ? err.message : String(err))
      throw err
    }
  }

  useEffect(() => {
    async function createSubscription() {
      try {
        await WebSocketManager.get().addStateSubscription(keyRef.current, {
          contractAddress: contractAddress,
          statePath: statePath,
          options: optionsMemo,
          onStateUpdate: (update: WebSocketStateUpdateInner) => {
            setLoading(false)

            // deep merge the update into the state
            // @ts-ignore
            setState((prev) => {
              const mergedState = deepMerge(prev, update.state)

              if (onStateChange) {
                onStateChange(mergedState)
              }

              return mergedState
            })
          },
        })
      } catch (err) {
        console.error("Error creating subscription", err)
        setError(err instanceof Error ? err.message : String(err))
        setLoading(false)
        return
      }
    }

    console.log("Creating subscription", contractAddress, statePath)
    createSubscription()

    return () => {
      WebSocketManager.get().removeSubscription(keyRef.current)
    }
  }, [contractAddress, statePath, optionsMemo, onStateChange])

  return { state, loading, error, update }
}
