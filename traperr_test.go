package traperr_test

import (
	"github.com/glyn/traperr"
	"strings"
	"testing"
)

const testMessage = "test"

func TestMessageCapture(t *testing.T) {

	e := traperr.New(testMessage)

	actual := e.Error()

	if actual != testMessage {
		t.Errorf("%q != %q", actual, testMessage)
	}
}

func TestStackTraceCapture(t *testing.T) {

	const stackPortion = "traperr_test.TestStackTraceCapture"

	e := traperr.New(testMessage)

	stack := e.StackTrace()

	if !strings.Contains(stack, stackPortion) {
		t.Errorf("%q does not contain %q", stack, stackPortion)
	}
}
