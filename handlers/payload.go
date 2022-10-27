package handlers

type helloRequest struct {
	Name string `json:"name" binding:"required"`
}
