import axios from 'axios'

const BASE = '/api/blog'

export function getAllBlogs() {
  return axios.get(`${BASE}/all`)
}

export function getBlog(id) {
  return axios.get(`${BASE}/${id}`)
}

export function addBlog(title, content) {
  return axios.post(`${BASE}/add`, { title, content })
}

export function updateBlog(id, title, content) {
  return axios.post(`${BASE}/update/${id}`, { title, content })
}

export function deleteBlog(id) {
  return axios.delete(`${BASE}/delete/${id}`)
} 