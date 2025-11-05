import { defineConfig, loadEnv } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vite.dev/config/
export default defineConfig(({ mode }) => {
  // 从项目根目录加载环境变量
  const env = loadEnv(mode, '../', '')
  
  // 使用环境变量或默认值
  const apiBase = env.VITE_API_BASE || 'http://localhost:8081'
  const devPort = Number(env.VITE_DEV_PORT || 5173)
  
  console.log('Vite Config - API Base:', apiBase)
  console.log('Vite Config - Dev Port:', devPort)
  
  return {
    plugins: [vue()],
    server: {
      port: devPort,
      proxy: {
        '/api': {
          target: apiBase,
          changeOrigin: true,
        },
      },
    },
  }
})