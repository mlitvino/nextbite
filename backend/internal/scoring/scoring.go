package scoring

import (
	"math"

	"github.com/mlitvino/nextbite/backend/internal/models"
)

type Scorer interface {
	Score(store models.Store, user models.User) float64
}

type DefaultScorer struct{}

func NewDefaultScorer() *DefaultScorer {
	return &DefaultScorer{}
}

func (s *DefaultScorer) Score(store models.Store, user models.User) float64 {
	score := store.RatingAvg*10 + math.Log1p(float64(store.Orders7D))
	if store.IsOpenNow {
		score += 0.5
	}
	return score
}
