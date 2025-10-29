package admin

import (
	"net/http"
	"time"

	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/apperrors"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/dto/showtimes"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/responses"
	"github.com/gin-gonic/gin"
)

func (h *AdminHandler) CreateShowtime(ctx *gin.Context) {
	var req showtimes.CreateShowtimeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		responses.Error(ctx, http.StatusBadRequest, "invalid input: "+err.Error())
		return
	}

	startTime, err := time.Parse(time.RFC3339, req.StartTime)
	if err != nil {
		responses.Error(ctx, http.StatusBadRequest, "invalid start time format")
		return
	}
	endTime, err := time.Parse(time.RFC3339, req.EndTime)
	if err != nil {
		responses.Error(ctx, http.StatusBadRequest, "invalid end time format")
		return
	}
	showtime, err := h.Services.Showtimes.CreateShowtime(ctx.Request.Context(), req.MovieID, req.TheaterID, startTime, endTime, req.BasePrice)
	if err != nil {
		if se, ok := err.(*apperrors.ServiceError); ok {
			responses.Error(ctx, se.Code, se.Message)
			return
		}
		responses.Error(ctx, http.StatusInternalServerError, "Failed to create showtime")
		return
	}

	responses.Success(ctx, gin.H{
		"message":  "showtime created successfully",
		"showtime": showtimes.BuildShowtimeResponse(showtime),
	}, http.StatusCreated)
}
