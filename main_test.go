package main

import (
	"testing"
)

func TestRunSend(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr bool
		errMsg  string
	}{
		{
			name:    "等号格式 --token=value",
			args:    []string{"--token=abc123", "--chat-id=123456", "--text=Hello"},
			wantErr: false,
		},
		{
			name:    "等号格式所有参数",
			args:    []string{"--token=YOUR_BOT_TOKEN", "--chat-id=YOUR_CHAT_ID", "--text=Hello World"},
			wantErr: false,
		},
		{
			name:    "空格格式 --token value",
			args:    []string{"--token", "abc123", "--chat-id", "123456", "--text", "Hello"},
			wantErr: false,
		},
		{
			name:    "混合格式",
			args:    []string{"--token=abc123", "--chat-id", "123456", "--text=Hello"},
			wantErr: false,
		},
		{
			name:    "缺少 token",
			args:    []string{"--chat-id=123456", "--text=Hello"},
			wantErr: true,
			errMsg:  "--token 是必填参数",
		},
		{
			name:    "缺少 chat-id",
			args:    []string{"--token=abc123", "--text=Hello"},
			wantErr: true,
			errMsg:  "--chat-id 是必填参数",
		},
		{
			name:    "缺少 text",
			args:    []string{"--token=abc123", "--chat-id=123456"},
			wantErr: true,
			errMsg:  "--text 是必填参数",
		},
		{
			name:    "未知选项",
			args:    []string{"--unknown=value", "--chat-id=123", "--text=Hello"},
			wantErr: true,
			errMsg:  "未知选项: --unknown",
		},
		{
			name:    "帮助选项 -h",
			args:    []string{"-h"},
			wantErr: false,
		},
		{
			name:    "帮助选项 --help",
			args:    []string{"--help"},
			wantErr: false,
		},
		{
			name:    "空等号值 --token=",
			args:    []string{"--token=", "--chat-id=123", "--text=Hello"},
			wantErr: true,
			errMsg:  "--token 是必填参数",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := runSend(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("runSend() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && tt.errMsg != "" && err != nil {
				if err.Error() != tt.errMsg {
					t.Errorf("runSend() error message = %v, want %v", err.Error(), tt.errMsg)
				}
			}
		})
	}
}

func TestRunSend_EmptyArgs(t *testing.T) {
	err := runSend([]string{})
	if err == nil {
		t.Error("runSend() with empty args should return error")
	}
}
