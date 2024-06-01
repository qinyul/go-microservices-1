package controllers

import(
	"barqi.com/user/utils"
	"barqi.com/user/database"
	"barqi.com/user/auth"
	"barqi.com/user/common"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type User struct {
	utils utils.Utils
	userAuth auth.User
}

// Authenticate godoc
// @Summary Check user authentication
// @Description Authenticate user
// @Tags admin
// @Security ApiKeyAuth
// @Accept  multipart/form-data
// @Param username formData string true "Username"
// @Param password formData string true "Password" format(password)
// @Failure 401 {object} database.Error
// @Failure 500 {object} database.Error
// @Success 200 {object} database.Token
// @Router /admin/auth [post]
func (u *User) Authenticate(ctx *gin.Context){
    username :=  ctx.PostForm("username")
    password :=  ctx.PostForm("password")

	if username != "" && password != "" {
		var err error
		log.Debug("starting login process")
		_, err = u.userAuth.Login(username,password)

		if err != nil {
			if err == mongo.ErrNoDocuments || err.Error() == common.WrongUsernameOrPassword {
				ctx.JSON(http.StatusNotFound,database.Error{common.StatusCodeNotFound,common.WrongUsernameOrPassword})	
			}else{
				ctx.JSON(http.StatusInternalServerError,database.Error{common.StatusCodeUnknown,err.Error()})
			}
		}else{
			var tokenString string
			tokenString, err = u.utils.GenerateJWT(username,"")
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, database.Error{common.StatusCodeUnknown,err.Error()})
				log.Debug("[ERROR]: ",err)
				return
			}

			token := database.Token{tokenString}
			ctx.JSON(http.StatusOK,token)
		}
	}else{
		log.Debug("password and username is empty, returning unauthorize status")
		ctx.JSON(http.StatusUnauthorized, database.Error{common.StatusCodeUnknown,common.UnautohrizeErrorMessage})
	}

}

// AddUser godoc
// @Summary Add a new user
// @Description Add a new user
// @Tags user
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Token"
// @Param user body database.AddUser true "Add user"
// @Failure 500 {object} database.Error
// @Failure 400 {object} database.Error
// @Success 200 {object} database.Message
// @Router /users [post]
func (u *User) AddUser(ctx *gin.Context){
	var addUser database.AddUser
	if err := ctx.ShouldBindJSON(&addUser); err != nil {
		ctx.JSON(http.StatusBadRequest, database.Error{common.StatusCodeBadRequest, err.Error()})
		return
	}
	err := u.userAuth.Insert(addUser.Username, addUser.Password)
	if err != nil {
		if mongo.IsDuplicateKeyError(err){
			log.Debug("User already exist")
			ctx.JSON(http.StatusBadRequest, database.Error{common.StatusCodeBadRequest, common.UserAlreadyExist})
			return
		}
		ctx.JSON(http.StatusInternalServerError, database.Error{common.StatusCodeUnknown, err.Error()})
		log.Debug("[ERROR]: ",err)
	}else {
		ctx.JSON(http.StatusOK, database.Message{common.UserCreated})
		log.Debug("Registered a new user: ", addUser.Username)
	}
}

// ListUsers godoc
// @Summary List all existing users
// @Description List all existing users
// @Tags user
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Token"
// @Failure 500 {object} database.Error
// @Success 200 {array} database.UserInformation
// @Router /users/list [get]
func (u *User) ListUsers(ctx *gin.Context){
	var users []database.UserInformation
	var err error
	users, err = u.userAuth.GetAll(common.EmptyString)

	if err != nil {
		log.Debug("[ERROR]: ",err.Error())
		ctx.JSON(http.StatusInternalServerError, database.Error{common.StatusCodeUnknown,err.Error()})
	}else {
		ctx.JSON(http.StatusOK,users)
	}
}

// GetUserByID godoc
// @Summary Get a user by ID
// @Description Get a user by ID
// @Tags user
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Token"
// @Param id path string true "User ID"
// @Failure 500 {object} database.Error
// @Failure 400 {object} database.Error
// @Success 200 {object} database.UserInformation
// @Router /users/detail/{id} [get]
func (u *User) GetUserByID(ctx *gin.Context){
	var user database.UserInformation
	var err error
	id := ctx.Param("id")
	user, err = u.userAuth.GetByID(id)

	if err != nil {
		log.Debug("[ERROR: ", err.Error())
		if err == mongo.ErrNoDocuments {
			ctx.JSON(http.StatusNotFound,database.Error{common.StatusCodeNotFound,common.UserNotExist})	
		}else {
			ctx.JSON(http.StatusInternalServerError,database.Error{common.StatusCodeUnknown, err.Error()})
		}
	}else{
		ctx.JSON(http.StatusOK,user)
	}
}

// GetUsersByParams godoc
// @Summary Get a users by username parameter
// @Description Get a users by username parameter
// @Tags user
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Token"
// @Param username query string true "Username"
// @Failure 500 {object} database.Error
// @Failure 400 {object} database.Error
// @Success 200 {array} database.UserInformation
// @Router /users [get]
func (u *User) GetUsersByParams(ctx *gin.Context){
	var users []database.UserInformation
	var err error
	username := ctx.Query("username")
	users, err = u.userAuth.GetAll(username)

	if err != nil {
		log.Debug("[ERROR]: ",err.Error())
		ctx.JSON(http.StatusInternalServerError, database.Error{common.StatusCodeUnknown,err.Error()})
	}else {
		ctx.JSON(http.StatusOK,users)
	}
}

// DeleteUserByID godoc
// @Summary Delete a user by ID
// @Description Delete a user by ID
// @Tags user
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Token"
// @Param id path string true "User ID"
// @Failure 500 {object} database.Error
// @Failure 400 {object} database.Error
// @Success 200 {object} database.Message
// @Router /users/{id} [delete]
func (u *User) DeleteUserByID(ctx *gin.Context) {
	id := ctx.Param("id")
	err := u.userAuth.DeleteByID(id)

	if err != nil {
		log.Debug("[ERROR]: ",err.Error())
		if err.Error() == common.UserNotExist {
			ctx.JSON(http.StatusNotFound, database.Error{common.StatusCodeNotFound, err.Error()})
		}else {
			ctx.JSON(http.StatusInternalServerError, database.Error{common.StatusCodeUnknown,err.Error()})
		}
	} else {
		ctx.JSON(http.StatusOK, database.Message{common.DeleteSuccess})
	}
}

// UpdateUser godoc
// @Summary Update an existing user
// @Description Update an existing user
// @Tags user
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Param Authorization header string true "Token"
// @Param user body database.User true "User ID"
// @Failure 500 {object} database.Error
// @Success 200 {object} database.UserInformation
// @Router /users/{id} [patch]
func (u *User) UpdateUser(ctx *gin.Context){
	var user database.User
	id := ctx.Param("id")
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, database.Error{common.StatusCodeUnknown, err.Error()})
		return
	}

	err := u.userAuth.Update(user,id)

	if err != nil {
		log.Debug("[ERROR]: ",err.Error())
		if err.Error() == common.UserNotExist {
			ctx.JSON(http.StatusNotFound, database.Error{common.StatusCodeNotFound, err.Error()})
		}else {
			ctx.JSON(http.StatusInternalServerError, database.Error{common.StatusCodeUnknown,err.Error()})
		}
	} else {
		var result database.UserInformation
		result, err = u.userAuth.GetByID(id)
		ctx.JSON(http.StatusOK, result)
	}
}