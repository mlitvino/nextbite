import { useState } from 'react'
import RecommendationsTabs from './components/RecommendationsTabs'
import SessionPanel from './components/SessionPanel'
import StoreItem from './components/StoreItem'
import { useRecommendations } from './hooks/useRecommendations'
import './App.css'

function App() {
  const { allStores, myStores, isLoading, error } = useRecommendations()
  const [activeTab, setActiveTab] = useState<'all' | 'my'>('all')

  return (
    <div className="page">
      <header className="top-bar">
        <RecommendationsTabs activeTab={activeTab} onChange={setActiveTab} />
        <SessionPanel />
      </header>
      <main className="content" role="tabpanel" aria-label="Recommendations">
        <section className="content-shell">
          <div className="store-list">
            {isLoading ? <p className="store-status">Loading stores...</p> : null}
            {error ? <p className="store-status">{error}</p> : null}
            {!isLoading && !error && activeTab === 'all'
              ? allStores.map((store) => <StoreItem key={store.id} store={store} />)
              : null}
            {!isLoading && !error && activeTab === 'my'
              ? myStores.map((store) => <StoreItem key={store.id} store={store} />)
              : null}
          </div>
        </section>
      </main>
    </div>
  )
}

export default App
