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
    mapping(uint256 => uint256) public index2Height;
    mapping(string => bool) public txUnlocked;
    mapping(bytes32 => EnumerableSet.AddressSet) addrsSet;
    uint256 public appchainIndex = 0;
    uint256 public relayIndex = 0;

    bytes32 public constant RELAYER_ROLE = "RELAYER_ROLE"; //0x52454c415945525f524f4c450000000000000000000000000000000000000000
    bytes32 public constant PIER_ROLE = "PIER_ROLE";   //0x504945525f524f4c450000000000000000000000000000000000000000000000

    event Lock(
        address ethToken,
        address relayToken,
        address locker,
        address recipient,
        uint256 amount,
        uint256 appchainIndex
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

    function init() public onlyAdmin {
        appchainIndex = 0;
        relayIndex = 0;
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
        appchainIndex = appchainIndex.add(1);
        index2Height[appchainIndex] = block.number;
        emit Lock(
            token,
            supportToken[token],
            msg.sender,
            recipient,
            amount,
            appchainIndex
        );
    }

    function unlock(
        address token,
        address from,
        address recipient,
        uint256 amount,
        string memory _txid,
        uint256 _relayIndex,
        bytes[] memory signatures
    ) public onlySupportToken(token) onlyCrosser whenNotUnlocked(_txid) {
        if (relayIndex != _relayIndex - 1) {
            revert("index not match");
        }
        uint N = getRoleMemberCount(RELAYER_ROLE);
        uint threshold = (N + (N - 1) / 3 + 2) / 2;
        if (signatures.length < threshold) {
            return;
        }
        bytes32 hash = keccak256(
            abi.encodePacked(token, from, recipient, amount, _txid)
        );

        for (uint256 i; i < signatures.length; i++) {
            address relayer = recover(ECDSA.toEthSignedMessageHash(hash), signatures[i]);
            if (hasRole(RELAYER_ROLE, relayer)) {
                EnumerableSet.add(addrsSet[hash], relayer);
            }
        }

        if (EnumerableSet.length(addrsSet[hash]) < threshold) {
            revert("signatures invalid");
        }

        txUnlocked[_txid] = true;
        relayIndex = relayIndex.add(1);
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

    function recover(bytes32 hash, bytes memory signature) internal pure returns (address) {
        // Check the signature length
        if (signature.length != 65) {
            //            revert("ECDSA: invalid signature length");
            return address(0);
        }

        // Divide the signature in r, s and v variables
        bytes32 r;
        bytes32 s;
        uint8 v;

        // ecrecover takes the signature parameters, and the only way to get them
        // currently is to use assembly.
        // solhint-disable-next-line no-inline-assembly
        assembly {
            r := mload(add(signature, 0x20))
            s := mload(add(signature, 0x40))
            v := byte(0, mload(add(signature, 0x60)))
        }
        v += 27;

        // EIP-2 still allows signature malleability for ecrecover(). Remove this possibility and make the signature
        // unique. Appendix F in the Ethereum Yellow paper (https://ethereum.github.io/yellowpaper/paper.pdf), defines
        // the valid range for s in (281): 0 < s < secp256k1n ÷ 2 + 1, and for v in (282): v ∈ {27, 28}. Most
        // signatures from current libraries generate a unique signature with an s-value in the lower half order.
        //
        // If your library generates malleable signatures, such as s-values in the upper range, calculate a new s-value
        // with 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141 - s1 and flip v from 27 to 28 or
        // vice versa. If your library also generates signatures with 0/1 for v instead 27/28, add 27 to v to accept
        // these malleable signatures as well.
        if (uint256(s) > 0x7FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF5D576E7357A4501DDFE92F46681B20A0) {
            //            revert("ECDSA: invalid signature 's' value");
            return address(0);
        }

        if (v != 27 && v != 28) {
            //            revert("ECDSA: invalid signature 'v' value");
            return address(0);
        }

        // If the signature is valid (and not malleable), return the signer address
        address signer = ecrecover(hash, v, r, s);
        return signer;
    }
}