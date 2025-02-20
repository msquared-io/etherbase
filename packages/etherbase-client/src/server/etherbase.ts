import type { EtherbaseConfig } from "../config"
import { fetchSourcesData } from "../etherbase"

export async function getSources(config: EtherbaseConfig) {
  return fetchSourcesData(config)
}
