"use client"

import { useCallback, useEffect, useMemo, useRef, useState } from "react"
import { useEtherbaseContext } from "../EtherbaseProvider"
import { WebSocketManager } from "./WebSocketManager"
import type {
  EtherbaseEvent,
  UseEtherbaseEventsProps,
  UseEtherbaseEventsReturn,
} from "./types/events"

export default function useEtherbaseEvents({
  contractAddress,
  events,
  onEvent,
}: UseEtherbaseEventsProps): UseEtherbaseEventsReturn {
  useEtherbaseContext()

  const [error, setError] = useState<string | null>(null)
  const localIdRef = useRef<string>(Math.random().toString(36).substring(2, 15))

  // Stabilize the error handling wrapper function
  const handleEvent = useCallback(
    (event: EtherbaseEvent) => {
      try {
        onEvent(event)
      } catch (err) {
        setError(err instanceof Error ? err.message : String(err))
      }
    },
    [onEvent],
  )

  // biome-ignore lint/correctness/useExhaustiveDependencies: want to minimize re-subscriptions
  useEffect(() => {
    if (!contractAddress) {
      setError("No contract address provided")
      return
    }

    const localId = localIdRef.current

    WebSocketManager.get()
      .addEventSubscription(localId, {
        contractAddress,
        events,
        onEvent: handleEvent,
      })
      .catch((err) => {
        setError(err instanceof Error ? err.message : String(err))
      })

    return () => {
      WebSocketManager.get().removeSubscription(localId)
    }
  }, [contractAddress, JSON.stringify(events), handleEvent])

  return { error }
}
