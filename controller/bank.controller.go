package controller

import (
	"net/http"

	"errors"

	"github.com/celpostgress-api/common"
	"github.com/celpostgress-api/dto"
	"github.com/celpostgress-api/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type BankController struct {
	BankService services.IBankService
}

func NewBankController(bankService services.IBankService) IBankController {
	return &BankController{
		BankService: bankService,
	}
}

// Create Bank
// @Summary  create bank
// @Description create new bank
// @Tags banks
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"ok"
// @Security ApiKeyAuth
// @Param   bank body  dto.CreateBankDto true "bank"
// @Router /bank/create [post]
func (controller *BankController) Create(c *gin.Context) {

	var createBankDto dto.CreateBankDto
	// err := c.ShouldBindJSON(&createBankDto)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, common.ResponseForm1Forbidden(err.Error()))
	// 	return
	// }
	if err := c.ShouldBindJSON(&createBankDto); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]common.ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = common.ErrorMsg{Field: fe.Field(), Message: common.GetErrorMsg(fe)}
			}
			c.JSON(http.StatusBadRequest, common.ResponseFinValidatorError("BankController", "Create", out))
		}
		return
	}

	bankResponse := controller.BankService.Create(c, createBankDto)
	response := common.ResponseOk(bankResponse, 1)

	c.JSON(http.StatusOK, response)
}

// Update Bank
// @Summary  update bank by oid
// @Description update existing bank
// @Tags banks
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"ok"
// @Security ApiKeyAuth
// @Param oid path string true  "oid"
// @Param   bank body  dto.UpdateBankDto true "bank"
// @Router /bank/{oid} [put]
func (controller *BankController) Update(c *gin.Context) {
	oid := c.Param("oid")
	bankUpdateRequest := dto.UpdateBankDto{}
	err := c.ShouldBindJSON(&bankUpdateRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.ResponseForm1Forbidden(err.Error()))
		return
	}

	// bankUpdateRequest.Oid = oid

	bankResponse := controller.BankService.Update(c, bankUpdateRequest, oid)
	response := common.ResponseOk(bankResponse, 1)

	c.JSON(http.StatusOK, response)
}

// Delete Bank
// @Summary  delete bank
// @Description delete bank by oid
// @Tags banks
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"ok"
// @Security ApiKeyAuth
// @Param oid path string true  "oid"
// @Router /bank/delete/{oid} [delete]
func (controller *BankController) Delete(c *gin.Context) {
	oid := c.Param("oid")

	controller.BankService.Delete(c, oid)
	response := common.ResponseDeleteOk(oid)

	c.JSON(http.StatusOK, response)
}

// Find Bank By Oid
// @Summary  find bank by oid
// @Description update new bank
// @Tags banks
// @Accept  json
// @Produce  json
// @Success 200 {object} common.DefaultResponse	"ok"
// @Security ApiKeyAuth
// @Param oid path string true  "oid"
// @Router /bank/{oid} [get]
func (controller *BankController) FindById(c *gin.Context) {
	oid := c.Param("oid")
	bank := controller.BankService.FindById(c, oid)
	response := common.ResponseOk(bank, 1)
	c.JSON(http.StatusOK, response)
}

// Get All Bank
// @Summary List existing banks
// @Description Get all the existing banks
// @Tags banks
// @Accept  json
// @Produce  json
// @Success 200 {object} common.DefaultResponse	"ok"
// @Security ApiKeyAuth
// @Param   bankDto body  dto.CreateBankDto true "bank"
// @Param        pageIndex    query     int  false  "page Index"  minimum(1)    maximum(10)
// @Param        pageSize    query     int  false  "page Size"  minimum(1)    maximum(10)
// @Router /bank/q [post]
func (controller *BankController) Find(c *gin.Context) {
	// @Param pageIndex  path int false "pageIndex"
	// @Param pageSize  path int true "pageSize"
	pageIndex := c.Query("pageIndex")
	pageSize := c.Query("pageSize")

	var bankDto dto.FilterBankDto
	err := c.ShouldBindJSON(&bankDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.ResponseForm1Forbidden(err.Error()))
		return
	}
	criteria := ""
	criteriaQ := ""
	if len(pageIndex) == 0 {
		pageIndex = "0"
	}
	if len(pageSize) == 0 {
		pageSize = "10"
	}

	criteriaQ += `where "bankCode" Like '%` + bankDto.BankCode + `%' and "bankName" Like '%` + bankDto.BankName + `%'`
	criteria += criteriaQ + ` order by "bankCode" asc LIMIT ` + pageSize + `  OFFSET ` + pageIndex

	banks := controller.BankService.Find(c, criteria)
	banksCount := controller.BankService.Find(c, criteriaQ)

	response := common.ResponseOk(banks, len(banksCount))
	c.JSON(http.StatusOK, response)
}
