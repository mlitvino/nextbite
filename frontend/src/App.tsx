import SessionPanel from './components/SessionPanel'
import './App.css'

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
        <div className="content-shell" />
      </main>
    </div>
  )
}

export default App
