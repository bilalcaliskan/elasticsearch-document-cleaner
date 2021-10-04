package banner

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/dimiro1/banner"
)

func TestInitBanner(t *testing.T) {
	t.Log("starting banner creation test based on ../../banner.txt")

	bannerBytes, err := ioutil.ReadFile("../../banner.txt")
	if err != nil {
		t.Errorf("%v\n", err.Error())
		return
	}

	banner.Init(os.Stdout, true, false, strings.NewReader(string(bannerBytes)))
}
