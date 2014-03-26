package traperr_test

import (
	"github.com/glyn/traperr"
	"testing"
)

func TestMessageCapture(t *testing.T) {

	e := traperr.New("test")

	actual := e.Error()

	if actual != "test" {
		t.Errorf("%q != %q", actual, "test")
	}

}
