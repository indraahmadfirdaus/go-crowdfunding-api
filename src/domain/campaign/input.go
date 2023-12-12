package campaign

import "crowdfunding-api/src/domain/user"

type GetByIdInput struct {
	ID int `uri:"id" binding:"required"`
}

type CreateCampaignInput struct {
	Name             string    `json:"name" binding:"required"`
	ShortDescription string    `json:"short_description" binding:"required"`
	Description      string    `json:"description" binding:"required"`
	Perks            []string  `json:"perks" binding:"required"`
	GoalAmount       int       `json:"goal_amount" binding:"required"`
	Slug             string    `json:"slug"`
	User             user.User `json:"user"`
}

type UpdateCampaignInput struct {
	ID   int `uri:"id" binding:"required"`
	Data struct {
		Name             string    `json:"name"`
		ShortDescription string    `json:"short_description"`
		Description      string    `json:"description"`
		Perks            []string  `json:"perks"`
		GoalAmount       int       `json:"goal_amount"`
		Slug             string    `json:"slug"`
		User             user.User `json:"user"`
	}
}
