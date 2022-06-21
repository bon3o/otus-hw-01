package main

import (
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadEnvValue(t *testing.T) {
	t.Run("test BAR", func(t *testing.T) {
		fileName := filepath.Join("testdata", "env", "BAR")
		envValue, err := processFile(fileName)
		require.Equal(t, EnvValue{Value: "bar", NeedRemove: false}, envValue)
		require.NoError(t, err)
	})
	t.Run("test EMPTY", func(t *testing.T) {
		fileName := filepath.Join("testdata", "env", "EMPTY")
		envValue, err := processFile(fileName)
		require.Equal(t, EnvValue{Value: "", NeedRemove: false}, envValue)
		require.NoError(t, err)
	})
	t.Run("test FOO", func(t *testing.T) {
		fileName := filepath.Join("testdata", "env", "FOO")
		envValue, err := processFile(fileName)
		require.Equal(t, EnvValue{Value: "   foo\nwith new line", NeedRemove: false}, envValue)
		require.NoError(t, err)
	})
	t.Run("test HELLO", func(t *testing.T) {
		fileName := filepath.Join("testdata", "env", "HELLO")
		envValue, err := processFile(fileName)
		require.Equal(t, EnvValue{Value: "\"hello\"", NeedRemove: false}, envValue)
		require.NoError(t, err)
	})
	t.Run("test UNSET", func(t *testing.T) {
		fileName := filepath.Join("testdata", "env", "UNSET")
		envValue, err := processFile(fileName)
		require.Equal(t, EnvValue{Value: "", NeedRemove: true}, envValue)
		require.NoError(t, err)
	})
}

func TestReadDir(t *testing.T) {
	t.Run("test incorrect env", func(t *testing.T) {
		dir, err := os.MkdirTemp("", "incorrect")
		if err != nil {
			log.Fatal(err)
		}
		defer os.RemoveAll(dir) // clean up

		file := filepath.Join(dir, "key=")
		if err := os.WriteFile(file, []byte("value"), 0o666); err != nil {
			log.Fatal(err)
		}

		environment, err := ReadDir(dir)
		require.Nil(t, environment)
		require.ErrorIs(t, ErrInvalidFileName, err)
	})
	t.Run("test env", func(t *testing.T) {
		dir := filepath.Join("testdata", "env")
		environment, err := ReadDir(dir)
		desiredResult := Environment{
			"BAR":   {Value: "bar"},
			"EMPTY": {Value: ""},
			"FOO":   {Value: "   foo\nwith new line"},
			"HELLO": {Value: "\"hello\""},
			"UNSET": {Value: "", NeedRemove: true},
		}
		require.Equal(t, desiredResult, environment)
		require.NoError(t, err)
	})
}
