pragma solidity ^0.6.8;
pragma experimental ABIEncoderV2;

import "./@openzeppelin/contracts/access/AccessControl.sol";
import "./@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "./@openzeppelin/contracts/token/ERC20/SafeERC20.sol";
import "./@openzeppelin/contracts/access/Ownable.sol";
import "./@openzeppelin/contracts/cryptography/ECDSA.sol";
import "./lib/SafeDecimalMath.sol";
import "./@openzeppelin/contracts/utils/EnumerableSet.sol";
import "./interchain.sol";
import "./IUniswapV2Router01.sol";


contract Proxy is AccessControl {
    using SafeERC20 for IERC20;
    using SafeMath for uint256;
    using SafeDecimalMath for uint256;
    using EnumerableSet for EnumerableSet.AddressSet;

    mapping(address => address) public supportToken;
    mapping(address => uint256) public lockAmount;
    mapping(uint256 => uint256) public index2Height;
    mapping(string => bool) public txUnlocked;
    mapping(bytes32 => EnumerableSet.AddressSet) addrsSet;
    address public _interchainAddr;
    address public _dexAddr;

    function proxy(address appToken,
        address relayToken,
        address from,
        address recipient,
        uint256 amount,
        string memory _txid,
        uint256 _appchainIndex,
        address dstChainId,
        address dstContract
    ) public {
        InterchainSwap(_interchainAddr).mint(appToken, relayToken,from, address(this),amount,_txid, _appchainIndex);
        IERC20(relayToken).approve(_dexAddr, 10000000000000000000);
        // todo addSupportToken
        address dstRelayToken = InterchainSwap(_interchainAddr).app2bxhToken(dstContract);
        //address[2] memory adds = [relayToken, dstRelayToken];
        address[] memory adds = new address[](2);
        adds[0] = relayToken;
        adds[1] = dstRelayToken;
        uint[] memory amounts = IUniswapV2Router01(_dexAddr).swapExactTokensForTokens(amount,0, adds, address(this), block.number + 10);
        IERC20(dstRelayToken).approve(_interchainAddr, 100000000000000000000000);
        InterchainSwap(_interchainAddr).burn(dstRelayToken, amounts[amounts.length-1], recipient);
    }

   constructor(address interchainAddr, address dexAddr) public {
        _setupRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _interchainAddr = interchainAddr;
        _dexAddr = dexAddr;
    }
}