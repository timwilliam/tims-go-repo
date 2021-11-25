package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	// before
	fmt.Println("*** Before running unit test")

	m.Run()

	// after
	fmt.Println("*** After running unit test")
}

func TestSayHello(t *testing.T) {
	var result string

	result = SayHello("Tim")
	assert.Equal(t, "Hello Tim", result)

	result = SayHello("William")
	assert.Equal(t, "Hello William", result)
}

func TestSkip(t *testing.T) {
	// $ go test -v -run=TestSkip ./...

	if runtime.GOOS == "darwin" {
		t.Skip("Test cannot be run on MacOS")
	}

	result := SayHello("Tim")
	assert.Equal(t, "Hello Tim", result)
}

func TestSubTest(t *testing.T) {
	// $ go test -v -run=TestSubTest/Subtest-1 ./...
	// $ go test -v -run=/Subtest-1 ./...

	t.Run("Subtest-1", func(t *testing.T) {
		result := SayHello("Tim")
		assert.Equal(t, "Hello Tim", result)
	})
	t.Run("Subtest-2", func(t *testing.T) {
		result := SayHello("William")
		assert.Equal(t, "Hello William", result)
	})
}

func TestTableSayHello(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Tim",
			request:  "Tim",
			expected: "Hello Tim",
		},
		{
			name:     "William",
			request:  "William",
			expected: "Hello William",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := SayHello(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}
