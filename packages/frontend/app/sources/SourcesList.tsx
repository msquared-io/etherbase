"use client"

import { useRouter } from "next/navigation"
import type { Source } from "@msquared/etherbase-client"

interface SourcesListProps {
  sources: Source[]
}

export default function SourcesList({ sources }: SourcesListProps) {
  const router = useRouter()
  console.log("sources", sources)

  return (
    <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-4">
      {sources.map((source) => (
        <div
          key={source.sourceAddress}
          onClick={() => router.push(`/sources/${source.sourceAddress}`)}
          className="cursor-pointer rounded-lg border border-gray-200 p-4 hover:shadow-lg transition"
        >
          <h2 className="text-lg font-semibold truncate">
            {source.sourceAddress}
          </h2>
          <p className="mt-1 text-xs text-gray-600">Owner: {source.owner}</p>
          <p className="mt-1 text-sm text-gray-600">Click to view events</p>
        </div>
      ))}
    </div>
  )
}
