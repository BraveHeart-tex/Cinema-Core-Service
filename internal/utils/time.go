package utils

import "time"

func ToRFC3339Ptr(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.Format(time.RFC3339)
}
