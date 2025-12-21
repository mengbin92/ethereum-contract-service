#!/bin/bash

# 创建 .env 文件的脚本
# 从 .env.example 复制或直接生成 .env

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"
ENV_EXAMPLE="$PROJECT_ROOT/.env.example"
ENV_FILE="$PROJECT_ROOT/.env"

if [ -f "$ENV_FILE" ]; then
    echo "⚠️  .env 文件已存在，跳过创建"
    echo "   如需重新创建，请先删除: rm $ENV_FILE"
    exit 0
fi

# 如果 .env.example 存在，则复制它
if [ -f "$ENV_EXAMPLE" ]; then
    cp "$ENV_EXAMPLE" "$ENV_FILE"
    echo "✅ 已从 .env.example 创建 .env 文件: $ENV_FILE"
else
    # 如果 .env.example 不存在，直接生成 .env 文件
    cat > "$ENV_FILE" << 'EOF'
# Server Configuration
HTTP_PORT=8000
GRPC_PORT=9000

# Database Configuration
DB_USER=root
DB_PASSWORD=jKBrZHGcsNG5fMc52EWz
DB_HOST=mysql
DB_PORT=3306
DB_NAME=eth_contract_service

# Redis Configuration
REDIS_HOST=redis
REDIS_PORT=6379
REDIS_PASSWORD=
REDIS_DB=0

# Ethereum Configuration
ETH_RPC_URL=http://localhost:8545
ETH_CHAIN_ID=1337
# 1 for mainnet, 5 for goerli, 11155111 for sepolia

# ERC20 Contract Address
ERC20_CONTRACT_ADDRESS=0x0000000000000000000000000000000000000000
EOF
    echo "✅ 已创建 .env 文件: $ENV_FILE"
fi

echo "📝 请根据实际情况修改 .env 文件中的配置"
echo ""
echo "主要配置项："
echo "  - ETH_RPC_URL: 以太坊 RPC 节点地址"
echo "  - ETH_CHAIN_ID: 链 ID (1=主网, 5=Goerli, 11155111=Sepolia)"
echo "  - ERC20_CONTRACT_ADDRESS: ERC20 合约地址"
echo "  - DB_PASSWORD: 数据库密码"
echo ""

