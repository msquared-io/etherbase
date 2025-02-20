"use client"

import { useState, useEffect, useCallback, useRef } from "react"
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogFooter,
} from "@/components/ui/dialog"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select"
import { TrashIcon } from "lucide-react"
import type { Address } from "viem"
import { Argument, useEtherbaseSource } from "@msquared/etherbase-client"

// Keep only Solidity types
const SOLIDITY_TYPES = [
  "address",
  "bool",
  "string",
  "bytes",
  "uint256",
  "int256",
  "bytes32",
  "uint8",
  "uint16",
  "uint32",
  "uint64",
  "uint128",
  "int8",
  "int16",
  "int32",
  "int64",
  "int128",
]

interface RegisterEventButtonProps {
  sourceAddress: Address
  onEventRegistered: () => void
}

function ArgumentInput({
  argument,
  index,
  showIndexed,
  onChange,
  onDelete,
}: {
  argument: Argument
  index: number
  showIndexed: boolean
  onChange: (newParam: Argument) => void
  onDelete?: () => void
}) {
  return (
    <div className="border p-2 rounded my-2">
      <div className="flex flex-col sm:flex-row gap-4 items-center">
        <div className="flex-1">
          <Label htmlFor={`param-name-${index}`} className="block mb-1">
            Parameter Name
          </Label>
          <Input
            id={`argument-name-${index}`}
            value={argument.name}
            onChange={(e) => onChange({ ...argument, name: e.target.value })}
            maxLength={32}
          />
        </div>
        <div className="flex-1">
          <Label htmlFor={`argument-type-${index}`} className="block mb-1">
            Argument Type
          </Label>
          <Select
            value={argument.argType}
            onValueChange={(value) => onChange({ ...argument, argType: value })}
          >
            <SelectTrigger id={`param-type-${index}`} className="w-full">
              <SelectValue placeholder="Select type" />
            </SelectTrigger>
            <SelectContent>
              {SOLIDITY_TYPES.map((type) => (
                <SelectItem key={type} value={type}>
                  {type}
                </SelectItem>
              ))}
            </SelectContent>
          </Select>
        </div>
        {showIndexed && (
          <div className="flex items-center space-x-2">
            <Input
              type="checkbox"
              id={`argument-indexed-${index}`}
              checked={argument.isIndexed}
              onChange={(e) =>
                onChange({ ...argument, isIndexed: e.target.checked })
              }
              className="w-4 h-4"
            />
            <Label htmlFor={`argument-indexed-${index}`}>Indexed</Label>
          </div>
        )}
        {onDelete && (
          <Button
            variant="outline"
            onClick={onDelete}
            className="text-white border-red-500"
          >
            <TrashIcon className="w-4 h-4" />
          </Button>
        )}
      </div>
    </div>
  )
}

export default function RegisterEventButton({
  sourceAddress,
  onEventRegistered,
}: RegisterEventButtonProps) {
  const [isOpen, setIsOpen] = useState(false)
  const [showIndexed, setShowIndexed] = useState(false)
  const [solidityABI, setSolidityABI] = useState("")
  const { registerEvent } = useEtherbaseSource({ sourceAddress })
  const isUpdatingFromABI = useRef(false)
  const [newEvent, setNewEvent] = useState<{
    name: string
    arguments: Argument[]
  }>({
    name: "userProximityEvent",
    arguments: [
      {
        name: "userId",
        argType: "string",
        isIndexed: false,
      },
      {
        name: "userAddress",
        argType: "string",
        isIndexed: false,
      },
      {
        name: "description",
        argType: "string",
        isIndexed: false,
      },
      {
        name: "proximity",
        argType: "uint256",
        isIndexed: false,
      },
      {
        name: "timestamp",
        argType: "uint256",
        isIndexed: false,
      },
    ],
  })

  // A simple parser to convert an ABI string into event name and parameters.
  const parseABI = useCallback(
    (abiString: string): { name: string; arguments: Argument[] } | null => {
      try {
        const cleaned = abiString.trim()
        const eventMatch = cleaned.match(
          /^event\s+([a-zA-Z_][a-zA-Z0-9_]*)\s*\((.*)\)$/,
        )
        if (!eventMatch) return null

        const [, eventName, paramsStr] = eventMatch
        if (!paramsStr.trim()) return { name: eventName, arguments: [] }

        const params = paramsStr.split(",").map((s) => s.trim())
        const parsedArguments: Argument[] = params.map((paramStr) => {
          const parts = paramStr.split(" ").filter(Boolean)
          const isIndexed = parts.includes("indexed")
          const filtered = parts.filter((p) => p !== "indexed")
          const [type, name] = filtered
          return { name, argType: type, isIndexed }
        })

        return { name: eventName, arguments: parsedArguments }
      } catch (e) {
        console.error("ABI parse error:", e)
        return null
      }
    },
    [],
  )

  // A simple formatter to convert event name and parameters into an ABI string.
  const formatABI = useCallback((name: string, args: Argument[]): string => {
    const argumentStr = args
      .filter((a) => a.name.trim() !== "" && a.argType.trim() !== "")
      .map((a) => `${a.argType}${a.isIndexed ? " indexed" : ""} ${a.name}`)
      .join(", ")
    return `event ${name}(${argumentStr})`
  }, [])

  // When the ABI changes (user manually edits), update the event name and parameters
  useEffect(() => {
    if (isUpdatingFromABI.current) {
      isUpdatingFromABI.current = false
      return
    }

    const parsed = parseABI(solidityABI)
    if (parsed) {
      setNewEvent(parsed)
    }
  }, [solidityABI, parseABI])

  // When event name or parameters change, update the ABI
  useEffect(() => {
    const abiString = formatABI(newEvent.name, newEvent.arguments)
    if (abiString !== solidityABI) {
      isUpdatingFromABI.current = true
      setSolidityABI(abiString)
    }
  }, [newEvent.name, newEvent.arguments, formatABI, solidityABI])

  const handleAddArgument = () => {
    const newArgument: Argument = {
      name: "",
      argType: "string",
      isIndexed: false,
    }

    setNewEvent((prev) => ({
      ...prev,
      arguments: [...prev.arguments, newArgument],
    }))
  }

  const handleRegister = async () => {
    try {
      // Validate event name
      if (!newEvent.name || newEvent.name.trim() === "") {
        throw new Error("Event name is required")
      }

      // Validate parameters
      if (
        !newEvent.arguments.some(
          (a) => a.name.trim() !== "" && a.argType.trim() !== "",
        )
      ) {
        throw new Error(
          "At least one parameter with a name and argType is required",
        )
      }

      const indexedCount = newEvent.arguments.filter((a) => a.isIndexed).length
      if (indexedCount > 3) {
        throw new Error(
          "Solidity events allow at most three indexed parameters",
        )
      }

      // Validate parameter types
      for (const arg of newEvent.arguments) {
        if (!SOLIDITY_TYPES.includes(arg.argType)) {
          throw new Error(
            `Invalid type "${arg.argType}". Allowed types: ${SOLIDITY_TYPES.join(", ")}`,
          )
        }
      }

      await registerEvent({
        name: newEvent.name,
        args: newEvent.arguments,
      })
      setIsOpen(false)
      setNewEvent({
        name: "",
        arguments: [{ name: "", argType: "string", isIndexed: false }],
      })
      setShowIndexed(false)
      setSolidityABI("")
      onEventRegistered()
    } catch (error) {
      console.error("Error registering event:", error)
      alert(error instanceof Error ? error.message : "Failed to register event")
    }
  }

  return (
    <>
      {/* Trigger Button */}
      <Button onClick={() => setIsOpen(true)} className="bg-white text-black">
        Register New Event
      </Button>

      {/* Modal Dialog */}
      <Dialog open={isOpen} onOpenChange={setIsOpen}>
        <DialogContent className="max-w-3xl p-6 bg-black">
          <DialogHeader>
            <DialogTitle>Register New Event</DialogTitle>
          </DialogHeader>
          <div className="space-y-4">
            {/* Event Name */}
            <div>
              <Label htmlFor="event-name" className="block mb-1">
                Event Name
              </Label>
              <Input
                id="event-name"
                value={newEvent.name}
                onChange={(e) =>
                  setNewEvent((prev) => ({ ...prev, name: e.target.value }))
                }
                maxLength={32}
              />
            </div>

            {/* Solidity ABI Editor & Toggle for Indexed Options */}
            {newEvent.name && (
              <>
                <div className="flex gap-4">
                  <Button
                    variant="secondary"
                    onClick={() => setShowIndexed((prev) => !prev)}
                    className="bg-white text-black"
                  >
                    {showIndexed
                      ? "Hide Indexed Options"
                      : "Show Indexed Options"}
                  </Button>
                </div>
                <div className="mt-4">
                  <Label htmlFor="solidity-abi" className="block mb-1">
                    Solidity ABI
                  </Label>
                  <textarea
                    id="solidity-abi"
                    className="w-full p-2 border rounded bg-gray-800 text-white"
                    rows={4}
                    value={solidityABI}
                    onChange={(e) => setSolidityABI(e.target.value)}
                    placeholder="e.g. event Transfer(address indexed from, address indexed to, uint256 amount)"
                  />
                </div>
              </>
            )}

            {/* Parameters */}
            <div>
              <Label className="block mb-1">Parameters</Label>
              {newEvent.arguments.map((arg, idx) => (
                <ArgumentInput
                  key={`${idx}-${arg.name}-${arg.argType}`}
                  argument={arg}
                  index={idx}
                  showIndexed={showIndexed}
                  onChange={(newArg) => {
                    const updated = [...newEvent.arguments]
                    updated[idx] = newArg
                    setNewEvent((prev) => ({
                      ...prev,
                      arguments: updated,
                    }))
                  }}
                  onDelete={() => {
                    const updated = [...newEvent.arguments]
                    updated.splice(idx, 1)
                    setNewEvent((prev) => ({
                      ...prev,
                      arguments: updated,
                    }))
                  }}
                />
              ))}
              <Button
                onClick={handleAddArgument}
                variant="secondary"
                className="bg-white text-black"
              >
                Add Parameter
              </Button>
            </div>
          </div>

          {/* Action buttons */}
          <DialogFooter className="mt-6 flex justify-end gap-3">
            <Button variant="outline" onClick={() => setIsOpen(false)}>
              Cancel
            </Button>
            <Button onClick={handleRegister} className="bg-white text-black">
              Register
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>
    </>
  )
}
