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
import type { UseEtherbaseContractProps } from "./types/etherbase"
import type { Argument } from "./types/events"
import type { EtherstoreState } from "./types/state"
import useWebThree from "./useWebThree"

export default function useEtherbaseContract({
  contractAddress,
}: UseEtherbaseContractProps) {
  useEtherbaseContext()
  const { httpReaderUrl, wsReaderUrl, wsWriterUrl, useBackend, privateKey } =
    getConfig()

  const emitEventBackend = useCallback(
    async ({
      methodName,
      args,
    }: { methodName: string; args: Record<string, unknown> }) => {
      if (!privateKey) {
        throw new Error("No private key configured")
      }

      if (wsWriterUrl) {
        WebSocketManager.get().executeContractMethod({
          contractAddress,
          methodName,
          args,
        })
      } else {
        throw new Error("No backend ws URL configured")
      }
    },
    [privateKey, wsWriterUrl, contractAddress],
  )

  const executeContractMethod = useCallback(
    async ({
      methodName,
      args,
    }: { methodName: string; args: Record<string, unknown> }) => {
      if (useBackend) {
        console.log("executing contract method backend", methodName, args)
        await emitEventBackend({ methodName, args })
      } else {
        console.error(
          "executing contract method via browser is not supported yet",
        )
      }
    },
    [useBackend, emitEventBackend],
  )

  return { execute: executeContractMethod }
}
