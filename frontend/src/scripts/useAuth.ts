import { ref } from 'vue'
import { meApi, loginApi, logoutApi, changePasswordApi, type User } from './api/auth'

const isAuthenticated = ref(false)
const curUser = ref<User | null>(null)
const isInitialized = ref(false)

export function useAuth() {
  async function updateMe(): Promise<User | null> {
    const user = await meApi()

    isAuthenticated.value = user?.login != undefined

    curUser.value = user
    isInitialized.value = true
    return user as User || null
  }

  async function me(): Promise<User | null> {
    if (!isInitialized.value) {
      await updateMe()
    }
    return curUser.value
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

  async function changePassword(oldPassword: string, newPassword: string): Promise<boolean> {
    const res = await changePasswordApi(oldPassword, newPassword)
    return res.ok
  }

  return { isAuthenticated, me, login, logout, changePassword }
}
