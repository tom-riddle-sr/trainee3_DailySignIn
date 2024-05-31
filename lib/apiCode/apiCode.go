package apiCode

type Code int

const (
	Success         Code = 10000000
	AlreadySignIn   Code = 20000000
	UnknownError    Code = 30000000
	ValidateError   Code = 40000000
	ActivityNotOpen Code = 50000000
)

var resList = map[Code]string{
	Success:         "Success",
	AlreadySignIn:   "Data Exist",
	UnknownError:    "Unknown Error",
	ValidateError:   "Validate Error",
	ActivityNotOpen: "Activity Not Open",
}

func (c Code) GetMsg() string {
	return resList[c]
}
