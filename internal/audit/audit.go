package audit

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuditEvent struct {
	Event     string                 `json:"event"`
	UserID    uint64                 `json:"user_id,omitempty"`
	Email     string                 `json:"email,omitempty"`
	Success   bool                   `json:"success"`
	ErrorMsg  string                 `json:"error,omitempty"`
	IP        string                 `json:"ip_address,omitempty"`
	UserAgent string                 `json:"user_agent,omitempty"`
	RequestID string                 `json:"request_id,omitempty"`
	Timestamp time.Time              `json:"timestamp"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

var logger *zap.Logger

func Init(l *zap.Logger) {
	logger = l
}

func LogEvent(ctx *gin.Context, e AuditEvent) {
	if logger == nil {
		return
	}

	// Auto-fill context fields if missing
	if e.IP == "" && ctx != nil {
		e.IP = ctx.ClientIP()
	}
	if e.UserAgent == "" && ctx != nil {
		e.UserAgent = ctx.Request.UserAgent()
	}
	if e.RequestID == "" && ctx != nil {
		e.RequestID = ctx.GetString("requestID")
	}
	if e.Timestamp.IsZero() {
		e.Timestamp = time.Now().UTC()
	}

	entry := logger.Info
	if !e.Success {
		entry = logger.Warn
	}

	fields := []zap.Field{
		zap.String("event", e.Event),
		zap.Uint64("user_id", e.UserID),
		zap.String("email", e.Email),
		zap.Bool("success", e.Success),
		zap.String("error", e.ErrorMsg),
		zap.String("ip_address", e.IP),
		zap.String("user_agent", e.UserAgent),
		zap.String("request_id", e.RequestID),
		zap.Time("timestamp", e.Timestamp),
	}

	for k, v := range e.Metadata {
		fields = append(fields, zap.Any(k, v))
	}

	entry("audit event", fields...)
}
