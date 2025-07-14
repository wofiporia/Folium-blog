<template>
  <div class="admin-panel-container">
    <div class="panel-header">
      <h2>博客管理后台</h2>
      <div>
        <button class="back-home" @click="goHome">返回主界面</button>
        <button class="logout" @click="logout">退出登录</button>
      </div>
    </div>
    <div class="tab-bar">
      <div :class="['tab', activeTab === 'upload' ? 'active' : '']" @click="activeTab = 'upload'">上传</div>
      <div :class="['tab', activeTab === 'manage' ? 'active' : '']" @click="activeTab = 'manage'">管理</div>
    </div>
    <div v-if="activeTab === 'upload'">
      <div class="upload-section">
        <label>上传 Markdown 文件：</label>
        <input type="file" accept=".md" @change="handleFileUpload" />
      </div>
      <div class="edit-section">
        <label>编辑/粘贴 Markdown 内容：</label>
        <textarea v-model="markdownContent" rows="12" placeholder="在这里编辑或粘贴 markdown 内容"></textarea>
      </div>
      <div class="preview-section">
        <label>预览：</label>
        <div class="preview" v-html="htmlContent"></div>
      </div>
      <button class="confirm-btn" @click="confirmUpload">确认上传</button>
      <div v-if="uploadMsg" class="upload-msg">{{ uploadMsg }}</div>
    </div>
    <div v-else>
      <div v-for="(blogs, tag) in groupedBlogs" :key="tag" class="tag-group">
        <h3>{{ tag }}</h3>
        <ul class="blog-list">
          <li v-for="blog in blogs" :key="blog.id" class="blog-item">
            <span class="blog-title">{{ blog.title }}</span>
            <span class="blog-date">{{ blog.date }}</span>
            <button class="edit-btn" @click="editBlog(blog)">修改</button>
            <button class="delete-btn" @click="deleteBlog(blog)">删除</button>
          </li>
        </ul>
      </div>
      <!-- 编辑弹窗 -->
      <div v-if="showEditModal" class="modal-mask">
        <div class="modal-box">
          <h4>编辑博客</h4>
          <textarea v-model="editContent" rows="10" style="width:100%"></textarea>
          <div class="modal-actions">
            <button @click="saveEdit">保存</button>
            <button @click="showEditModal = false">取消</button>
          </div>
        </div>
      </div>
      <!-- 删除确认弹窗 -->
      <div v-if="showDeleteModal" class="modal-mask">
        <div class="modal-box">
          <h4>确认删除？</h4>
          <div class="modal-actions">
            <button @click="confirmDelete">确定</button>
            <button @click="showDeleteModal = false">取消</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { marked } from 'marked'

const router = useRouter()
const activeTab = ref('upload')
const markdownContent = ref('')
const uploadMsg = ref('')

const htmlContent = computed(() => marked.parse(markdownContent.value))

// 假数据博客列表
const blogList = ref([
  { id: 1, title: '第一篇博客', date: '2024-06-01', tags: ['前端', '生活'], content: '内容1' },
  { id: 2, title: '第二篇博客', date: '2024-06-02', tags: ['后端'], content: '内容2' },
  { id: 3, title: '第三篇博客', date: '2024-06-03', tags: ['前端'], content: '内容3' },
])

const groupedBlogs = computed(() => {
  const groups = {}
  blogList.value.forEach(blog => {
    blog.tags.forEach(tag => {
      if (!groups[tag]) groups[tag] = []
      groups[tag].push(blog)
    })
  })
  return groups
})

// 编辑与删除弹窗逻辑
const showEditModal = ref(false)
const showDeleteModal = ref(false)
const editContent = ref('')
let editingBlog = null
let deletingBlog = null

function editBlog(blog) {
  editingBlog = blog
  editContent.value = blog.content
  showEditModal.value = true
}
function saveEdit() {
  if (editingBlog) {
    editingBlog.content = editContent.value
    showEditModal.value = false
  }
}
function deleteBlog(blog) {
  deletingBlog = blog
  showDeleteModal.value = true
}
function confirmDelete() {
  if (deletingBlog) {
    blogList.value = blogList.value.filter(b => b !== deletingBlog)
    showDeleteModal.value = false
  }
}

function handleFileUpload(e) {
  const file = e.target.files[0]
  if (file && file.name.endsWith('.md')) {
    const reader = new FileReader()
    reader.onload = (evt) => {
      markdownContent.value = evt.target.result
    }
    reader.readAsText(file)
  } else {
    uploadMsg.value = '请上传 .md 格式的文件'
  }
}

function confirmUpload() {
  if (!markdownContent.value.trim()) {
    uploadMsg.value = '内容不能为空！'
    return
  }
  uploadMsg.value = '上传成功！（此处为模拟，实际应调用后端接口）'
}

function logout() {
  localStorage.removeItem('admin_login')
  router.push('/admin')
}

function goHome() {
  router.push('/')
}

onMounted(() => {
  if (localStorage.getItem('admin_login') !== '1') {
    router.push('/admin')
  }
})
</script>

<style scoped>
.admin-panel-container {
  max-width: 900px;
  margin: 40px auto;
  background: #fff;
  border-radius: 10px;
  box-shadow: 0 2px 16px 0 rgba(0,0,0,0.10);
  padding: 32px 28px 28px 28px;
}
.panel-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
}
.logout {
  background: #eee;
  border: none;
  padding: 6px 16px;
  border-radius: 4px;
  cursor: pointer;
}
.logout:hover {
  background: #ddd;
}
.back-home {
  background: #eee;
  border: none;
  padding: 6px 16px;
  border-radius: 4px;
  cursor: pointer;
  margin-right: 12px;
}
.back-home:hover {
  background: #d6eaff;
}
.tab-bar {
  display: flex;
  margin-bottom: 24px;
  border-bottom: 1.5px solid #e0e3ea;
}
.tab {
  padding: 10px 32px;
  cursor: pointer;
  font-size: 1.1em;
  color: #888;
  border-bottom: 2.5px solid transparent;
  transition: color 0.2s, border-bottom 0.2s;
}
.tab.active {
  color: #646cff;
  border-bottom: 2.5px solid #646cff;
  font-weight: bold;
}
.upload-section, .edit-section, .preview-section {
  margin-bottom: 22px;
}
label {
  font-weight: bold;
  display: block;
  margin-bottom: 8px;
}
textarea {
  width: 100%;
  font-size: 1em;
  padding: 10px;
  border-radius: 4px;
  border: 1px solid #ccc;
  resize: vertical;
  min-height: 120px;
}
.preview {
  background: #fafbfc;
  border: 1px solid #e0e3ea;
  border-radius: 4px;
  padding: 16px;
  min-height: 80px;
  font-size: 1em;
  color: #444;
}
.confirm-btn {
  background: #646cff;
  color: #fff;
  border: none;
  border-radius: 4px;
  padding: 10px 32px;
  font-size: 1.1em;
  cursor: pointer;
  margin-top: 10px;
}
.confirm-btn:hover {
  background: #4b53c0;
}
.upload-msg {
  margin-top: 16px;
  color: #2ecc71;
  font-weight: bold;
}
.tag-group {
  margin-bottom: 32px;
}
.blog-list {
  list-style: none;
  padding: 0;
}
.blog-item {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 10px 0;
  border-bottom: 1px solid #f0f0f0;
}
.blog-title {
  flex: 1;
  font-weight: 500;
}
.blog-date {
  color: #aaa;
  margin: 0 18px;
  font-size: 0.98em;
}
.edit-btn, .delete-btn {
  background: #eee;
  border: none;
  border-radius: 4px;
  padding: 4px 12px;
  margin-left: 8px;
  cursor: pointer;
  font-size: 0.98em;
}
.edit-btn:hover {
  background: #c7e0ff;
}
.delete-btn:hover {
  background: #ffd6d6;
}
.modal-mask {
  position: fixed;
  left: 0; top: 0; right: 0; bottom: 0;
  background: rgba(0,0,0,0.18);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
}
.modal-box {
  background: #fff;
  border-radius: 8px;
  padding: 24px 20px 16px 20px;
  min-width: 320px;
  max-width: 90vw;
  box-shadow: 0 2px 16px 0 rgba(0,0,0,0.13);
}
.modal-actions {
  margin-top: 18px;
  text-align: right;
}
.modal-actions button {
  margin-left: 12px;
  padding: 6px 18px;
  border-radius: 4px;
  border: none;
  cursor: pointer;
}
</style> 