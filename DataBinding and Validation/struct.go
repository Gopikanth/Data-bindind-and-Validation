package detail

type Person struct {
	Name        string `json:"name" binding:"required"`
	Father_Name string `json:"fathername" binding:"required"`
	Age         uint8  `json:"age" binding:"gte=0,lte=130"`
	//greaterthan 0 and less than 130
	Email string `json:"email" validate:"required,email"`
}

type Celebrity struct {
	Title string `json:"title" binding:"min=2,max=10" validate:"is-celebrity"`
	// is celebrity is used for validation
	Description string `json:"description" binding:"max=20"`
	URL         string `json:"url" binding:"required,url"`
	Actors      int8   `json:"actors" binding:"gt=2,lt=20"`
	Author      Person `json:"author"`
}

// we can handle those errors which are not binded
