package middlewares

type Middleware struct {
	UserMiddleware            *userMiddleware
	DonationMiddleware        *donationMiddleware
	DonationRequestMiddleware *donationRequestMiddleware
	NotificationMiddleware    *notificationMiddleware
}

func NewMiddleware(userMiddleware *userMiddleware, donationMiddleware *donationMiddleware, donationRequestMiddleware *donationRequestMiddleware, notificationMiddleware *notificationMiddleware) Middleware {
	return Middleware{
		UserMiddleware:            userMiddleware,
		DonationMiddleware:        donationMiddleware,
		DonationRequestMiddleware: donationRequestMiddleware,
		NotificationMiddleware:    notificationMiddleware,
	}
}
