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
        <label>标题：</label>
        <input v-model="uploadTitle" placeholder="请输入标题" />
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
        <ul class="blog-list" v-if="Array.isArray(blogs) && blogs.length > 0">
          <li v-for="blog in blogs" :key="blog.id" class="blog-item">
            <span class="blog-title">{{ blog.title }}</span>
            <span class="blog-date">
              创建：{{ blog.uploadDate ? blog.uploadDate.replace('T', ' ').slice(0, 19) : '' }}<br>
              更新：{{ blog.updateDate ? blog.updateDate.replace('T', ' ').slice(0, 19) : '' }}
            </span>
            <button class="edit-btn" @click="editBlog(blog)">修改</button>
            <button class="delete-btn" @click="deleteBlogConfirm(blog)">删除</button>
          </li>
        </ul>
        <div v-else class="blog-item" style="text-align:center;color:#aaa;">暂无博客</div>
      </div>
      <!-- 编辑弹窗 -->
      <div v-if="showEditModal" class="modal-mask">
        <div class="modal-box">
          <h4>编辑博客</h4>
          <input v-model="editTitle" style="width:100%;margin-bottom:10px;" placeholder="标题" />
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
import { getAllBlogs, addBlog, updateBlog, deleteBlog } from '../api/blog'

const router = useRouter()
const activeTab = ref('upload')
const markdownContent = ref('')
const uploadTitle = ref('')
const uploadMsg = ref('')
const blogList = ref([])
const htmlContent = computed(() => marked.parse(markdownContent.value))

// 编辑与删除弹窗逻辑
const showEditModal = ref(false)
const showDeleteModal = ref(false)
const editContent = ref('')
let editingBlog = null
let deletingBlog = null
const editTitle = ref('')

// 加载博客列表
async function loadBlogs() {
  try {
    const res = await getAllBlogs()
    console.log('管理界面获取博客列表响应:', res)
    blogList.value = res.data.data
    console.log('管理界面博客列表:', blogList.value)
  } catch (error) {
    console.error('获取博客列表失败:', error)
  }
}

onMounted(() => {
  document.title = 'Folium-茯苓的博客园 - 控制台'
  if (!localStorage.getItem('token')) {
    router.push('/admin')
  } else {
    loadBlogs()
  }
})

// 上传Tab
async function handleFileUpload(e) {
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

async function confirmUpload() {
  if (!markdownContent.value.trim()) {
    uploadMsg.value = '内容不能为空！'
    return
  }
  
  let title = uploadTitle.value.trim()
  if (!title) {
    // 自动提取
    const titleMatch = markdownContent.value.match(/^#\s+(.+)$/m)
    title = titleMatch ? titleMatch[1] : '未命名博客'
  }
  
  try {
    await addBlog(title, markdownContent.value)
    uploadMsg.value = '上传成功！'
    markdownContent.value = ''
    uploadTitle.value = ''
    loadBlogs()
  } catch (error) {
    console.error('添加博客失败:', error)
    let errorMessage = '添加博客失败'
    
    if (error.response) {
      // 服务器返回了错误响应
      if (error.response.data && error.response.data.message) {
        errorMessage = error.response.data.message
      } else if (error.response.status === 401) {
        errorMessage = '登录已过期，请重新登录'
        localStorage.removeItem('token')
        router.push('/admin')
        return
      } else if (error.response.status === 403) {
        errorMessage = '没有权限添加博客'
      } else if (error.response.status >= 500) {
        errorMessage = '服务器错误，请稍后重试'
      }
    } else if (error.request) {
      // 网络错误
      errorMessage = '网络连接失败，请检查网络'
    }
    
    uploadMsg.value = errorMessage
  }
}

// 管理Tab
function editBlog(blog) {
  editingBlog = blog
  editTitle.value = blog.title
  editContent.value = blog.content
  showEditModal.value = true
}
async function saveEdit() {
  if (editingBlog) {
    try {
      await updateBlog(editingBlog.id, editTitle.value, editContent.value)
      showEditModal.value = false
      loadBlogs()
      alert('博客更新成功！')
    } catch (error) {
      console.error('更新博客失败:', error)
      let errorMessage = '更新博客失败'
      
      if (error.response) {
        // 服务器返回了错误响应
        if (error.response.data && error.response.data.message) {
          errorMessage = error.response.data.message
        } else if (error.response.status === 401) {
          errorMessage = '登录已过期，请重新登录'
          localStorage.removeItem('token')
          router.push('/admin')
          return
        } else if (error.response.status >= 500) {
          errorMessage = '服务器错误，请稍后重试'
        }
      } else if (error.request) {
        // 网络错误
        errorMessage = '网络连接失败，请检查网络'
      }
      
      alert(errorMessage)
    }
  }
}
function deleteBlogConfirm(blog) {
  deletingBlog = blog
  showDeleteModal.value = true
}
async function confirmDelete() {
  if (deletingBlog) {
    try {
      await deleteBlog(deletingBlog.id)
      showDeleteModal.value = false
      loadBlogs()
      alert('博客删除成功！')
    } catch (error) {
      console.error('删除博客失败:', error)
      let errorMessage = '删除博客失败'
      
      if (error.response) {
        // 服务器返回了错误响应
        if (error.response.data && error.response.data.message) {
          errorMessage = error.response.data.message
        } else if (error.response.status === 401) {
          errorMessage = '登录已过期，请重新登录'
          localStorage.removeItem('token')
          router.push('/admin')
          return
        } else if (error.response.status === 403) {
          errorMessage = '没有权限删除此博客'
        } else if (error.response.status === 404) {
          errorMessage = '博客不存在或已被删除'
        } else if (error.response.status >= 500) {
          errorMessage = '服务器错误，请稍后重试'
        }
      } else if (error.request) {
        // 网络错误
        errorMessage = '网络连接失败，请检查网络'
      }
      
      alert(errorMessage)
      // 删除失败时不关闭模态框，让用户可以重试
    }
  }
}

function logout() {
  localStorage.removeItem('token')
  router.push('/admin')
}
function goHome() {
  router.push('/')
}

// 标签分组（可选，按需实现）
const groupedBlogs = computed(() => {
  // 这里只做全部分组，如需按标签分组可扩展
  return { '全部博客': blogList.value }
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