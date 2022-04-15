package main

import (
	"elasticsearch-document-cleaner/internal/banner"
	"elasticsearch-document-cleaner/internal/logging"
	_ "go.uber.org/automaxprocs"
	"go.uber.org/zap"
)

var (
	logger *zap.Logger
)

func init() {
	logger = logging.GetLogger()
	banner.InitBanner()
}

func main() {
	logger.Info("started elasticsearch-document-cleaner")
}
