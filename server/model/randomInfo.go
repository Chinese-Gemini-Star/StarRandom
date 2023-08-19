package server

// {"left":"1","right":"10","number":"6","isRepeat":"true"}
type RandomInfo struct {
	Left     int  `json:"left"`
	Right    int  `json:"right"`
	Number   int  `json:"number"`
	IsRepeat bool `json:"isRepeat"`
}
