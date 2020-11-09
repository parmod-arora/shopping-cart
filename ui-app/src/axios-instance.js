import axios from 'axios'
import { AUTH_TOKEN_KEY } from "./config";
const instance = axios.create({
  headers: {
    'Content-Type': 'application/json'
  },
  timeout: 10000
})

// Set the AUTH token for any request
instance.interceptors.request.use((config) => {
  const token = sessionStorage.getItem(AUTH_TOKEN_KEY)
  config.headers.Authorization = token ? `Bearer ${token}` : ''
  return config
}, (error) => {
  return Promise.reject(error)
})

instance.interceptors.response.use(response => {
  return response.data
}, error => {
  console.log(error.response.data)
  if (error.response.status === 401) {
    sessionStorage.removeItem(AUTH_TOKEN_KEY)
    window.location.replace("/login");
    return Promise.reject(error)
  }
  if (error.response.data) {
    return Promise.reject(error.response.data)  
  }
  return Promise.reject(error)
})

export default instance