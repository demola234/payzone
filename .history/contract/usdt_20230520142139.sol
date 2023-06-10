pragma solidity ^0.8.0;

import "./ERC20.sol";

// create a contract to crerate usdt wallets and transfer between usdt wallets
contract Usdt is ERC20 {
    // create a mapping to store usdt wallets
    mapping(address => bool) public usdtWallets;

    // create a constructor to set the name, symbol and decimals of the usdt token
    constructor() ERC20("USDT", "USDT", 6) {
        // mint 1000000000 usdt tokens to the contract deployer
        _mint(msg.sender, 1000000000 * 10**decimals());
    }

    // create a function to create usdt wallets
    function createUsdtWallet() public {
        // check if the usdt wallet is already created
        require(usdtWallets[msg.sender] == false, "USDT: wallet already created");

        // set the usdt wallet to true
        usdtWallets[msg.sender] = true;
    }

    // create a function to transfer usdt tokens between usdt wallets
    function transferUsdt(address _to, uint256 _amount) public {
        // check if the usdt wallet is created
        require(usdtWallets[msg.sender] == true, "USDT: wallet not created");

        // transfer usdt tokens
        _transfer(msg.sender, _to, _amount);
    }
}