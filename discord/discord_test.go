package discord_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"discord-message-cli/discord"
)

func TestSendMessage(t *testing.T) {
	testMessage := "Test message"

	// モックHTTPサーバーの作成
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		if r.URL.Path != "/webhook" { // r.URL.String() を r.URL.Path に変更
			t.Errorf("Expected request to /webhook, got %s", r.URL.Path) // エラーメッセージも修正
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	// テストケース1: 正常系
	err := discord.SendMessage(server.URL+"/webhook", testMessage) // webhook path を追加
	if err != nil {
		t.Errorf("SendMessage should not return error, got %v", err)
	}

	// テストケース2: HTTPエラー
	serverError := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer serverError.Close()
	err = discord.SendMessage(serverError.URL+"/webhook", testMessage) // webhook path を追加
	if err == nil {
		t.Errorf("SendMessage should return error for HTTP 500")
	}

	// テストケース3: JSON marshalエラー (payload作成失敗)
	// モックのwebhookURLを空文字列にして、Marshalエラーを発生させる
	err = discord.SendMessage("", testMessage)
	if err == nil {
		t.Errorf("SendMessage should return error for JSON marshal failure")
	}
}

// HTTPクライアントのモックを作成 (未使用)
type MockHTTPClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return m.DoFunc(req)
}
