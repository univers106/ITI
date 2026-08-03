import { ref } from 'vue'
import { me as checkAuthApi, login as loginApi, logout as logoutApi } from './auth'

const isAuthenticated = ref(false)

export function useAuth() {
  async function checkAuth(): Promise<void> {
    isAuthenticated.value = await checkAuthApi()
  }

  async function login(loginValue: string, passwordValue: string): Promise<boolean> {
    const res = await loginApi(loginValue, passwordValue)
    if (res.ok) {
      isAuthenticated.value = true
    }
    return res.ok
  }

  async function logout(): Promise<boolean> {
    const res = await logoutApi()
    if (res.ok) {
      isAuthenticated.value = false
    }
    return res.ok
  }

  return { isAuthenticated, checkAuth, login, logout }
}
