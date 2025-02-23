"use client"

import { useRouter } from "next/navigation"
import { useEtherbase } from "@msquared/etherbase-client"
import { useEffect } from "react"
import CreateSourceButton from "@/components/CreateSourceButton"

export default function SourcesList() {
  const router = useRouter()
  const { sources, fetchSources } = useEtherbase()

  useEffect(() => {
    fetchSources()
  }, [fetchSources])

  console.log("sources", sources)

  return (
    <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-4 flex-col">
      <div className="mb-3 flex gap-4 flex-col items-start">
        <div className="flex gap-4 items-center">
          <h1 className="text-xl font-bold">Sources</h1>
          <CreateSourceButton fetchSources={fetchSources} />
        </div>
        {sources.map((source) => (
          <div
            key={source.sourceAddress}
            onClick={() => router.push(`/sources/${source.sourceAddress}`)}
            onKeyDown={(e) => {
              if (e.key === "Enter") {
                router.push(`/sources/${source.sourceAddress}`)
              }
            }}
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
    </div>
  )
}
