package domain

// BaseResponse represents the standard API response format
// @Description Standard API response format with generic data type
type BaseResponse[T any] struct {
	Message string `json:"message" example:"success" description:"Response message indicating the result of the operation"`
	Data    T      `json:"data" description:"Response data of generic type"`
}

// Response is a concrete type for Swagger documentation
// @Description Standard API response format for Swagger documentation
type Response struct {
	Message string      `json:"message" example:"success" description:"Response message indicating the result of the operation"`
	Data    interface{} `json:"data" description:"Response data (can be any type)"`
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
