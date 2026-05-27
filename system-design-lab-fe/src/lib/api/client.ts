import axios from 'axios'
import { config } from '@/config/env'

export interface ApiError extends Error {
  status: number
  code: string
}

function createApiError(status: number, code: string, message: string): ApiError {
  const err = new Error(message) as ApiError
  err.name = 'ApiError'
  err.status = status
  err.code = code
  return err
}

export const apiClient = axios.create({
  baseURL: config.apiBaseUrl,
  headers: { 'Content-Type': 'application/json' },
  timeout: 10_000,
})

apiClient.interceptors.response.use(
  (res) => res,
  (err) => {
    const status: number = err.response?.status ?? 0
    const message: string = err.response?.data?.error ?? err.message ?? 'Unknown error'
    return Promise.reject(createApiError(status, String(status), message))
  },
)
