package accountmodel

type GetAccountRequest struct {
	ID int64 `json:"id" uri:"id" binding:"required,min=1"`
}
