package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	for _, tc := range []struct {
		name   string
		from   string
		offset int64
		limit  int64
		result string
	}{
		{
			name:   "out_offset0_limit0.txt",
			from:   "testdata/input.txt",
			offset: 0,
			limit:  0,
			result: "testdata/out_offset0_limit0.txt",
		},
		{
			name:   "out_offset0_limit10.txt",
			from:   "testdata/input.txt",
			offset: 0,
			limit:  10,
			result: "testdata/out_offset0_limit10.txt",
		},
		{
			name:   "out_offset0_limit1000.txt",
			from:   "testdata/input.txt",
			offset: 0,
			limit:  1000,
			result: "testdata/out_offset0_limit1000.txt",
		},
		{
			name:   "out_offset0_limit10000.txt",
			from:   "testdata/input.txt",
			offset: 0,
			limit:  10000,
			result: "testdata/out_offset0_limit10000.txt",
		},
		{
			name:   "testdata/out_offset100_limit1000.txt",
			from:   "testdata/input.txt",
			offset: 100,
			limit:  1000,
			result: "testdata/out_offset100_limit1000.txt",
		},
		{
			name:   "out_offset6000_limit1000.txt",
			from:   "testdata/input.txt",
			offset: 6000,
			limit:  1000,
			result: "testdata/out_offset6000_limit1000.txt",
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tempFile, err := filepath.Abs("testdata")
			require.NoError(t, err)
			fileTo, err := os.CreateTemp(tempFile, "test_case")
			require.NoError(t, err)
			expResFile, err := os.Open(tc.result)
			require.NoError(t, err)
			defer func() {
				_ = fileTo.Close()
				_ = expResFile.Close()
				_ = os.Remove(fileTo.Name())
			}()
			expRes, err := ioutil.ReadAll(expResFile)
			require.NoError(t, err)

			err = Copy(tc.from, fileTo.Name(), tc.offset, tc.limit)
			require.NoError(t, err)

			res, err := ioutil.ReadAll(fileTo)
			require.NoError(t, err)
			assert.Equal(t, expRes, res)
		})
	}
}
