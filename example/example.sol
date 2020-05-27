pragma solidity ^0.5.6;

contract Store {
    event ItemSet(string key, string value);

    mapping (string => string) public items;

    function setItem(string memory key, string memory value) public {
        items[key] = value;
        emit ItemSet(key, value);
    }

    function getItem(string memory key) public view returns(string memory) {
        return items[key];
    }
}