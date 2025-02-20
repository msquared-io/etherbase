import { expect } from "chai";
import hre from "hardhat";
import { Hex, keccak256, pad, toBytes, toHex, parseAbi, stringToBytes } from "viem";
import { encodeFunctionData } from "viem";

/**
 * Deploys the `EtherbaseSource` contract (which inherits RoleControl).
 */
async function deployEtherbaseSourceFixture() {
  const [deployer, user1, user2] = await hre.viem.getWalletClients();

  const etherDatabaseLibrary = await hre.viem.deployContract(
    "EtherDatabaseLib",
    [deployer.account.address],
    {
      args: undefined,
    },
  );

  const etherbaseSource = await hre.viem.deployContract("EtherbaseSource", [deployer.account.address], {
    libraries: {
      EtherDatabaseLib: etherDatabaseLibrary.address,
    },
  });

  return { etherbaseSource, deployer, user1, user2 };
}

const parseEventAbi = (abi: string) => {
  const eventTopic = keccak256(stringToBytes(abi));
  const eventAbi = parseAbi([abi]);
  const eventName = eventAbi[0].name;
  const parameters = eventAbi[0].inputs.map((input: any) => ({
    name: input.name,
    parameterType: input.type,
    isIndexed: !!input?.indexed,
  }));
  return { eventName, eventTopic, parameters };
};

const registerEventSchema = async (etherbaseSource: any, eventAbi: string) => {
  const { eventName, eventTopic, parameters } = parseEventAbi(eventAbi);
  await etherbaseSource.write.registerEventSchema([eventName, eventTopic, parameters]);
  return { eventName, eventTopic };
};

const emitEvent = async (etherbaseSource: any, user: any, eventName: string, values: any[]) => {
  const eventSchema = await etherbaseSource.read.getEventSchemaFromName([eventName]);

  const topics: string[] = [eventSchema.eventTopic]; // Add event topic hash as first topic
  const dataTypes: any[] = [];
  const dataValues: any[] = [];

  // Validate number of parameters
  if (values.length !== eventSchema.parameters.length) {
    throw new Error("Event arguments/values length mismatch");
  }

  // Process each parameter
  eventSchema.parameters.forEach((parameter: any, index: number) => {
    const value = values[index];

    if (parameter.isIndexed) {
      if (parameter.parameterType === "string") {
        topics.push(keccak256(stringToBytes(value)));
      } else if (parameter.parameterType.startsWith("bytes")) {
        topics.push(keccak256(value));
      } else if (parameter.parameterType === "address") {
        topics.push(pad(value, { size: 32 }));
      } else if (parameter.parameterType === "uint256" || parameter.parameterType === "int256") {
        topics.push(pad(toHex(value), { size: 32 }));
      } else if (parameter.parameterType === "bool") {
        topics.push(pad(value ? "0x01" : "0x00", { size: 32 }));
      } else {
        topics.push(pad(toHex(value), { size: 32 }));
      }
    } else {
      dataTypes.push(parameter.parameterType);
      dataValues.push(value);
    }
  });

  // Encode non-indexed parameters as data
  const data =
    "0x" +
    (dataValues.length > 0
      ? encodeFunctionData({
          abi: [
            {
              type: "function",
              name: "dummy",
              inputs: dataTypes.map((type) => ({ type })),
              outputs: [],
              stateMutability: "nonpayable",
            },
          ],
          args: dataValues,
        }).slice(10) // Remove function selector (first 10 chars)
      : "");

  await etherbaseSource.write.emitEvent(
    [eventName, topics.slice(1), data], // Remove event topic from topics array
    {
      account: user.account,
    },
  );

  return { parameterTopics: topics.slice(1), data };
};

describe("EtherbaseSource", () => {
  describe("emitEvent", () => {
    it("should revert if caller does not have EMITTER role", async () => {
      const { etherbaseSource, user2 } = await deployEtherbaseSourceFixture();

      await etherbaseSource.write.createWalletIdentity([user2.account.address]);

      const { eventName } = await registerEventSchema(etherbaseSource, "event TestEvent()");

      await expect(emitEvent(etherbaseSource, user2, eventName, [])).to.be.revertedWithCustomError(
        etherbaseSource,
        "Unauthorized",
      );
    });

    it("should succeed if caller has EMITTER role (address-based)", async () => {
      const { etherbaseSource, deployer, user1 } = await deployEtherbaseSourceFixture();

      await etherbaseSource.write.createWalletIdentity([user1.account.address]);
      const user1Key = keccak256(user1.account.address);

      await etherbaseSource.write.grantRole([user1Key, 0]);

      const { eventName } = await registerEventSchema(etherbaseSource, "event TestEvent()");

      await emitEvent(etherbaseSource, user1, eventName, []);
    });

    it("should succeed if event has 1 parameter", async () => {
      const { etherbaseSource, deployer, user1 } = await deployEtherbaseSourceFixture();

      await etherbaseSource.write.createWalletIdentity([user1.account.address]);
      const user1Key = keccak256(user1.account.address);

      await etherbaseSource.write.grantRole([user1Key, 0]);

      const { eventName, eventTopic } = await registerEventSchema(
        etherbaseSource,
        "event TestEvent(uint256 id)",
      );

      const { parameterTopics, data } = await emitEvent(etherbaseSource, user1, eventName, [12345]);

      const publicClient = await hre.viem.getPublicClient();
      const logs = await publicClient.getLogs();
      expect(logs.length).to.eq(1, "Should emit exactly one log");
      expect(logs[0].topics.length).to.eq(1, "Should have 1 topic");
      expect(logs[0].topics[0]).to.be.equal(eventTopic);
      expect(logs[0].data).to.be.equal(data);
    });

    it("should succeed if event has 2 parameters", async () => {
      const { etherbaseSource, deployer, user1 } = await deployEtherbaseSourceFixture();

      await etherbaseSource.write.createWalletIdentity([user1.account.address]);
      const user1Key = keccak256(user1.account.address);

      await etherbaseSource.write.grantRole([user1Key, 0]);

      const { eventName, eventTopic } = await registerEventSchema(
        etherbaseSource,
        "event TestEvent(uint256 id, address indexed user)",
      );

      const { parameterTopics, data } = await emitEvent(etherbaseSource, user1, eventName, [
        12345,
        user1.account.address,
      ]);
      expect(parameterTopics.length).to.eq(1, "Should have 1 parameter topic");

      const publicClient = await hre.viem.getPublicClient();
      const logs = await publicClient.getLogs();
      expect(logs.length).to.eq(1, "Should emit exactly one log");
      expect(logs[0].topics.length).to.eq(2, "Should have 2 topics");
      expect(logs[0].topics[0]).to.be.equal(eventTopic);
      expect(logs[0].topics[1]).to.be.equal(parameterTopics[0]);
      expect(logs[0].data).to.be.equal(data);
    });
  });
});
