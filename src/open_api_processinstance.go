package dingtalk

const (
	// ProcessInstanceStatusNew 审批状态 新创建
	ProcessInstanceStatusNew = "New"
	// ProcessInstanceStatusRunning 审批状态 审批中
	ProcessInstanceStatusRunning= "RUNNING"
	// ProcessInstanceStatusTerminated 审批状态 被终止
	ProcessInstanceStatusTerminated = "TERMINATED"
	// ProcessInstanceStatusCompleted 审批状态 完成
	ProcessInstanceStatusCompleted = "COMPLETED"
	// ProcessInstanceStatusCanceled 审批状态 取消
	ProcessInstanceStatusCanceled = "CANCELED"
)

const (
	// ProcessInstanceResultAgree 审批结果 同意
	ProcessInstanceResultAgree = "agree"
	// ProcessInstanceResultRefuse 审批状态 拒绝
	ProcessInstanceResultRefuse = "refuse"
)

// *************************获取审批实例详情***********************************

// ProcessInstanceGetResp 审批实例详情响应
type ProcessInstanceGetResp struct {
	OpenAPIResponse
	ProcessInstanceTopVo ProcessInstanceTopVo `json:"process_instance"`

}

// ProcessInstanceTopVo 审批实例详情
type ProcessInstanceTopVo struct {
	Title string `json:"title"`
	CreateTime string `json:"create_time"`
	FinishTime string `json:"finish_time"`
	OriginatorUserID string `json:"originator_userid"`
	OriginatorDeptID string `json:"originator_dept_id"`
	Status string `json:"status"`
	ApproverUserids []string `json:"approver_userids"`
	CCUserids []string `json:"cc_userids"`
	Result string `json:"result"`
	BusinessID string `json:"business_id"`
	OriginatorDeptName string `json:"originator_dept_name"`
	MainProcessInstanceID string `json:"main_process_instance_id"`
	ORMComponentValues []FormComponentValueVo `json:"orm_component_values"`
}

// FormComponentValueVo 表单详情
type FormComponentValueVo struct {
	Name string `json:"name"`
	Value string `json:"value"`
	ExtValue string `json:"ext_value"`
	ID string `json:"id"`
	ComponentType string `json:"component_type"`
}

// *************************添加审批评论***********************************

// ProcessInstanceCommentAddReq 添加审批评论 请求
type ProcessInstanceCommentAddReq struct {
	Request AddCommentRequest `json:"request"`
}

// AddCommentRequest 请求对象
type AddCommentRequest struct {
	ProcessInstanceID string `json:"process_instance_id"`
	Text string `json:"text"`
	CommentUserid string `json:"comment_userid"`
	File File `json:"file"`
}

// File 文件
type File struct {
	Photos []string `json:"photos"`
	attachments []Attachment `json:"attachments"`
}

// Attachment 附件
type Attachment struct {
	SpaceID string `json:"space_id"`
	FileType string `json:"file_type"`
	FileName string `json:"file_name"`
	FileID string `json:"file_id"`
	FileSize string `json:"file_size"`
}

// ProcessInstanceCommentAddResp 添加审批评论 响应
type ProcessInstanceCommentAddResp struct {
	OpenAPIResponse
	Success bool `json:"success"`
	Result bool `json:"result"`
}

// ProcessInstanceGet 获取审批实例详情
func (dtc *DingTalkClient) ProcessInstanceGet(processInstanceID string) (ProcessInstanceGetResp, error) {
	var data ProcessInstanceGetResp
	reqData := map[string]string{"process_instance_id": processInstanceID}
	err := dtc.httpRPC("topapi/processinstance/get", nil, reqData, &data)
	return data, err
}

// ProcessInstanceCommentAdd 添加审批评论
func (dtc *DingTalkClient) ProcessInstanceCommentAdd(info *ProcessInstanceCommentAddReq) (ProcessInstanceCommentAddResp, error) {
	var data ProcessInstanceCommentAddResp
	err := dtc.httpRPC("/topapi/process/instance/comment/add", nil, info, &data)
	return data, err
}