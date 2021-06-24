pragma solidity ^0.6.8;
pragma experimental ABIEncoderV2;

import "./@openzeppelin/contracts/access/AccessControl.sol";
import "./@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "./@openzeppelin/contracts/token/ERC20/SafeERC20.sol";
import "./@openzeppelin/contracts/access/Ownable.sol";
import "./@openzeppelin/contracts/cryptography/ECDSA.sol";
import "./lib/SafeDecimalMath.sol";
import "./@openzeppelin/contracts/utils/EnumerableSet.sol";
import "./interface/IMintBurn.sol";


contract InterchainSwap is AccessControl {
    using SafeERC20 for IERC20;
    using SafeMath for uint256;
    using SafeDecimalMath for uint256;
    using EnumerableSet for EnumerableSet.AddressSet;

    mapping(address => address) public eth2bxhToken;
    mapping(address => address) public bxh2ethToken;
    mapping(address => uint256) public mintAmount;
    mapping(uint256 => uint256) public index2Height;
    mapping(string => bool) public txMinted;
    uint256 public appchainIndex = 0;
    uint256 public relayIndex = 0;

    bytes32 public constant PIER_ROLE = "PIER_ROLE"; //0x504945525f524f4c450000000000000000000000000000000000000000000000

    event Burn(
        address ethToken,
        address relayToken,
        address burner,
        address recipient,
        uint256 amount,
        uint256 relayIndex
    );
    event Mint(
        address ethToken,
        address relayToken,
        address from,
        address recipient,
        uint256 amount,
        string txid,
        uint256 appchainIndex
    );

    constructor(address[] memory _piers) public {
        _setupRole(DEFAULT_ADMIN_ROLE, msg.sender);
        for (uint256 i; i < _piers.length; i++) {
            _setupRole(PIER_ROLE, _piers[i]);
        }
    }

    function addSupportToken(address ethTokenAddr, address relayTokenAddr) public onlyAdmin {
        require(
            eth2bxhToken[ethTokenAddr] == address(0),
            "Token already Supported"
        );
        require(
            bxh2ethToken[relayTokenAddr] == address(0),
            "Token already Supported"
        );
        eth2bxhToken[ethTokenAddr] = relayTokenAddr;
        bxh2ethToken[relayTokenAddr] = ethTokenAddr;
    }

    function removeSupportToken(address ethTokenAddr) public onlyAdmin {
        require(eth2bxhToken[ethTokenAddr] != address(0), "Token not Supported");
        delete bxh2ethToken[eth2bxhToken[ethTokenAddr]];
        delete eth2bxhToken[ethTokenAddr];

    }

    function addSupportTokens(
        address[] memory ethTokenAddrs,
        address[] memory relayTokenAddrs
    ) public {
        require(
            ethTokenAddrs.length == relayTokenAddrs.length,
            "Token length not match"
        );
        for (uint256 i; i < ethTokenAddrs.length; i++) {
            addSupportToken(ethTokenAddrs[i], relayTokenAddrs[i]);
        }
    }

    function removeSupportTokens(address[] memory addrs) public {
        for (uint256 i; i < addrs.length; i++) {
            removeSupportToken(addrs[i]);
        }
    }


    function burn(address relayToken, uint256 amount, address recipient) public onlySupportToken(bxh2ethToken[relayToken]) {
        mintAmount[relayToken] = mintAmount[relayToken].sub(
            amount
        );
        IMintBurn(relayToken).burn(msg.sender, amount);
        relayIndex = relayIndex.add(1);
        index2Height[relayIndex]=block.number;
        emit Burn(bxh2ethToken[relayToken], relayToken, msg.sender, recipient, amount, relayIndex);
    }

    function mint(
        address ethToken,
        address relayToken,
        address from,
        address recipient,
        uint256 amount,
        string memory _txid,
        uint256 _appchainIndex
    ) public onlySupportToken(ethToken) onlyCrosser whenNotMinted(_txid) {
        require(eth2bxhToken[ethToken] == relayToken, "Burn::Not Support Token");
        if (appchainIndex != _appchainIndex - 1) {
            revert("index not match");
        }
        txMinted[_txid] = true;
        appchainIndex = appchainIndex.add(1);
        mintAmount[relayToken] = mintAmount[relayToken].add(amount);
        IMintBurn(relayToken).mint(recipient, amount);
        emit Mint(ethToken, relayToken, from, recipient, amount, _txid, appchainIndex);
    }

    modifier onlySupportToken(address token) {
        require(eth2bxhToken[token] != address(0), "Mint or Burn::Not Support Token");
        _;
    }

    modifier onlyAdmin {
        require(hasRole(DEFAULT_ADMIN_ROLE, msg.sender), "caller is not admin");
        _;
    }

    modifier onlyCrosser {
        require(hasRole(PIER_ROLE, msg.sender), "caller is not crosser");
        _;
    }

    modifier whenNotMinted(string memory _txid) {
        require(txMinted[_txid] == false, "tx minted");
        _;
    }
}