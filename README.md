# Wofiporia 极简博客系统

一个前后端分离、功能完整、安全美观的极简个人博客项目。

---

## 功能简介
- 博客首页展示所有文章，支持点击查看详情
- 博客详情页支持 Markdown 渲染，显示创建/更新时间
- 管理员登录（JWT 安全认证，账号密码存数据库）
- 后台支持博客的上传、编辑、删除
- 管理员入口、退出登录、路由守卫、token 自动校验
- 支持“加载中”提示、错误提示，体验流畅
- 首页右上角有管理员入口，左上角有 GitHub 主页链接

## 技术栈
- 前端：Vue 3 + Vite + Vue Router + Axios + Marked
- 后端：Spring Boot 3 + Spring Data JPA + MySQL + JWT (jjwt)

## 启动方式

### 后端
1. 配置数据库连接（`backend/blog/src/main/resources/application.yml`）
2. 在 `backend/blog` 目录下执行：
   ```bash
   mvn spring-boot:run
   ```
3. 确保数据库有管理员账号（如用户名 `wofiporia`，密码 `123456`）

### 前端
1. 在 `frontend` 目录下执行：
   ```bash
   npm install
   npm run dev
   ```
2. 访问 `http://localhost:5173` 查看博客首页

## 目录结构
```
wofiporia-blog/
  backend/
    blog/
      src/main/java/com/wofiporia/blog/...
      src/main/resources/application.yml
      pom.xml
  frontend/
    src/
      views/
      api/
      router/
      ...
    package.json
    vite.config.js
  README.md
```

## 开发环境
- Node.js 18+
- npm 9+ 或 yarn 1+
- JDK 17+（建议 21）
- Maven 3.8+
- MySQL 8+
- 推荐编辑器：VS Code / IntelliJ IDEA

---

如需二次开发或部署，建议先阅读源码，欢迎 star！
