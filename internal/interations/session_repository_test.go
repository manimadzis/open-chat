package interations

import (
	"testing"
)

func TestSessionRepo(t *testing.T) {
	pool := connect()
	create(pool)
	// drop(pool)
}
