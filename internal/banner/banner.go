package banner

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/dimiro1/banner"
)

func InitBanner() {
	bannerBytes, _ := ioutil.ReadFile("banner.txt")
	banner.Init(os.Stdout, true, false, strings.NewReader(string(bannerBytes)))
}
