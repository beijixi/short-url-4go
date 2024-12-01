package interfaces

import "net/http"

type IAccessLogService interface {
	// BatchQueryHits(shortIDs []string) (map[string]int64, error)
	RecordAccessLog(shortID string, headers http.Header) error
}
