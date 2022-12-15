// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC721/extensions/ERC721BurnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC721/extensions/ERC721EnumerableUpgradeable.sol";


interface ITokenuri {
    function tokenURI(uint256 _tokenId, string memory _app, uint256 _level) external view returns (string memory);
}


contract Permisson{

    address public MINTER;
    address public catOwner;


    modifier onlyMinter(){
        require(MINTER==msg.sender || catOwner==msg.sender, "Only Minter");
        _;
    }
}


contract BCAT is Initializable, ERC721Upgradeable, ERC721EnumerableUpgradeable, OwnableUpgradeable,
UUPSUpgradeable,Permisson {
    ITokenuri public tokenUri;
    uint256 tokenId;
    string public app;
    uint256 public maxLevel;
    uint256 public maxSupply;
    bool public canTransfer;

    mapping(uint256 => uint256) public levelCount;
    mapping(uint256 => uint256) public level;

    mapping(address => bool) public minted;



    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    function initialize(string memory _name, string memory _symbol, address _tokenUri, string memory _app,
        uint256 _maxSupply, uint256 _maxLevel, bool _canTransfer,
        address _upgrader) initializer public {
        __ERC721_init(_name, _symbol);
        __ERC721Enumerable_init();
        __Ownable_init();
        __UUPSUpgradeable_init();
        transferOwnership(_upgrader);
        tokenUri = ITokenuri(_tokenUri);
        catOwner = msg.sender;
        app = _app;
        if (_maxLevel > 5) {
            _maxLevel = 5;
        }
        maxLevel = _maxLevel;
        canTransfer = _canTransfer;
        maxSupply = _maxSupply;
    }


    function _authorizeUpgrade(address newImplementation)
    internal
    onlyOwner
    override
    {}

    modifier permitted(){
        require(owner()==msg.sender || catOwner==msg.sender, "Not permit");
        _;
    }

    function setMINTER(address _minter)public{
        if(MINTER==address(0)){
            MINTER=_minter;
        }else{
            require(owner()==msg.sender || catOwner==msg.sender,"only owner");
            MINTER=_minter;
        }
    }


    // The following functions are overrides required by Solidity.

    function _beforeTokenTransfer(address from, address to, uint256 _tokenId, uint256 _batchSize)
    internal
    override(ERC721Upgradeable, ERC721EnumerableUpgradeable)
    {
        if (MINTER!=msg.sender || catOwner!=msg.sender) {
            require(canTransfer, "Cannot transfer");
        }

        super._beforeTokenTransfer(from, to, _tokenId, _batchSize);
    }


    function supportsInterface(bytes4 interfaceId)
    public
    view
    override(ERC721Upgradeable, ERC721EnumerableUpgradeable)
    returns (bool)
    {
        return super.supportsInterface(interfaceId);
    }


    function changeTokenUri(address _tokenuri) public permitted {
        tokenUri = ITokenuri(_tokenuri);
    }

    function switchCanTransfer() public permitted {
        canTransfer = !canTransfer;
    }

    function setMaxSupply(uint256 _maxSupply) public permitted {
        maxSupply = _maxSupply;
    }

    function setMaxLevel(uint256 _maxLevel) public permitted {
        require(_maxLevel >= maxLevel, "Not support");
        maxLevel = _maxLevel;
    }


    function upgrade(uint256 _tokenId) public onlyMinter {
        safeUpgrade(_tokenId);
    }


    function mint(address _to) public onlyMinter returns (uint256){
        return safeMint(_to);
    }

    function safeUpgrade(uint256 _tokenId) internal {
        require(level[_tokenId] > 0, "Nonexist");
        require(level[_tokenId] < maxLevel, "Max level");
        levelCount[level[_tokenId]] -= 1;
        level[_tokenId] += 1;
        levelCount[level[_tokenId]] += 1;
    }

    function safeMint(address _to) internal returns (uint256){
        require(!minted[_to], "One chance");
        if (maxSupply > 0) {
            require(totalSupply() < maxSupply, "Exceeded");
        }
        tokenId += 1;
        _safeMint(_to, tokenId);
        level[tokenId] = 1;
        levelCount[1] += 1;
        minted[_to] = true;
        return tokenId;
    }


    function tokenURI(uint256 _tokenId)
    public
    view
    override
    returns (string memory)
    {
        return tokenUri.tokenURI(_tokenId, app, level[_tokenId]);
    }
}