package serializer

type ResponseUser struct {
	Status int    `json:"status" exmaple:"200"`
	Data   User   `json:"data"`
	Msg    string `json:"msg" exmaple:"ok"`
	Error  string `json:"error" exmaple:""`
}

type ResponseTask struct {
	Status int    `json:"status" exmaple:"200"`
	Data   Task   `json:"data"`
	Msg    string `json:"msg" exmaple:"ok"`
	Error  string `json:"error" exmaple:""`
}
