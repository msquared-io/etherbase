"use client"

import { useCallback, useMemo, useState } from "react"
import { createPublicClient, createWalletClient, custom } from "viem"
import type { Address, PublicClient, WalletClient } from "viem"
import { getConfig } from "../config"

type UseWebThreeReturn = {
  publicClient: PublicClient | null
  getWalletClient: () => Promise<WalletClient | undefined>
}

export default function useWebThree(): UseWebThreeReturn {
  const [walletClient, setWalletClient] = useState<WalletClient | null>(null)
  const { chain, privateKey } = getConfig()

  const publicClient = useMemo(() => {
    if (!window.ethereum) {
      return null
    }

    const pClient = createPublicClient({
      chain: chain,
      transport: custom(window.ethereum),
    })

    return pClient
  }, [chain])

  const getWalletClient = useCallback(async () => {
    if (walletClient) {
      // Already initialized
      console.log("Already initialized")
      return walletClient
    }

    if (!chain) {
      console.error(
        "No chain specified in config but is required for write operations",
      )
      return
    }

    if (!window.ethereum) {
      console.error("No Ethereum provider found")
      return
    }

    const accounts = (await window.ethereum.request({
      method: "eth_requestAccounts",
    })) as Address[]
    const account = accounts[0] as Address

    const wClient = createWalletClient({
      account: account,
      transport: custom(window.ethereum),
      chain: chain,
    })

    setWalletClient(wClient)
    return wClient
  }, [walletClient, chain])

  return { publicClient, getWalletClient }
}
