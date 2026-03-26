package main

import (
	"os"
	"path/filepath"
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
			name:    "缺少 text 和 text-file",
			args:    []string{"--token=abc123", "--chat-id=123456"},
			wantErr: true,
			errMsg:  "--text 或 --text-file 必须指定一个",
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

func TestRunSend_TextFile(t *testing.T) {
	// 创建临时目录
	tmpDir := t.TempDir()

	// 创建测试文件
	testFile := filepath.Join(tmpDir, "test_message.txt")
	testContent := "Hello from file!"
	if err := os.WriteFile(testFile, []byte(testContent), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// 创建不存在的文件路径
	nonExistentFile := filepath.Join(tmpDir, "non_existent.txt")

	tests := []struct {
		name    string
		args    []string
		wantErr bool
		errMsg  string
	}{
		{
			name:    "使用 --text-file 等号格式",
			args:    []string{"--token=abc123", "--chat-id=123456", "--text-file=" + testFile},
			wantErr: false,
		},
		{
			name:    "使用 --text-file 空格格式",
			args:    []string{"--token=abc123", "--chat-id", "123456", "--text-file", testFile},
			wantErr: false,
		},
		{
			name:    "--text 和 --text-file 同时使用",
			args:    []string{"--token=abc123", "--chat-id=123456", "--text=Hello", "--text-file=" + testFile},
			wantErr: true,
			errMsg:  "--text 和 --text-file 不能同时使用",
		},
		{
			name:    "--text-file 文件不存在",
			args:    []string{"--token=abc123", "--chat-id=123456", "--text-file=" + nonExistentFile},
			wantErr: true,
			errMsg:  "读取文件失败:",
		},
		{
			name:    "--text-file 值为空",
			args:    []string{"--token=abc123", "--chat-id=123456", "--text-file="},
			wantErr: true,
			errMsg:  "--text 或 --text-file 必须指定一个",
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
				// 对于文件错误，只检查前缀
				if !containsPrefix(err.Error(), tt.errMsg) {
					t.Errorf("runSend() error message = %v, want prefix %v", err.Error(), tt.errMsg)
				}
			}
		})
	}
}

func containsPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}
