package stringtransformer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringTransformation(t *testing.T) {
	testCases := [3]string{"abc", "a4bcd5e", "45"}
	shouldReturn := []string{"abc", "aaaabcddddde", ""}
	var result []string

	for _, v := range testCases {
		functionResult := Transform([]byte(v))
		result = append(result, functionResult.String())
	}

	assert.Equal(t, shouldReturn, result)
}
