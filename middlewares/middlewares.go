package middlewares

type Middleware struct {
	UserMiddleware            *userMiddleware
	DonationMiddleware        *donationMiddleware
	DonationRequestMiddleware *donationRequestMiddleware
}

func NewMiddleware(userMiddleware *userMiddleware, donationMiddleware *donationMiddleware, donationRequestMiddleware *donationRequestMiddleware) Middleware {
	return Middleware{
		UserMiddleware:            userMiddleware,
		DonationMiddleware:        donationMiddleware,
		DonationRequestMiddleware: donationRequestMiddleware,
	}
}
