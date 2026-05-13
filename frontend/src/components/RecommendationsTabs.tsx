type TabKey = 'all' | 'my'

type RecommendationsTabsProps = {
  activeTab: TabKey
  onChange: (tab: TabKey) => void
}

function RecommendationsTabs({ activeTab, onChange }: RecommendationsTabsProps) {
  return (
    <nav className="tabs" role="tablist" aria-label="Recommendation views">
      <button
        className={`tab ${activeTab === 'all' ? 'is-active' : ''}`}
        type="button"
        role="tab"
        aria-selected={activeTab === 'all'}
        onClick={() => onChange('all')}
      >
        All
      </button>
      <button
        className={`tab ${activeTab === 'my' ? 'is-active' : ''}`}
        type="button"
        role="tab"
        aria-selected={activeTab === 'my'}
        onClick={() => onChange('my')}
      >
        My recommendation
      </button>
    </nav>
  )
}

export default RecommendationsTabs
