"use client"

import SourcesList from "./SourcesList"
import CreateSourceButton from "@/components/CreateSourceButton"
import { useEtherbase } from "@msquared/etherbase-client"
import { useEffect } from "react"

export default function SourcesPage() {
  const { sources, fetchSources } = useEtherbase()
  
  useEffect(() => {
    fetchSources()
  }, [fetchSources])

  return (
    <div className="p-4">
      <div className="mb-3 flex justify-between items-center">
        <h1 className="text-xl font-bold">Sources</h1>
        <CreateSourceButton />
      </div>
      <SourcesList sources={sources} />
    </div>
  )
}
