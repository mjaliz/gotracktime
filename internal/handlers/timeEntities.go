package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/mjaliz/gotracktime/internal/models"
	"github.com/mjaliz/gotracktime/internal/utils"
	"net/http"
)

func (repo *DBRepo) CreateTimeEntity(c *gin.Context) {
	user := getUserFromContext(c)
	var timeEntityInput models.TimeEntityInput
	if err := c.ShouldBindJSON(&timeEntityInput); err != nil {
		validationErrs := utils.ParseValidationError(err)
		utils.FailedResponse(c, http.StatusBadRequest, validationErrs, "")
		return
	}
	timeEntityInput.UserID = user.UserID
	timeEntityDB, err := repo.DB.InsertTimeEntity(timeEntityInput)
	if err != nil {
		utils.FailedResponse(c, http.StatusInternalServerError, nil, "")
		return
	}
	var output models.TimeEntityOutput
	output.CreatedAt = timeEntityDB.CreatedAt
	output.StartedAt = timeEntityDB.StartedAt
	output.DescriptionID = timeEntityDB.DescriptionID
	output.ProjectID = timeEntityDB.ProjectID
	utils.SuccessResponse(c, http.StatusCreated, output, "")
}
