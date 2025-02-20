/// <reference path="./types/window.d.ts" />
"use client"

import { useCallback, useEffect, useState } from "react"
import type { Address } from "viem"
import { useEtherbaseContext } from "../EtherbaseProvider"
import { EtherbaseAbi } from "../abi/Etherbase"
import { getConfig } from "../config"
import { fetchSourcesData } from "../etherbase"
import type { Source } from "./types/etherbase"
import useWebThree from "./useWebThree"

export default function useEtherbase() {
  useEtherbaseContext()

  const { httpReaderUrl } = getConfig()
  const { publicClient, getWalletClient } = useWebThree()
  const [etherbaseAddress, setEtherbaseAddress] = useState<Address | null>(null)
  const [sources, setSources] = useState<Source[]>([])

  useEffect(() => {
    async function fetchEtherbaseAddress() {
      const etherbaseAddress = await fetch(`${httpReaderUrl}/etherbase-address`, {
        cache: 'no-store'
      }).then(
        (res) => res.json(),
      )
      console.log("Etherbase address:", etherbaseAddress)
      setEtherbaseAddress(etherbaseAddress)
    }
    void fetchEtherbaseAddress()
  }, [httpReaderUrl])

  const fetchSources = useCallback(async () => {
    const sources = await fetchSourcesData()
    setSources(sources)
  }, [])

  const getWalletClientInternal = useCallback(async () => {
    const walletClient = await getWalletClient()

    if (!walletClient) {
      throw new Error("Wallet client not found")
    }

    return walletClient
  }, [getWalletClient])

  const executeWrite = useCallback(
    async (writeFunctionName: string, args: unknown[]) => {
      if (!publicClient) {
        throw new Error("Public client not found")
      }

      const walletClient = await getWalletClientInternal()

      if (!etherbaseAddress) {
        throw new Error("Etherbase address not found")
      }

      if (
        !EtherbaseAbi.find((f: any) =>
          f.type === "function" && f.name === writeFunctionName
        )
      ) {
        throw new Error(`Invalid function name: ${writeFunctionName}`)
      }

      const hash = await walletClient.writeContract({
        address: etherbaseAddress,
        abi: EtherbaseAbi,
        // @ts-ignore
        functionName: writeFunctionName,
        // @ts-ignore
        args,
      })
      const receipt = await publicClient.waitForTransactionReceipt({ hash })
      console.log("Execute write receipt:", receipt)
      return receipt
    },
    [getWalletClientInternal, etherbaseAddress, publicClient],
  )

  const createSource = useCallback(async () => {
    if (!window.ethereum) {
      throw new Error("MetaMask not installed")
    }

    console.log("Creating source")
    const receipt = await executeWrite("createSource", [])

    // The new source address is in the first log
    const sourceAddress = receipt?.logs[0].address

    return { receipt, sourceAddress }
  }, [executeWrite])

  return { createSource, sources, fetchSources }
}
