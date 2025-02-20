import hre from "hardhat";
import { expect } from "chai";
import { keccak256 } from "viem";

// Helper function to convert address to identity key
const addressToKey = (address: string) => {
  return keccak256(address);
};

// Deployment fixture
const deployRoleControl = async () => {
  const [deployer, user1, user2] = await hre.viem.getWalletClients();

  const roleControl = await hre.viem.deployContract("RoleControlImpl", [
    deployer.account.address,
  ]);

  return { roleControl, deployer, user1, user2 };
};

describe("RoleControl", () => {
  describe("Role Granting", () => {
    it("should set up owner with all roles during deployment", async () => {
      const { roleControl, deployer } = await deployRoleControl();
      const ownerKey = addressToKey(deployer.account.address);

      // Check all roles
      expect(await roleControl.read.hasRole([ownerKey, 0])).to.be.true; // EMITTER
      expect(await roleControl.read.hasRole([ownerKey, 1])).to.be.true; // REGISTRAR
      expect(await roleControl.read.hasRole([ownerKey, 2])).to.be.true; // ADMIN
    });

    it("should allow admin to create new identity", async () => {
      const { roleControl, user1 } = await deployRoleControl();
      await roleControl.write.createWalletIdentity([user1.account.address]);

      const events = await roleControl.getEvents.IdentityCreated();
      expect(events.length).to.be.equal(1);
      const user1Key = addressToKey(user1.account.address);
      expect(events[0].args.identity).to.be.equal(user1Key);
    });

    it("should allow admin to grant roles", async () => {
      const { roleControl, user1 } = await deployRoleControl();
      // First create identity
      await roleControl.write.createWalletIdentity([user1.account.address]);
      const user1Key = addressToKey(user1.account.address);

      // Grant EMITTER role
      await roleControl.write.grantRole([user1Key, 0]);

      const events = await roleControl.getEvents.RoleGranted();
      expect(events.length).to.be.equal(1);
      expect(events[0].args.identity).to.be.equal(user1Key);
      expect(events[0].args.role).to.be.equal(0);

      expect(await roleControl.read.hasRole([user1Key, 0])).to.be.true;
    });

    it("should prevent non-admin from granting roles", async () => {
      const { roleControl, user1, user2 } = await deployRoleControl();

      await roleControl.write.createWalletIdentity([user1.account.address]);

      // Create identity for user2
      await roleControl.write.createWalletIdentity([user2.account.address]);
      const user2Key = addressToKey(user2.account.address);

      // Try to grant role from non-admin account
      await expect(
        roleControl.write.grantRole([user2Key, 0], {
          account: user1.account,
        })
      ).to.be.revertedWithCustomError(roleControl, "Unauthorized");
    });

    it("should prevent creating duplicate identities", async () => {
      const { roleControl, user1 } = await deployRoleControl();

      await roleControl.write.createWalletIdentity([user1.account.address]);
      const user1Key = addressToKey(user1.account.address);

      // add role to "make" the user
      await roleControl.write.grantRole([user1Key, 0]);

      await expect(
        roleControl.write.createWalletIdentity([user1.account.address])
      ).to.be.revertedWithCustomError(roleControl, "IdentityAlreadyExists");
    });

    it("should allow viewing all identities", async () => {
      const { roleControl, deployer, user1 } = await deployRoleControl();

      await roleControl.write.createWalletIdentity([user1.account.address]);
      const user1Key = addressToKey(user1.account.address);

      await roleControl.write.grantRole([user1Key, 0]); // Grant EMITTER role

      const allIdentities = await roleControl.read.getAllIdentities();
      expect(allIdentities.length).to.equal(2); // Owner + user1

      const ownerKey = addressToKey(deployer.account.address);
      expect(allIdentities[0].identityKey).to.equal(ownerKey);
      expect(allIdentities[1].identityKey).to.equal(user1Key);
    });
  });
});
