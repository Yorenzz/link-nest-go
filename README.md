# LinkNest

LinkNest 是一个现代化的个人主页和链接管理平台，允许用户创建个性化的主页，展示链接、社交媒体、项目等内容。

## 功能特性

- 🔐 **用户认证系统** - 支持用户名/邮箱注册/登录，JWT认证
- 🎨 **个性化主页** - 创建和自定义个人主页
- 🧩 **模块化设计** - 支持多种模块类型（链接、GitHub、Spotify、RSS等）
- 📱 **响应式设计** - 适配各种设备和屏幕尺寸
- 🎯 **模板系统** - 官方和用户自定义模板
- 💳 **付费功能** - Pro用户专属功能和支付系统
- 📊 **数据分析** - 页面访问统计和模块点击分析（规划中）
- 🌐 **自定义域名** - 支持绑定个人域名（规划中）

## 技术栈

- **后端**: Go 1.24.2
- **Web框架**: Gin
- **数据库**: PostgreSQL
- **ORM**: GORM
- **认证**: JWT (golang-jwt/jwt/v5)
- **配置管理**: Viper
- **数据类型**: GORM Datatypes (JSON支持)

## 项目结构

```
link-nest/
├── cmd/
│   └── server/          # 应用程序入口
├── configs/             # 配置文件
│   ├── config.yaml      # 应用配置
│   ├── config.go        # 配置结构定义
│   └── database_schema_full.sql  # 数据库架构
├── internal/
│   ├── api/             # HTTP处理器
│   ├── auth/            # 认证中间件和JWT处理
│   ├── database/        # 数据库连接和迁移
│   ├── models/          # 数据模型
│   ├── repository/      # 数据访问层
│   └── service/         # 业务逻辑层
├── go.mod
└── go.sum
```

## 数据模型

### 核心实体

- **Users** - 用户信息管理
- **Templates** - 页面模板系统
- **UserPages** - 用户个人主页
- **Modules** - 页面模块组件
- **Payments** - 支付和订阅记录

### 支持的模块类型

- 📎 **Link** - 链接模块
- 📝 **Text** - 文本内容
- 🖼️ **Image** - 图片展示
- 🐙 **GitHub** - GitHub贡献和仓库 (Pro)
- 🎵 **Spotify** - 音乐播放状态 (Pro)
- 📡 **RSS** - RSS订阅内容 (Pro)

## 快速开始

### 环境要求

- Go 1.24.2+
- PostgreSQL 12+

### 安装步骤

1. **克隆项目**
   ```bash
   git clone <repository-url>
   cd link-nest
   ```

2. **安装依赖**
   ```bash
   go mod download
   ```

3. **配置数据库**
   - 创建PostgreSQL数据库
   - 应用程序启动时会自动执行数据库迁移，无需手动执行SQL文件

4. **配置应用**
   ```yaml
   # configs/config.yaml
   server:
     port: ":8080"
   database:
     host: "localhost"
     port: "5432"
     user: "your_username"
     password: "your_password"
     dbname: "link_nest"
     sslmode: "disable"
   ```

5. **启动服务**
   ```bash
   go run cmd/server/main.go
   ```

## API 接口

### 公开接口

- `POST /api/v1/register` - 用户注册
  ```json
  {
    "username": "testuser",
    "email": "test@example.com",
    "password": "password123"
  }
  ```

- `POST /api/v1/login` - 用户登录
  ```json
  {
    "email": "test@example.com",
    "password": "password123"
  }
  ```

### 认证接口

需要在请求头中包含JWT token：`Authorization: Bearer <token>`

- 模板管理API（规划中）
- 用户主页管理API（规划中）
- 模块管理API（规划中）
- 支付管理API（规划中）

## 开发指南

### 添加新模块类型

1. 在 `internal/models/` 中定义模块配置结构
2. 在 `internal/service/` 中实现业务逻辑
3. 在 `internal/api/` 中添加HTTP处理器
4. 更新数据库schema中的模块类型

### 代码规范

- 使用Go标准代码格式 (`go fmt`)
- 遵循Go命名约定
- 为公开函数添加注释
- 使用依赖注入模式

## 部署

### Docker部署 (推荐)

```dockerfile
# Dockerfile示例
FROM golang:1.24.2-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main cmd/server/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/configs ./configs
CMD ["./main"]
```

### 环境变量

应用程序主要通过 `configs/config.yaml` 文件进行配置，也支持环境变量覆盖：

- `SERVER_PORT` - 服务器端口（默认：8080）
- `DB_HOST` - 数据库主机
- `DB_PORT` - 数据库端口
- `DB_USER` - 数据库用户名
- `DB_PASSWORD` - 数据库密码
- `DB_NAME` - 数据库名称

## 数据库结构

应用程序使用GORM自动迁移功能，启动时会自动创建以下表：

- `users` - 用户信息表
- `templates` - 模板表
- `user_pages` - 用户主页表
- `modules` - 页面模块表
- `payments` - 支付记录表

详细的数据库结构可以参考 `configs/database_schema_full.sql` 文件。

## 测试API

启动服务器后，可以使用以下命令测试API：

### 用户注册
```bash
curl -X POST http://localhost:8080/api/v1/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "email": "test@example.com",
    "password": "password123"
  }'
```

### 用户登录
```bash
curl -X POST http://localhost:8080/api/v1/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "password": "password123"
  }'
```

## 开发状态

### 已完成功能
- ✅ 用户认证系统（注册/登录）
- ✅ JWT token生成和验证
- ✅ 数据库模型和迁移
- ✅ 基础API架构

### 开发中功能
- 🚧 模板管理API
- 🚧 用户主页管理API
- 🚧 模块管理API
- 🚧 支付系统API

### 计划功能
- 📋 前端界面
- 📋 页面访问统计
- 📋 自定义域名支持
- 📋 第三方集成（GitHub、Spotify等）

## 贡献

欢迎提交Issue和Pull Request来改进项目。

## 许可证

[MIT License](LICENSE)
