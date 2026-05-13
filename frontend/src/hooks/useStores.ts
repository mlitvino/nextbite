import { useEffect, useState } from 'react'
import { listStores } from '../services/stores'
import type { Store } from '../types/store'

type UseStoresResult = {
  stores: Store[]
  isLoading: boolean
  error: string | null
}

export function useStores(): UseStoresResult {
  const [stores, setStores] = useState<Store[]>([])
  const [isLoading, setIsLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)

  useEffect(() => {
    let isMounted = true

    const fetchStores = async () => {
      setIsLoading(true)
      setError(null)
      try {
        const data = await listStores()
        if (!isMounted) return
        setStores(data)
      } catch (err) {
        if (!isMounted) return
        const message = err instanceof Error ? err.message : 'Failed to load stores.'
        setError(message)
      } finally {
        if (!isMounted) return
        setIsLoading(false)
      }
    }

    fetchStores()

    return () => {
      isMounted = false
    }
  }, [])

  return { stores, isLoading, error }
}
