package brand

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestQinghaiBrandTask11(t *testing.T) {
	s := NewService(NewRegistry(), time.Now)
	lot := IngredientLot{MinTemperature: 0, MaxTemperature: 4, Temperatures: []float64{1, 2, 3}}
	require.NoError(t, s.CheckColdChain(context.Background(), lot))
}
