package function

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunction(t *testing.T) {
	assert := assert.New(t)

	t.Run("deveria retornar dois números somados ao chamar a função de soma", func(t *testing.T) {
		resultado := Soma(10, 10)
		assert.Equal(20, resultado)
	})

	t.Run("deveria retornar zero se o number2 for zero na função de soma", func(t *testing.T) {
		resultado := Soma(10, 0)
		assert.Equal(0, resultado)
	})

	t.Run("deveria retornar dois números substraidos ao chamar a função de substração", func(t *testing.T) {
		resultado := Substracao(10, 10)
		assert.Equal(0, resultado)
	})

	t.Run("deveria retornar zero se o number1 for zero na função de substração", func(t *testing.T) {
		resultado := Substracao(0, 10)
		assert.Equal(resultado, 0)
	})

	t.Run("deveria retornar zero se o number2 for zero na função de substração", func(t *testing.T) {
		resultado := Substracao(10, 0)
		assert.Equal(resultado, 0)
	})
}
