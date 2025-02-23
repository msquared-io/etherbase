"use client"

import React from "react"
import { useEtherstore } from "@msquared/etherbase-client"
import type { Address } from "viem"
import { Card } from "@/components/ui/card"

interface StatePanelProps {
  sourceAddress: Address
}

// A component to recursively render nested state objects
function StateTree({ data }: { data: any }) {
  if (typeof data !== "object" || data === null) {
    return <span className="text-sm">{String(data)}</span>
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

  if (loading) {
    return <div>Loading state...</div>
  }

  if (error) {
    return <div className="text-red-500">Error: {error}</div>
  }

  return (
    <Card className="p-4">
      <h3 className="text-lg font-medium mb-4">Current State</h3>
      <StateTree data={state} />
    </Card>
  )
}
