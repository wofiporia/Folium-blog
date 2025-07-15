<template>
  <div class="container">
    <button class="back" @click="goBack">返回</button>
    <div class="date">
      创建：{{ blogUploadDate }}<br>
      更新：{{ blogUpdateDate }}
    </div>
    <div class="content" v-html="htmlContent"></div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { marked } from 'marked'
import { getBlog } from '../api/blog'

const route = useRoute()
const router = useRouter()

const blogId = route.params.id
const blogContent = ref('')
const blogTitle = ref('')
const blogUploadDate = ref('')
const blogUpdateDate = ref('')

onMounted(async () => {
  const res = await getBlog(blogId)
  const blog = res.data
  blogContent.value = blog.content
  blogTitle.value = blog.title
  blogUploadDate.value = blog.uploadDate ? blog.uploadDate.replace('T', ' ').slice(0, 19) : ''
  blogUpdateDate.value = blog.updateDate ? blog.updateDate.replace('T', ' ').slice(0, 19) : ''
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