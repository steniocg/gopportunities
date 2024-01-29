package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/steniocg/gopportunities/schemas"
)

func CreateOpeningHandler(ctx *gin.Context) {

	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
	}

	opening schemas.Opening := schemas.Opening {
		Role: request.Role,
		Company: request.Company,
		Lcation: request.Location,
		Remote: *request.Remote,
		Link: request.Link,
		Salary: request.Salary,
	} 

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("erro creating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error criating opening database")
		return
	}

	sandSuccess(ctx, "create-opening", opening)
}
