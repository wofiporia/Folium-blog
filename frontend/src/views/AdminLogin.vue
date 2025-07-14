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
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const username = ref('')
const password = ref('')
const error = ref('')
const router = useRouter()

const ADMIN_USER = 'admin'
const ADMIN_PASS = '123456'

function handleLogin() {
  if (username.value === ADMIN_USER && password.value === ADMIN_PASS) {
    localStorage.setItem('admin_login', '1')
    router.push('/admin/panel')
  } else {
    error.value = '账号或密码错误'
  }
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
</style> 