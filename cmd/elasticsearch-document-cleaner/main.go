package main

import (
	"elasticsearch-document-cleaner/internal/logging"
	"io/ioutil"
	"os"
	"strings"

	"github.com/dimiro1/banner"
	"go.uber.org/zap"
)

var (
	logger *zap.Logger
)

func init() {
	logger = logging.GetLogger()

	bannerBytes, _ := ioutil.ReadFile("banner.txt")
	banner.Init(os.Stdout, true, false, strings.NewReader(string(bannerBytes)))
}

func main() {
	logger.Info("started elasticsearch-document-cleaner")
}
