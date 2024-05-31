package output

import (
	"trainee3/model/output/retStatus"
)

type CommonResponse struct {
	RetStatus retStatus.RetStatus `json:"retStatus"`      //狀態
	Data      interface{}         `json:"data,omitempty"` //資料
}
