package utils

// BaseResponse represents the standard API response format
// @Description Standard API response format
type BaseResponse[T any] struct {
	Message string `json:"message" example:"success"`
	Data    T      `json:"data"`
}

// Response is a concrete type for Swagger documentation
// @Description Standard API response format
type Response struct {
	Message string      `json:"message" example:"success"`
	Data    interface{} `json:"data"`
}

// SuccessResponse creates a success response with data
func SuccessResponse[T any](data T) BaseResponse[T] {
	return BaseResponse[T]{
		Message: "success",
		Data:    data,
	}
}

// SuccessResponseWithMessage creates a success response with custom message
func SuccessResponseWithMessage[T any](message string, data T) BaseResponse[T] {
	return BaseResponse[T]{
		Message: message,
		Data:    data,
	}
}

// ErrorResponse creates an error response (data will be nil)
func ErrorResponse(message string) BaseResponse[interface{}] {
	return BaseResponse[interface{}]{
		Message: message,
		Data:    nil,
	}
}
