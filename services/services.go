package services

type Services struct {
	Activity IActivity
	Refresh  ICache
}

func New(activity IActivity, refresh ICache) *Services {
	return &Services{
		Activity: activity,
		Refresh:  refresh,
	}
}
