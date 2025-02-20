// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

error Unauthorized();
error IdentityAlreadyExists();
error InvalidIdentity();
error IdentityDoesNotExist();

abstract contract RoleControl {
    enum Role {
        EMITTER,
        REGISTRAR,
        ADMIN,
        DBSET
    }

    struct Identity {
        bool valid;
        Role[] roles;
    }

    struct IdentityView {
        address walletAddress;
        Role[] roles;
    }

    // State variables
    mapping(address => Identity) private identities;
    address[] private registeredWallets;
    address public owner;

    event RoleGranted(address indexed wallet, Role indexed role);
    event RoleRevoked(address indexed wallet, Role indexed role);
    event IdentityCreated(address indexed wallet);
    event IdentityDeleted(address indexed wallet);

    function initialize(address _owner) internal {
        owner = _owner;
        
        // Create identity for owner with all roles
        Role[] memory ownerRoles = new Role[](4);
        ownerRoles[0] = Role.ADMIN;
        ownerRoles[1] = Role.EMITTER;
        ownerRoles[2] = Role.REGISTRAR;
        ownerRoles[3] = Role.DBSET;
        
        _createIdentity(_owner, ownerRoles);
    }

    modifier auth(Role role) {
        if (!identities[msg.sender].valid) {
            revert IdentityDoesNotExist();
        }

        if (!hasRole(msg.sender, role)) {
            revert Unauthorized();
        }
        _;
    }

    // External Functions
    function getAllWallets() external view returns (address[] memory) {
        return registeredWallets;
    }

    function getAllIdentities() external view returns (IdentityView[] memory) {
        IdentityView[] memory arr = new IdentityView[](registeredWallets.length);
        for (uint256 i = 0; i < registeredWallets.length; i++) {
            address wallet = registeredWallets[i];
            Identity storage data = identities[wallet];
            arr[i] = IdentityView({
                walletAddress: wallet,
                roles: data.roles
            });
        }
        return arr;
    }
    
    function createWalletIdentity(address walletAddress, Role[] memory initialRoles) external auth(Role.ADMIN) {
        _createIdentity(walletAddress, initialRoles);
    }

    function grantRole(address wallet, Role role) external auth(Role.ADMIN) {
        _grantRole(wallet, role);
    }

    function revokeRole(address wallet, Role role) external auth(Role.ADMIN) {
        _revokeRole(wallet, role);
    }

    function deleteIdentity(address wallet) external auth(Role.ADMIN) {
        _deleteIdentity(wallet);
    }

    // Internal Logic
    function hasRole(address wallet, Role role) internal view returns (bool) {
        Role[] storage roles = identities[wallet].roles;
        for (uint256 i = 0; i < roles.length; i++) {
            if (roles[i] == role) return true;
        }
        return false;
    }

    function requireRole(address caller, Role role) internal view {
        if (!identities[caller].valid) {
            revert IdentityDoesNotExist();
        }

        if (!hasRole(caller, role)) {
            revert Unauthorized();
        }
    }

    function _createIdentity(address wallet, Role[] memory initialRoles) internal {
        if (identities[wallet].valid) {
            revert IdentityAlreadyExists();
        }

        identities[wallet].valid = true;
        registeredWallets.push(wallet);

        // Grant initial roles
        for (uint256 i = 0; i < initialRoles.length; i++) {
            _grantRole(wallet, initialRoles[i]);
        }

        emit IdentityCreated(wallet);
    }

    function _grantRole(address wallet, Role role) internal {
        if (!identities[wallet].valid) revert IdentityDoesNotExist();
        Identity storage id = identities[wallet];
        id.roles.push(role);
        emit RoleGranted(wallet, role);
    }

    function _revokeRole(address wallet, Role role) internal {
        if (!identities[wallet].valid) revert IdentityDoesNotExist();

        Identity storage id = identities[wallet];
        Role[] storage rolesArr = id.roles;
        for (uint256 i = 0; i < rolesArr.length; i++) {
            if (rolesArr[i] == role) {
                rolesArr[i] = rolesArr[rolesArr.length - 1];
                rolesArr.pop();
                emit RoleRevoked(wallet, role);
                return;
            }
        }
    }

    function _deleteIdentity(address wallet) internal {
        Identity storage id = identities[wallet];
        if (id.roles.length == 0) revert InvalidIdentity();

        delete identities[wallet];
        for (uint256 i = 0; i < registeredWallets.length; i++) {
            if (registeredWallets[i] == wallet) {
                registeredWallets[i] = registeredWallets[registeredWallets.length - 1];
                registeredWallets.pop();
                break;
            }
        }
        emit IdentityDeleted(wallet);
    }
}
