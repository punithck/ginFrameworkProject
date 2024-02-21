package entity

type Person struct {
	// Binding tag is used by gin
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Age       int8   `json:"age" binding:"gte=1,lte=130"`
	Email     string `json:"email" binding:"required,email"`
}

// Object model
// It is used to create video object
type Video struct {
	// adding custom validator, I am using validator.v9 package
	Title string `json:"title"  binding:"min=2,max=100" validate:"is-cool"`
	// We can add validation like this
	//Title       string `json:"title" xml:"title" form:"title" validate:"binding" binding:"required"`
	Description string `json:"description" binding:"max=200"`
	URL         string `json:"url" binding:"required,url"`
	Author      Person `json:"author" binding:"required"`
}
