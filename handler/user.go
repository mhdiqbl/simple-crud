package handler

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"ottoDigital/dto"
	"ottoDigital/model"
	"ottoDigital/utils"
)

func (h *handler) Login(c *gin.Context) {
	var userRequest *dto.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(400, err)
		return
	}

	user, err := h.repository.GetUserByEmail(c, userRequest.Email)

	if err != nil {
		c.JSON(http.StatusUnauthorized, dto.BaseResponse{
			Code:    401,
			Message: "Email & Password is not valid",
			Data:    nil,
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userRequest.Password))

	if err != nil {
		c.JSON(http.StatusUnauthorized, dto.BaseResponse{
			Code:    401,
			Message: "Email & Password is not valid",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, dto.BaseResponse{
		Code:    http.StatusOK,
		Message: "Berhasil login",
		Data:    &user,
	})
}

func (h *handler) Register(c *gin.Context) {
	var userRequest *dto.UserRequest

	if err := c.ShouldBind(&userRequest); err != nil {
		c.JSON(400, err)
		return
	}

	userEmail, _ := h.repository.GetUserByEmail(c, userRequest.Email)

	//fmt.Println(users.ID)

	if userEmail.ID != 0 {
		c.JSON(http.StatusConflict, dto.BaseResponse{
			Code:    409,
			Message: "Alamat email sudah terdaftar",
			Data:    nil,
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.DefaultCost)
	userRequest.Password = string(hashedPassword)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal server error",
			Data:    nil,
		})
		return
	}

	var user model.User
	user.Name = userRequest.Name
	user.Email = userRequest.Email
	user.Password = userRequest.Password

	_, err = h.repository.CreateUser(c, user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal server error",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusUnauthorized, dto.BaseResponse{
		Code:    http.StatusOK,
		Message: "Berhasil register user",
		Data:    nil,
	})
}

func (h *handler) GetUsers(c *gin.Context) {
	userResponse := make([]*dto.UserResponse, 0)
	users, err := h.repository.GetAllUsers(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.BaseResponse{
			Code:    1,
			Message: "Internal server error",
			Data:    nil,
		})
		return
	}

	for _, user := range users {
		userResponse = append(userResponse, &dto.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdateAt:  user.UpdateAt,
		})
	}

	c.JSON(http.StatusOK, dto.BaseResponse{
		Code:    http.StatusOK,
		Message: "Berhasil get data users",
		Data:    &userResponse,
	})
}

func (h *handler) GetUserByID(c *gin.Context) {
	var userResponse *dto.UserResponse
	userID, err := utils.ExtractorRequestParamID(&c.Params)

	user, err := h.repository.GetUserByID(c, int(userID))

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal server error",
			Data:    nil,
		})
		return
	}

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, dto.BaseResponse{
			Code:    http.StatusNotFound,
			Message: "User not found",
			Data:    nil,
		})
		return
	}

	userResponse = &dto.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	c.JSON(http.StatusOK, dto.BaseResponse{
		Code:    http.StatusOK,
		Message: "Berhasil get data user",
		Data:    &userResponse,
	})
	return
}

func (h *handler) UpdateUser(c *gin.Context) {
	var user model.User
	var userRequest *dto.UserRequest

	userId, err := utils.ExtractorRequestParamID(&c.Params)

	if err != nil {
		c.JSON(400, err)
		return
	}

	if err := c.ShouldBind(&userRequest); err != nil {
		c.JSON(400, err)
		return
	}

	user.Name = userRequest.Name
	user.Email = userRequest.Email

	var userFound bool
	//fmt.Println(user)
	userFound, err = h.repository.UpdateUser(c, int(userId), user)

	if err != nil {
		c.JSON(400, err)
		return
	}

	if !userFound {
		c.JSON(http.StatusNotFound, dto.BaseResponse{
			Code:    http.StatusNotFound,
			Message: "User not found",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, dto.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success update user",
		Data:    nil,
	})
	return
}

func (h *handler) DeleteUser(c *gin.Context) {
	userId, err := utils.ExtractorRequestParamID(&c.Params)

	if err != nil {
		c.JSON(400, err)
		return
	}

	var userFound bool
	userFound, err = h.repository.DeleteUser(c, int(userId))

	if !userFound {
		c.JSON(http.StatusNotFound, dto.BaseResponse{
			Code:    http.StatusNotFound,
			Message: "User not found",
			Data:    nil,
		})
		return
	}

	if err != nil {
		c.JSON(400, err)
		return
	}

	c.JSON(http.StatusNoContent, dto.BaseResponse{
		Code: http.StatusOK,
		Data: "Success delete user",
	})
	return
}
