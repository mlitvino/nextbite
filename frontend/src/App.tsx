import SessionPanel from './components/SessionPanel'
import StoreItem from './components/StoreItem'
import { useStores } from './hooks/useStores'
import './App.css'

function App() {
  const { stores, isLoading, error } = useStores()

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
            {isLoading ? <p className="store-status">Loading stores...</p> : null}
            {error ? <p className="store-status">{error}</p> : null}
            {!isLoading && !error
              ? stores.map((store) => <StoreItem key={store.id} store={store} />)
              : null}
          </div>
        </section>
      </main>
    </div>
  )
}

export default App
