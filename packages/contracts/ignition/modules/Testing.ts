import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";

const TestingModule = buildModule("Testing", (m) => {
  // Get the contract name from environment variable
  const contractToDeployName = process.env.CONTRACT_NAME;

  if (!contractToDeployName) {
    throw new Error("CONTRACT_NAME environment variable must be set");
  }

  console.log("Deploying contract:", contractToDeployName);

  // Deploy only the specified contract
  const contract = m.contract(contractToDeployName, []);

  return { [contractToDeployName]: contract };
});

export default TestingModule;
