package synapse

import (
	"context"
	"fmt"
	"net/http"

	"github.com/xh-polaris/inno_agent/biz/conf"
	"github.com/xh-polaris/inno_agent/biz/pkg/httpcli"
)

type httpClientImpl struct{}

func NewHTTPClient() Client {
	return &httpClientImpl{}
}

func synapseURL() string {
	return conf.GetConfig().SynapseURL
}

func buildHeader() http.Header {
	header := http.Header{}
	header.Set("content-type", "application/json")
	if conf.GetConfig().State != "test" {
		header.Set("X-Xh-Env", "test")
	}
	return header
}

func (c *httpClientImpl) Login(ctx context.Context, authType, authId, verify string) (*LoginResult, error) {
	header := buildHeader()
	body := map[string]any{
		"authType": authType,
		"authId":   authId,
		"verify":   verify,
		"app":      map[string]any{"name": "InnoAgent"},
	}
	url := synapseURL() + "/basic_user/login"
	resp, err := httpcli.GetHttpClient().Req("POST", url, header, body)
	if err != nil {
		return nil, fmt.Errorf("synapse login: %w", err)
	}
	if code, _ := resp["code"].(float64); code != 0 {
		msg, _ := resp["msg"].(string)
		return nil, fmt.Errorf("synapse login error %d: %s", int(code), msg)
	}
	basicUser, _ := resp["basicUser"].(map[string]any)
	basicUserId, _ := basicUser["basicUserId"].(string)
	token, _ := resp["token"].(string)
	isNew, _ := resp["new"].(bool)
	return &LoginResult{
		BasicUserId: basicUserId,
		Token:       token,
		IsNew:       isNew,
	}, nil
}

func (c *httpClientImpl) Register(ctx context.Context, authType, authId, verify, password string) (*RegisterResult, error) {
	header := buildHeader()
	body := map[string]any{
		"authType": authType,
		"authId":   authId,
		"verify":   verify,
		"password": password,
		"app":      map[string]any{"name": "InnoAgent"},
	}
	url := synapseURL() + "/basic_user/register"
	resp, err := httpcli.GetHttpClient().Req("POST", url, header, body)
	if err != nil {
		return nil, fmt.Errorf("synapse register: %w", err)
	}
	if code, _ := resp["code"].(float64); code != 0 {
		msg, _ := resp["message"].(string)
		return nil, fmt.Errorf("synapse register error %d: %s", int(code), msg)
	}
	token, _ := resp["token"].(string)
	return &RegisterResult{Token: token}, nil
}

func (c *httpClientImpl) ResetPassword(ctx context.Context, authHeader, newPassword string) error {
	header := buildHeader()
	header.Set("Authorization", authHeader)
	body := map[string]any{
		"newPassword": newPassword,
		"app":         map[string]any{"name": "InnoAgent"},
	}
	url := synapseURL() + "/basic_user/reset_password"
	resp, err := httpcli.GetHttpClient().Req("POST", url, header, body)
	if err != nil {
		return fmt.Errorf("synapse reset_password: %w", err)
	}
	if code, _ := resp["code"].(float64); code != 0 {
		msg, _ := resp["message"].(string)
		return fmt.Errorf("synapse reset_password error %d: %s", int(code), msg)
	}
	return nil
}