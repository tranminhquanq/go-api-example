package models

type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

type Response struct {
	Data []Movie `json:"data"`
	Message string `json:"message"`
	Code int8 `json:"code"`
}