package response

type Error struct {
	ErrorMessage string `json:"error_message"`
}

type ResponseTemplate[T any] struct {
	Success bool    `json:"success"`
	Message string  `json:"message"`
	Errors  []Error `json:"errors,omitempty"`
	Data    any     `json:"any"`
}

func SuccessBuilder(data any) ResponseTemplate[any] {
	return ResponseTemplate[any]{
		Success: true,
		Message: "success",
		Errors:  nil,
		Data:    data,
	}
}

func ErrorBuilder(errs ...error) ResponseTemplate[any] {
	var errors []Error

	for _, v := range errs {
		errors = append(errors, Error{
			ErrorMessage: v.Error(),
		})
	}

	return ResponseTemplate[any]{
		Success: false,
		Message: "error",
		Errors:  errors,
		Data:    nil,
	}

}
