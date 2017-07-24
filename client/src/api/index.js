import axios from 'axios'

const jwt = window.localStorage.getItem('token')

export const http = axios.create({
  baseURL: 'http://localhost:8081/api',
  timeout: 5000,
  headers: {
    'Authorization': jwt
  }
})
