package types

import (
	"encoding/json"

	"github.com/soylok/filmshelf/internal/core/domain"
)

type Nullable[T any] struct {
	domain.Nullable[T]
}

func (n *Nullable[T]) UnmarshalJSON(data []byte) error {
	var value *T
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	n.Value = value
	n.Set = true
	return nil
}
