package destination_city

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestDestCity(t *testing.T){
    t.Run("Sao Paulo", func(t *testing.T) {
        paths := [][]string{{"London","New York"},{"New York","Lima"},{"Lima","Sao Paulo"}}
        assert.Equal(t, "Sao Paulo", destCity(paths))
    })
    t.Run("A", func(t *testing.T) {
        paths := [][]string{{"B","C"},{"D","B"},{"C","A"}}
        assert.Equal(t, "A", destCity(paths))
    })
    t.Run("Z", func(t *testing.T) {
        paths := [][]string{{"A","Z"}}
        assert.Equal(t, "Z", destCity(paths))
    })
}
