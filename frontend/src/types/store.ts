export type Geo = {
  latitude: number
  longitude: number
}

export type Store = {
  id: string
  name: string
  primary_cuisine: string
  cuisines: string[]
  price_tier: number
  rating_avg: number
  rating_count: number
  orders_7d: number
  is_open_now: boolean
  created_at: string
  geo: Geo
}

export type StoreInput = {
  name: string
  primary_cuisine: string
  cuisines?: string[]
  price_tier?: number
  rating_avg?: number
  rating_count?: number
  orders_7d?: number
  is_open_now?: boolean
  created_at?: string
  geo?: Geo
}
