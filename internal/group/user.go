package group

import (
	"golab8/internal/domain/model"
	"golab8/internal/domain/usecase"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	userUseCase usecase.User
}

func NewUser(userUseCase usecase.User) *User {
	return &User{
		userUseCase: userUseCase,
	}
}

// Get retrieves users based on query parameters.
// @Tags songs
// @Summary Get users
// @Description Retrieve a list of users with optional filters.
// @Param page query int false "Page number"
// @Param limit query int false "Number of users per page"
// @Param age query int false "Filter by age"
// @Param name query string false "Filter by name"
// @Success 200 {array} model.User "List of users"
// @Failure 400 {object} string "Invalid query parameters"
// @Failure 500 {object} string "Internal server error"
// @Router /api/v1/users/ [get]
func (u *User) Get(c *gin.Context) {
	page := 0
	limit := 0
	age := 0

	if pageStr := c.Query("page"); pageStr != "" {
		var err error
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, "invalid page parameter")
			return
		}
	}

	if limitStr := c.Query("limit"); limitStr != "" {
		var err error
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, "invalid limit parameter")
			return
		}
	}

	if ageStr := c.Query("age"); ageStr != "" {
		var err error
		age, err = strconv.Atoi(ageStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, "invalid age parameter")
			return
		}
	}

	name := c.Query("name")

	filter := model.GetUserFilter{
		Name:  name,
		Age:   age,
		Page:  page,
		Limit: limit,
	}

	songs, err := u.userUseCase.Get(c, filter)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "something went wrong")
		return
	}

	c.JSON(http.StatusOK, songs)
}

// GetById retrieves a user by their ID.
// @Tags songs
// @Summary Get user by ID
// @Description Retrieve a user based on the provided user ID.
// @Param id path uint64 true "User ID"
// @Success 200 {object} model.User "User details"
// @Failure 400 {object} string "Invalid user ID"
// @Failure 500 {object} string "Internal server error"
// @Router /api/v1/users/{id} [get]
func (u *User) GetById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "invalid song ID")
		return
	}

	user, err := u.userUseCase.GetById(c, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "something went wrong")
		return
	}

	c.JSON(http.StatusOK, user)
}

// Post adds a new user.
// @Tags songs
// @Summary Add a new user
// @Description This function accepts a JSON object of the user and adds it to the system.
// @Param user body model.AddUser true "User data"
// @Success 200 {integer} int "ID of the added user"
// @Failure 400 {object} string "Incorrect fields in request"
// @Failure 500 {object} string "Internal server error"
// @Router /api/v1/users/ [post]
func (u *User) Post(c *gin.Context) {
	user := model.AddUser{}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Incorrect fields in request")
		return
	}

	id, err := u.userUseCase.Add(c, user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "something went wrong")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// Put updates an existing user.
// @Tags songs
// @Summary Update an existing user
// @Description This function updates user information based on the provided ID and data.
// @Param id path int true "User ID"
// @Param user body model.UpdateUserSwagger true "Updated user data"
// @Success 200 {object} model.User "Updated user information"
// @Failure 400 {object} string "Incorrect fields or invalid user ID"
// @Failure 500 {object} string "Internal server error"
// @Router /api/v1/users/{id} [put]
func (u *User) Put(c *gin.Context) {
	updateUser := model.UpdateUser{}

	if err := c.Bind(&updateUser); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "incorrect fields")
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "invalid song ID")
		return
	}

	updateUser.ID = id

	user, err := u.userUseCase.Update(c, updateUser)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "something went wrong")
		return
	}

	c.JSON(http.StatusOK, user)
}

// Delete removes an existing user.
// @Tags songs
// @Summary Delete an existing user
// @Description This function deletes a user based on the provided ID.
// @Param id path int true "User ID"
// @Success 200 {boolean} bool "Indicates if the user was successfully deleted"
// @Failure 400 {object} string "Invalid user ID"
// @Failure 500 {object} string "Internal server error"
// @Router /api/v1/users/{id} [delete]
func (u *User) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "invalid song ID")
		return
	}

	isDeleted, err := u.userUseCase.Delete(c, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "something went wrong")
	}

	c.JSON(http.StatusOK, isDeleted)
}
