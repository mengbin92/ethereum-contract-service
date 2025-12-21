# Nginx 配置说明

本目录包含用于反向代理和 CORS 处理的 Nginx 配置文件。

## 功能

1. **HTTP API 反向代理** - 将 HTTP 请求代理到后端服务（端口 8000）
2. **gRPC 反向代理** - 将 gRPC 请求代理到后端服务（端口 9000）
3. **CORS 处理** - 在 Nginx 层面处理跨域请求，后端服务无需关心
4. **负载均衡** - 支持多个后端实例（通过 upstream 配置）

## 端口映射

- **80** - HTTP API（带 CORS，开发环境）
- **443** - HTTPS API（生产环境，需要 SSL 证书）
- **9090** - gRPC（HTTP/2）

## 配置文件说明

- `nginx.conf` - 开发环境配置（允许所有来源的 CORS）
- `nginx.conf.production` - 生产环境配置（HTTPS + 限流 + 安全域名）

## CORS 配置

### 开发环境（nginx.conf）

当前配置允许所有来源（`*`），适合本地开发：

```nginx
set $cors_origin '*';
```

### 生产环境（nginx.conf.production）

限制为特定域名，更安全：

```nginx
set $cors_origin 'https://yourdomain.com';
if ($http_origin = 'https://yourdomain.com') {
    set $cors_origin $http_origin;
}
```

## 使用方式

### 开发环境（Docker Compose）

```bash
# 启动所有服务（包括 Nginx）
docker-compose up -d

# 仅启动 Nginx
docker-compose up -d nginx

# 查看 Nginx 日志
docker-compose logs -f nginx
```

### 访问服务

- HTTP API: `http://localhost/api/v1/erc20/...`
- gRPC: `localhost:9090`

### 生产环境部署

1. **使用生产配置**：
   ```bash
   cp nginx/nginx.conf.production nginx/nginx.conf
   ```

2. **配置 SSL 证书**：
   - 将证书文件放到 `nginx/ssl/cert.pem`
   - 将私钥文件放到 `nginx/ssl/key.pem`

3. **修改域名**：
   - 在 `nginx.conf` 中替换 `yourdomain.com` 为实际域名

4. **测试配置**：
   ```bash
   docker-compose exec nginx nginx -t
   ```

5. **重启服务**：
   ```bash
   docker-compose restart nginx
   ```

## 架构说明

```
客户端
  ↓
Nginx (端口 80/443) - 处理 CORS、SSL、限流
  ↓
后端服务 (端口 8000/9000) - 专注业务逻辑
```

## 优势

1. **关注点分离**：后端服务专注于业务逻辑，不处理网络层问题
2. **性能优化**：Nginx 可以处理静态文件、压缩、缓存等
3. **安全性**：在 Nginx 层面统一处理安全策略
4. **可扩展性**：可以轻松添加多个后端实例进行负载均衡

## 安全建议

1. **生产环境**：
   - 使用 `nginx.conf.production` 配置
   - 配置 SSL 证书（Let's Encrypt 免费证书）
   - 限制 CORS 为特定域名
   - 启用 rate limiting

2. **限流配置**：
   - API: 10 请求/秒，突发 20
   - gRPC: 50 请求/秒，突发 100

3. **监控**：
   - 配置日志收集
   - 监控 Nginx 性能指标

## 故障排查

### 检查 Nginx 配置

```bash
docker-compose exec nginx nginx -t
```

### 查看 Nginx 日志

```bash
docker-compose logs nginx
```

### 测试后端连接

```bash
# 测试 HTTP API
curl http://localhost/api/v1/erc20/info?contract_address=0x...

# 测试健康检查
curl http://localhost/health
```
