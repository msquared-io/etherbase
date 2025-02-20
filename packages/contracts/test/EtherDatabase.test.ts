import hre from "hardhat"
import { expect } from "chai"
import {
  decodeAbiParameters,
  encodeAbiParameters,
  keccak256,
  type Log,
} from "viem"
import "./helpers"

import { parseAbi, stringToBytes, toHex, fromHex, encodePacked } from "viem"
import type { Hex } from "viem"

const deployEtherDatabase = async () => {
  const [deployer] = await hre.viem.getWalletClients()

  // Deploy validator first
  const validator = await hre.viem.deployContract("DataValidator", [])

  // Deploy library
  const etherDatabaseLibrary = await hre.viem.deployContract(
    "EtherDatabaseLib",
    [deployer.account.address],
    {
      args: undefined,
    },
  )

  // Deploy implementation with validator
  const etherDatabase = await hre.viem.deployContract(
    "EtherDatabaseImpl",
    [deployer.account.address],
    {
      libraries: {
        EtherDatabaseLib: etherDatabaseLibrary.address,
      },
      args: [validator.address], // Pass validator address to constructor
    },
  )

  return { etherDatabase, deployer }
}

const contractAbi = parseAbi([
  "event NewSegment(string segment, uint256 indexed id)",
])

const findEtherbaseStateChangeLog = (logs: Log[]) => {
  const eventNameTopic = keccak256(stringToBytes("EtherbaseStateChange"))
  return logs.find((log) => log.topics[0] === eventNameTopic)
}

const isLogForPath = (log: Log, path: string) => {
  const pathTopic = keccak256(stringToBytes(path))
  return log.topics[1] === pathTopic
}

enum DataType {
  NONE = 0,
  STRING = 1,
  BOOL = 2,
  UINT = 3,
  INT = 4,
  BYTES = 5,
}

enum DataChangeType {
  UPSERT = 0,
  DELETE = 1,
}

const isLogData = (
  log: Log,
  data: Hex,
  dataType: DataType,
  dataChangeType: DataChangeType,
) => {
  return (
    log.data ===
    encodePacked(["bytes", "uint8", "uint8"], [data, dataType, dataChangeType])
  )
}

describe.only("EtherDatabase", function () {
  it("should deploy correctly", async () => {
    const { etherDatabase } = await deployEtherDatabase()
    expect(etherDatabase.address).to.be.not.undefined
  })

  it("should add a new value and emit events (NewSegment, Upsert)", async () => {
    const { etherDatabase } = await deployEtherDatabase()
    const segments = ["foo", "bar"]
    const dataType = 1 // DataType.STRING
    const data = stringToBytes("Hello")
    const hexData = toHex(data)

    await etherDatabase.write.setValue([segments, dataType, hexData])

    const events = await (await hre.viem.getPublicClient()).getContractEvents({
      address: etherDatabase.address,
      abi: contractAbi,
      eventName: "NewSegment",
    })

    expect(events.length).to.eq(
      2,
      "Should have two NewSegment events (foo, bar)",
    )

    expect(events[0].args.segment).to.be.equal("foo")
    expect(events[1].args.segment).to.be.equal("bar")
    expect(events[0].args.id).to.be.equal(1n)
    expect(events[1].args.id).to.be.equal(2n)

    // check for log data change
    const publicClient = await hre.viem.getPublicClient()
    const logs = await publicClient.getLogs()
    console.log(logs)
    expect(logs.length).to.be.equal(3)
    const log = findEtherbaseStateChangeLog(logs)
    expect(log).to.not.be.undefined
    expect(log!.topics.length).to.be.equal(2)
    expect(isLogForPath(log!, "store[foo.bar]")).to.be.true
    expect(isLogData(log!, hexData, dataType, DataChangeType.UPSERT)).to.be.true
  })

  it("should retrieve the newly set value", async () => {
    const { etherDatabase } = await deployEtherDatabase()
    const segments = ["alpha"]
    const dataType = 1 // DataType.STRING
    const data = stringToBytes("world")
    const hexData = toHex(data)

    await etherDatabase.write.setValue([segments, dataType, hexData])

    const [storedDataType, storedData] = await etherDatabase.read.getValue([
      segments,
    ])
    expect(storedDataType).to.eq(dataType)

    const storedStr = fromHex(storedData, "string")
    expect(storedStr).to.eq("world")
  })

  it("should update an existing value (no new NewSegment event)", async () => {
    const { etherDatabase } = await deployEtherDatabase()

    const segments = ["myKey"]
    const firstValue = stringToBytes("initial")
    const hexFirstValue = toHex(firstValue)
    await etherDatabase.write.setValue([
      segments,
      DataType.STRING,
      hexFirstValue,
    ])

    // update
    const secondValue = stringToBytes("updated")
    const hexSecondValue = toHex(secondValue)
    await etherDatabase.write.setValue([
      segments,
      DataType.STRING,
      hexSecondValue,
    ])

    const events = await (await hre.viem.getPublicClient()).getContractEvents({
      address: etherDatabase.address,
      abi: contractAbi,
      eventName: "NewSegment",
    })
    expect(events.length).to.eq(0, "No new segment should be emitted")

    // Check the stored value
    const [storedDataType, storedData] = await etherDatabase.read.getValue([
      segments,
    ])
    expect(storedDataType).to.eq(1)
    const storedStr = fromHex(storedData, "string")
    expect(storedStr).to.eq("updated")
  })

  it("should remove a value and emit a DELETE event log", async () => {
    const { etherDatabase } = await deployEtherDatabase()
    const segments = ["deleteMe"]
    const dataType = DataType.STRING
    await etherDatabase.write.setValue([
      segments,
      dataType,
      toHex(stringToBytes("gone soon")),
    ])

    // Now remove
    await etherDatabase.write.removeValue([segments])

    const publicClient = await hre.viem.getPublicClient()
    const logs = await publicClient.getLogs()
    expect(logs.length).to.eq(1, "Should emit exactly one DELETE log")
    const log = findEtherbaseStateChangeLog(logs)
    expect(log).to.not.be.undefined
    expect(isLogForPath(log!, "store[deleteMe]")).to.be.true
    expect(isLogData(log!, "0x", DataType.NONE, DataChangeType.DELETE)).to.be
      .true

    // Now ensure the value is actually removed
    await expect(
      etherDatabase.read.getValue([segments]),
    ).to.be.revertedWithCustomError(etherDatabase, "PathNotFound")
  })

  it("should revert when removing a path that does not exist", async () => {
    const { etherDatabase } = await deployEtherDatabase()
    const segments = ["no", "such", "key"]
    await expect(
      etherDatabase.write.removeValue([segments]),
    ).to.be.revertedWithCustomError(etherDatabase, "PathNotFound")
  })

  describe("Type-specific setters and getters", () => {
    it("should store and retrieve a string value", async () => {
      const { etherDatabase } = await deployEtherDatabase()
      const segments = ["test", "string"]
      const value = "Hello World"

      await etherDatabase.write.setString([segments, value])
      const result = await etherDatabase.read.getString([segments])
      expect(result).to.equal(value)
    })

    it("should store and retrieve a uint value", async () => {
      const { etherDatabase } = await deployEtherDatabase()
      const segments = ["test", "uint"]
      const value = 12345n

      await etherDatabase.write.setUint([segments, value])
      const result = await etherDatabase.read.getUint([segments])
      expect(result).to.equal(value)
    })

    it("should store and retrieve a bool value", async () => {
      const { etherDatabase } = await deployEtherDatabase()
      const segments = ["test", "bool"]
      const value = true

      await etherDatabase.write.setBool([segments, value])
      const result = await etherDatabase.read.getBool([segments])
      expect(result).to.equal(value)
    })

    it("should store and retrieve an int value", async () => {
      const { etherDatabase } = await deployEtherDatabase()
      const segments = ["test", "int"]
      const value = -12345n

      await etherDatabase.write.setInt([segments, value])
      const result = await etherDatabase.read.getInt([segments])
      expect(result).to.equal(value)
    })

    it("should store and retrieve bytes value", async () => {
      const { etherDatabase } = await deployEtherDatabase()
      const segments = ["test", "bytes"]
      const value = "0x1234567890"

      await etherDatabase.write.setBytes([segments, value])
      const result = await etherDatabase.read.getBytes([segments])
      expect(result).to.equal(value)
    })
  })

  describe("Error handling", () => {
    it("should revert with PathNotFound when getting non-existent path", async () => {
      const { etherDatabase } = await deployEtherDatabase()
      const segments = ["nonexistent", "path"]

      await expect(
        etherDatabase.read.getString([segments]),
      ).to.be.revertedWithCustomError(etherDatabase, "PathNotFound")
    })

    it("should revert with InvalidDataType when getting wrong type", async () => {
      const { etherDatabase } = await deployEtherDatabase()
      const segments = ["test", "type"]

      // Store as string
      await etherDatabase.write.setString([segments, "test"])

      // Try to read as uint
      await expect(
        etherDatabase.read.getUint([segments]),
      ).to.be.revertedWithCustomError(etherDatabase, "InvalidDataType(3, 1)")
    })

    it("should allow updating value type", async () => {
      const { etherDatabase } = await deployEtherDatabase()
      const segments = ["test", "update"]

      // First store as string
      await etherDatabase.write.setString([segments, "test"])

      // Then update to uint
      await etherDatabase.write.setUint([segments, 123n])

      // Should now be readable as uint
      const result = await etherDatabase.read.getUint([segments])
      expect(result).to.equal(123n)

      // But not as string
      await expect(
        etherDatabase.read.getString([segments]),
      ).to.be.revertedWithCustomError(etherDatabase, "InvalidDataType(1, 3)")
    })

    it("should revert with InvalidDataEncoding when data is malformed", async () => {
      const { etherDatabase } = await deployEtherDatabase()
      const segments = ["test", "malformed"]

      // Set malformed string data using generic setter
      await etherDatabase.write.setValue([
        segments,
        DataType.STRING,
        "0x1234", // Invalid string encoding
      ])

      // Try to read as string
      await expect(
        etherDatabase.read.getString([segments]),
      ).to.be.revertedWithCustomError(etherDatabase, "InvalidDataEncoding(1)")
    })
  })

  describe("Generic setValue/getValue", () => {
    it("should work alongside type-specific setters", async () => {
      const { etherDatabase } = await deployEtherDatabase()
      const segments = ["test", "generic"]

      // Set using generic setter
      const dataType = DataType.STRING
      const value = "test value"
      await etherDatabase.write.setValue([
        segments,
        dataType,
        encodeAbiParameters(
          [
            {
              type: "string",
            },
          ],
          [value],
        ),
      ])

      // Should be readable by type-specific getter
      const result = await etherDatabase.read.getString([segments])
      expect(result).to.equal(value)
    })

    it("should return correct type and data with getValue", async () => {
      const { etherDatabase } = await deployEtherDatabase()
      const segments = ["test", "get"]

      // Set using type-specific setter
      await etherDatabase.write.setUint([segments, 456n])

      // Read using generic getter
      const [dataType, data] = await etherDatabase.read.getValue([segments])
      expect(dataType).to.equal(DataType.UINT)
      expect(decodeAbiParameters([{ type: "uint256" }], data)[0]).to.equal(456n)
    })
  })
})
