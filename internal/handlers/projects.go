package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/mjaliz/gotracktime/internal/models"
	"github.com/mjaliz/gotracktime/internal/utils"
	"net/http"
)

func (repo *DBRepo) CreateProject(c *gin.Context) {
	user := getUserFromContext(c)
	var projectInput models.ProjectInput
	if err := c.ShouldBindJSON(&projectInput); err != nil {
		validationErrs := utils.ParseValidationError(err)
		utils.FailedResponse(c, http.StatusBadRequest, validationErrs, "")
		return
	}
	projectInput.UserID = user.UserID
	projectDB, err := repo.DB.InsertProject(projectInput)
	if err != nil {
		utils.FailedResponse(c, http.StatusInternalServerError, nil, "")
		return
	}
	var output models.ProjectOutput
	output.CreatedAt = projectDB.CreatedAt
	output.Title = projectDB.Title
	utils.SuccessResponse(c, http.StatusCreated, output, "")
}
