package banner

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/dimiro1/banner"
)

func TestInitBanner(t *testing.T) {
	bannerBytes, err := ioutil.ReadFile("../../banner.txt")
	if err != nil {
		t.Errorf("%v\n", err.Error())
	}

	banner.Init(os.Stdout, true, false, strings.NewReader(string(bannerBytes)))
}
