package helpers_test

import (
	"testing"

	"github.com/centrypoint/refrigerator/back/go/main/server/helpers"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	Test     string
	Expected int
}

func TestRepresentShelflife(t *testing.T) {
	var cases = []testCase{
		{
			Test:     "60 дней",
			Expected: 60,
		},
		{
			Test:     "1 день",
			Expected: 1,
		},
		{
			Test:     "12 месяцев",
			Expected: 365,
		},
		{
			Test:     "6 месяцев",
			Expected: 182,
		},
		{
			Test:     "1 год",
			Expected: 365,
		},
	}

	for _, v := range cases {
		assert.Equal(t, v.Expected, helpers.RepresentShelflife(v.Test))
	}
}
