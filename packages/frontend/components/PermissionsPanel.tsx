"use client"

import { useState, useEffect, Fragment } from "react"
import type { Address } from "viem"
import { formatEther } from "viem"
import { generatePrivateKey, privateKeyToAccount } from "viem/accounts"
import { useEtherbasePermissions } from "@msquared/etherbase-client"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { 
  CheckIcon, 
  ChevronDownIcon, 
  TrashIcon, 
  CopyIcon, 
  InfoIcon,
  WalletIcon,
  PlusIcon,
  ShieldIcon
} from "lucide-react"
import { Listbox, Transition } from "@headlessui/react"
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog"
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"
import { Badge } from "@/components/ui/badge"
import { Separator } from "@/components/ui/separator"
import { ScrollArea } from "@/components/ui/scroll-area"

// Roles data with added color information
const ROLES = [
  {
    key: 0,
    name: "Emit",
    description: "Can emit events",
    color: "bg-accent/80",
    textColor: "text-accent-foreground",
  },
  {
    key: 1,
    name: "Define",
    description: "Can define new events",
    color: "bg-accent/80",
    textColor: "text-accent-foreground",
  },
  {
    key: 2,
    name: "Grant",
    description: "Can grant or revoke permissions",
    color: "bg-accent/80",
    textColor: "text-accent-foreground",
  },
  {
    key: 3,
    name: "Database Setter",
    description: "Can set values in the database",
    color: "bg-accent/80",
    textColor: "text-accent-foreground",
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

// Helper function to truncate wallet addresses
function truncateAddress(address: string): string {
  if (!address) return "";
  return `${address.substring(0, 6)}...${address.substring(address.length - 4)}`;
}

// Role badge component for consistent role display
function RoleBadge({ roleKey }: { roleKey: number }) {
  const role = ROLES.find(r => r.key === roleKey);
  if (!role) return null;
  
  return (
    <Badge 
      variant="secondary"
      className="font-mono"
      title={role.description}
    >
      {role.name}
    </Badge>
  );
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
        <Listbox.Button className="relative w-full cursor-default rounded-md border border-input bg-background text-foreground py-2 pl-3 pr-10 text-left focus:outline-none focus:ring-2 focus:ring-ring shadow-sm transition-all hover:border-accent">
          <span className="block truncate">
            {selected.length
              ? selected
                  .map(
                    (roleKey) =>
                      options.find((r) => r.key === roleKey)?.name ||
                      roleKey.toString(),
                  )
                  .join(", ")
              : "Select roles..."}
          </span>
          <span className="pointer-events-none absolute inset-y-0 right-0 flex items-center pr-2">
            <ChevronDownIcon className="h-5 w-5 text-muted-foreground" />
          </span>
        </Listbox.Button>
        <Transition
          as={Fragment}
          leave="transition ease-in duration-100"
          leaveFrom="opacity-100"
          leaveTo="opacity-0"
        >
          <Listbox.Options className="absolute z-10 mt-1 max-h-60 w-full overflow-auto rounded-md border border-input bg-background text-foreground py-1 text-base shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none sm:text-sm">
            {options.map((role) => (
              <Listbox.Option
                key={role.key}
                value={role.key}
                disabled={false}
                className={({ active, disabled }) =>
                  `relative cursor-default select-none py-2 pl-10 pr-4 ${
                    active ? "bg-accent/5" : ""
                  } ${disabled ? "opacity-50" : ""} transition-colors`
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
                          active ? "text-foreground" : "text-muted-foreground"
                        }`}
                      >
                        <CheckIcon className="h-5 w-5" />
                      </span>
                    ) : null}
                    <span className="block text-xs text-muted-foreground">
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
      <DialogContent className="sm:max-w-[425px] bg-background text-foreground">
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
      <DialogContent className="sm:max-w-[425px] bg-background text-foreground">
        <DialogHeader>
          <DialogTitle>Save Your Wallet Details</DialogTitle>
          <DialogDescription className="text-destructive">
            Important: Save these details now. The private key will never be
            shown again!
          </DialogDescription>
        </DialogHeader>
        <div className="grid gap-4 py-4">
          <div className="space-y-4">
            <div>
              <Label>Wallet Address</Label>
              <div className="relative">
                <Input
                  readOnly
                  value={walletAddress}
                  className="font-mono text-sm pr-10"
                />
                <Button
                  size="sm"
                  variant="ghost"
                  className="absolute right-1 top-1/2 -translate-y-1/2 h-7 w-7 p-0"
                  onClick={() => {
                    navigator.clipboard.writeText(walletAddress)
                    alert("Wallet address copied to clipboard!")
                  }}
                >
                  <CopyIcon className="h-4 w-4" />
                </Button>
              </div>
            </div>
            <div>
              <Label>Private Key</Label>
              <div className="relative">
                <Input
                  readOnly
                  value={privateKey}
                  className="font-mono text-sm pr-10"
                  type="password"
                />
                <Button
                  size="sm"
                  variant="ghost"
                  className="absolute right-1 top-1/2 -translate-y-1/2 h-7 w-7 p-0"
                  onClick={() => {
                    navigator.clipboard.writeText(privateKey)
                    alert("Private key copied to clipboard!")
                  }}
                >
                  <CopyIcon className="h-4 w-4" />
                </Button>
              </div>
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
      <div className="flex items-center gap-2 mb-4">
        <ShieldIcon className="h-6 w-6 text-muted-foreground" />
        <h2 className="text-2xl font-semibold">Permissions Management</h2>
      </div>

      {canGrantAnyRole && (
        <Card>
          <CardHeader>
            <CardTitle className="text-lg font-semibold">Grant Permissions</CardTitle>
            <CardDescription>
              Add new wallets or grant roles to existing wallets
            </CardDescription>
          </CardHeader>
          <CardContent>
            <div className="flex flex-col gap-4">
              <div className="mb-4">
                <Label className="mb-2">Wallet Type</Label>
                <div className="flex gap-4">
                  <Button
                    variant={!createNewWallet ? "default" : "outline"}
                    onClick={() => setCreateNewWallet(false)}
                  >
                    <WalletIcon className="h-4 w-4 mr-2" />
                    Existing Wallet
                  </Button>
                  <Button
                    variant={createNewWallet ? "default" : "outline"}
                    onClick={() => setCreateNewWallet(true)}
                  >
                    <PlusIcon className="h-4 w-4 mr-2" />
                    Create New Wallet
                  </Button>
                </div>
              </div>
              
              <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div>
                  {!createNewWallet ? (
                    <>
                      <Label htmlFor="new-address" className="mb-2 block">Wallet Address</Label>
                      <Input
                        id="new-address"
                        value={newAddress}
                        onChange={(e) => setNewAddress(e.target.value)}
                        placeholder="Enter wallet address (0x...)"
                        className="font-mono"
                      />
                      <p className="text-sm text-muted-foreground mt-1">
                        Enter the wallet address to grant permissions
                      </p>
                    </>
                  ) : (
                    <div className="flex items-center h-full">
                      <div className="text-sm text-muted-foreground bg-accent/5 p-4 rounded-md border border-accent/10">
                        <InfoIcon className="h-5 w-5 text-muted-foreground inline mr-2" />
                        A new wallet will be generated with a private key. Make sure to save the private key when prompted.
                      </div>
                    </div>
                  )}
                </div>
                
                <div>
                  <Label className="mb-2 block">Roles to Grant</Label>
                  <MultiSelect
                    options={ROLES}
                    selected={selectedRoles}
                    onChange={setSelectedRoles}
                  />
                  <div className="flex flex-wrap gap-2 mt-2">
                    {selectedRoles.map(roleKey => (
                      <RoleBadge key={roleKey} roleKey={roleKey} />
                    ))}
                  </div>
                </div>
              </div>
            </div>
          </CardContent>
          <CardFooter className="justify-end">
            <Button
              onClick={handleGrantRoles}
              disabled={
                (!createNewWallet && !newAddress) || selectedRoles.length === 0
              }
            >
              <ShieldIcon className="h-4 w-4 mr-2" />
              Grant Roles
            </Button>
          </CardFooter>
        </Card>
      )}

      <Card>
        <CardHeader>
          <div className="flex justify-between items-center">
            <CardTitle className="text-lg font-semibold">
              Wallet Permissions
              <Badge variant="outline" className="ml-2">
                {permissions.length} wallets
              </Badge>
            </CardTitle>
          </div>
          <CardDescription>
            Manage existing wallet permissions and balances
          </CardDescription>
        </CardHeader>
        <CardContent>
          <div className="space-y-4">
            {permissions.length === 0 ? (
              <div className="flex items-center justify-center h-[200px] border rounded-md bg-accent/5">
                <p className="text-sm text-muted-foreground">
                  No wallets with permissions found
                </p>
              </div>
            ) : (
              <ScrollArea className="h-[400px] border rounded-md">
                <div className="p-4 space-y-4">
                  {permissions.map((permission, index) => (
                    <div key={permission.walletAddress}>
                      <div className="bg-accent/5 rounded-lg p-4 border border-accent/10 hover:border-accent/20 transition-all">
                        <div className="flex flex-col md:flex-row justify-between gap-4">
                          <div className="space-y-2">
                            <div className="flex flex-wrap items-center gap-2">
                              {permission.isOwner && (
                                <Badge>Owner</Badge>
                              )}
                              <div className="font-mono text-sm flex items-center">
                                {truncateAddress(permission.walletAddress)}
                                <Button
                                  size="sm"
                                  variant="ghost"
                                  className="h-6 w-6 p-0 ml-1"
                                  onClick={() => {
                                    navigator.clipboard.writeText(permission.walletAddress)
                                    alert("Address copied to clipboard!")
                                  }}
                                >
                                  <CopyIcon className="h-3 w-3" />
                                </Button>
                              </div>
                            </div>
                            
                            <div className="flex flex-wrap gap-2 mt-1">
                              {permission.roles.map(roleKey => (
                                <RoleBadge key={roleKey} roleKey={roleKey} />
                              ))}
                            </div>
                          </div>
                          
                          <div className="flex items-center gap-3">
                            <div className="text-sm bg-background px-3 py-1 rounded-md border border-input">
                              <span className="text-muted-foreground">Balance: </span>
                              <span className="font-mono">
                                {permission.balance ? formatEther(permission.balance) : "0"} SOM
                              </span>
                            </div>
                            
                            <Button
                              onClick={() => {
                                setDepositTargetAddress(permission.walletAddress as Address)
                                setShowDepositModal(true)
                              }}
                              size="sm"
                            >
                              Deposit
                            </Button>
                            
                            {!permission.isOwner && (
                              <Button
                                onClick={() => {
                                  if (confirm("Are you sure you want to revoke this identity? This action cannot be undone.")) {
                                    revokeIdentity(permission.walletAddress as Address)
                                  }
                                }}
                                variant="destructive"
                                title="Revoke Identity"
                                size="icon"
                              >
                                <TrashIcon className="h-4 w-4" />
                              </Button>
                            )}
                          </div>
                        </div>
                        
                        {!permission.isOwner && (
                          <div className="mt-4">
                            <Label className="text-sm text-muted-foreground mb-2 block">Manage Roles</Label>
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
                        )}
                      </div>
                      {index < permissions.length - 1 && <Separator className="my-4" />}
                    </div>
                  ))}
                </div>
              </ScrollArea>
            )}
          </div>
        </CardContent>
      </Card>

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
