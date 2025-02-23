import { defineChain } from "viem"

export const somnia = defineChain({
  id: 50312,
  name: "Somnia",
  network: "somnia",
  nativeCurrency: {
    decimals: 18,
    name: "STT",
    symbol: "STT",
  },
  rpcUrls: {
    default: { http: ["https://dream-rpc.somnia.network/"] },
    public: { http: ["https://dream-rpc.somnia.network/"] },
  },
})
