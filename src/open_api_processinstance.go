package dingtalk

import (
	"fmt"
	"strings"
	"time"
)

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
	// TaskResultAgree 同意
	TaskResultAgree = "AGREE"
	// TaskResultRefuse 拒绝
	TaskResultRefuse = "REFUSE"
	// TaskResultRedirected 转交
	TaskResultRedirected = "REDIRECTED"
)

const (
	// ProcessInstanceResultAgree 审批结果 同意
	ProcessInstanceResultAgree = "agree"
	// ProcessInstanceResultRefuse 审批状态 拒绝
	ProcessInstanceResultRefuse = "refuse"
)

const (
	// TaskStatusNew 未启动
	TaskStatusNew = "NEW"
	// TaskStatusRunning 处理中
	TaskStatusRunning = "RUNNING"
	// TaskStatusPaused 暂停
	TaskStatusPaused = "PAUSED"
	// TaskStatusCanceled 取消
	TaskStatusCanceled = "CANCELED"
	// TaskStatusCompleted 完成
	TaskStatusCompleted = "COMPLETED"
	// TaskStatusTerminated 终止
	TaskStatusTerminated = "TERMINATED"
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
	FormComponentValues []FormComponentValueVo `json:"form_component_values"`
	Tasks []TaskTopVo `json:"tasks"`
}

// TaskTopVo 任务
type TaskTopVo struct {
	UserID string `json:"userid"`
	TaskStatus string `json:"task_status"`
	TaskResult string `json:"task_result"`
	CreateTime string `json:"create_time"`
	FinishTime string `json:"finish_time"`
	TaskID string `json:"taskid"`
	URL string `json:"url"`
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

// *************************下载审批附件***********************************

// ProcessInstanceFileURLGetResp 下载审批附件 响应
type ProcessInstanceFileURLGetResp struct {
	OpenAPIResponse
	Success bool `json:"success"`
	Result AppSpaceResponse `json:"result"`
}

// AppSpaceResponse 文件内容
type AppSpaceResponse struct {
	FileID string `json:"file_id"`
	SpaceID string `json:"space_id"`
	DownloadURI string `json:"download_uri"`
}

// *************************获取审批实例ID列表***********************************

// ProcessInstanceListIdsResp 获取审批实例id列表响应
type ProcessInstanceListIdsResp struct {
	OpenAPIResponse
	Result ProcessInstanceListIdsPageResult `json:"result"`
}

// ProcessInstanceListIdsPageResult 分页数据
type ProcessInstanceListIdsPageResult struct {
	List []string `json:"list"`
	NextCursor int `json:"next_cursor"`
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

// ProcessInstanceFileURLGet 下载审批附件
func (dtc *DingTalkClient) ProcessInstanceFileURLGet(processInstanceID, fileID string) (ProcessInstanceFileURLGetResp, error) {
	var data ProcessInstanceFileURLGetResp
	reqData := map[string]map[string]string{
		"request": {
			"process_instance_id": processInstanceID,
			"file_id": fileID,
		},
	}
	err := dtc.httpRPC("/topapi/processinstance/file/url/get", nil, reqData, &data)
	return data, err
}

// ProcessInstanceListIds 获取审批实例id
func (dtc *DingTalkClient) ProcessInstanceListIds(processCode string, startTime, endTime *time.Time, size, cursor int, userIds ...string) (ProcessInstanceListIdsResp, error) {
	var data ProcessInstanceListIdsResp
	if processCode == "" || startTime == nil {
		return data, fmt.Errorf("no required params")
	}
	reqData := map[string]interface{}{
		"process_code": processCode,
		"start_time": startTime.Unix()*1000,
	}
	if endTime != nil {
		reqData["end_time"] = endTime.Unix()*1000
	}
	if size != 0 {
		reqData["size"] = size
	}
	if cursor != 0 {
		reqData["cursor"] = cursor
	}
	if len(userIds) > 0 {
		reqData["userid_list"] = strings.Join(userIds, ",")
	}
	err := dtc.httpRPC("/topapi/processinstance/listids", nil, reqData, &data)
	return data, err
}