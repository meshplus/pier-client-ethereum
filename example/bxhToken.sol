// SPDX-License-Identifier: MIT

pragma solidity ^0.6.12;

import "@openzeppelin/contracts/token/ERC20/ERC20Pausable.sol";
import "@openzeppelin/contracts/utils/Pausable.sol";
import "@openzeppelin/contracts/access/AccessControl.sol";
import "./interface/IMintBurn.sol";
import "./interface/IPause.sol";

contract BaseToken is ERC20Pausable, AccessControl, IMintBurn, IPause{
    bytes32 public constant MINTER_ROLE ="MINTER_ROLE"; //0x4d494e5445525f524f4c45000000000000000000000000000000000000000000
    bytes32 public constant BURNER_ROLE ="BURNER_ROLE";//0x4255524e45525f524f4c45000000000000000000000000000000000000000000
    bytes32 public constant LIQUIDATION = "LIQUIDATION";
    constructor(
        string memory _name,
        string memory _symbol,
        uint8 decimal_,
        address admin
    ) public ERC20(_name, _symbol) {
        _setupRole(DEFAULT_ADMIN_ROLE, admin);
        _setupDecimals(decimal_);
    }

    function mint(address account, uint amount) public virtual override onlyMinter{
        _mint(account, amount);
    }

    function burn(address account, uint amount) public virtual override onlyBurner{
        _burn(account, amount);
    }

    function pause() public override onlyLiquidation {
        _pause();
    }

    function unpause() public override onlyLiquidation {
        _unpause();
    }

    // minter will only be tunnel
    modifier onlyMinter {
        require(hasRole(MINTER_ROLE, msg.sender), "Caller is not a minter");
        _;
    }

    modifier onlyBurner {
        require(hasRole(BURNER_ROLE, msg.sender), "Caller is not a burner");
        _;
    }

    modifier onlyLiquidation {
        require(hasRole(LIQUIDATION, msg.sender), "Caller is not liquidation contract");
        _;
    }
}


contract BxhToken is BaseToken {

    constructor(
        string memory name_,
        string memory symbol_,
        uint8 decimal_,
        address admin
    ) public BaseToken(name_, symbol_, decimal_, admin) {}

}