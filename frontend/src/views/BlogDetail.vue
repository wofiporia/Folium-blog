<template>
  <div class="container">
    <button class="back" @click="goBack">返回</button>
    <div class="date">{{ blogDate }}</div>
    <div class="content" v-html="htmlContent"></div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { marked } from 'marked'

const route = useRoute()
const router = useRouter()

const blogId = route.params.id
const modules = import.meta.glob('../blog-md/*.md', { as: 'raw' })

const blogContent = ref('')
const blogTitle = ref('')
const blogDate = ref('')

onMounted(async () => {
  for (const path in modules) {
    if (path.includes(`${blogId}.md`)) {
      const content = await modules[path]()
      blogContent.value = content
      const titleMatch = content.match(/^#\s+(.+)$/m)
      const dateMatch = content.match(/\*日期：([\d\-]+)\*/)
      blogTitle.value = titleMatch ? titleMatch[1] : '未命名博客'
      blogDate.value = dateMatch ? dateMatch[1] : ''
      break
    }
  }
})

const htmlContent = computed(() => marked.parse(blogContent.value))

function goBack() {
  router.push('/')
}
</script>

<style scoped>
.container {
  max-width: 1000px;
  width: 96%;
  margin: 40px auto;
  padding: 20px;
  font-family: 'Segoe UI', 'PingFang SC', Arial, sans-serif;
  position: relative;
}
.back {
  position: fixed;
  top: 24px;
  right: 32px;
  margin-bottom: 0;
  background: #eee;
  border: none;
  padding: 6px 16px;
  border-radius: 4px;
  cursor: pointer;
  z-index: 1000;
}
.back:hover {
  background: #ddd;
}
.date {
  color: #aaa;
  font-size: 0.95em;
  margin-bottom: 18px;
}
.content {
  font-size: 1.1em;
  color: #444;
  line-height: 1.8;
  margin-top: 16px;
}
</style> 