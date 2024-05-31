package handlers

type Handlers struct {
	Auth  ISignIn
	Cache ICache
}

func New(signIn ISignIn, refresh ICache) *Handlers {
	return &Handlers{
		Auth:  signIn,
		Cache: refresh,
	}
}
