import type { HardhatUserConfig } from "hardhat/config"
import "@nomicfoundation/hardhat-viem"
import "@nomicfoundation/hardhat-ignition-viem"
const privateKey =
  process.env.PRIVATE_KEY ||
  "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"

console.log(privateKey)
const config: HardhatUserConfig = {
  solidity: {
    compilers: [
      {
        version: "0.8.28",
        settings: {
          optimizer: {
            enabled: true,
            runs: 10,
          },
        },
      },
      {
        version: "0.8.12",
      },
    ],
  },
  networks: {
    localhost: {
      url: "http://127.0.0.1:8545",
    },
    somnia: {
      url: "https://dream-rpc.somnia.network/",
      accounts: [privateKey],
    },
  },
}

export default config
