<template>
  <div class="admin-login-container">
    <h2>管理员登录</h2>
    <form @submit.prevent="handleLogin">
      <div class="form-group">
        <label>账号</label>
        <input v-model="username" type="text" required />
      </div>
      <div class="form-group">
        <label>密码</label>
        <input v-model="password" type="password" required />
      </div>
      <button type="submit">登录</button>
      <div v-if="error" class="error">{{ error }}</div>
    </form>
    <button class="back-home" @click="goHome">返回主界面</button>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import axios from 'axios'

const username = ref('')
const password = ref('')
const error = ref('')
const router = useRouter()

async function handleLogin() {
  error.value = ''
  try {
    const res = await axios.post('/api/admin/login', {
      username: username.value,
      password: password.value
    })
    if (res.data.success) {
      localStorage.setItem('token', res.data.token)
      router.push('/admin/panel')
    } else {
      error.value = res.data.message || '账号或密码错误'
    }
  } catch (e) {
    error.value = '登录失败，请检查网络或服务器！'
  }
}

function goHome() {
  router.push('/')
}
</script>

<style scoped>
.admin-login-container {
  max-width: 340px;
  margin: 120px auto;
  padding: 32px 28px 24px 28px;
  background: #fff;
  border-radius: 10px;
  box-shadow: 0 2px 16px 0 rgba(0,0,0,0.10);
  text-align: center;
}
h2 {
  margin-bottom: 24px;
}
.form-group {
  margin-bottom: 18px;
  text-align: left;
}
label {
  display: block;
  margin-bottom: 6px;
  color: #555;
}
input[type="text"], input[type="password"] {
  width: 100%;
  padding: 8px 10px;
  border: 1px solid #ccc;
  border-radius: 4px;
  font-size: 1em;
}
button[type="submit"] {
  width: 100%;
  padding: 10px 0;
  background: #646cff;
  color: #fff;
  border: none;
  border-radius: 4px;
  font-size: 1.1em;
  cursor: pointer;
  margin-top: 10px;
}
button[type="submit"]:hover {
  background: #4b53c0;
}
.error {
  color: #e74c3c;
  margin-top: 12px;
}
.back-home {
  margin-top: 18px;
  background: #eee;
  border: none;
  padding: 8px 20px;
  border-radius: 4px;
  cursor: pointer;
  font-size: 1em;
}
.back-home:hover {
  background: #f8faff;
  box-shadow: 0 4px 24px 0 rgba(0,0,0,0.12);
  transform: translateY(-2px) scale(1.01);
}
</style> 