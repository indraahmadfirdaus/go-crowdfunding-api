package campaign

type Service interface {
	GetCampaigns(userID int) ([]Campaign, error)
	GetCampaignDetail(ID int) (Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) GetCampaigns(userID int) ([]Campaign, error) {
	campaigns := []Campaign{}

	if userID == 0 {
		campaignsRes, err := s.repository.FindAll()

		if err != nil {
			return campaigns, err
		}

		campaigns = campaignsRes
	} else {

		campaignsRes, err := s.repository.FindByUserID(userID)

		if err != nil {
			return campaigns, err
		}

		campaigns = campaignsRes
	}

	return campaigns, nil
}

func (s *service) GetCampaignDetail(ID int) (Campaign, error) {
	campaign, err := s.repository.FindById(ID)

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}
