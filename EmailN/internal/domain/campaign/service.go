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
	Start(id string) error
}

type ServiceImp struct {
	Repository Repository
	SendMail   func(campaign *Campaign) error
}

func (s *ServiceImp) Create(newCampaign contract.NewCampaign) (string, error) {
	campaign, err := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails, newCampaign.CreatedBy)
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
		CreatedBy:            campaign.CreatedBy,
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

func (s *ServiceImp) SendEmailAndUpdateStatus(campaignSaved *Campaign) {
	err := s.SendMail(campaignSaved)
	if err != nil {
		campaignSaved.Fail()
	} else {
		campaignSaved.Done()
	}
	err = s.Repository.Update(campaignSaved)
}

func (s *ServiceImp) Start(id string) error {
	campaignSaved, err := s.Repository.GetById(id)
	if err != nil {
		return internalerrors.ProcessErrorToReturn(err)
	}

	if campaignSaved.Status != Pending {
		return errors.New("campaignSaved status invalid")
	}

	go s.SendEmailAndUpdateStatus(campaignSaved)

	campaignSaved.Started()
	err = s.Repository.Update(campaignSaved)
	if err != nil {
		return internalerrors.ErrInternal
	}

	return nil
}
