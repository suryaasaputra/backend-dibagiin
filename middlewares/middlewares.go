package middlewares

type Middleware struct {
	UserMiddleware            *userMiddleware
	DonationMiddleware        *donationMiddleware
	DonationRequestMiddleware *donationRequestMiddleware
	DonationHistoryMiddleware *donationHistoryMiddleware
}

func NewMiddleware(userMiddleware *userMiddleware, donationMiddleware *donationMiddleware, donationRequestMiddleware *donationRequestMiddleware, donationHistoryMiddleware *donationHistoryMiddleware) Middleware {
	return Middleware{
		UserMiddleware:            userMiddleware,
		DonationMiddleware:        donationMiddleware,
		DonationRequestMiddleware: donationRequestMiddleware,
		DonationHistoryMiddleware: donationHistoryMiddleware,
	}
}
