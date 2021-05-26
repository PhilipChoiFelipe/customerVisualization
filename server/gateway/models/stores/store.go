package stores

import "fmt"

// The space in which the user provide product or services.
// Not to be confused with database storage.
type Store struct {
	ID        int64  `json:"id"`
	UserID		int64 `json:"userId"`
	StoreName     string `json:"storeName"`
	StoreLocation string `json:"storeLocation"`
}

type StoreUpdate struct {
	StoreName     string `json:"storeName"`
	StoreLocation string `json:"storeLocation"`
}

func (s *Store) ApplyUpdates(updates *StoreUpdate) error {
	if updates.StoreName == "" || updates.StoreLocation == "" {
		return fmt.Errorf("empty updating value")
	}
	s.StoreName = updates.StoreName
	s.StoreLocation = updates.StoreLocation
	return nil
}