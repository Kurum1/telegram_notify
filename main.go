package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const (
	version = "1.0.0"
)

func main() {
	if len(os.Args) < 2 {
		showMainHelp()
		os.Exit(1)
	}

	cmd := os.Args[1]

	switch cmd {
	case "-h", "--help":
		showMainHelp()
		os.Exit(0)
	case "-v", "--version":
		fmt.Printf("telegram_notify version %s\n", version)
		os.Exit(0)
	case "send":
		if err := runSend(os.Args[2:]); err != nil {
			fmt.Fprintf(os.Stderr, "错误: %v\n", err)
			os.Exit(1)
		}
	default:
		if strings.HasPrefix(cmd, "-") {
			fmt.Fprintf(os.Stderr, "未知选项: %s\n\n", cmd)
		} else {
			fmt.Fprintf(os.Stderr, "未知命令: %s\n\n", cmd)
		}
		showMainHelp()
		os.Exit(1)
	}
}

func showMainHelp() {
	fmt.Fprintf(os.Stderr, "用法: %s <命令> [选项]\n\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "一个用于发送 Telegram 消息的命令行工具\n\n")
	fmt.Fprintf(os.Stderr, "全局选项:\n")
	fmt.Fprintf(os.Stderr, "  -h|--help      显示帮助信息\n")
	fmt.Fprintf(os.Stderr, "  -v|--version   显示版本信息\n\n")
	fmt.Fprintf(os.Stderr, "命令:\n")
	fmt.Fprintf(os.Stderr, "  send           发送消息到 Telegram\n\n")
	fmt.Fprintf(os.Stderr, "使用 '%s send --help' 查看 send 命令的详细用法\n", os.Args[0])
}

func runSend(args []string) error {
	var token, chatID, text string
	var showHelp bool

	// 解析参数
	for i := 0; i < len(args); i++ {
		arg := args[i]
		switch arg {
		case "-h", "--help":
			showHelp = true
		case "--token":
			if i+1 >= len(args) {
				return fmt.Errorf("--token 需要参数值")
			}
			token = args[i+1]
			i++
		case "--chat-id":
			if i+1 >= len(args) {
				return fmt.Errorf("--chat-id 需要参数值")
			}
			chatID = args[i+1]
			i++
		case "--text":
			if i+1 >= len(args) {
				return fmt.Errorf("--text 需要参数值")
			}
			text = args[i+1]
			i++
		default:
			if strings.HasPrefix(arg, "-") {
				return fmt.Errorf("未知选项: %s", arg)
			}
			// 如果 text 未设置，将剩余参数作为 text
			if text == "" {
				text = arg
			}
		}
	}

	if showHelp {
		showSendHelp()
		return nil
	}

	// 验证必填参数
	if token == "" {
		return fmt.Errorf("--token 是必填参数")
	}
	if chatID == "" {
		return fmt.Errorf("--chat-id 是必填参数")
	}
	if text == "" {
		return fmt.Errorf("--text 是必填参数")
	}

	return sendMessage(token, chatID, text)
}

func showSendHelp() {
	fmt.Fprintf(os.Stderr, "用法: %s send [选项]\n\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "发送消息到 Telegram Bot\n\n")
	fmt.Fprintf(os.Stderr, "选项:\n")
	fmt.Fprintf(os.Stderr, "  --token <token>      Bot Token (必填)\n")
	fmt.Fprintf(os.Stderr, "  --chat-id <id>       Chat ID (必填)\n")
	fmt.Fprintf(os.Stderr, "  --text <message>     消息内容 (必填)\n")
	fmt.Fprintf(os.Stderr, "  -h|--help            显示帮助信息\n\n")
	fmt.Fprintf(os.Stderr, "示例:\n")
	fmt.Fprintf(os.Stderr, "  %s send --token=\"YOUR_BOT_TOKEN\" --chat-id=\"YOUR_CHAT_ID\" --text=\"Hello World\"\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "  %s send --token \"TOKEN\" --chat-id \"ID\" --text \"Hello World\"\n", os.Args[0])
}

func sendMessage(token, chatID, text string) error {
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)

	// 构造 form data（与 curl -d 相同）
	data := url.Values{}
	data.Set("chat_id", chatID)
	data.Set("text", text)

	resp, err := http.Post(apiURL, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
	if err != nil {
		return fmt.Errorf("请求失败: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("读取响应失败: %w", err)
	}

	// 直接输出 Telegram API 响应内容
	fmt.Println(string(body))
	return nil
}
