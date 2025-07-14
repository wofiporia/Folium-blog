import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import Home from '../views/Home.vue'
import BlogDetail from '../views/BlogDetail.vue'

const routes: Array<RouteRecordRaw> = [
  { path: '/', name: 'Home', component: Home },
  { path: '/blog/:id', name: 'BlogDetail', component: BlogDetail }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router 