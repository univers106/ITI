export async function me(): Promise<boolean> {
  try {
    const res = await fetch('api/private/hello')
    console.log(res)
    return res.ok
  } catch {
    return false
  }
}

export async function login(loginValue: string, passwordValue: string): Promise<Response> {
  return fetch('/api/public/login', {
    method: 'POST',
    headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
    body: new URLSearchParams({ login: loginValue, password: passwordValue }),
  })
}

export async function logout(): Promise<Response> {
  return fetch('/api/private/logout', { method: 'POST' })
}
