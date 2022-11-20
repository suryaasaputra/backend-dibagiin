package controllers

import (
	"dibagi/helpers"
	"dibagi/models"
	"dibagi/repository"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type userController struct {
	UserRepository repository.IUserRepository
}

func NewUserController(userRepository repository.IUserRepository) *userController {
	return &userController{
		UserRepository: userRepository,
	}
}

func (u userController) Register(ctx *gin.Context) {
	var registerUser = models.UserRegisterRequest{}

	contentType := helpers.GetRequestHeaders(ctx).ContentType
	if contentType == "application/json" {
		err := ctx.ShouldBindJSON(&registerUser)
		if err != nil {
			response := helpers.GetResponse(true, http.StatusBadRequest, "error binding request", nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
	} else {
		err := ctx.ShouldBind(&registerUser)
		if err != nil {
			response := helpers.GetResponse(true, http.StatusBadRequest, "error binding request", nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
	}

	// profilPhotoUrl, err := helpers.UploadToBucket(ctx.Request, "profil_photo", registerUser.ID)

	// if err != nil {
	// 	response := helpers.GetResponse(true, http.StatusInternalServerError, err.Error(), nil)
	// 	ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
	// 	return
	// }

	var user = models.User{
		Email:       registerUser.Email,
		FullName:    registerUser.FullName,
		UserName:    registerUser.UserName,
		Password:    registerUser.Password,
		Gender:      registerUser.Gender,
		PhoneNumber: registerUser.PhoneNumber,
		Address:     registerUser.Address,
	}

	resp, err := u.UserRepository.RegisterUser(user)
	if err != nil {
		response := helpers.GetResponse(true, http.StatusBadRequest, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.GetResponse(false, http.StatusOK, "Register Success", resp)
	ctx.JSON(http.StatusCreated, response)

}

func (u userController) Login(ctx *gin.Context) {
	var request = models.UserLoginRequest{}

	contentType := helpers.GetRequestHeaders(ctx).ContentType
	if contentType == "application/json" {
		err := ctx.ShouldBindJSON(&request)
		if err != nil {
			response := helpers.GetResponse(true, http.StatusBadRequest, "error binding request", nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
	} else {
		err := ctx.ShouldBind(&request)
		if err != nil {
			response := helpers.GetResponse(true, http.StatusBadRequest, "error binding request", nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
	}

	userResult := u.UserRepository.GetUserByEmail(request.Email)

	if userResult.Email == "" {
		response := helpers.GetResponse(true, http.StatusUnauthorized, "Email not registered", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	isPasswordCorrect := helpers.ComparePassword([]byte(userResult.Password), []byte(request.Password))

	if !isPasswordCorrect {
		response := helpers.GetResponse(true, http.StatusUnauthorized, "Incorrect password", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	token, err := helpers.GenerateToken(userResult.ID, userResult.UserName)

	if err != nil {
		response := helpers.GetResponse(true, http.StatusInternalServerError, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

	response := helpers.GetResponse(false, http.StatusOK, "Login Success", gin.H{
		"id":               userResult.ID,
		"email":            userResult.Email,
		"user_name":        userResult.UserName,
		"full_name":        userResult.FullName,
		"profil_photo_url": userResult.ProfilPhotoUrl,
		"token":            token,
	})

	ctx.JSON(http.StatusOK, response)
}

func (u userController) GetUser(ctx *gin.Context) {
	userNameURL := ctx.Param("userName")
	userResult := u.UserRepository.GetUserByUserName(userNameURL)
	if userResult.UserName == "" {
		response := helpers.GetResponse(true, http.StatusNotFound, "User not found", nil)
		ctx.AbortWithStatusJSON(http.StatusNotFound, response)
		return
	}
	response := helpers.GetResponse(false, http.StatusOK, "Success", userResult)
	ctx.JSON(http.StatusOK, response)
}

func (u userController) Update(ctx *gin.Context) {
	userNameURL := ctx.Param("userName")
	var request = models.EditUserRequest{}

	contentType := helpers.GetRequestHeaders(ctx).ContentType
	if contentType == "application/json" {
		err := ctx.ShouldBindJSON(&request)
		if err != nil {
			response := helpers.GetResponse(true, http.StatusBadRequest, "error binding request", nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
	} else {
		err := ctx.ShouldBind(&request)
		if err != nil {
			response := helpers.GetResponse(true, http.StatusBadRequest, "error binding request", nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
	}

	var newUserData = models.User{
		Email:       request.Email,
		UserName:    request.UserName,
		FullName:    request.FullName,
		Gender:      request.Gender,
		PhoneNumber: request.PhoneNumber,
		Address:     request.Address,
	}
	result, err := u.UserRepository.EditUser(userNameURL, newUserData)
	if err != nil {
		response := helpers.GetResponse(true, http.StatusBadRequest, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.GetResponse(false, http.StatusOK, "Success Update User Data", result)
	ctx.JSON(http.StatusOK, response)
}

func (u userController) SetProfilePhoto(ctx *gin.Context) {
	request := models.SetProfilePhotoRequest{}
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	id := fmt.Sprintf("%v", userData["id"])

	err := ctx.ShouldBind(&request)
	if err != nil {
		response := helpers.GetResponse(true, http.StatusBadRequest, "error binding request", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	profilPhotoUrl, err := helpers.UploadToBucket(ctx.Request, "profil_photo", id)

	if err != nil {
		response := helpers.GetResponse(true, http.StatusInternalServerError, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

	err = u.UserRepository.SetProfilePhoto(id, profilPhotoUrl)
	if err != nil {
		response := helpers.GetResponse(true, http.StatusInternalServerError, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

	response := helpers.GetResponse(false, http.StatusOK, "Success Update Profile Photo", nil)
	ctx.JSON(http.StatusOK, response)
}

func (u userController) CheckUser(ctx *gin.Context) {
	email := ctx.Query("email")
	userName := ctx.Query("user_name")
	if userName != "" && email == "" {
		result := u.UserRepository.GetUserByUserName(userName)
		if result.UserName == userName {
			response := helpers.GetResponse(true, http.StatusBadRequest, "Username Already Taken", nil)
			ctx.JSON(http.StatusBadRequest, response)
			return
		}
		response := helpers.GetResponse(false, http.StatusOK, "Username Available", nil)
		ctx.JSON(http.StatusOK, response)

	} else if email != "" && userName == "" {
		result := u.UserRepository.GetUserByEmail(email)
		if result.Email == email {
			response := helpers.GetResponse(true, http.StatusBadRequest, "Email Already Taken", nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		response := helpers.GetResponse(false, http.StatusOK, "Email Available", nil)
		ctx.JSON(http.StatusOK, response)
	}
}

func (u userController) Delete(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := fmt.Sprintf("%v", userData["id"])

	err := u.UserRepository.DeleteUser(userID)
	if err != nil {
		response := helpers.GetResponse(true, http.StatusInternalServerError, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}
	response := helpers.GetResponse(false, http.StatusOK, "Success Delete Account", nil)
	ctx.JSON(http.StatusOK, response)
}
