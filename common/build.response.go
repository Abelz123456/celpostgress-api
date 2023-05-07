package common

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-playground/validator/v10"
)

// type WebResponse struct {
// 	Code   int         `json:"code"`
// 	Status string      `json:"status"`
// 	Data   interface{} `json:"data"`
// }

// DefaultResponse default payload response
type DefaultResponse struct {
	ResultCode string      `json:"resultCode"`
	HttpStatus string      `json:"http_status"`
	Message    string      `json:"developerMessage"`
	Total      int         `json:"total"`
	Data       interface{} `json:"data"`
}

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type NotFoundError struct {
	Error string
}

func NewNotFoundError(error string) NotFoundError {
	return NotFoundError{Error: error}
}

type ServerError struct {
	Error string
}

func NewServerError(error string) ServerError {
	return ServerError{Error: error}
}

func ResponseOk(response interface{}, total int) DefaultResponse {
	return DefaultResponse{
		"200",
		"200",
		"success", total,
		response,
	}
}

func ResponseDeleteOk(response interface{}) DefaultResponse {
	return DefaultResponse{
		"200",
		"200",
		"success", 1,
		response,
	}
}

// func ResponseOkWithCondition(err string) DefaultResponse {
// 	return DefaultResponse{
// 		"20001",
// 		"200",
// 		"success_with_condition",
// 	}
// }

func GetErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "lte":
		return "Should be less than " + fe.Param()
	case "gte":
		return "Should be greater than " + fe.Param()
	}
	return "Unknown error"
}

func ResponseOkDataNotFound(response interface{}) DefaultResponse {

	return DefaultResponse{
		"200",
		"200",
		"data not found", 0,
		response,
	}
}

func ResponseInternalServerError(response interface{}) DefaultResponse {
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	log.SetOutput(file)
	log.Println(err)
	return DefaultResponse{
		"40401",
		"404",
		"internal server error", 0,
		response,
	}
}

// func ResponseOkDataCreatedSuccess(err string) DefaultResponse {
// 	return DefaultResponse{
// 		"20100",
// 		"201",
// 		"created_success",
// 	}
// }

func ResponseBadRequest(err string) DefaultResponse {
	return DefaultResponse{
		"400",
		"400",
		"bad_request", 0, err,
	}
}

func ResponseUnAuthorized(err string) DefaultResponse {
	return DefaultResponse{
		"401",
		"401",
		"unauthorized", 0, err,
	}
}

func ResponseForm1Forbidden(err string) DefaultResponse {
	return DefaultResponse{
		"40300",
		"403",
		"missing_or_invalid_parameter", 0, err,
	}
}

func ResponseFinValidatorError(controller string, funcName string, fe interface{}) DefaultResponse {
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer file.Close()
	log.SetOutput(file)
	logStr := controller + funcName + fmt.Sprintf("%v", fe)
	log.Println(logStr)

	return DefaultResponse{
		"40300",
		"403",
		"missing_or_invalid_parameter", 0, fe,
	}
}

// func ResponseDataExist(err string) DefaultResponse {
// 	return DefaultResponse{
// 		"40301",
// 		"403",
// 		"data_exist",
// 	}
// }

// func ResponseQuotaExceded(err string) DefaultResponse {
// 	return DefaultResponse{
// 		"42200",
// 		"422",
// 		"quota_exceded",
// 	}
// }

// func ResponsDataNotFound(err string) DefaultResponse {
// 	return DefaultResponse{
// 		"40401",
// 		"404",
// 		"created_success",
// 	}
// }

// func ResponsIse1SystemError(err string) DefaultResponse {
// 	return DefaultResponse{
// 		"50000",
// 		"500",
// 		"system_error",
// 	}
// }

// func ResponsIse2SystemError(err string) DefaultResponse {
// 	return DefaultResponse{
// 		"50001",
// 		"500",
// 		"created_success",
// 	}
// }

// func ResponsConnectionTimeOut(err string) DefaultResponse {
// 	return DefaultResponse{
// 		"5002",
// 		"500",
// 		"connection_timeout",
// 	}
// }

// func ResponsConnectionError(err string) DefaultResponse {
// 	return DefaultResponse{
// 		"50003",
// 		"500",
// 		"connection_error",
// 	}
// }

// func ResponsIseQueryError(err string) DefaultResponse {
// 	return DefaultResponse{
// 		"50004",
// 		"500",
// 		"execute_query_error",
// 	}
// }

// func ResponsInsertError(err string) DefaultResponse {
// 	return DefaultResponse{
// 		"50005",
// 		"500",
// 		"execute_insert_error",
// 	}
// }

// func ResponsUpdateError(err string) DefaultResponse {
// 	return DefaultResponse{
// 		"50006",
// 		"500",
// 		"execute_update_error",
// 	}
// }

// func ResponsDeleteError(err string) DefaultResponse {
// 	return DefaultResponse{
// 		"50007",
// 		"500",
// 		"execute_delete_error",
// 	}
// }

// func ResponsIse8UnknowError(err string) DefaultResponse {
// 	return DefaultResponse{
// 		"50060",
// 		"500",
// 		"unknow_error",
// 	}
// }

// func ResponsSu1ServerBusy(err string) DefaultResponse {
// 	return DefaultResponse{
// 		"50300",
// 		"503",
// 		"server_busy",
// 	}
// }

// func ResponsServerUnvailable(err string) DefaultResponse {
// 	return DefaultResponse{
// 		"50301",
// 		"503",
// 		"server_unavailable",
// 	}
// }

// func ResponsGateWayTimeout(err string) DefaultResponse {
// 	return DefaultResponse{
// 		"50400",
// 		"504",
// 		"gateway_timeout_error",
// 	}
// }

func PanicIfError(err error) {
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer file.Close()
	log.SetOutput(file)
	log.Println(err)
}

// func ReadFromRequestBody(request *http.Request, result interface{}) {
// 	decoder := json.NewDecoder(request.Body)
// 	err := decoder.Decode(result)
// 	PanicIfError(err)
// }

// func WriteToResponseBody(c *gin.Context, response interface{}) {
// 	decoder := json.NewDecoder(response.Body)
// 	err := decoder.Decode(result)
// 	PanicIfError(err)
// 	// c.JSON(http.StatusOK, response)
// 	// PanicIfError(err)
// }

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	// sentry.CaptureMessage(fmt.Sprintf("%v", err))
	if err != nil {
		errorRollback := tx.Rollback()
		PanicIfError(errorRollback)
		panic(err)
	} else {
		errorCommit := tx.Commit()
		PanicIfError(errorCommit)
	}
}
