package controllers

import "dibagi/repository"

type donationHistoryController struct {
	DonationHistoryRepository repository.IDonationHistoryRepository
}

func NewDonationHistoryController(dr repository.IDonationHistoryRepository) *donationHistoryController {
	return &donationHistoryController{
		DonationHistoryRepository: dr,
	}
}
