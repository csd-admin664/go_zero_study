//  SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract FundToken {
    //通证的名称
    //通证的简称
    //发行总量
    //onwer 地址
    //每个地址有多少通证
    //mint 获取通证
    //转移通证
    //查看地址下通证的数量
    string  public tokenName;
    string public tokenSymbol;
    uint256 public totalSupply;
    address public owner;
    mapping(address => uint256) public balanceOf;

    constructor(string  memory _tokenName, string memory _tokenSymbol) {
        tokenName = _tokenName;
        tokenSymbol = _tokenSymbol;
        owner = msg.sender;
    }

    function mint(unint256 amountToMint)  {
        balanceOf[msg.sender] += amountToMint;
        totalSupply += amountToMint;
    }
}