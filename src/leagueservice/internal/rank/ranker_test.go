package rank_test

import (
	"testing"

	"github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/rank"

	"github.com/stretchr/testify/assert"
)

func TestRanker(t *testing.T) {
	t.Run("returns correct rank", func(t *testing.T) {
		ranker := rank.Ranker{}
		ranker.Insert(5)
		ranker.Insert(1)
		ranker.Insert(4)
		ranker.Insert(4)
		ranker.Insert(5)
		ranker.Insert(9)
		ranker.Insert(7)
		ranker.Insert(13)
		ranker.Insert(3)

		rank, ok := ranker.Rank(1)
		assert.True(t, ok)
		assert.Equal(t, 9, rank)

		rank, ok = ranker.Rank(3)
		assert.True(t, ok)
		assert.Equal(t, 8, rank)

		rank, ok = ranker.Rank(4)
		assert.True(t, ok)
		assert.Equal(t, 6, rank)

		rank, ok = ranker.Rank(13)
		assert.True(t, ok)
		assert.Equal(t, 1, rank)

		rank, ok = ranker.Rank(9)
		assert.True(t, ok)
		assert.Equal(t, 2, rank)

		rank, ok = ranker.Rank(111)
		assert.False(t, ok)
	})

	t.Run("returns correct rank with duplicates (dense rank)", func(t *testing.T) {
		ranker := rank.Ranker{}
		ranker.Insert(1)
		ranker.Insert(1)
		ranker.Insert(1)

		ranker.Insert(2)

		ranker.Insert(3)
		ranker.Insert(3)

		ranker.Insert(4)
		ranker.Insert(4)
		ranker.Insert(4)
		ranker.Insert(4)

		ranker.Insert(5)

		rank, ok := ranker.Rank(5)
		assert.True(t, ok)
		assert.Equal(t, 1, rank)

		rank, ok = ranker.Rank(4)
		assert.True(t, ok)
		assert.Equal(t, 2, rank)

		rank, ok = ranker.Rank(3)
		assert.True(t, ok)
		assert.Equal(t, 6, rank)

		rank, ok = ranker.Rank(2)
		assert.True(t, ok)
		assert.Equal(t, 8, rank)

		rank, ok = ranker.Rank(1)
		assert.True(t, ok)
		assert.Equal(t, 9, rank)
	})
}
