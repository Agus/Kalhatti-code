pragma solidity ^0.4.24;

contract Loyalty{
    mapping(address => uint256) LoyaltyBalance;
    address owner;

    constructor() public{
        owner = msg.sender;
    }

    modifier isOwner{
        require(msg.sender == owner);
        _;
    }

    function makePayment() public payable{
        uint LoyaltyPoints = msg.value / 10;
        //Check if sender has balance and for overflows
        require(LoyaltyBalance[msg.sender] + LoyaltyPoints >= LoyaltyBalance[msg.sender]);
        LoyaltyBalance[msg.sender] = LoyaltyBalance[msg.sender] + LoyaltyPoints;
    }

    function refundPayment(uint amount, address receiver) public isOwner{
        uint LoyaltyPoints = amount / 10;
        //Check if sender has balance and for underflows
        require(LoyaltyBalance[receiver] - LoyaltyPoints >= 0);
        require(receiver.send(amount));
        LoyaltyBalance[receiver] = LoyaltyBalance[receiver] - LoyaltyPoints;
    }

}
