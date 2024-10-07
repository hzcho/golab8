package group

import (
	"golab8/internal/domain/model"
	"golab8/internal/domain/usecase"

	"net/http"

	"github.com/gin-gonic/gin"
)

type Auth struct {
	authUseCase usecase.Auth
}

func NewAuth(authUseCase usecase.Auth) *Auth {
	return &Auth{
		authUseCase: authUseCase,
	}
}

// @Summary Sign in user
// @Description Authenticate user and return a JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param request body model.SignInBody true "User login and password"
// @Success 200 {object} model.SignInResponse "token"
// @Failure 400 {object} string "incorrect request structure"
// @Failure 500 {object} string "couldn't create the token"
// @Router /api/v1/auth/login [post]
func (g *Auth) SignIn(c *gin.Context) {
	var request model.SignInBody

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, "incorrect request structure")
		return
	}

	token, err := g.authUseCase.GenerateToken(c.Request.Context(), request.Login, request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "couldn't create the token")
		return
	}

	c.JSON(http.StatusOK, model.SignInResponse{
		Token: token,
	})
}

// @Summary Register a new user
// @Description Create a new user account and return the user ID
// @Tags auth
// @Accept json
// @Produce json
// @Param request body model.RegisterBody true "User login and password"
// @Success 200 {object} model.RegisterResponse "account id"
// @Failure 400 {object} string "incorrect request structure"
// @Failure 500 {object} string "couldn't create the user"
// @Router /api/v1/auth/register [post]
func (g *Auth) Register(c *gin.Context) {
	var request model.RegisterBody

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, "incorrect request structure")
		return
	}

	accountId, err := g.authUseCase.CreateAccount(c.Request.Context(), request.Login, request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "couldn't create the user")
		return
	}

	c.JSON(http.StatusOK, model.RegisterResponse{
		AccountId: accountId,
	})
}
