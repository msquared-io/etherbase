/// <reference path="./types/window.d.ts" />
"use client"

import { useCallback, useEffect, useState } from "react"
import type { Abi, Address } from "viem"
import { useEtherbaseContext } from "../EtherbaseProvider"
import { EtherbaseAbi } from "../abi/Etherbase"
import { getConfig } from "../config"
import { fetchCustomContractsData, fetchSourcesData } from "../etherbase"
import type { CustomContract, Source } from "./types/etherbase"
import useWebThree from "./useWebThree"

export default function useEtherbase() {
  useEtherbaseContext()

  const { httpReaderUrl } = getConfig()
  const { publicClient, getWalletClient } = useWebThree()
  const [etherbaseAddress, setEtherbaseAddress] = useState<Address | null>(null)
  const [sources, setSources] = useState<Source[]>([])
  const [customContracts, setCustomContracts] = useState<CustomContract[]>([])

  useEffect(() => {
    async function fetchEtherbaseAddress() {
      const etherbaseAddress = await fetch(
        `${httpReaderUrl}/etherbase-address`,
        {
          cache: "no-store",
        },
      ).then((res) => res.json())
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
        !EtherbaseAbi.find(
          (f) => f.type === "function" && f.name === writeFunctionName,
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
        // gas: 2000000n,
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

    // Poll for the source to appear in the sources list
    const checkForSource = async (): Promise<void> => {
      const currentSources = await fetchSourcesData()
      setSources(currentSources)

      const sourceFound = currentSources.some(
        (source: Source) =>
          source.sourceAddress.toLowerCase() === sourceAddress?.toLowerCase(),
      )

      if (!sourceFound) {
        await new Promise((resolve) => setTimeout(resolve, 250))
        return checkForSource()
      }
    }

    if (sourceAddress) {
      await checkForSource()
    }

    return { receipt, sourceAddress }
  }, [executeWrite])

  const fetchCustomContracts = useCallback(async () => {
    const customContracts = await fetchCustomContractsData()
    setCustomContracts(customContracts)
  }, [])

  // Poll for custom contracts every second
  useEffect(() => {
    fetchCustomContracts()

    const intervalId = setInterval(() => {
      fetchCustomContracts()
    }, 1000)

    return () => clearInterval(intervalId)
  }, [fetchCustomContracts])

  const addCustomContract = useCallback(
    async (contractAddress: Address, contractAbi: string) => {
      if (!window.ethereum) {
        throw new Error("MetaMask not installed")
      }

      // validate contract is valid abi by checking there is at least one function and event
      console.log("contractAbi", contractAbi)
      const abi = JSON.parse(contractAbi) as Abi
      if (
        !abi.find((f) => f.type === "function") ||
        !abi.find((f) => f.type === "event")
      ) {
        throw new Error(
          "Invalid contract abi as it does not contain any functions or events. Make sure the abi you are providing is valid.",
        )
      }

      console.log("Adding custom contract")
      const receipt = await executeWrite("addCustomContract", [
        contractAddress,
        contractAbi,
      ])

      // Immediately fetch updated contracts
      await fetchCustomContracts()

      return receipt.status === "success"
    },
    [executeWrite, fetchCustomContracts],
  )

  const deleteCustomContract = useCallback(
    async (contractAddress: Address) => {
      const receipt = await executeWrite("deleteCustomContract", [
        contractAddress,
      ])

      // Immediately fetch updated contracts
      await fetchCustomContracts()

      return receipt.status === "success"
    },
    [executeWrite, fetchCustomContracts],
  )

  return {
    createSource,
    sources,
    fetchSources,
    customContracts,
    fetchCustomContracts,
    addCustomContract,
    deleteCustomContract,
  }
}
