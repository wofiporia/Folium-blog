import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import Home from '../views/Home.vue'
import BlogDetail from '../views/BlogDetail.vue'
import AdminLogin from '../views/AdminLogin.vue'
import AdminPanel from '../views/AdminPanel.vue'

const routes: Array<RouteRecordRaw> = [
  { path: '/', name: 'Home', component: Home },
  { path: '/blog/:id', name: 'BlogDetail', component: BlogDetail },
  { path: '/admin', name: 'AdminLogin', component: AdminLogin },
  { path: '/admin/panel', name: 'AdminPanel', component: AdminPanel },
]


const router = createRouter({
  history: createWebHistory(),
  routes
})

// 路由守卫：保护 /admin/panel，未登录跳转到 /admin
router.beforeEach((to, from, next) => {
  if (to.path === '/admin/panel') {
    if (localStorage.getItem('admin_login') === '1') {
      next()
    } else {
      next('/admin')
    }
  } else {
    next()
  }
})

export default router 