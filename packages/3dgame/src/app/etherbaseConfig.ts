import { type EtherbaseConfig, somnia } from "@msquared/etherbase-client"
import { hardhat } from "viem/chains"

const isLocal = process.env.NEXT_PUBLIC_ENV === "local"
const useLocalBackend = process.env.NEXT_PUBLIC_USE_LOCAL_BACKEND === "true"

const localUrl = "http://localhost"
// const localUrl = "http://172.16.128.93"

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
    "0x0c9a44c9e7778f9f3132ab2ad581b1473f84683e6b42da3938160dc602ee29d0",
  useBackend: true,
}

export const sourceAddress = "0x2e30b662c4Df268edA9efce596CDF3896b50B43C"
