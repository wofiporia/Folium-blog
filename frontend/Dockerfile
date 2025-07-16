# 第一阶段：构建前端静态文件
FROM node:20 AS build
WORKDIR /app
COPY . .
RUN npm install && chmod +x node_modules/.bin/vite && npm run build

# 第二阶段：用 nginx 服务静态文件
FROM nginx:alpine
COPY --from=build /app/dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf
EXPOSE 80 