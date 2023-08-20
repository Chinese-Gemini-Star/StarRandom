package server

// {"left":"1","right":"10","number":"6","isRepeat":"true"}
type RandomInfo struct {
	Left     int  `json:"left"`
	Right    int  `json:"right"`
	Number   int  `json:"number"`
	IsRepeat bool `json:"isRepeat"`
}

// Check 表单验证
func (r *RandomInfo) Check() bool {
	if r.Left > r.Right {
		// 范围不成立
		return false
	}
	if r.Left < -1000 || r.Right > 1000 {
		// 范围不成立
		return false
	}
	if r.Number > 100 || r.Number < 1 {
		// 数量不成立
		return false
	}
	if !r.IsRepeat && r.Number > r.Right-r.Left+1 {
		// 范围内不可能不重复
		return false
	}
	return true
}
