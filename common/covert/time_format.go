package covert

import "time"

// TimeFormatMMDD covert time to format "mm-DD"
func TimeFormatMMDD(t time.Time) string {
	return t.Format("01-02")
}
