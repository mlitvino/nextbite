import SessionPanel from './components/SessionPanel'
import StoreItem from './components/StoreItem'
import type { Store } from './types/store'
import './App.css'

const stores: Store[] = [
  {
    id: 'store-1',
    name: 'Golden Wok Express',
    primary_cuisine: 'Chinese',
    cuisines: ['Chinese', 'Asian'],
    price_tier: 2,
    rating_avg: 4.6,
    rating_count: 214,
    orders_7d: 132,
    is_open_now: true,
    created_at: '2024-05-10T12:00:00Z',
    geo: { latitude: 37.789, longitude: -122.401 },
  },
  {
    id: 'store-2',
    name: 'Bella Vita Pizza',
    primary_cuisine: 'Italian',
    cuisines: ['Italian', 'Pizza'],
    price_tier: 3,
    rating_avg: 4.3,
    rating_count: 98,
    orders_7d: 76,
    is_open_now: false,
    created_at: '2024-04-18T08:30:00Z',
    geo: { latitude: 37.781, longitude: -122.412 },
  },
  {
    id: 'store-3',
    name: 'Cedar & Spice',
    primary_cuisine: 'Middle Eastern',
    cuisines: ['Middle Eastern', 'Mediterranean'],
    price_tier: 2,
    rating_avg: 4.8,
    rating_count: 156,
    orders_7d: 164,
    is_open_now: true,
    created_at: '2024-03-05T15:15:00Z',
    geo: { latitude: 37.774, longitude: -122.431 },
  },
]

function App() {
  return (
    <div className="page">
      <header className="top-bar">
        <nav className="tabs" role="tablist" aria-label="Recommendation views">
          <button className="tab is-active" type="button" role="tab" aria-selected="true">
            All
          </button>
          <button className="tab" type="button" role="tab" aria-selected="false">
            My recommendation
          </button>
        </nav>
          <SessionPanel />
      </header>
      <main className="content" role="tabpanel" aria-label="Recommendations">
        <section className="content-shell">
          <div className="store-list">
            {stores.map((store) => (
              <StoreItem key={store.id} store={store} />
            ))}
          </div>
        </section>
      </main>
    </div>
  )
}

export default App
