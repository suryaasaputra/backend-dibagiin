package middlewares

import "dibagi/repository"

type donationHistoryMiddleware struct {
	DonationHistoryRepository repository.IDonationHistoryRepository
}

func NewDonationHistoryMiddleware(donationHistoryRepository repository.IDonationHistoryRepository) *donationHistoryMiddleware {
	return &donationHistoryMiddleware{
		DonationHistoryRepository: donationHistoryRepository,
	}
}
