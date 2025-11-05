<template>
  <div class="container">
    <header>
      <a class="github-link" href="https://github.com/wofiporia" target="_blank" title="GitHub主页">
        <svg class="github-icon" viewBox="0 0 1024 1024" width="32" height="32" fill="currentColor">
          <path d="M512 76C273.6 76 76 273.6 76 512c0 192.8 125.2 356.4 298.8 414.2 21.8 4 29.8-9.4 29.8-20.8 0-10.2-0.4-44.2-0.6-80.2-121.6 26.4-147.4-51.6-147.4-51.6-19.8-50.2-48.4-63.6-48.4-63.6-39.6-27 3-26.4 3-26.4 43.8 3.2 66.8 45 66.8 45 38.8 66.6 101.8 47.4 126.6 36.2 3.8-28 15.2-47.4 27.6-58.4-97-11-199-48.4-199-215.6 0-47.6 17-86.6 44.8-117.2-4.4-11-19.4-55.2 4.2-115 0 0 36.6-11.8 120 44.8 34.8-9.6 72.2-14.4 109.4-14.6 37.2 0.2 74.6 5 109.4 14.6 83.2-56.6 119.8-44.8 119.8-44.8 23.6 59.8 8.6 104 4.2 115 27.8 30.6 44.8 69.6 44.8 117.2 0 167.6-102.2 204.4-199.6 215.2 15.6 13.4 29.4 39.8 29.4 80.2 0 57.8-0.6 104.4-0.6 118.6 0 11.6 8 24.8 30 20.6C822.8 868.2 948 704.8 948 512 948 273.6 750.4 76 512 76z" />
        </svg>
      </a>
      <button class="admin-btn" @click="goAdmin">管理员入口</button>
      <h1>Folium-茯苓的博客园</h1>
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
  document.title = 'Folium-茯苓的博客园'
  loading.value = true
  try {
    const res = await getAllBlogs()
    console.log('获取博客列表响应:', res)
    blogs.value = res.data.data // 修正：使用res.data.data而不是res.data
    console.log('博客列表:', blogs.value)
  } catch (error) {
    console.error('获取博客列表失败:', error)
  } finally {
    loading.value = false
  }
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
.github-link {
  position: absolute;
  top: 18px;
  left: 24px;
  z-index: 10;
  color: #222;
  opacity: 0.85;
  transition: opacity 0.2s;
}
.github-link:hover {
  opacity: 1;
  color: #646cff;
}
.github-icon {
  display: block;
}
</style> 