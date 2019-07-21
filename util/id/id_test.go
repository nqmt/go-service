package id

import (
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)


func TestId(t *testing.T) {
	t.Run(" singleton", func(t *testing.T) {
		tid1 := New()
		tid2 := New()

		require.Equal(t, reflect.ValueOf(tid1).Pointer(), reflect.ValueOf(tid2).Pointer())
	})

	t.Run("always generate new id", func(t *testing.T) {
		tid := New()

		id1 := tid.Gen()
		id2 := tid.Gen()

		require.NotEqual(t, id1, id2)
	})

	t.Run("freeze id", func(t *testing.T) {
		tid := New()

		tid.Freeze("11")

		id1 := tid.Gen()
		id2 := tid.Gen()

		expect := "11"

		require.Equal(t, expect, id1)
		require.Equal(t, expect, id2)
	})
}
