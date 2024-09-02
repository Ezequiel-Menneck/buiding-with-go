package contract

type NewCampaign struct {
	Name      string
	Content   string
	Emails    []string
	CreatedBy string
}

type CampaignResponse struct {
	Id                   string
	Name                 string
	Status               string
	Content              string
	AmountOfEmailsToSend int
	CreatedBy            string
}
