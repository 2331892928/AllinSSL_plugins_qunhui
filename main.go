package main

import (
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"os"
)

type ActionInfo struct {
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Params      map[string]any `json:"params,omitempty"` // 可选参数
}

type Request struct {
	Action string                 `json:"action"`
	Params map[string]interface{} `json:"params"`
}

type Response struct {
	Status  string                 `json:"status"`
	Message string                 `json:"message"`
	Result  map[string]interface{} `json:"result"`
}

var pluginMeta = map[string]interface{}{
	"name":        "qunhui",
	"description": "部署到群晖",
	"version":     "1.0.0",
	"author":      "AMEN",
	"config": map[string]interface{}{
		"SYNO_SCHEME":   "群晖管理面板URL的协议：http、https，用于群晖API调用",
		"SYNO_HOSTNAME": "群晖管理面板URL的域名或者IP，用于群晖API调用",
		"SYNO_PORT":     "群晖管理面板URL的端口号，用于群晖API调用",
		"SYNO_USERNAME": "群晖NAS用户账号，由于证书自动部署需要使用API上传证书至群晖系统中，而群晖系统中只有管理员有权限上传证书，所以该用户账号必须要在administrators群组中，并且不能开启双重验证。建议单独创建一个专有用户账号用于证书的自动部署，并为该用户账号设置16位以上复杂的密码。建议对该账号禁用所有文件夹读写权限以及所有应用的访问权限，提升NAS数据的安全性",
		"SYNO_PASSWORD": "群晖NAS用户密码，建议使用16位以上复杂的密码",
	},
	"actions": []ActionInfo{
		{
			Name:        "deploy",
			Description: "部署到群晖",
			Params: map[string]any{
				"AS_DEFAULT": "是否删除重复的域名证书，默认 false",
			},
		},
	},
}

// **解析 PEM 格式的证书**
func ParseCertificate(certPEM []byte) (*x509.Certificate, error) {
	block, _ := pem.Decode(certPEM)
	if block == nil {
		return nil, fmt.Errorf("无法解析证书 PEM")
	}
	return x509.ParseCertificate(block.Bytes)
}

func outputJSON(resp *Response) {
	_ = json.NewEncoder(os.Stdout).Encode(resp)
}

func outputError(msg string, err error) {
	outputJSON(&Response{
		Status:  "error",
		Message: fmt.Sprintf("%s: %v", msg, err),
	})
}

func main() {
	var req Request
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		outputError("读取输入失败", err)
		return
	}

	if err := json.Unmarshal(input, &req); err != nil {
		outputError("解析请求失败", err)
		return
	}

	switch req.Action {
	case "get_metadata":
		outputJSON(&Response{
			Status:  "success",
			Message: "插件信息",
			Result:  pluginMeta,
		})
	case "list_actions":
		outputJSON(&Response{
			Status:  "success",
			Message: "支持的动作",
			Result:  map[string]interface{}{"actions": pluginMeta["actions"]},
		})
	case "deploy":
		rep, err := deploy(req.Params)
		if err != nil {
			outputError("部署失败", err)
			return
		}
		outputJSON(rep)
	default:
		outputJSON(&Response{
			Status:  "error",
			Message: "未知 action: " + req.Action,
		})
	}
}
