package filter

import (
	"testing"
)

func TestFilter(t *testing.T) {
	fm := NewFilterManager(new(Target))
	fm.SetFilter(new(AuthenticationFilter))
	fm.SetFilter(new(DebugFilter))

	c := NewClient()
	c.SetFilterManager(fm)
	c.SendRequest("home")
}
