package main

import (
	"flag"

	"github.com/VictoriaMetrics/VictoriaMetrics/lib/logger"
)

var (
	from, to      string
	limit, offset int64
)

func init() {
	flag.StringVar(&from, "from", "", "file to read from")
	flag.StringVar(&to, "to", "", "file to write to")
	flag.Int64Var(&limit, "limit", 0, "limit of bytes to copy")
	flag.Int64Var(&offset, "offset", 0, "offset in input file")
}

func main() {
	flag.Parse()
	if from == "" || to == "" {
		logger.Fatalf("from or to path is not passed")
	}
	err := Copy(from, to, offset, limit)
	if err != nil {
		logger.Fatalf("error while copying file: %s", err)
	}
}
