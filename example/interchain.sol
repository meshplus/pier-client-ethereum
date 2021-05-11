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
    mapping(address => mapping(address => uint256)) public mintAmount;
    mapping(string => bool) public txUnlocked;
    mapping(bytes32 => EnumerableSet.AddressSet) addrsSet;

    bytes32 public constant RELAYER_ROLE = "RELAYER_ROLE"; //0x52454c415945525f524f4c450000000000000000000000000000000000000000
    bytes32 public constant PIER_ROLE = "PIER_ROLE";

    event Burn(
        address ethToken,
        address relayToken,
        address burner,
        address recipient,
        uint256 amount
    );
    event Mint(
        address ethToken,
        address relayToken,
        address from,
        address recipient,
        uint256 amount,
        string txid
    );

    constructor(address[] memory _relayers) public {
        _setupRole(DEFAULT_ADMIN_ROLE, msg.sender);
        for (uint256 i; i < _relayers.length; i++) {
            _setupRole(RELAYER_ROLE, _relayers[i]);
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


    function burn(address token, uint256 amount, address recipient) public onlySupportToken(token) {
        mintAmount[token][msg.sender] = mintAmount[token][msg.sender].sub(
            amount
        );
        IMintBurn(token).burn(recipient, amount);
        emit Burn(bxh2ethToken[token], token, msg.sender, recipient, amount);
    }

    function mint(
        address ethToken,
        address relayToken,
        address from,
        address recipient,
        uint256 amount,
        string memory _txid
    ) public onlySupportToken(ethToken) onlyCrosser whenNotUnlocked(_txid) {
        require(eth2bxhToken[ethToken] == bxh2ethToken[relayToken], "Burn::Not Support Token");

        txUnlocked[_txid] = true;
        mintAmount[relayToken][recipient] = mintAmount[relayToken][recipient].add(amount);
        IMintBurn(relayToken).mint(recipient, amount);
        emit Mint(ethToken, relayToken, from, recipient, amount, _txid);
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

    modifier whenNotUnlocked(string memory _txid) {
        require(txUnlocked[_txid] == false, "tx unlocked");
        _;
    }
}