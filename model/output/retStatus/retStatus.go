package retStatus

import apicode "trainee3/lib/apiCode"

type RetStatus struct {
	Code apicode.Code `json:"code"` //狀態碼
	Msg  string       `json:"msg"`  //訊息
}

func New(code apicode.Code) RetStatus {
	return RetStatus{
		Code: code,
		Msg:  code.GetMsg(),
	}
}
