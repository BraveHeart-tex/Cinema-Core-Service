package audit

import (
	"github.com/gin-gonic/gin"
)

type AdminAuditParams struct {
	Action           string
	TargetUserID     uint64
	PerformedByID    uint64
	PerformedByEmail string
	Success          bool
	ErrorMsg         string
	Metadata         map[string]any
}

func LogAdminAction(ctx *gin.Context, params AdminAuditParams) {
	eventName := "admin." + params.Action
	if !params.Success {
		eventName += ".failure"
	} else {
		eventName += ".success"
	}

	if params.Metadata == nil {
		params.Metadata = make(map[string]any)
	}
	params.Metadata["performedByID"] = params.PerformedByID
	params.Metadata["performedByEmail"] = params.PerformedByEmail

	LogEvent(ctx, AuditEvent{
		Event:    eventName,
		UserID:   params.TargetUserID,
		Success:  params.Success,
		ErrorMsg: params.ErrorMsg,
		Metadata: params.Metadata,
	})
}
