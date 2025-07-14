<template>
  <div class="container">
    <header>
      <h1>Wofiporia 的博客</h1>
      <p class="subtitle">一个极简的个人博客</p>
    </header>
    <main>
      <ul class="blog-list">
        <li v-for="blog in blogs" :key="blog.id" @click="goToDetail(blog.id)" class="blog-card">
          <h2>{{ blog.title }}</h2>
          <p class="summary">{{ blog.summary }}</p>
          <span class="date">{{ blog.date }}</span>
        </li>
      </ul>
    </main>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'

const modules = import.meta.glob('../blog-md/*.md', { as: 'raw' })

const blogs = ref([])
const router = useRouter()

onMounted(async () => {
  const blogList = []
  let idx = 1
  for (const path in modules) {
    const content = await modules[path]()
    const titleMatch = content.match(/^#\s+(.+)$/m)
    const dateMatch = content.match(/\*日期：([\d\-]+)\*/)
    const summaryMatch = content.match(/^[^#>\-\*\d\s][^\n]{5,}/m)
    blogList.push({
      id: path.split('/').pop().replace('.md', '') || String(idx),
      title: titleMatch ? titleMatch[1] : `未命名博客${idx}`,
      summary: summaryMatch ? summaryMatch[0].slice(0, 40) + '...' : '暂无摘要',
      date: dateMatch ? dateMatch[1] : ''
    })
    idx++
  }
  blogs.value = blogList.sort((a, b) => a.id.localeCompare(b.id))
})

function goToDetail(id) {
  router.push(`/blog/${id}`)
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
</style> 