package recommendation

import (
	"sort"

	"github.com/mlitvino/nextbite/backend/internal/models"
	"github.com/mlitvino/nextbite/backend/internal/scoring"
)

type Engine struct {
	scorer scoring.Scorer
}

func NewEngine(scorer scoring.Scorer) *Engine {
	return &Engine{scorer: scorer}
}

type scoredStore struct {
	store models.Store
	score float64
}

func (e *Engine) Recommend(user models.User, stores []models.Store, limit int) []models.Store {
	if limit <= 0 {
		limit = len(stores)
	}

	scored := make([]scoredStore, 0, len(stores))
	for _, store := range stores {
		scored = append(scored, scoredStore{store: store, score: e.scorer.Score(store, user)})
	}

	sort.SliceStable(scored, func(i, j int) bool {
		if scored[i].score == scored[j].score {
			if scored[i].store.Name == scored[j].store.Name {
				return scored[i].store.ID < scored[j].store.ID
			}
			return scored[i].store.Name < scored[j].store.Name
		}
		return scored[i].score > scored[j].score
	})

	if limit > len(scored) {
		limit = len(scored)
	}

	out := make([]models.Store, 0, limit)
	for i := 0; i < limit; i++ {
		out = append(out, scored[i].store)
	}

	return out
}
