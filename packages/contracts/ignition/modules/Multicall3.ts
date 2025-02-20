import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";

const Multicall3Module = buildModule("Multicall3", (m) => {
  const multicall3 = m.contract("Multicall3", []);

  return { multicall3 };
});

export default Multicall3Module;
