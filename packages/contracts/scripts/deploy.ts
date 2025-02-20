import { spawn } from "node:child_process"
import type { Address } from "viem"
import fs from "node:fs/promises"
import path from "node:path"
import hre from "hardhat"

const network = hre.network.name
const isLocal = network !== "somnia"

async function updateConfig(
  multicallAddress: Address,
  etherbaseAddress: Address,
) {
  const configPath = path.join(
    __dirname,
    "../../../packages/backend/src/config.ts",
  )
  const configContent = await fs.readFile(configPath, "utf-8")

  const newConfig = configContent.replace(
    isLocal ? /const Dev = {[\s\S]*?};/ : /const Somnia = {[\s\S]*?};/,
    isLocal
      ? `const Dev = {
  ETHERBASE_ADDRESS: "${etherbaseAddress}" as Address,
  MULTICALL_ADDRESS: "${multicallAddress}" as Address,
};`
      : `const Somnia = {
  ETHERBASE_ADDRESS: "${etherbaseAddress}" as Address,
  MULTICALL_ADDRESS: "${multicallAddress}" as Address,
};`,
  )

  await fs.writeFile(configPath, newConfig)
}

async function copyABIs() {
  const sourceDir = path.join(__dirname, "../artifacts/contracts")

  // Define mapping of contract files to their destination directories
  const abiMapping = {
    "EtherbaseSource.sol": [
      path.join(__dirname, "../../../packages/etherbase-client/src/abi"),
    //   path.join(__dirname, "../../../packages/frontend/hooks/abi"),
    ],
    "Etherbase.sol": [
      path.join(__dirname, "../../../packages/etherbase-client/src/abi"),
    ],
  }

  // Ensure all destination directories exist
  const allDests = [...new Set(Object.values(abiMapping).flat())]
  await Promise.all(allDests.map((dir) => fs.mkdir(dir, { recursive: true })))

  for (const [sourceFile, destinations] of Object.entries(abiMapping)) {
    const baseName = sourceFile.replace(".sol", "")
    const jsonFileName = `${baseName}.json`
    const tsFileName = `${baseName}.ts`
    const sourcePath = path.join(sourceDir, sourceFile, jsonFileName)

    try {
      // Read and parse the source JSON
      const sourceContent = await fs.readFile(sourcePath, "utf-8")
      const parsedContent = JSON.parse(sourceContent)

      // Create TypeScript content
      const tsContent = `export const ${baseName}Abi = ${JSON.stringify(parsedContent.abi, null, 2)} as const;\n`

      // Write to all destination paths
      await Promise.all(
        destinations.map(async (dest) => {
          await fs.writeFile(path.join(dest, tsFileName), tsContent)
          console.log(`Copied ABI for ${sourceFile} to ${dest}`)
        }),
      )
    } catch (error) {
      console.error(`Failed to copy ABI for ${sourceFile}:`, error)
    }
  }
}

async function main() {
//   const deployCmd = isLocal ? "local" : "somnia"

//   const deployProcessMulticall = spawn(
//     "npm",
//     ["run", `${deployCmd}:deploy:multicall3`],
//     {
//       stdio: "pipe",
//       env: {
//         ...process.env,
//         HARDHAT_IGNITION_CONFIRM_DEPLOYMENT: "false",
//       },
//     },
//   )

//   let allOutputMulticall = ""
//   const multicallAddress = await new Promise<Address>((resolve, reject) => {
//     deployProcessMulticall.stdout.on("data", (data) => {
//       allOutputMulticall += data.toString()
//     })

//     deployProcessMulticall.on("close", (code) => {
//       if (code !== 0) {
//         reject(new Error(`Multicall deployment failed with code ${code}`))
//         return
//       }
//       const matches = allOutputMulticall.match(/0x[a-fA-F0-9]{40}/g)
//       if (!matches || matches.length === 0) {
//         reject(new Error("Could not find multicall address in deploy output"))
//         return
//       }
//       resolve(matches[matches.length - 1] as Address)
//     })
//   })

//   console.log("Multicall deployed at:", multicallAddress)

//   // Deploy Etherbase
//   const deployProcess = spawn("npm", ["run", `${deployCmd}:deploy`], {
//     stdio: "pipe",
//     env: {
//       ...process.env,
//       HARDHAT_IGNITION_CONFIRM_DEPLOYMENT: "false",
//     },
//   })

//   let allOutput = ""
//   const etherbaseAddress = await new Promise<Address>((resolve, reject) => {
//     deployProcess.stdout.on("data", (data) => {
//       allOutput += data.toString()
//     })

//     deployProcess.on("close", (code) => {
//       if (code !== 0) {
//         reject(new Error(`Etherbase deployment failed with code ${code}`))
//         return
//       }
//       const matches = allOutput.match(/0x[a-fA-F0-9]{40}/g)
//       if (!matches || matches.length === 0) {
//         reject(new Error("Could not find etherbase address in deploy output"))
//         return
//       }
//       resolve(matches[matches.length - 1] as Address)
//     })
//   })

//   console.log("Etherbase deployed at:", etherbaseAddress)

  // Update config file
//   await updateConfig(multicallAddress, etherbaseAddress)
//   console.log("Config updated successfully")

  // Copy ABIs
  await copyABIs()
  console.log("ABIs copied successfully")
}

main().catch((error) => {
  console.error(error)
  process.exit(1)
})