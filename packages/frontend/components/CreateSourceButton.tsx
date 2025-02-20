"use client"

import { useCallback, useState } from "react"
import { useRouter } from "next/navigation"
import { Button } from "@/components/ui/button"
import { useEtherbase } from "@msquared/etherbase-client"

export default function CreateSourceButton() {
  const [isCreating, setIsCreating] = useState(false)
  const router = useRouter()
  const { fetchSources, createSource } = useEtherbase()

  const handleCreate = useCallback(async () => {
    try {
      setIsCreating(true)
      const { sourceAddress } = await createSource()
      console.log("Created source at:", sourceAddress)

      await fetchSources()

      // Refresh the page to show new emitter
      router.refresh()
    } catch (error) {
      console.error("Error creating source:", error)
    } finally {
      setIsCreating(false)
    }
  }, [router, createSource, fetchSources])

  return (
    <Button
      onClick={handleCreate}
      disabled={isCreating}
      className="bg-white text-black"
    >
      {isCreating ? "Creating..." : "Create New Source"}
    </Button>
  )
}
