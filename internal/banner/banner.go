package banner

import (
	"os"
	"strings"

	"github.com/dimiro1/banner"
)

func InitBanner() {
	bannerBytes, _ := os.ReadFile("banner.txt")
	banner.Init(os.Stdout, true, false, strings.NewReader(string(bannerBytes)))
}
