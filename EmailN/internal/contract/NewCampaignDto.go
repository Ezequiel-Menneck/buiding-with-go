package contract

type NewCampaign struct {
	Name    string
	Content string
	Emails  []string
}

type CampaignResponse struct {
	Id      string
	Name    string
	Status  string
	Content string
}
