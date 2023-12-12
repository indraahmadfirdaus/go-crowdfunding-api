package campaign

import (
	"errors"
)

type Service interface {
	GetCampaigns(userID int) ([]Campaign, error)
	GetCampaignDetail(ID int) (Campaign, error)
	CreateCampaign(input CreateCampaignInput) error
	UpdateCampaign(input UpdateCampaignInput) error
}

type service struct {
	repository Repository
}

func NewService() *service {
	repository := NewRepository()
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

func (s *service) CreateCampaign(input CreateCampaignInput) error {
	payload := Campaign{}

	payload.Name = input.Name
	payload.UserId = input.User.ID
	payload.Description = input.Description
	payload.ShortDescription = input.ShortDescription
	payload.GoalAmount = input.GoalAmount
	payload.CurrentAmount = 0

	_, err := s.repository.Save(payload)

	if err != nil {
		return err
	}

	return nil
}

func (s *service) UpdateCampaign(input UpdateCampaignInput) error {
	payload, err := s.repository.FindById(input.ID)

	if err != nil {
		return err
	}

	if payload.ID == 0 {
		return errors.New("Not Found Campaign")
	}

	data := input.Data

	if data.Name != "" {
		payload.Name = data.Name
	}

	if data.Description != "" {
		payload.Description = data.Description
	}

	if data.ShortDescription != "" {
		payload.ShortDescription = data.ShortDescription
	}

	if data.GoalAmount != 0 {
		payload.GoalAmount = data.GoalAmount
	}

	_, err = s.repository.Update(payload)

	if err != nil {
		return err
	}

	return nil
}
