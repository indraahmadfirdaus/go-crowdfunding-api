package campaign

type GetByIdInput struct {
	ID int `uri:"id" binding:"required"`
}
