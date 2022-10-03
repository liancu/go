package handlers

type helloRequest struct {
	Name string `json:"name" binding:"required"`
}

type productRequest struct {
	Name  string `json:"name" binding:"required"`
	Price uint   `json:"price" binding:"required"`
}
