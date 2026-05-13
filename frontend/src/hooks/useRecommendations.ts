import { useEffect, useState } from 'react'
import { listStores } from '../services/stores'
import type { Store } from '../types/store'

type UseRecommendationsResult = {
  allStores: Store[]
  myStores: Store[]
  isLoading: boolean
  error: string | null
}

const mockMyStores: Store[] = [
  {
    id: 'mock-1',
    name: 'Saffron & Lime',
    primary_cuisine: 'Indian',
    cuisines: ['Indian', 'Vegetarian'],
    price_tier: 2,
    rating_avg: 4.7,
    rating_count: 142,
    orders_7d: 118,
    is_open_now: true,
    created_at: '2024-04-02T10:00:00Z',
    geo: { latitude: 37.776, longitude: -122.423 },
  },
  {
    id: 'mock-2',
    name: 'Portside Poke',
    primary_cuisine: 'Hawaiian',
    cuisines: ['Hawaiian', 'Seafood'],
    price_tier: 2,
    rating_avg: 4.5,
    rating_count: 87,
    orders_7d: 64,
    is_open_now: false,
    created_at: '2024-02-14T09:30:00Z',
    geo: { latitude: 37.792, longitude: -122.396 },
  },
]

export function useRecommendations(): UseRecommendationsResult {
  const [allStores, setAllStores] = useState<Store[]>([])
  const [isLoading, setIsLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)

  useEffect(() => {
    let isMounted = true

    const fetchAll = async () => {
      setIsLoading(true)
      setError(null)
      try {
        const data = await listStores()
        if (!isMounted) return
        setAllStores(data)
      } catch (err) {
        if (!isMounted) return
        const message = err instanceof Error ? err.message : 'Failed to load stores.'
        setError(message)
      } finally {
        if (!isMounted) return
        setIsLoading(false)
      }
    }

    fetchAll()

    return () => {
      isMounted = false
    }
  }, [])

  return { allStores, myStores: mockMyStores, isLoading, error }
}
