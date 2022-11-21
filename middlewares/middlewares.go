package middlewares

type Middleware struct {
	UserMiddleware     *userMiddleware
	DonationMiddleware *donationMiddleware
}

func NewMiddleware(userMiddleware *userMiddleware, donationMiddleware *donationMiddleware) Middleware {
	return Middleware{
		UserMiddleware:     userMiddleware,
		DonationMiddleware: donationMiddleware,
	}
}
