package pvztest

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// RequireEqualTime asserts that the postgres timestamp is within its accuracy limitations (micro-seconds).
func RequireEqualTime(t *testing.T, local time.Time, postgres time.Time, msgAndArgs ...interface{}) {
	delta := local.Sub(postgres)
	require.LessOrEqual(t, delta.Nanoseconds(), int64(1000), msgAndArgs...)
	require.GreaterOrEqual(t, delta.Nanoseconds(), int64(0), msgAndArgs...)
}
