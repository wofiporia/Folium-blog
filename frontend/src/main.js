import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import router from './router'
import axios from 'axios'

// 添加 axios 请求拦截器，自动加 token
axios.interceptors.request.use(config => {
  const token = localStorage.getItem('token')
  if (token) {
    // 如果token已经包含Bearer前缀，直接使用；否则添加Bearer前缀
    if (token.startsWith('Bearer ')) {
      config.headers.Authorization = token
    } else {
      config.headers.Authorization = 'Bearer ' + token
    }
  }
  return config
})

// 添加 axios 响应拦截器，处理401错误
axios.interceptors.response.use(
  response => {
    return response
  },
  error => {
    // 只在特定情况下自动处理401错误，避免与组件内错误处理冲突
    if (error.response && error.response.status === 401) {
      console.log('检测到401错误，token可能无效')
      
      // 只有在非管理页面时才自动跳转，管理页面让组件自己处理
      if (!window.location.pathname.startsWith('/admin')) {
        localStorage.removeItem('token')
        window.location.href = '/admin/login'
        alert('登录已过期，请重新登录')
      }
    }
    return Promise.reject(error)
  }
)

createApp(App).use(router).mount('#app')
