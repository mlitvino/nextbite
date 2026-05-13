import type { Store, StoreInput } from '../types/store'

const API_BASE = '/api'

type ListResponse = {
  items: Store[]
}

export async function listStores(): Promise<Store[]> {
  const response = await fetch(`${API_BASE}/stores`, {
    credentials: 'include',
  })

  if (!response.ok) {
    const body = await response.json().catch(() => ({ error: 'failed to list stores' }))
    throw new Error(body.error ?? 'failed to list stores')
  }

  const data: ListResponse = await response.json()
  return data.items
}

export async function getStoreById(id: string): Promise<Store> {
  const response = await fetch(`${API_BASE}/stores/${id}`, {
    credentials: 'include',
  })

  if (!response.ok) {
    const body = await response.json().catch(() => ({ error: 'failed to fetch store' }))
    throw new Error(body.error ?? 'failed to fetch store')
  }

  return response.json()
}

export async function createStore(payload: StoreInput): Promise<Store> {
  const response = await fetch(`${API_BASE}/stores`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    credentials: 'include',
    body: JSON.stringify(payload),
  })

  if (!response.ok) {
    const body = await response.json().catch(() => ({ error: 'failed to create store' }))
    throw new Error(body.error ?? 'failed to create store')
  }

  return response.json()
}

export async function updateStore(id: string, payload: StoreInput): Promise<Store> {
  const response = await fetch(`${API_BASE}/stores/${id}`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    credentials: 'include',
    body: JSON.stringify(payload),
  })

  if (!response.ok) {
    const body = await response.json().catch(() => ({ error: 'failed to update store' }))
    throw new Error(body.error ?? 'failed to update store')
  }

  return response.json()
}

export async function deleteStore(id: string): Promise<void> {
  const response = await fetch(`${API_BASE}/stores/${id}`, {
    method: 'DELETE',
    credentials: 'include',
  })

  if (!response.ok) {
    const body = await response.json().catch(() => ({ error: 'failed to delete store' }))
    throw new Error(body.error ?? 'failed to delete store')
  }
}
