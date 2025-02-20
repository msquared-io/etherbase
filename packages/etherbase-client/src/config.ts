import type { Chain, Hex } from "viem"

export interface EtherbaseConfig {
  httpReaderUrl?: string
  wsReaderUrl?: string
  wsWriterUrl?: string
  privateKey?: string // Optional private key for write operations
  chain?: Chain // Optional chain for write operations inside the browser
  useBackend?: boolean  // Optional flag to use the backend for write operations
}

let config: EtherbaseConfig | null = null

export function initializeApp(userConfig: EtherbaseConfig) {
  if (!userConfig.wsReaderUrl && !userConfig.wsWriterUrl && !userConfig.httpReaderUrl) {
    throw new Error(
      "wsReaderUrl or wsWriterUrl or httpReaderUrl must be provided in etherbase config",
    )
  }

  config = userConfig
}

export function getConfig(): EtherbaseConfig {
  if (!config) {
    throw new Error(
      "Etherbase must be initialized with initializeApp() before use, or use the EtherbaseProvider",
    )
  }

  return config
}
