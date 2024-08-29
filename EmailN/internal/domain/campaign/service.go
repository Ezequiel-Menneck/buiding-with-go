package campaign

import (
	"emailn/internal/contract"
	internalerrors "emailn/internal/internal-errors"
	"errors"
)

type Service interface {
	Create(newCampaign contract.NewCampaign) (string, error)
	GetById(id string) (*contract.CampaignResponse, error)
	Delete(id string) error
}

type ServiceImp struct {
	Repository Repository
}

func (s *ServiceImp) Create(newCampaign contract.NewCampaign) (string, error) {
	campaign, err := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)
	if err != nil {
		return "", err
	}
	err = s.Repository.Create(campaign)
	if err != nil {
		return "", internalerrors.ErrInternal
	}

	return campaign.ID, nil
}

func (s *ServiceImp) GetById(id string) (*contract.CampaignResponse, error) {
	campaign, err := s.Repository.GetById(id)
	if err != nil {
		return nil, internalerrors.ProcessErrorToReturn(err)
	}

	if campaign == nil {
		return nil, nil
	}
	return &contract.CampaignResponse{
		Id:                   campaign.ID,
		Name:                 campaign.Name,
		Status:               campaign.Status,
		Content:              campaign.Content,
		AmountOfEmailsToSend: len(campaign.Contacts),
	}, nil
}

func (s *ServiceImp) Delete(id string) error {
	campaign, err := s.Repository.GetById(id)
	if err != nil {
		return internalerrors.ProcessErrorToReturn(err)
	}

	if campaign.Status != Pending {
		return errors.New("campaign status invalid")
	}

	campaign.Delete()
	if err = s.Repository.Delete(campaign); err != nil {
		return internalerrors.ErrInternal
	}

	return err
}
