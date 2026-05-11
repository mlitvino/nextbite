const API_BASE = import.meta.env.VITE_API_BASE ?? 'http://localhost:3000/api'

type LoginPayload = {
  username: string
  password: string
}

type SignupPayload = {
  name: string
  username: string
  password: string
}

type User = {
  id: string
  name: string
  username: string
}

export async function login(payload: LoginPayload): Promise<User> {
  const response = await fetch(`${API_BASE}/auth/login`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    credentials: 'include',
    body: JSON.stringify(payload),
  })

  if (!response.ok) {
    const body = await response.json().catch(() => ({ error: 'login failed' }))
    throw new Error(body.error ?? 'login failed')
  }

  return response.json()
}

export async function logout(): Promise<void> {
  const response = await fetch(`${API_BASE}/auth/logout`, {
    method: 'POST',
    credentials: 'include',
  })

  if (!response.ok) {
    const body = await response.json().catch(() => ({ error: 'logout failed' }))
    throw new Error(body.error ?? 'logout failed')
  }
}

export async function signup(payload: SignupPayload): Promise<User> {
  const response = await fetch(`${API_BASE}/auth/signup`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    credentials: 'include',
    body: JSON.stringify(payload),
  })

  if (!response.ok) {
    const body = await response.json().catch(() => ({ error: 'signup failed' }))
    throw new Error(body.error ?? 'signup failed')
  }

  return response.json()
}
