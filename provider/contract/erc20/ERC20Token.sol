// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

/**
 * @title ERC20Token
 * @dev 基于 OpenZeppelin 实现的 ERC20 代币合约，支持自定义代币名称和符号
 * @notice 部署时可以指定代币名称、符号、小数位数和初始供应量
 */
contract ERC20Token is ERC20 {
    /**
     * @dev 代币的小数位数
     */
    uint8 private _decimals;

    /**
     * @dev 构造函数，初始化代币
     * @param name_ 代币名称（例如："My Token"）
     * @param symbol_ 代币符号（例如："MTK"）
     * @param decimals_ 代币小数位数（通常为 18）
     * @param initialSupply_ 初始供应量（以最小单位计算，例如：1000000 * 10^18）
     * @param owner_ 初始代币接收地址（通常是部署者地址）
     */
    constructor(
        string memory name_,
        string memory symbol_,
        uint8 decimals_,
        uint256 initialSupply_,
        address owner_
    ) ERC20(name_, symbol_) {
        require(
            owner_ != address(0),
            "ERC20Token: owner cannot be zero address"
        );
        require(
            decimals_ > 0 && decimals_ <= 18,
            "ERC20Token: decimals must be between 1 and 18"
        );

        _decimals = decimals_;

        // 如果初始供应量大于 0，则铸造代币给指定地址
        if (initialSupply_ > 0) {
            _mint(owner_, initialSupply_);
        }
    }

    /**
     * @dev 返回代币的小数位数
     * @return 代币的小数位数
     */
    function decimals() public view virtual override returns (uint8) {
        return _decimals;
    }

    /**
     * @dev 铸造新代币（仅合约所有者可以调用）
     * @param to 接收代币的地址
     * @param amount 铸造的代币数量
     */
    function mint(address to, uint256 amount) public {
        // 注意：这里需要根据你的需求添加访问控制
        // 例如使用 OpenZeppelin 的 Ownable 或 AccessControl
        _mint(to, amount);
    }

    /**
     * @dev 销毁代币
     * @param amount 销毁的代币数量
     */
    function burn(uint256 amount) public {
        _burn(msg.sender, amount);
    }

    /**
     * @dev 从指定地址销毁代币（需要先授权）
     * @param from 代币来源地址
     * @param amount 销毁的代币数量
     */
    function burnFrom(address from, uint256 amount) public {
        _spendAllowance(from, msg.sender, amount);
        _burn(from, amount);
    }
}
