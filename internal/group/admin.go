package group

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Admin struct{}

func NewAdmin() *Admin {
	return &Admin{}
}

// @Summary Get admin information
// @Description Retrieves admin status and greets the admin or informs if the user is not an admin.
// @Tags Admin
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {string} string "Greeting message"
// @Failure 500 {string} string "Something went wrong."
// @Router /api/v1/admins/ [get]
func (a *Admin) Get(c *gin.Context) {
	admin, exists := c.Get("admin")
	if !exists {
		c.JSON(http.StatusInternalServerError, "Something went wrong.")
		return
	}

	isAdmin := admin.(bool)

	if isAdmin {
		c.JSON(http.StatusOK, "привет, admin")
		return
	} else {
		c.JSON(http.StatusOK, "ты не admin!")
		return
	}
}
