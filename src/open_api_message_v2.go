package dingtalk

// *************************发送工作通知***********************************

// MessageCorpConversationAsyncSendReq 发送工作通知  请求
type MessageCorpConversationAsyncSendReq struct {
	AgentID int `json:"agent_id"`
	UserIDList string `json:"userid_list"`
	DeptIDList string `json:"deptid_list"`
	ToAllUser bool `json:"to_all_user"`
	Msg map[string]interface{} `json:"msg"`
}

// MessageCorpConversationAsyncSendResp 发送工作通知 响应
type MessageCorpConversationAsyncSendResp struct {
	OpenAPIResponse
	TaskID int `json:"task_id"`
}

// MessageCorpConversationAsyncSend 发送工作通知
func (dtc *DingTalkClient) MessageCorpConversationAsyncSend(info *MessageCorpConversationAsyncSendReq) (MessageCorpConversationAsyncSendResp, error) {
	var data MessageCorpConversationAsyncSendResp
	err := dtc.httpRPC("topapi/message/corpconversation/asyncsend_v2", nil, info, &data)
	return data, err
}