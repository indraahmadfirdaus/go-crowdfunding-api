package campaign

type campaignDetailResponse struct {
	ID               int             `json:"id"`
	Name             string          `json:"name"`
	ShortDescription string          `json:"short_description"`
	Description      string          `json:"description"`
	Perks            []string        `json:"perks"`
	BackerCount      int             `json:"backer_count"`
	CurrentAmount    int             `json:"current_amount"`
	GoalAmount       int             `json:"goal_amount"`
	Slug             string          `json:"slug"`
	CampaignImages   []imageResponse `json:"campaign_images"`
	User             userResponse    `json:"user"`
}

type userResponse struct {
	Name           string `json:"name"`
	AvatarFileName string `json:"avatar_file_name"`
}

type imageResponse struct {
	FileName  string `json:"file_name"`
	IsPrimary bool   `json:"is_primary"`
}

func FormatCampaignDetailResponse(c Campaign) campaignDetailResponse {
	resp := campaignDetailResponse{}

	resp.ID = c.ID
	resp.Name = c.Name
	resp.ShortDescription = c.ShortDescription
	resp.Description = c.Description
	resp.Perks = c.Perks
	resp.BackerCount = c.BackerCount
	resp.CurrentAmount = c.CurrentAmount
	resp.GoalAmount = c.GoalAmount
	resp.Slug = c.Slug
	resp.User.Name = c.User.Name
	resp.User.AvatarFileName = c.User.AvatarFileName

	campaignImages := []imageResponse{}

	for _, v := range c.CampaignImages {
		img := imageResponse{}
		img.FileName = v.FileName
		img.IsPrimary = v.IsPrimary
		campaignImages = append(campaignImages, img)
	}

	resp.CampaignImages = campaignImages

	return resp
}

func FormatGetListCampaignResponse(c []Campaign) []campaignDetailResponse {
	resp := []campaignDetailResponse{}

	for _, v := range c {
		single := FormatCampaignDetailResponse(v)
		resp = append(resp, single)
	}

	return resp
}
