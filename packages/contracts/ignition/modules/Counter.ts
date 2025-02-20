import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";

const CounterModule = buildModule("Counter", (m) => {
  const counter = m.contract("Counter", []);

  return { counter };
});

export default CounterModule;
