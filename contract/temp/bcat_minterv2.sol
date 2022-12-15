// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";


interface IERC20 {
    /**
     * @dev Returns the amount of tokens in existence.
     */
    function totalSupply() external view returns (uint256);

    /**
     * @dev Returns the amount of tokens owned by `account`.
     */
    function balanceOf(address account) external view returns (uint256);

    /**
     * @dev Moves `amount` tokens from the caller's account to `recipient`.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * Emits a {Transfer} event.
     */
    function transfer(address recipient, uint256 amount) external returns (bool);

    /**
     * @dev Returns the remaining number of tokens that `spender` will be
     * allowed to spend on behalf of `owner` through {transferFrom}. This is
     * zero by default.
     *
     * This value changes when {approve} or {transferFrom} are called.
     */
    function allowance(address owner, address spender) external view returns (uint256);

    /**
     * @dev Sets `amount` as the allowance of `spender` over the caller's tokens.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * IMPORTANT: Beware that changing an allowance with this method brings the risk
     * that someone may use both the old and the new allowance by unfortunate
     * transaction ordering. One possible solution to mitigate this race
     * condition is to first reduce the spender's allowance to 0 and set the
     * desired value afterwards:
     * https://github.com/ethereum/EIPs/issues/20#issuecomment-263524729
     *
     * Emits an {Approval} event.
     */
    function approve(address spender, uint256 amount) external returns (bool);

    /**
     * @dev Moves `amount` tokens from `sender` to `recipient` using the
     * allowance mechanism. `amount` is then deducted from the caller's
     * allowance.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * Emits a {Transfer} event.
     */
    function transferFrom(
        address sender,
        address recipient,
        uint256 amount
    ) external returns (bool);

    /**
     * @dev Emitted when `value` tokens are moved from one account (`from`) to
     * another (`to`).
     *
     * Note that `value` may be zero.
     */
    event Transfer(address indexed from, address indexed to, uint256 value);

    /**
     * @dev Emitted when the allowance of a `spender` for an `owner` is set by
     * a call to {approve}. `value` is the new allowance.
     */
    event Approval(address indexed owner, address indexed spender, uint256 value);
}


interface ICAT {
    function mint(address _to) external returns (uint256);

    function upgrade(uint256 _tokenId) external;

    function totalSupply() external view returns (uint256);

    function level(uint256 _tokenId) external view returns (uint256);

    function setMINTER(address _minter) external;

    function catOwner() external view returns (address);

}

contract Permisson {
    mapping(address => bool) public operators;


    modifier onlyOperator(){
        require(operators[msg.sender], "Only operator");
        _;
    }
}


contract BCatMinter is Initializable, OwnableUpgradeable, UUPSUpgradeable, Permisson {

    ICAT public icat;

    struct Token {
        IERC20 erc20Address;
        uint256 amount;
    }

    uint256 public freeMintCount;
    uint256 public freeMintProcessCount;

    mapping(uint256 => Token[]) public upgradeCosts;

    IERC20[] public receiveTokens;

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() initializer {}

    function initialize(uint256 _freeMintCount, Token[][] memory _upgradeCosts, address _upgrader, address _icat) public initializer {
        __Ownable_init();
        __UUPSUpgradeable_init();
        transferOwnership(_upgrader);
        operators[msg.sender] = true;
        operators[_upgrader] = true;
        freeMintCount = _freeMintCount;
        for (uint i = 0; i < _upgradeCosts.length; i++) {
            _pushTokens(upgradeCosts[i + 1], _upgradeCosts[i]);
        }
        icat = ICAT(_icat);
        icat.setMINTER(address(this));
    }

    function _authorizeUpgrade(address newImplementation)
    internal
    onlyOwner
    override
    {}



    modifier onlyOwners(){
        require(icat.catOwner() == msg.sender || owner() == msg.sender, "Only owner");
        _;
    }

    function addOperator(address _operator) public onlyOwners {
        operators[_operator] = true;
    }

    function delOperator(address _operator) public onlyOwners {
        operators[_operator] = false;
    }

    function changeIcat(address _cat) public onlyOwners {
        icat = ICAT(_cat);
    }


    function setUpgradeCost(uint256 _level, Token[] memory _cost) public onlyOperator {
        delete upgradeCosts[_level];
        _pushTokens(upgradeCosts[_level], _cost);
    }

    function setFreeMintCount(uint256 _count) public onlyOperator {
        freeMintCount = _count;
    }


    function _pushTokens(Token[] storage _self, Token[] memory _cost) internal {
        for (uint i = 0; i < _cost.length; i++) {
            _self.push(Token(_cost[i].erc20Address, _cost[i].amount));
        }
    }


    function getUpgradeCost(uint256 _level) public view returns (Token[] memory){
        return upgradeCosts[_level];
    }

    function freeMint(address _to) public onlyOperator returns (uint256){
        require(freeMintProcessCount < freeMintCount, "Exceeded");
        freeMintProcessCount += 1;
        return icat.mint(_to);
    }

    function publicMint() public payable returns (uint256){
        _receiveToken(upgradeCosts[1]);
        return icat.mint(msg.sender);
    }

    function upgrade(uint256 _tokenId) public payable {
        _receiveToken(upgradeCosts[icat.level(_tokenId) + 1]);
        return icat.upgrade(_tokenId);
    }


    function _receiveToken(Token[] memory _costs) public payable {
        for (uint i = 0; i < _costs.length; i++) {
            if (address(_costs[i].erc20Address) == address(0)) {
                if (_costs[i].amount > 0) {
                    require(msg.value >= _costs[i].amount, "Insufficient");
                }
            } else {
                _costs[i].erc20Address.transferFrom(msg.sender, address(this), _costs[i].amount);
                bool hasToken = false;
                for (uint j = 0; j < receiveTokens.length; j++) {
                    if (receiveTokens[j] == _costs[i].erc20Address) {
                        hasToken = true;
                        break;
                    }
                }
                if (!hasToken) {
                    receiveTokens.push(_costs[i].erc20Address);
                }
            }
        }
    }


    function withdraw(address payable _to, IERC20 _token, uint256 _amount) public onlyOperator {
        if (address(_token) == address(0)) {
            require(address(this).balance >= _amount, "Insufficient balance");
            _to.transfer(_amount);
        } else {
            require(_token.balanceOf(address(this)) >= _amount, "Insufficient balance");
            _token.transfer(_to, _amount);
        }
    }


}