package sdk_baidu

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/baidubce/app-builder/go/appbuilder"
	"github.com/kysion/base-library/utility/kconv"
	"io"
	"os"
)

// 大模型平台： [百度智能云千帆AppBuilder (baidu.com)](https://console.bce.baidu.com/ai_apaas/dialogHome)

// GetAiSummary 文字AI分析总结
func (s *sSdkBaidu) GetAiSummary(ctx context.Context, text string, identifier string) (string, error) {

	// 根据业务标识查找应用配置
	conf, err := s.GetBaiduSdkConf(ctx, identifier)
	if err != nil {
		return "", err
	}

	result, err := s.StartConversation(ctx, conf.AppID, conf.AESKey, text, "")
	if err != nil {
		return "", sys_service.SysLogs().ErrorSimple(ctx, err, "百度智能体-获取文字AI分析总结失败", "SDK-Baidu-GetAiSummary")
	}

	return result, nil
}

// StartConversation 开始AI会话，并返回智能体的回答
func (s *sSdkBaidu) StartConversation(ctx context.Context, appId string, appBuilderToken string, text string, filePath string) (string, error) {
	// 设置APPBUILDER_TOKEN、GATEWAY_URL_V2环境变量
	_ = os.Setenv("APPBUILDER_TOKEN", appBuilderToken)
	// 默认可不填，默认值是 https://qianfan.baidubce.com
	_ = os.Setenv("GATEWAY_URL_V2", "")
	config, err := appbuilder.NewSDKConfig("", "")
	if err != nil {
		fmt.Println("new config failed: ", err)
		return "", err
	}

	// 初始化实例
	appID := appId
	builder, err := appbuilder.NewAppBuilderClient(appID, config)
	if err != nil {
		fmt.Println("new agent builder failed: ", err)
		return "", err
	}

	// 创建对话ID
	conversationID, err := builder.CreateConversation()
	if err != nil {
		fmt.Println("create conversation failed: ", err)
		return "", err
	}

	// 上传文件 (如果不需要引用特定的文档，则不需要) 【可选】
	// 与创建应用时绑定的知识库不同之处在于，
	// 所上传文件仅在本次会话ID下发生作用，如果创建新的会话ID，上传的文件自动失效
	// 而知识库在不同的会话ID下均有效
	var fileID string
	if filePath != "" {
		fileID, err = builder.UploadLocalFile(conversationID, filePath)
		if err != nil {
			fmt.Println("upload local file failed:", err)
			return "", err
		}
	}

	// 执行流式对话
	// 注意file_ids不是必填项，如果不需要引用特定的文档，则将[]string{fileID}更换为nil即可
	// 同时还需要将上文的fileID, err := builder.UploadLocalFile(conversationID,  "/path/to/cv.pdf")代码
	// 更换为 _, err = client.UploadLocalFile(conversationID,  "/path/to/cv.pdf"),否则会报错
	var i appbuilder.AppBuilderClientIterator
	if fileID != "" {
		i, err = builder.Run(conversationID, text, []string{fileID}, true)
	} else {
		i, err = builder.Run(conversationID, text, nil, true)
	}
	if err != nil {
		fmt.Println("run failed: ", err)
		return "", err
	}

	// 解析智能体响应数据，解析成为本的回答
	completedAnswer := ""
	var answer *appbuilder.AppBuilderClientAnswer
	for answer, err = i.Next(); err == nil; answer, err = i.Next() {
		completedAnswer = completedAnswer + answer.Answer

		// 打印响应数据 【可注释】
		//for _, ev := range answer.Events {
		//	printResponse(&ev)
		//}
	}

	// 迭代正常结束err应为io.EOF
	if errors.Is(err, io.EOF) {
		fmt.Println("run success")
		fmt.Println("智能体回答内容： ", completedAnswer)
	} else {
		fmt.Println("run failed:", err)
	}

	return completedAnswer, nil
}

// printResponse 解析智能体响应数据，解析成为本的回答
func printResponse(ev *appbuilder.Event) {
	if ev.ContentType == appbuilder.TextContentType {
		var detail appbuilder.TextDetail
		rawData, _ := getRawData(ev)
		_ = json.Unmarshal(rawData, &detail)
		fmt.Println(detail)

	} else if ev.ContentType == appbuilder.CodeContentType {
		var detail appbuilder.CodeDetail
		rawData, _ := getRawData(ev)
		_ = json.Unmarshal(rawData, &detail)
		fmt.Println(detail)

	} else if ev.ContentType == appbuilder.ImageContentType {
		var detail appbuilder.ImageDetail
		rawData, _ := getRawData(ev)
		_ = json.Unmarshal(rawData, &detail)
		fmt.Println(detail)

	} else if ev.ContentType == appbuilder.RAGContentType {
		var detail appbuilder.RAGDetail
		rawData, _ := getRawData(ev)
		_ = json.Unmarshal(rawData, &detail)
		fmt.Println(detail)

	} else if ev.ContentType == appbuilder.FunctionCallContentType {
		var detail appbuilder.FunctionCallDetail
		rawData, _ := getRawData(ev)
		_ = json.Unmarshal(rawData, &detail)
		fmt.Println(detail)

	} else if ev.ContentType == appbuilder.AudioContentType {
		var detail appbuilder.AudioDetail
		rawData, _ := getRawData(ev)
		_ = json.Unmarshal(rawData, &detail)
		fmt.Println(detail)

	} else if ev.ContentType == appbuilder.VideoContentType {
		var detail appbuilder.VideoDetail
		rawData, _ := getRawData(ev)
		_ = json.Unmarshal(rawData, &detail)

	} else if ev.ContentType == appbuilder.StatusContentType {
		var detail appbuilder.StatusDetail
		rawData, _ := getRawData(ev)
		_ = json.Unmarshal(rawData, &detail)
		fmt.Println(detail)

	} else { // 默认detail
		var detail appbuilder.DefaultDetail
		rawData, _ := getRawData(ev)
		_ = json.Unmarshal(rawData, &detail)
		fmt.Println(detail)

	}
}

func getRawData(ev *appbuilder.Event) ([]byte, error) {
	var rawData []byte
	switch v := ev.Detail.(type) {
	case json.RawMessage:
		rawData = []byte(v)
	case interface{}:
		rawData, _ = json.Marshal(v)
	default:
		return nil, fmt.Errorf("unsupported type: %T", v)
	}
	return rawData, nil
}

// GetAppList 获取应用列表 【暂未对接】
func (s *sSdkBaidu) GetAppList(ctx context.Context) (*sys_model.AppBuilderAppListRes, error) {

	config := appbuilder.SDKConfig{
		GatewayURL:            "",
		GatewayURLV2:          "",
		ConsoleOpenAPIVersion: "",
		ConsoleOpenAPIPrefix:  "",
		SecretKey:             "",
		HTTPClient:            nil,
	}

	apps, err := appbuilder.GetAppList(appbuilder.GetAppListRequest{
		Limit: 100, // 	返回结果的最大数量，默认值为10, 最大值为100
	}, &config)
	if err != nil {
		fmt.Printf("get apps failed: %v", err)
	}
	fmt.Println(len(apps))

	res := kconv.Struct(apps, &sys_model.AppBuilderAppListRes{})

	return res, err
}
