package pvztest

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// RequireEqualTime asserts that the postgres timestamp is within its accuracy limitations (micro-seconds).
func RequireEqualTime(t *testing.T, local time.Time, postgres time.Time, msgAndArgs ...interface{}) {
	require.InDelta(t, local.Nanosecond(), postgres.Nanosecond(), 1001, msgAndArgs...)
}
