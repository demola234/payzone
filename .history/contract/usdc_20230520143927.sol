pragma solidity ^0.8.0;

import "./ERC20.sol";

// create a contract to crerate usdc wallets and transfer between usdc wallets
contract Usdc is ERC20 {
    // create a mapping to store usdc wallets
    mapping(address => bool) public usdcWallets;

    // create a constructor to set the name, symbol and decimals of the usdc token
    constructor() ERC20("USDC", "USDC", 6) {
        // mint 1000000000 usdc tokens to the contract deployer
        _mint(msg.sender, 1000000000 * 10**decimals());
    }

    // create a function to create usdc wallets
    function createUsdcWallet() public {
        // check if the usdc wallet is already created
        require(usdcWallets[msg.sender] == false, "USDC: wallet already created");

        // set the usdc wallet to true
        usdcWallets[msg.sender] = true;
    }

    // create a function to transfer usdc tokens between usdc wallets
    function transferUsdc(address _to, uint256 _amount) public {
        // check if the usdc wallet is created
        require(usdcWallets[msg.sender] == true, "USDC: wallet not created");
        /

        // transfer usdc tokens
        _transfer(msg.sender, _to, _amount);
    }
}