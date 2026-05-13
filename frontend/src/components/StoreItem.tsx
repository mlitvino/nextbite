import type { Store } from '../types/store'

type StoreItemProps = {
  store: Store
}

function StoreItem({ store }: StoreItemProps) {
  return (
    <article className="store-card">
      <div className="store-main">
        <div>
          <h3 className="store-name">{store.name}</h3>
          <p className="store-meta">{store.primary_cuisine}</p>
        </div>
        <div className="store-rating">
          <span>{store.rating_avg.toFixed(1)}</span>
          <span className="store-count">({store.rating_count})</span>
        </div>
      </div>
      <div className="store-details">
        <span>Price {store.price_tier}</span>
        <span>{store.is_open_now ? 'Open now' : 'Closed'}</span>
        <span>{store.orders_7d} orders (7d)</span>
      </div>
    </article>
  )
}

export default StoreItem
