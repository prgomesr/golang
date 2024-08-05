package teste_area

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Retangulo", func(t *testing.T) {
		ret := Retangulo{10, 12}
		areaEsperada := float64(120)
		areaRecebida := ret.Area()

		assert.Equal(t, areaEsperada, areaRecebida)
	})

	t.Run("Circulo", func(t *testing.T) {
		circ := Circulo{10}
		areaEsperada := math.Pi * 100
		areaRecebida := circ.Area()

		assert.Equal(t, areaEsperada, areaRecebida)
	})
}
