package courseschedule

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanFinish(t *testing.T) {
	t.Run("impossible case", func(t *testing.T) {
		numCourses := 3
		prereqs := [][]int{{0, 1}, {1, 2}, {2, 0}}
		assert.False(t, canFinish(numCourses, prereqs))
	})
	t.Run("possible case", func(t *testing.T) {
		numCourses := 3
		prereqs := [][]int{{0, 1}, {1, 2}}
		assert.True(t, canFinish(numCourses, prereqs))
	})
}
