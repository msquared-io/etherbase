import { type EtherbaseConfig, somnia } from "@msquared/etherbase-client"
import { hardhat } from "viem/chains"

const isLocal = process.env.NEXT_PUBLIC_ENV === "local"
const useLocalBackend = process.env.NEXT_PUBLIC_USE_LOCAL_BACKEND === "true"

const localUrl = "http://localhost"

export const etherbaseConfig: EtherbaseConfig = {
  chain: isLocal ? hardhat : somnia,
  httpReaderUrl: useLocalBackend
    ? `${localUrl}:8082`
    : "https://etherbase-reader-496683047294.europe-west2.run.app",
  wsReaderUrl: useLocalBackend
    ? `${localUrl}:8082`
    : "wss://etherbase-reader-496683047294.europe-west2.run.app",
  wsWriterUrl: useLocalBackend
    ? `${localUrl}:8081`
    : "wss://etherbase-writer-496683047294.europe-west2.run.app",
  useBackend: false,
}
