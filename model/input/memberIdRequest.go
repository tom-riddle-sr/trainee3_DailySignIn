package input

type MemberIdRequest struct {
	MemberID int32 `json:"memberId" validate:"required"`
}
