# Ethereum Contract Service

一个基于 [Kratos](https://github.com/go-kratos/kratos) 框架的以太坊合约调用服务，提供常见以太坊合约（如 ERC20、ERC721 等）的 Go 语言调用接口。

## 项目定位

本项目专注于提供常见以太坊智能合约的 Go 语言调用服务，通过 HTTP RESTful API 和 gRPC 接口，方便开发者快速集成以太坊合约功能。

## 功能特性

- ✅ **HTTP/gRPC 双协议支持**：同时支持 HTTP RESTful API 和 gRPC
- ✅ **ERC20 合约支持**：完整的 ERC20 代币操作（查询余额、转账、授权、铸造、销毁等）
- ✅ **合约部署**：支持部署新的 ERC20 合约
- ✅ **多链支持**：支持主网、测试网和本地开发链
- ✅ **Nginx 反向代理**：通过 Nginx 处理 CORS、SSL 等网络层问题
- ✅ **结构化日志**：基于 zap 的日志系统
- ✅ **健康检查**：内置健康检查端点
- ✅ **配置管理**：支持环境变量覆盖

## 支持的合约类型

### ERC20 代币合约

- ✅ 查询代币信息（名称、符号、精度、总供应量）
- ✅ 查询余额
- ✅ 转账
- ✅ 授权和查询授权额度
- ✅ 授权转账（transferFrom）
- ✅ 铸造代币（mint）
- ✅ 销毁代币（burn、burnFrom）
- ✅ 部署新合约

### 未来计划

- 🔲 ERC721 NFT 合约支持
- 🔲 ERC1155 多代币标准支持
- 🔲 其他常见合约标准

## 项目结构

```
.
├── api/                    # API 定义（protobuf）
│   └── erc20/v1/          # ERC20 API 定义
├── cmd/                    # 应用入口
│   └── app/               # 主程序
├── configs/               # 配置文件
├── internal/              # 内部代码
│   ├── conf/             # 配置定义
│   ├── global/           # 全局变量
│   ├── server/           # 服务器初始化
│   └── service/          # 业务服务
├── provider/             # 基础设施提供者
│   ├── contract/         # 合约绑定（Go bindings）
│   │   └── erc20/       # ERC20 合约绑定
│   ├── eth/              # 以太坊客户端
│   ├── cache/            # Redis 缓存
│   ├── db/               # 数据库
│   └── logger/           # 日志
├── nginx/                 # Nginx 配置
└── third_party/          # 第三方 proto 文件
```

## 快速开始

### 1. 安装依赖工具

```bash
make init
```

### 2. 生成代码

```bash
make all
```

### 3. 配置环境变量

创建 `.env` 文件：

```bash
./scripts/setup-env.sh
```

或手动创建 `.env` 文件，配置以下关键项：

```bash
# 以太坊 RPC 节点
ETH_RPC_URL=http://localhost:8545  # 或使用 Infura/Alchemy
ETH_CHAIN_ID=1337  # 1=主网, 5=Goerli, 11155111=Sepolia

# 数据库配置
DB_NAME=eth_contract_service
DB_PASSWORD=your_password
```

### 4. 启动服务

使用 Docker Compose（推荐）：

```bash
docker-compose up -d
```

或本地运行：

```bash
make build
./bin/app -conf configs
```

## API 端点

### ERC20 接口

#### 查询接口

- `GET /api/v1/erc20/info?contract_address=0x...` - 查询代币信息
- `GET /api/v1/erc20/balance?contract_address=0x...&address=0x...` - 查询余额
- `GET /api/v1/erc20/allowance?contract_address=0x...&owner=0x...&spender=0x...` - 查询授权额度

#### 交易接口

- `POST /api/v1/erc20/transfer` - 转账
- `POST /api/v1/erc20/approve` - 授权
- `POST /api/v1/erc20/transfer-from` - 授权转账
- `POST /api/v1/erc20/mint` - 铸造代币
- `POST /api/v1/erc20/burn` - 销毁代币
- `POST /api/v1/erc20/burn-from` - 从指定地址销毁代币
- `POST /api/v1/erc20/deploy` - 部署新合约

#### 健康检查

- `GET /health` - 健康检查端点

详细的 API 文档请参考生成的 `openapi.yaml` 文件。

## 配置说明

### 以太坊配置

在 `configs/config.yaml` 或环境变量中配置：

```yaml
ethereum:
  rpc_url: http://localhost:8545  # 以太坊 RPC 节点
  chain_id: 1337                  # 链 ID
  timeout: 30s
  max_retries: 3
  contracts:
    erc20: 0x...  # ERC20 合约地址（可选，可通过 API 动态指定）
```

### 环境变量

所有配置项都支持通过环境变量覆盖：

- `ETH_RPC_URL` - 以太坊 RPC 节点地址
- `ETH_CHAIN_ID` - 链 ID
- `ERC20_CONTRACT_ADDRESS` - 默认 ERC20 合约地址
- `SERVER_HTTP_ADDR` - HTTP 服务地址
- `SERVER_GRPC_ADDR` - gRPC 服务地址
- `DB_NAME` - 数据库名称
- `DB_PASSWORD` - 数据库密码

详细配置说明请参考 [ENV_SETUP.md](./ENV_SETUP.md)。

## 开发指南

### 添加新的合约类型

1. 在 `provider/contract/` 目录下添加合约 Solidity 文件
2. 编译合约生成 Go bindings（使用 `solc` 和 `abigen`）
3. 在 `api/` 目录下定义新的 proto 文件
4. 运行 `make api` 生成代码
5. 在 `internal/service/` 中实现服务逻辑
6. 在 `internal/server/` 中注册服务

### 合约编译

合约编译已集成到 Makefile 中，参考 `provider/contract/erc20/` 目录下的示例。

## 网络架构

```
客户端
  ↓
Nginx (端口 80/9090) - 处理 CORS、SSL、限流
  ↓
后端服务 (端口 8000/9000) - 业务逻辑
  ↓
以太坊节点 (RPC)
```

## 构建和部署

### 本地构建

```bash
make build
```

### Docker 构建

```bash
docker build -t eth-contract-service .
```

### 生产环境

1. 使用 `nginx/nginx.conf.production` 配置
2. 配置 SSL 证书
3. 限制 CORS 为特定域名
4. 启用限流和监控
5. 使用强密码和安全配置

## 许可证

MIT

## 贡献

欢迎提交 Issue 和 Pull Request！
