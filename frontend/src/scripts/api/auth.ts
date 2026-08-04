export type User = {
  login: string;
  name: string;
  permissions: string[];
}

export async function meApi(): Promise<null | User> {
  try {
    const res = await fetch('/api/auth/me')
    if (!res.ok) return null;
    return res.json()
  } catch {
    return null
  }
}

export async function loginApi(loginValue: string, passwordValue: string): Promise<Response> {
  return fetch('/api/auth/login', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ login: loginValue, password: passwordValue }),
  })
}

export async function logoutApi(): Promise<Response> {
  return fetch('/api/auth/logout', { method: 'POST' })
}

export async function changePasswordApi(oldPassword: string, newPassword: string): Promise<Response> {
  return fetch('/api/auth/change-password', {
    method: 'PATCH',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ oldPassword, newPassword }),
  })
}
