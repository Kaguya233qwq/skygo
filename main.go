package main

import (
	"os"
	_ "skygo/plugins"
	"skygo/utils"

	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/driver"
)

func main() {
	botName := os.Getenv(utils.BotName)
	superUserIDs := utils.GetSuperUsers()
	accessToken := os.Getenv(utils.AccessToken)
	CommandPrefix := os.Getenv(utils.CommandPrefix)

	// 动态选择 WebSocket 驱动
	var wsDriver zero.Driver
	serverURL := os.Getenv(utils.WebSocketServerUrl)
	if serverURL != "" {
		wsDriver = driver.NewWebSocketServer(16, serverURL, string(accessToken))
	} else {
		clientURL := os.Getenv(utils.WebSocketClientUrl)
		wsDriver = driver.NewWebSocketClient(clientURL, string(accessToken))
	}

	zero.RunAndBlock(&zero.Config{
		NickName:      []string{botName},
		CommandPrefix: CommandPrefix,
		SuperUsers:    superUserIDs,
		Driver: []zero.Driver{
			wsDriver,
		},
	}, nil)
}
