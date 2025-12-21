# 环境变量配置说明

## 创建 .env 文件

项目使用 `.env` 文件来管理环境变量。如果 `.env` 文件不存在，请按以下步骤创建：

### 方法一：使用脚本（推荐）

```bash
# 运行设置脚本（会自动创建 .env 文件）
./scripts/setup-env.sh
```

**注意**：脚本会检查 `.env.example` 是否存在，如果存在则复制它，否则会直接生成默认的 `.env` 文件。

### 方法二：手动创建

```bash
# 复制示例文件
cp .env.example .env

# 或者手动创建
touch .env
```

### 方法三：使用命令行

```bash
cat > .env << 'EOF'
# Server Configuration
HTTP_PORT=8000
GRPC_PORT=9000

# Database Configuration
DB_USER=root
DB_PASSWORD=jKBrZHGcsNG5fMc52EWz
DB_HOST=mysql
DB_PORT=3306
DB_NAME=demo_project

# Redis Configuration
REDIS_HOST=redis
REDIS_PORT=6379
REDIS_PASSWORD=
REDIS_DB=0

# Ethereum Configuration
ETH_RPC_URL=http://localhost:8545
ETH_CHAIN_ID=1337

# ERC20 Contract Address
ERC20_CONTRACT_ADDRESS=0x0000000000000000000000000000000000000000
EOF
```

## 环境变量说明

### 服务器配置

- `HTTP_PORT`: HTTP 服务端口（默认: 8000）
- `GRPC_PORT`: gRPC 服务端口（默认: 9000）

### 数据库配置

- `DB_USER`: 数据库用户名（默认: root）
- `DB_PASSWORD`: 数据库密码
- `DB_HOST`: 数据库主机（Docker 环境使用: mysql）
- `DB_PORT`: 数据库端口（默认: 3306）
- `DB_NAME`: 数据库名称（默认: demo_project）

### Redis 配置

- `REDIS_HOST`: Redis 主机（Docker 环境使用: redis）
- `REDIS_PORT`: Redis 端口（默认: 6379）
- `REDIS_PASSWORD`: Redis 密码（可选）
- `REDIS_DB`: Redis 数据库编号（默认: 0）

### 以太坊配置

- `ETH_RPC_URL`: 以太坊 RPC 节点地址
  - 本地: `http://localhost:8545`
  - Infura: `https://mainnet.infura.io/v3/YOUR_API_KEY`
  - Alchemy: `https://eth-mainnet.g.alchemy.com/v2/YOUR_API_KEY`
  
- `ETH_CHAIN_ID`: 链 ID
  - `1`: 以太坊主网
  - `5`: Goerli 测试网
  - `11155111`: Sepolia 测试网
  - `1337`: 本地开发链（Ganache/Hardhat）

- `ERC20_CONTRACT_ADDRESS`: ERC20 合约地址（部署后更新）

## 注意事项

1. **`.env` 文件已添加到 `.gitignore`**，不会被提交到版本控制
2. **生产环境**：请使用强密码和安全配置
3. **私钥安全**：不要在 `.env` 文件中存储私钥，使用密钥管理服务

## 验证配置

创建 `.env` 文件后，可以验证配置：

```bash
# 检查环境变量
docker-compose config

# 启动服务
docker-compose up -d
```

