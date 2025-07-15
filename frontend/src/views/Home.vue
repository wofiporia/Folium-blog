<template>
  <div class="container">
    <header>
      <button class="admin-btn" @click="goAdmin">管理员入口</button>
      <h1>Wofiporia 的博客</h1>
      <p class="subtitle">一个极简的个人博客</p>
    </header>
    <main>
      <ul class="blog-list" v-if="!loading && Array.isArray(blogs) && blogs.length > 0">
        <li v-for="blog in blogs" :key="blog.id" @click="goToDetail(blog.id)" class="blog-card">
          <h2>{{ blog.title }}</h2>
          <p class="summary">{{ getSummary(blog.content) }}</p>
          <span class="date">
            创建：{{ blog.uploadDate ? blog.uploadDate.replace('T', ' ').slice(0, 19) : '' }}<br>
            更新：{{ blog.updateDate ? blog.updateDate.replace('T', ' ').slice(0, 19) : '' }}
          </span>
        </li>
      </ul>
      <div v-else-if="loading" class="blog-card" style="text-align:center;color:#aaa;">加载中...</div>
      <div v-else class="blog-card" style="text-align:center;color:#aaa;">暂无博客</div>
    </main>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { getAllBlogs } from '../api/blog'

const blogs = ref([])
const loading = ref(true)
const router = useRouter()

onMounted(async () => {
  loading.value = true
  const res = await getAllBlogs()
  blogs.value = res.data
  loading.value = false
})

function goToDetail(id) {
  router.push(`/blog/${id}`)
}

function getSummary(content) {
  if (!content) return ''
  // 取正文第一段（去掉markdown标题和空行）
  const lines = content.split('\n').filter(line => line.trim() && !line.trim().startsWith('#'))
  return lines.length ? lines[0].slice(0, 60) + (lines[0].length > 60 ? '...' : '') : ''
}

function goAdmin() {
  router.push('/admin')
}
</script>

<style scoped>
.container {
  max-width: 1000px;
  width: 96%;
  margin: 0 auto;
  padding: 0 20px 40px 20px;
  font-family: 'Segoe UI', 'PingFang SC', Arial, sans-serif;
}
header {
  text-align: center;
  margin-top: 0;
  margin-bottom: 32px;
  position: relative;
}
.subtitle {
  color: #888;
  font-size: 1.1em;
}
.blog-list {
  list-style: none;
  padding: 0;
  margin: 0;
}
.blog-card {
  background: #fff;
  border-radius: 12px;
  box-shadow: 0 4px 24px 0 rgba(0,0,0,0.10);
  border: 1.5px solid #e0e3ea;
  padding: 22px 24px 18px 24px;
  margin-bottom: 24px;
  cursor: pointer;
  transition: box-shadow 0.2s, transform 0.2s;
}
.blog-card:hover {
  box-shadow: 0 4px 24px 0 rgba(0,0,0,0.12);
  transform: translateY(-2px) scale(1.01);
  background: #f8faff;
}
.blog-card h2 {
  margin: 0 0 8px 0;
  font-size: 1.3em;
}
.summary {
  color: #666;
  margin: 8px 0 12px 0;
}
.date {
  font-size: 0.9em;
  color: #aaa;
}
.header-bar {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: 32px;
}
.admin-btn {
  position: absolute;
  top: 18px;
  right: 24px;
  background: #eee;
  border: none;
  padding: 8px 20px;
  border-radius: 4px;
  cursor: pointer;
  font-size: 1em;
}
.admin-btn:hover {
  background: #f8faff;
  box-shadow: 0 4px 24px 0 rgba(0,0,0,0.12);
  transform: translateY(-2px) scale(1.01);
}
</style> 