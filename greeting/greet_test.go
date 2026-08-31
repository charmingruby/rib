package greeting_test

import (
	"fmt"
	"testing"

	"github.com/charmingruby/rib/greeting"
	"github.com/stretchr/testify/assert"
)

func TestGreet(t *testing.T) {
	name := "rib"

	msg := greeting.Greet(name)
	expectedMsg := fmt.Sprintf("hi, %s!", name)

	assert.Equal(t, msg, expectedMsg)
}
