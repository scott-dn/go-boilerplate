package request

type UpdateBook struct {
	Name        string `json:"name"        validate:"required,max=255"`
	Author      string `json:"author"      validate:"required,max=255"`
	Description string `json:"description" validate:"required,max=5000"`
	Version     int    `json:"version"     validate:"required,gt=0"`
}

type AddBook struct {
	Name        string `json:"name"        validate:"required,max=255"`
	Author      string `json:"author"      validate:"required,max=255"`
	Description string `json:"description" validate:"required,max=5000"`
}
