package greeting_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/charmingruby/rib/greeting"
)

func TestGreet(t *testing.T) {
	name := "rib"

	msg := greeting.Greet(name)
	expectedMsg := fmt.Sprintf("hi, %s!", name)

	assert.Equal(t, expectedMsg, msg)
}
