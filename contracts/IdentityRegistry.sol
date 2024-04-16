// SPDX-License-Identifier: MIT

pragma solidity ^0.8.20;

import "./IdentityInterface.sol";

contract IdentityRegistry is IdentityInterface {
    struct Circuit {
        string deploymentType;
        address deploymentAddress;
        string ipfsHash;
    }

    address public owner;
    mapping(address => address) public identities;
    mapping(string => Circuit) public circuits;
    bool public running = true;

    modifier onlyRunning {
        require(running, "Contract is not running");
        _;
    }

    modifier authorised {
        address node = msg.sender;
        require(identities[node] == msg.sender);
        _;
    }

    modifier onlyOwner {
        require(msg.sender == owner);
        _;
    }


    constructor() {
        owner = msg.sender;
    }

    function destroy() external onlyOwner {
        running = false;
    }

    function register(address identityManagerContract, bytes memory signature) external onlyRunning {
        //check if user is already registered, no need to calculate hash if its already registered
        if (identities[msg.sender] != address(0x0)) {
            revert("Already registered. Use setManager call to update manager address.");
        }
        bytes32 messageHash = _hash(identityManagerContract, msg.sender);
        if (!_verify(messageHash, owner, signature)) {
            revert("Invalid signature");
        }
        identities[msg.sender] = identityManagerContract;
        // check identity exist
        emit NewManager(msg.sender, identityManagerContract);

    }


    function setManager(address manager) public onlyRunning authorised {
        identities[msg.sender] = manager;
        emit NewManager(msg.sender, manager);
    }


    function manager(address node) public view returns (address) {
        return identities[node];
    }


    function exists(address node) public view returns (bool) {
        return identities[node] != address(0x0);
    }

    function ownerOf(address node) external view returns (address){
        return identities[node];
    }

    function deregister(address node) external onlyRunning authorised {
        delete identities[node];
    }

    function setCircuit(string memory _circuitId, string memory _deploymentType,
        address _deploymentAddress, string memory _ipfsHash) external onlyRunning onlyOwner {
        circuits[_circuitId] = Circuit(_deploymentType, _deploymentAddress, _ipfsHash);
    }

    function getCircuit(string memory _circuitId) external view returns (string memory, address, string memory) {
        return (circuits[_circuitId].deploymentType, circuits[_circuitId].deploymentAddress, circuits[_circuitId].ipfsHash);
    }

    function splitSignature(
        bytes memory sig
    ) pure internal returns (bytes32 r, bytes32 s, uint8 v) {
        require(sig.length == 65, "invalid signature length");
        assembly {
        // first 32 bytes, after the length prefix
            r := mload(add(sig, 32))
        // second 32 bytes
            s := mload(add(sig, 64))
        // final byte
            v := byte(0, mload(add(sig, 96)))
        }
        // implicitly return (r, s, v)
    }

    function _hash(address identityManagerContract, address sender) pure internal returns (bytes32) {
        return keccak256(abi.encodePacked(identityManagerContract, sender));

    }

    function _verify(bytes32 _hashedMessage, address signer, bytes memory signature) pure internal returns (bool) {
        bytes memory prefix = "\x19Ethereum Signed Message:\n32";
        bytes32 prefixedHashMessage = keccak256(abi.encodePacked(prefix, _hashedMessage));
        (bytes32 _r, bytes32 _s, uint8 _v) = splitSignature(signature);
        address recovered = ecrecover(prefixedHashMessage, _v, _r, _s);
        return signer == recovered;
    }

}

