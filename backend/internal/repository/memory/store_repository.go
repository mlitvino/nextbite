package memory

import (
	"context"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/mlitvino/nextbite/backend/internal/models"
	"github.com/mlitvino/nextbite/backend/internal/repository"
)

type StoreRepository struct {
	mu     sync.RWMutex
	stores []models.Store
}

func NewStoreRepository() *StoreRepository {
	return &StoreRepository{stores: []models.Store{}}
}

func (s *StoreRepository) List(ctx context.Context) ([]models.Store, error) {
	_ = ctx
	s.mu.RLock()
	defer s.mu.RUnlock()
	return append([]models.Store(nil), s.stores...), nil
}

func (s *StoreRepository) Create(ctx context.Context, item models.Store) (models.Store, error) {
	_ = ctx
	if item.ID == "" {
		item.ID = uuid.NewString()
	}
	s.mu.Lock()
	s.stores = append(s.stores, item)
	s.mu.Unlock()
	return item, nil
}

func (s *StoreRepository) GetByID(ctx context.Context, id string) (models.Store, error) {
	_ = ctx
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, item := range s.stores {
		if item.ID == id {
			return item, nil
		}
	}
	return models.Store{}, store.ErrStoreNotFound
}

func (s *StoreRepository) Update(ctx context.Context, id string, item models.Store) (models.Store, error) {
	_ = ctx
	s.mu.Lock()
	defer s.mu.Unlock()
	for i := range s.stores {
		if s.stores[i].ID == id {
			item.ID = id
			s.stores[i] = item
			return item, nil
		}
	}
	return models.Store{}, store.ErrStoreNotFound
}

func (s *StoreRepository) Delete(ctx context.Context, id string) error {
	_ = ctx
	s.mu.Lock()
	defer s.mu.Unlock()
	for i := range s.stores {
		if s.stores[i].ID == id {
			s.stores = append(s.stores[:i], s.stores[i+1:]...)
			return nil
		}
	}
	return store.ErrStoreNotFound
}

func (s *StoreRepository) LoadFromCSV(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	start := 0
	if len(records) > 0 && len(records[0]) > 0 && strings.EqualFold(records[0][0], "id") {
		start = 1
	}

	items := make([]models.Store, 0, len(records))
	for i := start; i < len(records); i++ {
		row := records[i]
		if len(row) < 12 {
			return fmt.Errorf("invalid store csv row %d: expected 12 columns", i+1)
		}

		priceTier, err := strconv.Atoi(row[4])
		if err != nil {
			return fmt.Errorf("invalid price_tier at row %d: %w", i+1, err)
		}
		ratingAvg, err := strconv.ParseFloat(row[5], 64)
		if err != nil {
			return fmt.Errorf("invalid rating_avg at row %d: %w", i+1, err)
		}
		ratingCount, err := strconv.Atoi(row[6])
		if err != nil {
			return fmt.Errorf("invalid rating_count at row %d: %w", i+1, err)
		}
		orders7D, err := strconv.Atoi(row[7])
		if err != nil {
			return fmt.Errorf("invalid orders_7d at row %d: %w", i+1, err)
		}
		isOpenNow, err := strconv.ParseBool(row[8])
		if err != nil {
			return fmt.Errorf("invalid is_open_now at row %d: %w", i+1, err)
		}
		createdAt, err := time.Parse(time.RFC3339, row[9])
		if err != nil {
			return fmt.Errorf("invalid created_at at row %d: %w", i+1, err)
		}
		latitude, err := strconv.ParseFloat(row[10], 64)
		if err != nil {
			return fmt.Errorf("invalid latitude at row %d: %w", i+1, err)
		}
		longitude, err := strconv.ParseFloat(row[11], 64)
		if err != nil {
			return fmt.Errorf("invalid longitude at row %d: %w", i+1, err)
		}

		id := row[0]
		if id == "" {
			id = uuid.NewString()
		}

		var cuisines []string
		if row[3] != "" {
			cuisines = strings.Split(row[3], "|")
		}

		items = append(items, models.Store{
			ID:             id,
			Name:           row[1],
			PrimaryCuisine: row[2],
			Cuisines:       cuisines,
			PriceTier:      priceTier,
			RatingAvg:      ratingAvg,
			RatingCount:    ratingCount,
			Orders7D:       orders7D,
			IsOpenNow:      isOpenNow,
			CreatedAt:      createdAt,
			Geo: models.Geo{
				Latitude:  latitude,
				Longitude: longitude,
			},
		})
	}

	s.mu.Lock()
	s.stores = append(s.stores, items...)
	s.mu.Unlock()
	return nil
}
