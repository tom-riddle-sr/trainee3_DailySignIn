package output

type SignInReward struct {
	ID   int32  `json:"Id"`
	Num  int32  `json:"num"`
	Type string `json:"type"`
}

type SignInRewards []SignInReward
