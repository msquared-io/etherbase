import { getConfig } from "./config"
import type { EtherbaseConfig } from "./config"

export async function fetchSourcesData(config?: EtherbaseConfig) {
  const { httpReaderUrl } = config || getConfig()

  const response = await fetch(`${httpReaderUrl}/sources`, {
    cache: "no-store",
  })
  const data = await response.json()
  console.log("response", data)
  return data
}

export async function fetchCustomContractsData(config?: EtherbaseConfig) {
  const { httpReaderUrl } = config || getConfig()

  const response = await fetch(`${httpReaderUrl}/custom-contracts`, {
    cache: "no-store",
  })
  const data = await response.json()
  console.log("response", data)
  return data
}
