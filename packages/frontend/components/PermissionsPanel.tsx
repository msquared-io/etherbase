"use client"

import { useState, useEffect, Fragment } from "react"
import type { Address } from "viem"
import { formatEther } from "viem"
import { generatePrivateKey, privateKeyToAccount } from "viem/accounts"
import { useEtherbasePermissions } from "@msquared/etherbase-client"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { CheckIcon, ChevronDownIcon, TrashIcon } from "lucide-react"
import { Listbox, Transition } from "@headlessui/react"
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog"

// Roles data
const ROLES = [
  {
    key: 0,
    name: "Emit",
    description: "Can emit events",
  },
  {
    key: 1,
    name: "Define",
    description: "Can define new events",
  },
  {
    key: 2,
    name: "Grant",
    description: "Can grant or revoke permissions",
  },
  {
    key: 3,
    name: "Database Setter",
    description: "Can set values in the database",
  },
] as const

interface Permission {
  walletAddress: string
  roles: number[]
  isOwner?: boolean
  balance?: bigint
}

interface PermissionsPanelProps {
  sourceAddress: string
}

/**
 * A simple multi-select component built with Headless UI's Listbox.
 */
interface MultiSelectProps {
  options: typeof ROLES
  selected: number[]
  onChange: (value: number[]) => void
  disabled?: boolean
}

function MultiSelect({
  options,
  selected,
  onChange,
  disabled,
}: MultiSelectProps) {
  return (
    <Listbox value={selected} onChange={onChange} multiple disabled={disabled}>
      <div className="relative">
        <Listbox.Button className="relative w-full cursor-default rounded border border-gray-300 bg-black text-white py-2 pl-3 pr-10 text-left focus:outline-none focus:ring-1 focus:ring-blue-500">
          <span className="block truncate">
            {selected.length
              ? selected
                  .map(
                    (roleKey) =>
                      options.find((r) => r.key === roleKey)?.name ||
                      roleKey.toString(),
                  )
                  .join(", ")
              : "None"}
          </span>
          <span className="pointer-events-none absolute inset-y-0 right-0 flex items-center pr-2">
            <ChevronDownIcon className="h-5 w-5 text-gray-400" />
          </span>
        </Listbox.Button>
        <Transition
          as={Fragment}
          leave="transition ease-in duration-100"
          leaveFrom="opacity-100"
          leaveTo="opacity-0"
        >
          <Listbox.Options className="absolute z-10 mt-1 max-h-60 w-full overflow-auto rounded border border-gray-300 bg-black text-white py-1 text-base shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none sm:text-sm">
            {options.map((role) => (
              <Listbox.Option
                key={role.key}
                value={role.key}
                disabled={false}
                className={({ active, disabled }) =>
                  `relative cursor-default select-none py-2 pl-10 pr-4 text-white ${
                    active ? "bg-gray-600" : ""
                  }`
                }
              >
                {({ selected: isSelected, active }) => (
                  <>
                    <span
                      className={`block truncate ${isSelected ? "font-medium" : "font-normal"}`}
                    >
                      {role.name}
                    </span>
                    {isSelected ? (
                      <span
                        className={`absolute inset-y-0 left-0 flex items-center pl-3 ${
                          active ? "text-white" : "text-blue-600"
                        }`}
                      >
                        <CheckIcon className="h-5 w-5" />
                      </span>
                    ) : null}
                    <span className="block text-xs text-gray-500">
                      {role.description}
                    </span>
                  </>
                )}
              </Listbox.Option>
            ))}
          </Listbox.Options>
        </Transition>
      </div>
    </Listbox>
  )
}

interface PrivateKeyModalProps {
  open: boolean
  onClose: () => void
  privateKey: string
  walletAddress: string
}

interface DepositModalProps {
  open: boolean
  onClose: () => void
  onDeposit: (amount: string) => void
}

function DepositModal({ open, onClose, onDeposit }: DepositModalProps) {
  const [amount, setAmount] = useState("")

  return (
    <Dialog open={open} onOpenChange={onClose}>
      <DialogContent className="sm:max-w-[425px] bg-black text-white">
        <DialogHeader>
          <DialogTitle>Deposit SOM</DialogTitle>
          <DialogDescription>
            Enter the amount of SOM to deposit
          </DialogDescription>
        </DialogHeader>
        <div className="grid gap-4 py-4">
          <div className="space-y-4">
            <div>
              <Label>Amount (SOM)</Label>
              <Input
                type="number"
                step="0.000000000000000001"
                min="0"
                value={amount}
                onChange={(e) => setAmount(e.target.value)}
                placeholder="Enter amount"
                className="font-mono text-sm"
              />
            </div>
          </div>
        </div>
        <DialogFooter>
          <Button
            onClick={() => {
              onDeposit(amount)
              onClose()
            }}
            className="bg-white text-black"
            disabled={!amount || parseFloat(amount) <= 0}
          >
            Deposit
          </Button>
          <Button onClick={onClose} variant="outline">
            Cancel
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  )
}

function PrivateKeyModal({
  open,
  onClose,
  privateKey,
  walletAddress,
}: PrivateKeyModalProps) {
  return (
    <Dialog open={open} onOpenChange={onClose}>
      <DialogContent className="sm:max-w-[425px] bg-black text-white">
        <DialogHeader>
          <DialogTitle>Save Your Wallet Details</DialogTitle>
          <DialogDescription className="text-red-500">
            Important: Save these details now. The private key will never be
            shown again!
          </DialogDescription>
        </DialogHeader>
        <div className="grid gap-4 py-4">
          <div className="space-y-4">
            <div>
              <Label>Wallet Address</Label>
              <Input
                readOnly
                value={walletAddress}
                className="font-mono text-sm"
              />
            </div>
            <div>
              <Label>Private Key</Label>
              <Input
                readOnly
                value={privateKey}
                className="font-mono text-sm"
              />
            </div>
          </div>
        </div>
        <DialogFooter>
          <Button
            onClick={() => {
              const text = `Wallet Address: ${walletAddress}\nPrivate Key: ${privateKey}`
              navigator.clipboard.writeText(text)
              alert("Wallet details copied to clipboard!")
            }}
            className="bg-white text-black"
          >
            Copy All to Clipboard
          </Button>
          <Button onClick={onClose} variant="outline">
            Close
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  )
}

export default function PermissionsPanel({
  sourceAddress,
}: PermissionsPanelProps) {
  const [newAddress, setNewAddress] = useState("")
  const [selectedRoles, setSelectedRoles] = useState<number[]>([])
  const [permissions, setPermissions] = useState<Permission[]>([])
  const [createNewWallet, setCreateNewWallet] = useState(false)
  const [showPrivateKeyModal, setShowPrivateKeyModal] = useState(false)
  const [generatedPrivateKey, setGeneratedPrivateKey] = useState("")
  const [generatedWalletAddress, setGeneratedWalletAddress] = useState("")
  const [showDepositModal, setShowDepositModal] = useState(false)
  const [depositTargetAddress, setDepositTargetAddress] = useState<Address>()

  const {
    grantRoles,
    revokeIdentity,
    grantRole,
    revokeRole,
    deposit,
    fetchPermissions: fetchPermissionsData,
    canGrantRoles,
  } = useEtherbasePermissions({ sourceAddress: sourceAddress as Address })

  const handleGrantRoles = async () => {
    try {
      let targetAddress: string
      let newPrivateKey: string | undefined

      if (createNewWallet) {
        // Generate new wallet
        newPrivateKey = generatePrivateKey()
        const account = privateKeyToAccount(newPrivateKey as `0x${string}`)
        targetAddress = account.address
      } else {
        targetAddress = newAddress
      }

      // Create wallet identity with initial roles
      try {
        await grantRoles(targetAddress as Address, selectedRoles)

        // Only show private key modal after successful registration
        if (createNewWallet && newPrivateKey) {
          setGeneratedPrivateKey(newPrivateKey)
          setGeneratedWalletAddress(targetAddress)
          setShowPrivateKeyModal(true)
        }

        await fetchPermissionsData()
        setNewAddress("")
        setSelectedRoles([])
        setCreateNewWallet(false)
      } catch (error) {
        throw error
      }
    } catch (error) {
      console.error("Error granting roles:", error)
      alert(error instanceof Error ? error.message : "Failed to grant roles")
    }
  }

  useEffect(() => {
    fetchPermissionsData().then(setPermissions)
  }, [sourceAddress, fetchPermissionsData])

  const canGrantAnyRole = Object.values(canGrantRoles).some(Boolean)

  return (
    <div className="mt-8 space-y-6">
      <h2 className="text-2xl font-semibold">Permissions</h2>

      {canGrantAnyRole && (
        <div className="flex flex-col gap-4">
          <div className="flex flex-col sm:flex-row gap-4 items-end">
            <div className="flex-1">
              <div className="mb-4">
                <Label className="mb-2">Wallet Type</Label>
                <div className="flex gap-4">
                  <Button
                    variant={!createNewWallet ? "default" : "outline"}
                    onClick={() => setCreateNewWallet(false)}
                    className={!createNewWallet ? "bg-white text-black" : ""}
                  >
                    Existing Wallet
                  </Button>
                  <Button
                    variant={createNewWallet ? "default" : "outline"}
                    onClick={() => setCreateNewWallet(true)}
                    className={createNewWallet ? "bg-white text-black" : ""}
                  >
                    Create New Wallet
                  </Button>
                </div>
              </div>
              {!createNewWallet && (
                <>
                  <Label htmlFor="new-address">Wallet Address</Label>
                  <Input
                    id="new-address"
                    value={newAddress}
                    onChange={(e) => setNewAddress(e.target.value)}
                    placeholder="Enter wallet address"
                  />
                  <p className="text-sm text-gray-500">
                    Enter wallet address to grant permissions
                  </p>
                </>
              )}
            </div>
            <div className="flex-1">
              <Label>Roles</Label>
              <MultiSelect
                options={ROLES}
                selected={selectedRoles}
                onChange={setSelectedRoles}
              />
            </div>
            <Button
              onClick={handleGrantRoles}
              disabled={
                (!createNewWallet && !newAddress) || selectedRoles.length === 0
              }
              className="whitespace-nowrap bg-white text-black"
            >
              Grant Roles
            </Button>
          </div>
        </div>
      )}

      <div className="border border-gray-300 rounded p-4">
        {permissions.map((permission) => (
          <div
            key={permission.walletAddress}
            className={`mb-4 p-4 border-b border-gray-200 flex justify-between items-center`}
          >
            <div>
              <div className="flex items-center gap-2">
                <p className="text-base">
                  {permission.walletAddress}{" "}
                  {permission.isOwner && (
                    <span className="ml-2 rounded bg-blue-500 px-2 py-1 text-xs text-black">
                      Owner
                    </span>
                  )}
                </p>
                <span className="text-sm text-gray-400">
                  Balance:{" "}
                  {permission.balance ? formatEther(permission.balance) : "0"}{" "}
                  SOM
                </span>
                <Button
                  onClick={() => {
                    setDepositTargetAddress(permission.walletAddress as Address)
                    setShowDepositModal(true)
                  }}
                  className="ml-2 bg-green-600 hover:bg-green-700"
                  size="sm"
                >
                  Deposit SOM
                </Button>
              </div>
              <div className="mt-2">
                {/* For each identity, we allow updating their roles via a MultiSelect. */}
                <MultiSelect
                  options={ROLES}
                  selected={permission.roles}
                  onChange={async (newRoles: number[]) => {
                    if (permission.isOwner) return
                    const addedRoles = newRoles.filter(
                      (r) => !permission.roles.includes(r),
                    )
                    const removedRoles = permission.roles.filter(
                      (r) => !newRoles.includes(r),
                    )

                    try {
                      for (const role of addedRoles) {
                        await grantRole(
                          permission.walletAddress as Address,
                          role,
                        )
                      }
                      for (const role of removedRoles) {
                        await revokeRole(
                          permission.walletAddress as Address,
                          role,
                        )
                      }

                      await fetchPermissionsData()
                    } catch (error) {
                      console.error("Error updating roles:", error)
                      alert(
                        error instanceof Error
                          ? error.message
                          : "Failed to update roles",
                      )
                    }
                  }}
                  disabled={permission.isOwner}
                />
              </div>
            </div>
            {!permission.isOwner && (
              <Button
                onClick={() =>
                  revokeIdentity(permission.walletAddress as Address)
                }
                variant="destructive"
                title="Revoke Identity"
                size="icon"
              >
                <TrashIcon className="h-5 w-5" />
              </Button>
            )}
          </div>
        ))}
      </div>

      <PrivateKeyModal
        open={showPrivateKeyModal}
        onClose={() => setShowPrivateKeyModal(false)}
        privateKey={generatedPrivateKey}
        walletAddress={generatedWalletAddress}
      />

      <DepositModal
        open={showDepositModal}
        onClose={() => {
          setShowDepositModal(false)
          setDepositTargetAddress(undefined)
        }}
        onDeposit={(amount) => {
          if (depositTargetAddress) {
            deposit(depositTargetAddress, parseFloat(amount))
          }
        }}
      />
    </div>
  )
}
