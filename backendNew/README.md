# BackendNew - Go语言博客后端服务

## 项目概述

BackendNew是基于Go语言和Gin框架开发的博客系统后端服务，采用现代化的微服务架构设计，提供完整的博客管理功能。

## 技术栈

- **编程语言**: Go 1.24+
- **Web框架**: Gin
- **数据库**: MySQL + GORM
- **认证**: JWT (JSON Web Token)
- **容器化**: Docker

## 项目结构

```
backendNew/
├── cmd/                    # 应用程序入口
│   ├── server/            # 主服务入口
│   └── init-admin/        # 管理员初始化工具
├── internal/              # 内部包（不对外暴露）
│   ├── db/               # 数据库连接和初始化
│   ├── model/            # 数据模型定义
│   ├── repo/             # 数据访问层
│   ├── handler/          # HTTP请求处理器
│   ├── middleware/       # 中间件
│   └── util/            # 工具函数
├── go.mod                # Go模块定义
├── go.sum               # 依赖校验
└── Dockerfile           # Docker容器配置
```

## 核心功能

### 1. 博客管理
- 获取所有博客列表
- 根据ID获取博客详情
- 创建新博客（需要管理员权限）
- 更新博客内容（需要管理员权限）
- 删除博客（需要管理员权限）

### 2. 用户认证
- 管理员登录
- JWT Token生成和验证
- 路由权限控制

### 3. 数据管理
- 自动数据库迁移
- 默认管理员账户初始化
- 数据持久化操作

## API接口文档

### 公开接口（无需认证）

#### 获取所有博客
```http
GET /api/blog/all
```

**响应示例：**
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "title": "博客标题",
      "content": "博客内容",
      "uploadDate": "2024-01-01T00:00:00Z",
      "updateDate": "2024-01-01T00:00:00Z"
    }
  ]
}
```

#### 获取博客详情
```http
GET /api/blog/:id
```

**参数：**
- `id`: 博客ID

### 管理员接口（需要JWT认证）

#### 管理员登录
```http
POST /api/admin/login
```

**请求体：**
```json
{
  "username": "admin",
  "password": "admin123"
}
```

**响应示例：**
```json
{
  "success": true,
  "message": "登录成功",
  "data": {
    "token": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "username": "admin"
  }
}
```

#### 创建博客
```http
POST /api/blog/add
Authorization: Bearer <token>
```

**请求体：**
```json
{
  "title": "新博客标题",
  "content": "博客内容"
}
```

#### 更新博客
```http
POST /api/blog/update/:id
Authorization: Bearer <token>
```

#### 删除博客
```http
DELETE /api/blog/delete/:id
Authorization: Bearer <token>
```

### 系统接口

#### 健康检查
```http
GET /api/health
```

**响应：**
```json
{
  "ok": true
}
```

#### 根路径检查
```http
GET /
```

**响应：**
```
backendNew OK
```

## 数据模型

### Blog（博客）
```go
type Blog struct {
    ID         uint      `gorm:"primaryKey" json:"id"`
    Title      string    `json:"title"`
    Content    string    `gorm:"type:longtext" json:"content"`
    UploadDate time.Time `gorm:"column:upload_date" json:"uploadDate"`
    UpdateDate time.Time `gorm:"column:update_date" json:"updateDate"`
}
```

### Admin（管理员）
```go
type Admin struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    Username  string    `gorm:"uniqueIndex;not null;size:50" json:"username"`
    Password  string    `gorm:"not null;size:255" json:"-"`
    CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt"`
}
```

## 环境配置

### 环境变量

| 变量名 | 默认值 | 说明 |
|--------|--------|------|
| `APP_PORT` | `8081` | 服务端口 |
| `DB_DSN` | `root:123456@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local&allowNativePasswords=true` | 数据库连接字符串 |
| `JWT_SECRET` | `WfP@2025!Blog#Secure*Key&Salt123456789` | JWT密钥 |

### 数据库配置

项目使用MySQL数据库，支持以下配置方式：
1. 通过环境变量 `DB_DSN` 设置
2. 使用默认的本地MySQL配置

## 快速开始

### 1. 环境准备

确保已安装：
- Go 1.24+
- MySQL 8.0+

### 2. 克隆项目
```bash
git clone <repository-url>
cd backendNew
```

### 3. 安装依赖
```bash
go mod tidy
```

### 4. 配置数据库

创建数据库并配置连接：
```sql
CREATE DATABASE blog CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

### 5. 运行服务
```bash
# 设置环境变量（可选）
export DB_DSN="username:password@tcp(localhost:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
export JWT_SECRET="your-secret-key"

# 启动服务
go run cmd/server/main.go
```

### 6. 初始化管理员（可选）
```bash
go run cmd/init-admin/main.go
```

## Docker部署

### 构建镜像
```bash
docker build -t folium-backend-new .
```

### 运行容器
```bash
docker run -d \
  --name folium-backend \
  -p 8081:8081 \
  -e DB_DSN="mysql-dsn" \
  -e JWT_SECRET="your-secret" \
  folium-backend-new
```

## 开发指南

### 代码结构说明

- **cmd/**: 应用程序入口点
- **internal/db/**: 数据库连接和初始化逻辑
- **internal/model/**: 数据模型定义
- **internal/repo/**: 数据访问层，封装数据库操作
- **internal/handler/**: HTTP请求处理器
- **internal/middleware/**: 中间件（认证、CORS等）
- **internal/util/**: 工具函数（JWT处理等）

### 添加新功能

1. 在 `model/` 中定义数据模型
2. 在 `repo/` 中实现数据访问方法
3. 在 `handler/` 中实现HTTP处理器
4. 在 `main.go` 中注册路由

### 测试

运行测试：
```bash
go test ./...
```

## 安全特性

- JWT Token认证
- CORS跨域支持
- 密码不返回给客户端
- Token过期时间控制（24小时）
- 输入参数验证

## 故障排除

### 常见问题

1. **数据库连接失败**
   - 检查MySQL服务是否运行
   - 验证数据库连接字符串
   - 确认数据库用户权限

2. **JWT Token无效**
   - 检查Token格式是否正确
   - 验证Token是否过期
   - 确认JWT密钥配置

3. **端口被占用**
   - 更改 `APP_PORT` 环境变量
   - 检查是否有其他服务占用8081端口

### 日志查看

服务启动时会输出详细的日志信息，包括：
- 数据库连接状态
- 服务启动信息
- 请求处理日志

## 贡献指南

1. Fork项目
2. 创建功能分支
3. 提交更改
4. 推送到分支
5. 创建Pull Request

## 许可证

本项目采用MIT许可证。

## 联系方式

如有问题或建议，请联系项目维护者。