package request

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type ReqMethod string
type ContentType string

const (
	timeOut                   = 5
	Get           ReqMethod   = "GET"
	Post          ReqMethod   = "POST"
	Put           ReqMethod   = "PUT"
	Delete        ReqMethod   = "DELETE"
	JsonType      ContentType = "application/json"
	FormType      ContentType = "application/x-www-form-urlencoded"
	MultipartType ContentType = "multipart/form-data"
)

type HttpClient struct {
	http.Client
	Method  ReqMethod         // 请求方法
	ReqType ContentType       // 请求类型 json form raw
	RspType ContentType       // 返回类型
	Headers map[string]string // 设置header头
	Body    map[string]interface{}
}

// NewHttpClient 获取http实例，默认为get方法，json请求类型
func NewHttpClient() *HttpClient {
	return &HttpClient{http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       time.Second * timeOut,
	}, Get, JsonType, JsonType, nil, nil}
}

// Cal 发起请求
func (h *HttpClient) Cal(ctx context.Context, requestUrl string, target interface{}) error {
	reqBody, err := json.Marshal(h.Body)
	if err != nil {
		return err
	}
	request, err := http.NewRequest(string(h.Method), requestUrl, bytes.NewReader(reqBody))
	if err != nil {
		return err
	}
	request.Header.Set("Accept", string(h.RspType))
	request.Header.Set("Content-Type", string(h.ReqType))
	for k, v := range h.Headers {
		request.Header.Set(k, v)
	}
	response, err := h.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	rspBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(rspBody, &target); err != nil {
		return err
	}

	return nil
}

func (h *HttpClient) SetMethod(method ReqMethod) {
	h.Method = method
}

// 设置请求body
func (h *HttpClient) SetBody(body map[string]interface{}) {
	h.Body = body
}

// 设置超时时间
func (h *HttpClient) SetTimeOut(timeOut int) {
	h.Client.Timeout = time.Second * time.Duration(timeOut)
}

// 设置请求头类型
func (h *HttpClient) SetReqContentType(contentType ContentType) {
	h.ReqType = contentType
}

// 设置response返回类型
func (h *HttpClient) SetRspContentType(contentType ContentType) {
	h.RspType = contentType
}
