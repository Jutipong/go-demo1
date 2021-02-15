package entity

type Customer struct {
	ID    string `json:"id" binding:"required"`
	Code  string `json:"code" binding:"required"`
	Name  string `json:"name"`
	Age   int8   `json:"age" binding:"gte=1,lte=130"`
	Email string `json:"email" binding:"required,email"`
}


