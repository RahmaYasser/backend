package handlers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/tamiat/backend/pkg/domain/tamiat_user"
	"github.com/tamiat/backend/pkg/errs"
	"github.com/tamiat/backend/pkg/middleware"
	"github.com/tamiat/backend/pkg/service"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)

type TamiatUserHandlers struct {
	Service service.TamiatUserService
}
type CreateTUser struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
	Role     string `json:"role_name" form:"role_name" binding:"required"`
	Name     string `json:"name" form:"name" binding:"required"`
}
type UpdateUsr struct {
	Name string `json:"role" form:"name" binding:"required"`
	Role string `json:"role_name" form:"role_name" binding:"required"`
}
type ResetPass struct {
	NewPass string `json:"new_pass" form:"new_pass" binding:"required"`
}
type IDResponse struct {
	Id int `json:"id" form:"id"`
}

//
// @Summary Login endpoint
// @Description Provide email and password to login, response is JWT
// @Consume application/x-www-form-urlencoded
// @Produce application/json
// @Param email formData string true "Email"
// @Param password formData string true "Password"
// @Success 200 {object} handlers.JWT
// @Failure 400  {object}  errs.ErrResponse "Bad Request"
// @Failure 404  {object}  errs.ErrResponse "User not found"
// @Failure 401  {object}  errs.ErrResponse "Unauthorizes"
// @Failure 500  {object}  errs.ErrResponse "Internal server error"
// @Router /tamiatUser/login [post]
func (receiver TamiatUserHandlers) Login(ctx *gin.Context) {
	var userObj tamiat_user.TamiatUser
	var loginRequestData Login
	//decoding request body

	if err := ctx.ShouldBind(&loginRequestData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userObj.Email = loginRequestData.Email
	userObj.Password = loginRequestData.Password

	// retrieving hashed password from database
	hashedPassword, err := receiver.Service.Login(userObj)
	if err != nil {
		if err == errs.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"error": errs.ErrRecordNotFound.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": errs.ErrDb.Error()})
		return
	}
	// authentication process
	password := loginRequestData.Password
	if err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": errs.ErrInvalidPassword.Error()})
		return
	}
	//generating token
	token, err := middleware.GenerateToken(userObj.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": errs.ErrTokenErr.Error()})
		return
	}
	jwtObj := JWT{Token: token}
	ctx.JSON(http.StatusOK, jwtObj)
}

//
// @Summary Create Tamiat User endpoint
// @Description role name should be one of the following(Admin, Editor, Author, Contributor, Subscriber. Admins only can use this endpoint.
// @Consume application/x-www-form-urlencoded
// @Produce application/json
// @Param email formData string true "Email"
// @Param password formData string true "Password"
// @Param role_name formData string true "role name"
// @Param name formData string true "name"
// @Success 200 {object} handlers.IDResponse
// @Failure 400  {object}  errs.ErrResponse "Bad Request"
// @Failure 500  {object}  errs.ErrResponse "Internal server error"
// @Router /tamiatUser/create [post]
func (receiver TamiatUserHandlers) Create(ctx *gin.Context) {
	var userObj tamiat_user.TamiatUser
	var createRequestData CreateTUser
	//decoding request body
	if err := ctx.ShouldBind(&createRequestData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userObj.Name = createRequestData.Name
	userObj.Email = createRequestData.Email
	//encrypting password
	hash, err := bcrypt.GenerateFromPassword([]byte(createRequestData.Password), 10)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	userObj.Password = string(hash)
	//database connection
	userObj.RoleId, err = receiver.Service.GetRoleId(createRequestData.Role)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	userObj.Id, err = receiver.Service.Create(userObj)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//w.WriteHeader(http.StatusOK)
	userObj.Password = ""
	//json.NewEncoder(w).Encode(userObj)
	ctx.JSON(http.StatusOK, userObj)
}

//
// @Summary ReadAll Tamiat Users endpoint
// @Description Admins only can use this endpoint.
// @Accept application/json
// @Produce application/json
// @Success 200 {array} tamiat_user.TamiatUser
// @Success 404 {object} errs.ErrResponse "No users found"
// @Failure 500  {object}  errs.ErrResponse "Internal server error"
// @Router /tamiatUser/index [get]
func (receiver TamiatUserHandlers) ReadAll(ctx *gin.Context) {
	allUsers, err := receiver.Service.ReadAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if allUsers == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": errs.ErrRecordNotFound.Error()})
		return
	}
	ctx.JSON(http.StatusOK, allUsers)
}

//
// @Summary ReadUserByID  for Tamiat User endpoint
// @Description Admins only can use this endpoint.
// @Accept application/json
// @Produce application/json
// @Param  id path int true "user ID"
// @Success 200 {object} tamiat_user.TamiatUser
// @Failure 400  {object}  errs.ErrResponse "Bad Request"
// @Success 404 {object} errs.ErrResponse "No user found"
// @Failure 500  {object}  errs.ErrResponse "Internal server error"
// @Router /tamiatUser/read/{id} [get]
func (receiver TamiatUserHandlers) ReadUserByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errs.ErrParsingID.Error()})
		return
	}
	usrObj, err := receiver.Service.ReadUserByID(id)
	if err == sql.ErrNoRows {
		ctx.JSON(http.StatusNotFound, gin.H{"error": errs.ErrRecordNotFound.Error()})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, usrObj)
}

//
// @Summary Update name or role for Tamiat User endpoint
// @Description Admins only can use this endpoint.
// @Consume application/x-www-form-urlencoded
// @Produce application/json
// @Param  id path int true "user ID"
// @Param name formData string true "name"
// @Param role_name formData string true "role name"
// @Success 200 {object} handlers.IDResponse
// @Failure 400  {object}  errs.ErrResponse "Bad Request"
// @Success 404 {object} errs.ErrResponse "No users found"
// @Failure 500  {object}  errs.ErrResponse "Internal server error"
// @Router /tamiatUser/update/{id} [put]
func (receiver TamiatUserHandlers) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errs.ErrParsingID.Error()})
		return
	}
	var usrObj tamiat_user.TamiatUser
	usrObj.Id = id
	var RequestUpdateUsr UpdateUsr
	err = ctx.ShouldBind(&RequestUpdateUsr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errs.ErrParsingID)
		return
	}
	usrObj.Name = RequestUpdateUsr.Name
	usrObj.RoleId, err = receiver.Service.GetRoleId(RequestUpdateUsr.Role)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errs.ErrParsingID)
		return
	}
	err = receiver.Service.Update(usrObj)
	if err == sql.ErrNoRows {
		ctx.JSON(http.StatusNotFound, gin.H{"error": errs.ErrRecordNotFound.Error()})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"id": id})
}

//
// @Summary ResetPassword Tamiat endpoint
// @Description Admins only can use this endpoint.
// @Consume application/x-www-form-urlencoded
// @Produce application/json
// @Param  id path int true "user ID"
// @Param new_pass formData string true "new_pass"
// @Success 200 {object} handlers.IDResponse
// @Failure 400  {object}  errs.ErrResponse "Bad Request"
// @Success 404 {object} errs.ErrResponse "No users found"
// @Failure 500  {object}  errs.ErrResponse "Internal server error"
// @Router /tamiatUser/resetpass/{id} [put]
func (receiver TamiatUserHandlers) ResetPassword(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errs.ErrParsingID.Error()})
		return
	}
	var usrObj tamiat_user.TamiatUser
	usrObj.Id = id
	var RequestResetPass ResetPass
	err = ctx.ShouldBind(&RequestResetPass)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(RequestResetPass.NewPass), 10)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	usrObj.Password = string(hash)
	err = receiver.Service.ResetPassword(usrObj)
	if err == sql.ErrNoRows {
		ctx.JSON(http.StatusNotFound, gin.H{"error": errs.ErrRecordNotFound.Error()})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"id": id})
}

//
// @Summary Delete Tamiat endpoint
// @Description Admins only can use this endpoint.
// @Accept application/json
// @Produce application/json
// @Param  id path int true "user ID"
// @Success 200 {object} handlers.IDResponse
// @Failure 400  {object}  errs.ErrResponse "Bad Request"
// @Success 404 {object} errs.ErrResponse "No users found"
// @Failure 500  {object}  errs.ErrResponse "Internal server error"
// @Router /tamiatUser/delete/{id} [delete]
func (receiver TamiatUserHandlers) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errs.ErrParsingID.Error()})
		return
	}
	err = receiver.Service.Delete(id)
	if err == sql.ErrNoRows {
		ctx.JSON(http.StatusNotFound, gin.H{"error": errs.ErrRecordNotFound.Error()})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"id": id})
}
