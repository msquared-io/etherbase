import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";

const EtherbaseModule = buildModule("Etherbase", (m) => {
  const validator = m.contract("DataValidator");

  const etherbase = m.contract("Etherbase", [validator]);

  return { etherbase, validator };
});

export default EtherbaseModule;
