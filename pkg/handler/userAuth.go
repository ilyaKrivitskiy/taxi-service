package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ilyaKrivitskiy/taxi-service/models"
	"net/http"
)

func (h *Handler) userSignUp(c *gin.Context) {
	var input models.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.service.UserAuthorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) userSignIn(c *gin.Context) {

}
