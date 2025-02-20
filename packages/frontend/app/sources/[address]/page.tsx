"use server"

import { getEventDefinitions } from "@msquared/etherbase-client/server"
import SourceDetails from "./SourceDetails"
import type { Address, Hex } from "viem"
import { etherbaseConfig } from "@/app/etherbaseConfig"

export default async function SourcePage({
  params,
}: {
  params: { address: Address }
}) {
  const eventDefinitions = await getEventDefinitions(
    params.address as Hex,
    etherbaseConfig,
  )

  return (
    <SourceDetails
      address={params.address}
      initialEventDefinitions={eventDefinitions}
    />
  )
}
