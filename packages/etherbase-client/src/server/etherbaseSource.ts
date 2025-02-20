import type { Address } from "viem"
import type { EtherbaseConfig } from "../config"
import { fetchEventDefinitionsData } from "../etherbaseSource"

export async function getEventDefinitions(
  sourceAddress: Address,
  config: EtherbaseConfig,
) {
  return fetchEventDefinitionsData(sourceAddress, config)
}
