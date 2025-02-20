import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";

const MyContractModule = buildModule("MyContract", (m) => {
  const myContract = m.contract("MyContract", [1000000000000000000n]);

  return { myContract };
});

export default MyContractModule;
