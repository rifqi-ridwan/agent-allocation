package domain

type Agent struct {
	ID                   int  `json:"id"`
	IsAvailable          bool `json:"is_available"`
	CurrentCustomerCount bool `json:"current_customer_count"`
}

type AllAgent struct {
	Data struct {
		Agent struct {
			Data []Agent `json:"data"`
		} `json:"agents"`
	} `json:"data"`
}
