# SkyGo ✨

[![Go Version](https://img.shields.io/badge/Go-1.20%2B-blue.svg)](https://go.dev/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Release](https://img.shields.io/github/v/release/Kaguya233qwq/skygo)](https://github.com/Kaguya233qwq/skygo/releases)
[![Platform](https://img.shields.io/badge/Platform-Windows%20%7C%20Linux%20%7C%20macOS-lightgrey)](https://github.com/Kaguya233qwq/skygo/releases)

一个基于 [ZeroBot](https://github.com/wdvxdr1123/ZeroBot)，使用 Go 语言开发的《光·遇》游戏机器人。

旨在提供一个快速、轻量、跨平台的光遇机器人解决方案，易于部署和使用。

## 🚀 特性

-   **快速响应**: 基于 Go 语言和高性能的 ZeroBot 框架，性能卓越，资源占用低。
-   **简单易用**: 只需填写几个配置项即可启动，无需复杂的部署流程。
-   **跨平台**: 得益于go的交叉编译功能，项目集成一键跨平台编译脚本，可直接在 Windows, Linux, macOS 上编译和运行。

## ⚡️ 快速开始

对于大多数用户，推荐直接从 **Releases** 页面下载预编译好的程序。

#### 1. 下载程序

前往项目 [Releases 页面](https://github.com/YOUR_USERNAME/skygo/releases)，根据你的操作系统下载最新的压缩包。

-   `..._windows_amd64.zip`: 适用于 64 位的 Windows 系统。
-   `..._linux_amd64.tar.gz`: 适用于 64 位的 Linux 系统。
-   `..._darwin_amd64.tar.gz`: 适用于 Intel 芯片的 macOS。
-   `..._darwin_arm64.tar.gz`: 适用于 Apple M 系列芯片的 macOS。

#### 2. 解压文件

将下载的压缩包解压到。

#### 3. 首次运行 (生成配置)

直接运行解压后的可执行文件。
-   **Windows**: 双击 `skygo_windows_xxx.exe`。
-   **Linux / macOS**: 

授予软件可执行权限：`chmod x+ skygo_xxx_xxx`

在终端中运行 `./skygo_xxx_xxx`

程序会自动创建一个名为 `.env` 的文件并退出。

#### 4. 编辑配置文件

使用任意文本编辑器打开刚刚生成的 `.env` 文件，根据下面的说明填写所有必要信息。

#### ⚙️ 配置说明

所有配置项都在 `.env` 文件中。请确保在启动前已正确填写。

| 变量 (Variable)      | 说明                                                                                                                                                             | 示例                                     |
| -------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------- |
| `YT_API_KEY`         | **必需。** 用于调用相关功能的API Key。请前往 **[应天API](https://api.t1qq.com/user/key)** 获取。                                                                   | `abcde12345`                             |
| `BOT_NAME`           | **必需。** 机器人显示的昵称，默认为skygo                                        |  `skygo`                              |
|  `COMMAND_PREFIX`                   | **必需。** 触发机器人指令的指令前缀，默认为'.'                  |  `.`                                                                      | `skygo`                               |
| `SUPER_USERS`        | **必需。** 超级用户的QQ号列表，拥有管理员指令的执行权限。**多个用户请使用英文逗号(,)分隔。**                                                                              | `123456789,987654321`                    |
| `WS_SERVER_URL`      | **(二选一)** 反向 WebSocket 地址。如果使用此项，请将 `WS_CLIENT_URL` 留空。                                                                          | `ws://127.0.0.1:8000`                    |
| `WS_CLIENT_URL`      | **(二选一)** 正向 WebSocket 地址。如果使用此项，请将 `WS_SERVER_URL` 留空。                                                                          | `ws://127.0.0.1:6700`                    |
| `ACCESS_TOKEN`         | **可选。** 对接协议端时通信所需的令牌，需要和协议端保持一致，默认为空                                                                            | `12345`                           |

#### 5. 再次运行

配置好 `.env` 文件后保存，再次运行程序。如果其他错误，那么skygo此时应该就成功启动并连接到您的协议端了,您可以开始使用机器人内置的指令。

#### 指令列表

`status`: 查看bot运行设备的占用状态

`光遇菜单`：获取光遇有关的指令列表

`今日国服`：获取国服每日任务信息(图)

`国服复刻`：获取国服旅行先祖信息(图)

`季节蜡烛`：获取国服每日季节蜡烛位置(图)

`季节状态`：获取当前季节的信息

`活动日历`：获取游戏内官方的活动日历

`天气预报`：获取游戏内每日的天气预报

#### 常用onebot协议端推荐

[Lagrange](https://github.com/LagrangeDev/Lagrange.Core)

[NapcatQQ](https://github.com/NapNeko/NapCatQQ)

[LLOneBot](https://github.com/LLOneBot/LLOneBot)


## 🛠️ 从源码构建 (开发者)

如果您希望自行修改代码或进行二次开发，可以按照以下步骤从源码构建。

1.  **准备环境**
    确保您已安装 [Go](https://go.dev/dl/) (版本 1.20+) 和 [Git](https://git-scm.com/)。

2.  **克隆仓库**
    ```bash
    git clone https://github.com/Kaguya233qwq/skygo.git
    cd skygo
    ```

3.  **下载依赖**
    ```bash
    go mod tidy
    ```

4.  **构建**
    - **构建当前平台版本:**
      ```bash
      go build
      ```
    - **使用脚本交叉编译所有平台:**

      项目内提供了一个便捷的构建脚本：
      ```bash
      # Linux / macOS
      chmod +x build.sh
      ./build.sh
      
      # Windows
      .\build.bat
      ```
    编译好的文件生成在 `dist` 目录中。

## 🙏 致谢

-   **[ZeroBot](https://github.com/wdvxdr1123/ZeroBot)**: 本项目基于这个强大而简洁的go机器人框架开发，感谢各位开发者所贡献出的智慧与汗水。

## 📜 许可证

本项目采用 [MIT](LICENSE) 许可证。
