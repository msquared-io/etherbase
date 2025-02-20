import { type EtherbaseConfig, somnia } from "@msquared/etherbase-client"
import { hardhat } from "viem/chains"

const isLocal = process.env.NEXT_PUBLIC_ENV === "local"
const useLocalBackend = process.env.NEXT_PUBLIC_USE_LOCAL_BACKEND === "true"

// const localUrl = "http://localhost"
const localUrl = "http://172.16.128.93"

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
  privateKey:
    "0x1f7c04305fd6f99506c372b8dae42f559645bc86015df34442e9c0fa6efca5f3",
  useBackend: true,
}

export const sourceAddress = "0x1eF2faBc48E75DA9019470Ae7A7D7C3d19d6bd60"
