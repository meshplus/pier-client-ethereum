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

    mapping(address => address) public appToken2Pier;
    mapping(address => address) public app2bxhToken;
    mapping(address => address) public bxh2appToken;
    mapping(address => uint256) public mintAmount;
    mapping(address => mapping(uint256 => uint256)) public index2Height;
    mapping(string => bool) public txMinted;
    mapping(address => uint256) public appchainIndex;
    mapping(address => uint256) public relayIndex;

    bytes32 public constant PIER_ROLE = "PIER_ROLE"; //0x504945525f524f4c450000000000000000000000000000000000000000000000

    event Burn(
        address pier,
        address appToken,
        address relayToken,
        address burner,
        address recipient,
        uint256 amount,
        uint256 relayIndex
    );
    event Mint(
        address appToken,
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

    function addSupportToken(address pierId, address appTokenAddr, address relayTokenAddr) public onlyAdmin {
        require(
            app2bxhToken[appTokenAddr] == address(0),
            "Token already Supported"
        );
        require(
            bxh2appToken[relayTokenAddr] == address(0),
            "Token already Supported"
        );
        appToken2Pier[appTokenAddr] = pierId;
        app2bxhToken[appTokenAddr] = relayTokenAddr;
        bxh2appToken[relayTokenAddr] = appTokenAddr;
    }

    function removeSupportToken(address appTokenAddr) public onlyAdmin {
        require(app2bxhToken[appTokenAddr] != address(0), "Token not Supported");
        delete bxh2appToken[app2bxhToken[appTokenAddr]];
        delete app2bxhToken[appTokenAddr];

    }

    function addSupportTokens(
        address[] memory pierIds,
        address[] memory appTokenAddrs,
        address[] memory relayTokenAddrs
    ) public {
        require(
            appTokenAddrs.length == relayTokenAddrs.length,
            "Token length not match"
        );
        for (uint256 i; i < appTokenAddrs.length; i++) {
            addSupportToken(pierIds[i], appTokenAddrs[i], relayTokenAddrs[i]);
        }
    }

    function removeSupportTokens(address[] memory addrs) public {
        for (uint256 i; i < addrs.length; i++) {
            removeSupportToken(addrs[i]);
        }
    }

    function lockRollback(
        address appToken,
        address relayToken,
        address from,
        address recipient,
        uint256 amount,
        string memory _txid,
        uint256 _appchainIndex
    ) public {
        mint(appToken, relayToken, from, address(this), amount,_txid, _appchainIndex);
        IERC20(relayToken).approve(address(this), amount);
        burnSelf(relayToken, amount, from);
    }

    function burnSelf(address relayToken, uint256 amount, address recipient) internal onlySupportToken(bxh2appToken[relayToken]) {
        mintAmount[relayToken] = mintAmount[relayToken].sub(
            amount
        );
        IMintBurn(relayToken).burn(address(this), amount);
        address appchainToken = bxh2appToken[relayToken];
        address pierAddr = appToken2Pier[appchainToken];
        relayIndex[pierAddr] = relayIndex[pierAddr].add(1);
        index2Height[pierAddr][relayIndex[pierAddr]]=block.number;
        emit Burn(pierAddr, appchainToken, relayToken, address(this), recipient, amount, relayIndex[pierAddr]);
    }

    function burn(address relayToken, uint256 amount, address recipient) public onlySupportToken(bxh2appToken[relayToken]) {
        mintAmount[relayToken] = mintAmount[relayToken].sub(
            amount
        );
        IMintBurn(relayToken).burn(msg.sender, amount);
        address appchainToken = bxh2appToken[relayToken];
        address pierAddr = appToken2Pier[appchainToken];
        relayIndex[pierAddr] = relayIndex[pierAddr].add(1);
        index2Height[pierAddr][relayIndex[pierAddr]]=block.number;
        emit Burn(pierAddr, appchainToken, relayToken, msg.sender, recipient, amount, relayIndex[pierAddr]);
    }

    function mint(
        address appToken,
        address relayToken,
        address from,
        address recipient,
        uint256 amount,
        string memory _txid,
        uint256 _appchainIndex
    ) public onlySupportToken(appToken) onlyCrosser whenNotMinted(_txid) {
        require(app2bxhToken[appToken] == relayToken, "Burn::Not Support Token");
        address pierAddr = appToken2Pier[appToken];
        if (appchainIndex[pierAddr] != _appchainIndex - 1) {
            revert("index not match");
        }
        txMinted[_txid] = true;
        appchainIndex[pierAddr] = appchainIndex[pierAddr].add(1);
        mintAmount[relayToken] = mintAmount[relayToken].add(amount);
        IMintBurn(relayToken).mint(recipient, amount);
        emit Mint(appToken, relayToken, from, recipient, amount, _txid, appchainIndex[pierAddr]);
    }

    modifier onlySupportToken(address token) {
        require(app2bxhToken[token] != address(0), "Mint or Burn::Not Support Token");
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