pragma solidity ^0.6.8;
pragma experimental ABIEncoderV2;

import "./@openzeppelin/contracts/access/AccessControl.sol";
import "./@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "./@openzeppelin/contracts/token/ERC20/SafeERC20.sol";
import "./@openzeppelin/contracts/access/Ownable.sol";
import "./@openzeppelin/contracts/cryptography/ECDSA.sol";
import "./lib/SafeDecimalMath.sol";
import "./@openzeppelin/contracts/utils/EnumerableSet.sol";


contract Escrows is AccessControl {
    using SafeERC20 for IERC20;
    using SafeMath for uint256;
    using SafeDecimalMath for uint256;
    using EnumerableSet for EnumerableSet.AddressSet;

    mapping(address => address) public supportToken;
    mapping(address => mapping(address => uint256)) public lockAmount;
    mapping(string => bool) public txUnlocked;
    mapping(bytes32 => EnumerableSet.AddressSet) addrsSet;

    bytes32 public constant RELAYER_ROLE = "RELAYER_ROLE"; //0x52454c415945525f524f4c450000000000000000000000000000000000000000
    bytes32 public constant PIER_ROLE = "PIER_ROLE";

    event Lock(
        address ethToken,
        address relayToken,
        address locker,
        address recipient,
        uint256 amount
    );
    event Unlock(
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
            supportToken[ethTokenAddr] == address(0),
            "Token already Supported"
        );
        supportToken[ethTokenAddr] = relayTokenAddr;
    }

    function removeSupportToken(address ethTokenAddr) public onlyAdmin {
        require(supportToken[ethTokenAddr] != address(0), "Token not Supported");
        delete supportToken[ethTokenAddr];
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


    function lock(address token, uint256 amount, address recipient) public onlySupportToken(token) {
        lockAmount[token][msg.sender] = lockAmount[token][msg.sender].add(
            amount
        );
        IERC20(token).safeTransferFrom(msg.sender, address(this), amount);
        emit Lock(
            token,
            supportToken[token],
            msg.sender,
            recipient,
            amount
        );
    }

    function unlock(
        address token,
        address from,
        address recipient,
        uint256 amount,
        string memory _txid,
        bytes[] memory signatures
    ) public onlySupportToken(token) onlyCrosser whenNotUnlocked(_txid) {
        uint N = getRoleMemberCount(RELAYER_ROLE);
        uint f = (N -1)/3;
        uint threshold = (N + f + 2) / 2;
        if (signatures.length < threshold) {
            return;
        }
        bytes32 hash = keccak256(
            abi.encodePacked(token, from, recipient, amount, _txid)
        );

        for (uint256 i; i < signatures.length; i++) {
            address relayer = ECDSA.recover(ECDSA.toEthSignedMessageHash(hash), signatures[i]);
            if (hasRole(RELAYER_ROLE, relayer)) {
                EnumerableSet.add(addrsSet[hash], relayer);
            }
        }

        if (EnumerableSet.length(addrsSet[hash]) < threshold) {
            revert("signatures invaild");
        }

        txUnlocked[_txid] = true;
        lockAmount[token][recipient] = lockAmount[token][recipient].sub(amount);
        IERC20(token).safeTransfer(recipient, amount);
        emit Unlock(token, supportToken[token], from, recipient, amount, _txid);
    }

    modifier onlySupportToken(address token) {
        require(supportToken[token] != address(0), "Lock::Not Support Token");
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