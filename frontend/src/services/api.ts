type ApiErrorBody = {
  error?: string
}

type RequestOptions = {
  method?: string
  body?: unknown
}

const API_BASE = '/api'

async function request<T>(path: string, options: RequestOptions = {}): Promise<T> {
  const response = await fetch(`${API_BASE}${path}`, {
    method: options.method ?? 'GET',
    headers: options.body ? { 'Content-Type': 'application/json' } : undefined,
    credentials: 'include',
    body: options.body ? JSON.stringify(options.body) : undefined,
  })

  if (!response.ok) {
    const body = await response.json().catch(() => ({}))
    const message = (body as ApiErrorBody).error ?? 'request failed'
    throw new Error(message)
  }

  if (response.status === 204) {
    return undefined as T
  }

  return response.json() as Promise<T>
}

export function get<T>(path: string): Promise<T> {
  return request<T>(path)
}

export function post<T>(path: string, body?: unknown): Promise<T> {
  return request<T>(path, { method: 'POST', body })
}

export function put<T>(path: string, body?: unknown): Promise<T> {
  return request<T>(path, { method: 'PUT', body })
}

export function del<T>(path: string): Promise<T> {
  return request<T>(path, { method: 'DELETE' })
}
