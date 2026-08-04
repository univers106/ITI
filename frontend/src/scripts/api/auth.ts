export type User = {
  login: string;
  name: string;
  permissions: string[];
}

export async function meApi(): Promise<null | User> {
  try {
    const res = await fetch('/api/auth/me')
    return res.json()
  } catch {
    return null
  }
}

export async function loginApi(loginValue: string, passwordValue: string): Promise<Response> {
  return fetch('/api/auth/login', {
    method: 'POST',
    headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
    body: new URLSearchParams({ login: loginValue, password: passwordValue }),
  })
}

export async function logoutApi(): Promise<Response> {
  return fetch('/api/auth/logout', { method: 'POST' })
}

export async function changePasswordApi(oldPassword: string, newPassword: string): Promise<Response> {
  return fetch('/api/auth/change-password', {
    method: 'POST',
    headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
    body: new URLSearchParams({ oldPassword: oldPassword, newPassword: newPassword }),
  })
}
