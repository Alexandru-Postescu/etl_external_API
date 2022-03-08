package etl

// Interfaces are describing the behaviour of the application. We can distinguish three different behaviours based on our application
import (
	"context"

	"example.com/go-api/model"
)

type Extractor interface {
	Extract(ctx context.Context, numPeople int) ([]model.Person, error)
}

type Transformer interface {
	Transform(input []model.Person) (map[string][]model.Person, error)
}

type Loader interface {
	Load(m map[string][]model.Person) error
}
