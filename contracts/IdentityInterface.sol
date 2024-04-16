// SPDX-License-Identifier: MIT

pragma solidity ^0.8.20;

interface IdentityInterface {
    event NewManager(address indexed node, address manager);
    function register(address identityManagerContract, bytes memory signature) external;
    function setManager(address manager) external;
    function manager(address node) external view returns (address);
    function exists(address node) external view returns (bool);
}