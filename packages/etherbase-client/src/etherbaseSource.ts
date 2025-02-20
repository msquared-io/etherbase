import type { Address } from "viem"
import { getConfig } from "./config"
import type { EtherbaseConfig } from "./config"

export async function fetchEventDefinitionsData(
  sourceAddress: Address,
  config?: EtherbaseConfig,
) {
  const { httpReaderUrl } = config || getConfig()

  const response = await fetch(
    `${httpReaderUrl}/event-definitions?sourceAddress=${sourceAddress}`,
    { cache: "no-store" }
  )
  console.log("response", response)
  return response.json()
}
