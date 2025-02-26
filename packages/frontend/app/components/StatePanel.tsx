"use client"

import React, { useState } from "react"
import { useEtherstore } from "@msquared/etherbase-client"
import type { Address } from "viem"
import { Card, CardContent, CardHeader, CardTitle, CardDescription } from "@/components/ui/card"
import { Badge } from "@/components/ui/badge"
import { ScrollArea } from "@/components/ui/scroll-area"
import { DatabaseIcon } from "lucide-react"
import { Separator } from "@/components/ui/separator"

interface StatePanelProps {
  sourceAddress: Address
}

// A component to recursively render nested state objects
function StateTree({ data }: { data: any }) {
  if (typeof data !== "object" || data === null) {
    return <span className="text-sm font-mono">{String(data)}</span>
  }

  return (
    <div className="pl-4 space-y-1">
      {Object.entries(data).map(([key, value]) => (
        <div key={key}>
          <span className="font-medium text-sm">{key}: </span>
          <StateTree data={value} />
        </div>
      ))}
    </div>
  )
}

export default function StatePanel({ sourceAddress }: StatePanelProps) {
  const { state, loading, error } = useEtherstore({
    contractAddress: sourceAddress,
    path: [],
  })

  // Count the number of top-level state entries
  const stateEntryCount = state ? Object.keys(state).length : 0

  return (
    <div className="mt-8 space-y-6">
      <div className="flex items-center gap-2 mb-4">
        <DatabaseIcon className="h-6 w-6 text-muted-foreground" />
        <h2 className="text-2xl font-semibold">State Management</h2>
      </div>

      <Card>
        <CardHeader>
          <div className="flex justify-between items-center">
            <CardTitle className="text-lg font-semibold">
              Current State
              <Badge variant="outline" className="ml-2">
                {stateEntryCount} entries
              </Badge>
            </CardTitle>
          </div>
          <CardDescription>
            View and monitor the current state of your contract
          </CardDescription>
        </CardHeader>
        <CardContent>
          {loading ? (
            <div className="flex items-center justify-center h-[200px] border rounded-md bg-accent/5">
              <p className="text-sm text-muted-foreground">Loading state...</p>
            </div>
          ) : error ? (
            <div className="flex items-center justify-center h-[200px] border rounded-md bg-accent/5">
              <p className="text-sm text-destructive">Error: {error}</p>
            </div>
          ) : !state || Object.keys(state).length === 0 ? (
            <div className="flex items-center justify-center h-[200px] border rounded-md bg-accent/5">
              <p className="text-sm text-muted-foreground">No state found</p>
            </div>
          ) : (
            <ScrollArea className="h-[400px] border rounded-md p-4">
              <StateTree data={state} />
            </ScrollArea>
          )}
        </CardContent>
      </Card>
    </div>
  )
}
