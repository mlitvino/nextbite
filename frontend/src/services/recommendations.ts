import type { Store } from '../types/store'
import { get } from './api'

type ListResponse = {
  items: Store[]
}

type RecommendationsOptions = {
  limit?: number
}

export async function listMyRecommendations(options: RecommendationsOptions = {}): Promise<Store[]> {
  const params = new URLSearchParams()
  if (options.limit) {
    params.set('limit', String(options.limit))
  }

  const query = params.toString()
  const path = query ? `/me/recommendations?${query}` : '/me/recommendations'
  const data = await get<ListResponse>(path)
  return data.items
}
