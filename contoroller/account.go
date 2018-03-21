package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/celler/model"
)

// ShowAccount godoc
// @Summary Show a account
// @Description get string by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param id path int true "Account ID"
// @Success 200 {object} model.Account
// @Failure 400 {object} controller.HTTPError
// @Failure 404 {object} controller.HTTPError
// @Failure 500 {object} controller.HTTPError
// @Router /accounts/{id} [get]
func (c *Controller) ShowAccount(ctx *gin.Context) {
	id := ctx.Param("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		NewError(ctx, http.StatusBadRequest, err)
		return
	}
	account, err := model.AccountOne(aid)
	if err != nil {
		NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, account)
}

// ListAccounts godoc
// @Summary List accounts
// @Description get accounts
// @Accept  json
// @Produce  json
// @Param q query string false "name search by q"
// @Success 200 {array} model.Account
// @Failure 400 {object} controller.HTTPError
// @Failure 404 {object} controller.HTTPError
// @Failure 500 {object} controller.HTTPError
// @Router /accounts [get]
func (c *Controller) ListAccounts(ctx *gin.Context) {
	q := ctx.Request.URL.Query().Get("q")
	accounts, err := model.AccountsAll(q)
	if err != nil {
		NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, accounts)
}

// AddAccount godoc
// @Summary Add a account
// @Description add by json account
// @Accept  json
// @Produce  json
// @Param account body model.AddAccount true "Add account"
// @Success 200 {object} model.Account
// @Failure 400 {object} controller.HTTPError
// @Failure 404 {object} controller.HTTPError
// @Failure 500 {object} controller.HTTPError
// @Router /accounts [post]
func (c *Controller) AddAccount(ctx *gin.Context) {
	var addAccount model.AddAccount
	if err := ctx.ShouldBindJSON(&addAccount); err != nil {
		NewError(ctx, http.StatusBadRequest, err)
		return
	}
	if err := addAccount.Validation(); err != nil {
		NewError(ctx, http.StatusBadRequest, err)
		return
	}
	account := model.Account{
		Name: addAccount.Name,
	}
	lastID, err := account.Insert()
	if err != nil {
		NewError(ctx, http.StatusBadRequest, err)
		return
	}
	account.ID = lastID
	ctx.JSON(http.StatusOK, account)
}

// UpdateAccount godoc
// @Summary Update a account
// @Description Update by json account
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param  id path int true "Account ID"
// @Param  account body model.UpdateAccount true "Update account"
// @Success 200 {object} model.Account
// @Failure 400 {object} controller.HTTPError
// @Failure 404 {object} controller.HTTPError
// @Failure 500 {object} controller.HTTPError
// @Router /accounts/{id} [patch]
func (c *Controller) UpdateAccount(ctx *gin.Context) {
	id := ctx.Param("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		NewError(ctx, http.StatusBadRequest, err)
		return
	}
	var updateAccount model.UpdateAccount
	if err := ctx.ShouldBindJSON(&updateAccount); err != nil {
		NewError(ctx, http.StatusBadRequest, err)
		return
	}
	account := model.Account{
		ID:   aid,
		Name: updateAccount.Name,
	}
	err = account.Update()
	if err != nil {
		NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, account)
}

// DeleteAccount godoc
// @Summary Update a account
// @Description Delete by account ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param  id path int true "Account ID"
// @Success 204 {object} model.Account
// @Failure 400 {object} controller.HTTPError
// @Failure 404 {object} controller.HTTPError
// @Failure 500 {object} controller.HTTPError
// @Router /accounts/{id} [delete]
func (c *Controller) DeleteAccount(ctx *gin.Context) {
	id := ctx.Param("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		NewError(ctx, http.StatusBadRequest, err)
		return
	}
	err = model.Delete(aid)
	if err != nil {
		NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{})
}

// UploadAccountImage godoc
// @Summary Upload account image
// @Description Upload file
// @ID account id
// @Accept  multipart/form-data
// @Produce  json
// @Param file formData file true "account image"
// @Success 200 {object} controller.Message
// @Failure 400 {object} controller.HTTPError
// @Failure 404 {object} controller.HTTPError
// @Failure 500 {object} controller.HTTPError
// @Router /accounts/1/images [post]
func (c *Controller) UploadAccountImage(ctx *gin.Context) {
	id := ctx.Param("id")
	_, err := strconv.Atoi(id)
	if err != nil {
		NewError(ctx, http.StatusBadRequest, err)
		return
	}
	file, err := ctx.FormFile("file")
	if err != nil {
		NewError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, Message{Message: fmt.Sprintf("upload compleate finename=%s", file.Filename)})
}
