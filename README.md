# Folium-茯苓的博客园

## 项目简介

Folium博客是一个基于Go语言和Vue.js构建的前后端分离个人博客系统。

## 功能特性

- 博客列表展示
- 博客详情查看（支持Markdown渲染）
- 管理员登录认证（JWT）
- 博客增删改操作
- 响应式界面设计

## 技术栈

### 前端
- Vue 3
- Vite
- Vue Router
- Axios
- Marked

### 后端
- Go 1.24
- Gin框架
- GORM
- MySQL
- JWT认证

### 基础设施
- Docker
- Docker Compose
- Nginx

## 快速开始

### 环境要求
- Docker
- Docker Compose

### 部署步骤

1. 克隆项目
```bash
git clone https://github.com/你的用户名/Folium-blog.git
cd Folium-blog
```

2. 启动服务
```bash
docker compose up --build -d
```

3. 等待服务启动完成

4. 访问应用


### 服务端口
- 前端：5713端口
- 后端：8081端口

### 停止服务
```bash
docker compose down
```

## 项目结构

```
Folium-blog/
├── backend/blog/          # Go后端
├── frontend/             # Vue前端
├── docker-compose.yml
└── README.md
```

## 开发指南

### 前端开发
```bash
cd frontend
npm install
npm run dev
```

### 后端开发
```bash
cd backend/blog
# 安装依赖
go mod tidy
# 启动服务（使用热重载）
fresh
# 或直接运行
go run cmd/server/main.go
```

### 数据库
项目使用MySQL数据库，连接信息：
- 主机：db
- 端口：3306
- 数据库：folium
- 用户名：folium
- 密码：folium

## 注意事项

- 首次部署需要等待镜像和依赖下载完成
- 建议配置国内镜像源加速下载