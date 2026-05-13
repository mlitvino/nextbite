import type { Store, StoreInput } from '../types/store'
import { del, get, post, put } from './api'

type ListResponse = {
  items: Store[]
}

export async function listStores(): Promise<Store[]> {
  const data = await get<ListResponse>('/stores')
  return data.items
}

export async function getStoreById(id: string): Promise<Store> {
  return get<Store>(`/stores/${id}`)
}

export async function createStore(payload: StoreInput): Promise<Store> {
  return post<Store>('/stores', payload)
}

export async function updateStore(id: string, payload: StoreInput): Promise<Store> {
  return put<Store>(`/stores/${id}`, payload)
}

export async function deleteStore(id: string): Promise<void> {
  await del<void>(`/stores/${id}`)
}
