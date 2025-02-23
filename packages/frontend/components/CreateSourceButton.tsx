"use client"

import { useCallback, useState } from "react"
import { Button } from "@/components/ui/button"
import { useEtherbase } from "@msquared/etherbase-client"
import { useRouter } from "next/navigation"

type CreateSourceButtonProps = Readonly<{
  fetchSources: () => Promise<void>
}>

export default function CreateSourceButton({
  fetchSources,
}: CreateSourceButtonProps) {
  const [isCreating, setIsCreating] = useState(false)
  const { createSource } = useEtherbase()
  const router = useRouter()

  const handleCreate = useCallback(async () => {
    try {
      setIsCreating(true)
      const { sourceAddress } = await createSource()
      console.log("Created source at:", sourceAddress)

      router.push(`/sources/${sourceAddress}`)
    } catch (error) {
      console.error("Error creating source:", error)
    } finally {
      setIsCreating(false)
    }
  }, [createSource, router])

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
