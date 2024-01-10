package entities

import "net/http"

type StatusCode struct {
	code     string
	message  string
	httpCode int
}

func (receiver StatusCode) GetCode() string {
	return receiver.code
}

func (receiver StatusCode) GetMessage() string {
	return receiver.message
}

func (receiver StatusCode) GetHTTPCode() int {
	return receiver.httpCode
}

var (
	SuccessStatusCode = StatusCode{
		code:     "SUCCESS",
		message:  "Success",
		httpCode: http.StatusOK,
	}

	BadRequestStatusCode = StatusCode{
		code:     "BAD_REQUEST",
		message:  "Invalid data",
		httpCode: http.StatusBadRequest,
	}

	InternalServerErrorStatusCode = StatusCode{
		code:     "STATUS_INTERNAL_SERVER_ERROR",
		message:  "Invalid data",
		httpCode: http.StatusInternalServerError,
	}
)

const (
	MgoErrNotFound  string = "not found"
	MsgGeneralError string = "Có lỗi xảy ra trong quá trình xử lý, vui lòng thực hiện lại sau!"

	MessageTypeSuccess int = 1
	MessageTypeError   int = 2
	MessageTypeWarning int = 3

	ErrorCodeCommonStart int = 100000

	ErrorCodeIncreaseStart    int = 200000
	ErrorCodeCURLStart        int = 210000
	ErrorCodeFileStart        int = 220000
	ErrorCodeRepoProductStart int = 230000
)

const (
	InfoLogLevel                   = "info"
	WarnLogLevel                   = "warn"
	DebugLogLevel                  = "debug"
	ErrorLogLevel                  = "error"
	Step                           = "step"
	SlowResponseLogTag             = "slow_response"
	ThresholdWarnTriggerNumLogTag  = "threshold_warn_trigger_num"
	ThresholdWarnTriggerTimeLogTag = "threshold_warn_trigger_time"
)

const (
	ErrorCodeSystem int = iota + ErrorCodeCommonStart
	ErrorCodeParseJSONInput
	ErrorCodeInputNull
	ErrorCodeFindMongoDB
)
const (
	ErrorCodeIncreaseCacheLastID int = iota + ErrorCodeIncreaseStart
	ErrorCodeIncreaseCacheSetNX
	ErrorCodeIncreaseCacheIncBy
	ErrorCodeIncreaseCacheLTDefault
	ErrorCodeIncreaseLastIdByDBFind
	ErrorCodeIncreaseLastIdByDBFieldNoExist
	ErrorCodeIncreaseLastIdByDBMarshalField
	ErrorCodeIncreaseLastIdByDBUnmarshalField
)
const (
	ErrorCodeCURLMethod int = iota + ErrorCodeCURLStart
	ErrorCodeCURLRequest
	ErrorCodeCURLNoResp
	ErrorCodeCURLRespStatusCode
	ErrorCodeCURLResp
	ErrorCodeCURLRespBody
	ErrorCodeCURLNewRequestGet
	ErrorCodeCURLNewRequestNoGet
	ErrorCodeCURLParseJsonData
)
const (
	ErrorCodeFileDownloadFileGet int = iota + ErrorCodeFileStart
	ErrorCodeFileDownloadFileCreate
	ErrorCodeFileDownloadFileCopy
	ErrorCodeFileGzipOpen
	ErrorCodeFileGzipCreate
	ErrorCodeFileGzipCopy
	ErrorCodeFileUnGzipOpen
	ErrorCodeFileUnGzipRead
	ErrorCodeFileUnGzipCreate
	ErrorCodeFileUnGzipCopy
)

const (
	ErrorCodeRepoProductGetOverlapProductsFind int = iota + ErrorCodeRepoProductStart
)

var MapError = map[int]string{
	ErrorCodeSystem:         "Lỗi hệ thống.Vui lòng thử lại",
	ErrorCodeParseJSONInput: "Dữ liệu không hợp lệ",
	ErrorCodeFindMongoDB:    "Lỗi khi tìm kiếm dữ liệu từ Mongodb",
}
