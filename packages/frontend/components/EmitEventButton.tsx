"use client"

import { useState } from "react"
import { Button } from "@/components/ui/button"
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogFooter,
} from "@/components/ui/dialog"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { useEtherbaseSource } from "@msquared/etherbase-client"
import type { Address } from "viem"

interface EmitEventButtonProps {
  sourceAddress: string
  eventName: string
  args: { name: string; argType: string }[]
}

export default function EmitEventButton({
  sourceAddress,
  eventName,
  args,
}: EmitEventButtonProps) {
  const [isOpen, setIsOpen] = useState(false)
  const [paramValues, setParamValues] = useState<Record<string, string>>({})
  const { emitEvent } = useEtherbaseSource({
    sourceAddress: sourceAddress as Address,
  })

  const handleEmit = async () => {
    try {
      // Convert parameter values based on their types
      const convertedValues: Record<string, any> = {}

      args.forEach((arg) => {
        const value = paramValues[arg.name]

        // Basic type conversion
        switch (arg.argType) {
          case "uint256":
            convertedValues[arg.name] = BigInt(value)
            break
          case "bool":
            convertedValues[arg.name] = value.toLowerCase() === "true"
            break
          case "address":
            convertedValues[arg.name] = value as Address
            break
          default:
            convertedValues[arg.name] = value
        }
      })

      await emitEvent({
        name: eventName,
        args: convertedValues,
      })

      setIsOpen(false)
      setParamValues({})
    } catch (error) {
      console.error("Error emitting event:", error)
      alert(error instanceof Error ? error.message : "Failed to emit event")
    }
  }

  return (
    <>
      <Button onClick={() => setIsOpen(true)} variant="outline" size="sm">
        Emit Event
      </Button>

      <Dialog open={isOpen} onOpenChange={setIsOpen}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>Emit {eventName}</DialogTitle>
          </DialogHeader>

          <div className="space-y-4 py-4">
            {args.map((arg) => (
              <div key={arg.name} className="space-y-2">
                <Label htmlFor={arg.name}>
                  {arg.name} ({arg.argType})
                </Label>
                <Input
                  id={arg.name}
                  value={paramValues[arg.name] || ""}
                  onChange={(e) =>
                    setParamValues((prev) => ({
                      ...prev,
                      [arg.name]: e.target.value,
                    }))
                  }
                  placeholder={`Enter ${arg.argType} value`}
                />
              </div>
            ))}
          </div>

          <DialogFooter>
            <Button variant="outline" onClick={() => setIsOpen(false)}>
              Cancel
            </Button>
            <Button onClick={handleEmit}>Emit</Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>
    </>
  )
}
