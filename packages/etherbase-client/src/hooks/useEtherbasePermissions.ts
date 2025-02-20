"use client"

import { useCallback, useState } from "react"
import type { Address } from "viem"
import { parseEther } from "viem"
import { useEtherbaseContext } from "../EtherbaseProvider"
import { EtherbaseSourceAbi } from "../abi/EtherbaseSource"
import useWebThree from "./useWebThree"

export type UseEtherbasePermissionsProps = Readonly<{
  sourceAddress: Address
}>

interface Permission {
  walletAddress: string
  roles: number[]
  isOwner?: boolean
  balance?: bigint
}

export default function useEtherbasePermissions({
  sourceAddress,
}: UseEtherbasePermissionsProps) {
  useEtherbaseContext()
  const { publicClient, getWalletClient } = useWebThree()
  const [canGrantRoles, setCanGrantRoles] = useState<Record<number, boolean>>({})

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
      })
      const receipt = await publicClient.waitForTransactionReceipt({ hash })
      console.log("Execute write receipt:", receipt)
    },
    [sourceAddress, publicClient, getWalletClientInternal],
  )

  const grantRoles = useCallback(
    async (walletAddress: Address, roles: number[]) => {
      await executeWriteBrowser("createWalletIdentity", [walletAddress, roles])
    },
    [executeWriteBrowser],
  )

  const revokeIdentity = useCallback(
    async (walletAddress: Address) => {
      await executeWriteBrowser("deleteIdentity", [walletAddress])
    },
    [executeWriteBrowser],
  )

  const grantRole = useCallback(
    async (walletAddress: Address, role: number) => {
      await executeWriteBrowser("grantRole", [walletAddress, role])
    },
    [executeWriteBrowser],
  )

  const revokeRole = useCallback(
    async (walletAddress: Address, role: number) => {
      await executeWriteBrowser("revokeRole", [walletAddress, role])
    },
    [executeWriteBrowser],
  )

  const deposit = useCallback(
    async (targetAddress: Address) => {
      const walletClient = await getWalletClientInternal()
      const hash = await walletClient.sendTransaction({
        to: targetAddress,
        value: parseEther("1000"), // 10 SOM
        account: walletClient.account!,
        chain: walletClient.chain,
      })
      await publicClient?.waitForTransactionReceipt({ hash })
    },
    [publicClient, getWalletClientInternal],
  )

  const fetchPermissions = useCallback(async () => {
    if (!publicClient) {
      throw new Error("Public client not found")
    }

    // Get owner address
    const owner = await publicClient.readContract({
      address: sourceAddress,
      abi: EtherbaseSourceAbi,
      functionName: "owner",
    })

    // Fetch all identities with roles
    const identityViews = (await publicClient.readContract({
      address: sourceAddress,
      abi: EtherbaseSourceAbi,
      functionName: "getAllIdentities",
    })) as { walletAddress: string; roles: number[] }[]

    // Fetch balances for all addresses
    const balancePromises = identityViews.map((identity) =>
      publicClient.getBalance({
        address: identity.walletAddress as `0x${string}`,
      }),
    )
    const balances = await Promise.all(balancePromises)

    // Get current user's wallet to check permissions
    const accounts = await window.ethereum?.request({
      method: "eth_requestAccounts",
    }) as string[]
    const userWallet = accounts[0]
    const userIdentity = identityViews.find(
      (identity) =>
        identity.walletAddress.toLowerCase() === userWallet.toLowerCase(),
    )

    // For this example, we assume that only an identity with role 2 ("Grant") can grant permissions.
    const roleChecks = Array.from({ length: 4 }, (_, i) => [
      i,
      userIdentity ? userIdentity.roles.includes(2) : false,
    ])
    setCanGrantRoles(Object.fromEntries(roleChecks))

    const permissionsList = identityViews.map((identity, index) => ({
      walletAddress: identity.walletAddress,
      roles: identity.roles,
      isOwner: identity.walletAddress.toLowerCase() === (owner as string).toLowerCase(),
      balance: balances[index],
    }))

    return permissionsList
  }, [publicClient, sourceAddress])

  return {
    grantRoles,
    revokeIdentity,
    grantRole,
    revokeRole,
    deposit,
    fetchPermissions,
    canGrantRoles,
  }
} 