package main

import (
	"errors"
	"io"
	"os"

	"github.com/cheggaaa/pb"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	var err error
	src, err := os.Open(fromPath)
	if err != nil {
		return err
	}
	defer src.Close()
	fromStat, err := src.Stat()
	if err != nil {
		return err
	}
	fromLen := fromStat.Size()
	if offset > 0 {
		if offset > fromLen {
			return ErrOffsetExceedsFileSize
		}
		src.Seek(offset, io.SeekStart)
		fromLen -= offset
	}
	if fromLen == 0 {
		return ErrUnsupportedFile
	}
	if limit == 0 || limit > fromLen {
		limit = fromLen
	}
	dst, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer dst.Close()
	if err != nil {
		return err
	}
	bar := pb.New(int(limit)).SetUnits(pb.U_BYTES)
	bar.Start()
	_, err = io.CopyN(dst, bar.NewProxyReader(src), limit)
	bar.Finish()
	return err
}
