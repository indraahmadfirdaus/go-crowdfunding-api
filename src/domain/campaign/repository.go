package campaign

import (
	"crowdfunding-api/src/kernel"
	"errors"
	"fmt"
)

type Repository interface {
	FindAll() ([]Campaign, error)
	FindByUserID(userID int) ([]Campaign, error)
	FindById(ID int) (Campaign, error)
	Save(campaign Campaign) (Campaign, error)
	Update(campaign Campaign) (Campaign, error)
}

type repository struct {
}

func NewRepository() *repository {
	return &repository{}
}

func (r *repository) FindAll() ([]Campaign, error) {
	var campaigns []Campaign
	fmt.Println("here 3")

	if kernel.DB == nil {
		return nil, errors.New("nil database connection")
	}

	err := kernel.DB.
		Preload("CampaignImages", "campaign_images.is_primary = true").
		Preload("User").
		Find(&campaigns).
		Error

	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r *repository) FindByUserID(userID int) ([]Campaign, error) {
	var campaigns []Campaign
	err := kernel.DB.
		Where("user_id = ?", userID).
		Preload("CampaignImages", "campaign_images.is_primary = true").
		Preload("User").
		Find(&campaigns).
		Error

	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r *repository) FindById(ID int) (Campaign, error) {
	var campaign Campaign
	err := kernel.DB.
		Where("id = ?", ID).
		Preload("CampaignImages").
		Preload("User").
		Find(&campaign).
		Error

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (r *repository) Save(campaign Campaign) (Campaign, error) {
	err := kernel.DB.Create(&campaign).Error
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (r *repository) Update(campaign Campaign) (Campaign, error) {
	err := kernel.DB.Save(&campaign).Error
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}
