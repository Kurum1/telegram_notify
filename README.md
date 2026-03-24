# telegram_notify

一个用于发送 Telegram 消息的命令行工具，支持 Windows 可执行文件编译。

## 功能

- 发送消息到指定 Telegram Chat
- 子命令模式设计

## 安装

```bash
go build -o telegram_notify.exe
```

## 用法

```bash
# 显示帮助
telegram_notify --help

# 显示版本
telegram_notify --version

# 发送消息
telegram_notify send --token="YOUR_BOT_TOKEN" --chat-id="YOUR_CHAT_ID" --text="Hello World"

# 或者使用空格分隔参数
telegram_notify send --token "TOKEN" --chat-id "ID" --text "Hello World"
```

## 命令

| 命令 | 说明 |
|------|------|
| `send` | 发送消息到 Telegram |

## send 命令参数

| 参数 | 说明 | 必填 |
|------|------|------|
| `--token` | Bot Token | 是 |
| `--chat-id` | Chat ID | 是 |
| `--text` | 消息内容 | 是 |
| `-h\|--help` | 显示帮助 | 否 |

## 编译 Windows 可执行文件

```bash
# 在 macOS/Linux 上交叉编译 Windows exe
make build-windows

# 或者手动编译
GOOS=windows GOARCH=amd64 go build -o telegram_notify.exe
```
