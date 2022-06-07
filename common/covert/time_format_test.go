package covert

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestTimeFormatMMDD(t *testing.T) {
	timex := "2018-11-25"
	// to time.Time
	parse, err := time.Parse("2006-01-02", timex)
	require.NoError(t, err)
	ret := TimeFormatMMDD(parse)
	require.Equal(t, "11-25", ret)
}
