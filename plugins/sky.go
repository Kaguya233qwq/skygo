package plugins

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"skygo/gofer"
	"strings"

	logger "github.com/sirupsen/logrus"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

var API_KEY = os.Getenv("YT_API_KEY")
var VERSION = "1.0.3"

var ErrInvalidAPIKey = errors.New("无效的API Key")
var ErrApiCodeUnusually = errors.New("Api状态码异常")

func init() {
	err := checkingApiKey()
	if err != nil {
		logger.Error(err)
		logger.Error("请检查您的api_key是否配置正确")
		os.Exit(1)
	}

	getChineseServerDailyTask()
	getLocationOfSeasonCandles()
	getSeasonStatus()
	getEventSchedule()
	getChineseServerTravelingSpirit()
	getSkyWeatherForecast()
	getSkyMenu()
	logger.Info("Plugin:Sky loaded successfully")
}

// 检查api_key有效性
func checkingApiKey() error {
	type Response struct {
		Code any    `json:"code"`
		Msg  string `json:"msg"`
		Data string `json:"data"`
	}

	api := "https://api.t1qq.com/api/tool/yiyan"
	headers := map[string]string{}
	params := map[string]string{"key": API_KEY}
	resp, err := gofer.Get(api, headers, params)
	if err != nil {
		panic(err)
	}
	var response Response
	err = json.Unmarshal([]byte(resp.Body), &response)
	if err != nil {
		panic(err)
	}
	//对response.Code进行类型判断
	switch code := response.Code.(type) {
	case float64:
		if int(code) == 403 {
			return ErrInvalidAPIKey
		}
	case string:
		if code == "200" {
			logger.Info("调用成功，ApiKey已正确配置")
			return nil
		}
		return ErrApiCodeUnusually
	default:
		// 未知的Code类型(不太可能有)
		return fmt.Errorf("unexpected type for 'code' field: %T", code)
	}
	logger.Info("调用成功，ApiKey已正确配置")
	return nil
}

func getChineseServerDailyTask() {
	zero.OnCommand("今日国服").Handle(func(ctx *zero.Ctx) {
		api := "https://api.t1qq.com/api/sky/sc/scrw"
		headers := map[string]string{}
		params := map[string]string{"key": API_KEY}

		resp, err := gofer.Get(api, headers, params)
		if err != nil {
			go ctx.Send("请求失败: " + err.Error())
			return
		}

		// 检查响应状态
		if resp.Code != 200 {
			go ctx.Send(fmt.Sprintf("API返回错误状态码: %d", resp.Code))
			return
		}

		// 将字符串Body转换为字节切片
		imgBytes := []byte(resp.Body)

		// 使用base64发送图片
		imgBase64 := base64.StdEncoding.EncodeToString(imgBytes)
		go ctx.Send(message.Image("base64://" + imgBase64))
	})
}

func getLocationOfSeasonCandles() {
	zero.OnCommand("季节蜡烛").Handle(func(ctx *zero.Ctx) {
		api := "https://api.t1qq.com/api/sky/sc/scjl"
		headers := map[string]string{}
		params := map[string]string{"key": API_KEY}

		resp, err := gofer.Get(api, headers, params)
		if err != nil {
			go ctx.Send("请求失败: " + err.Error())
			return
		}

		// 检查响应状态
		if resp.Code != 200 {
			go ctx.Send(fmt.Sprintf("API返回错误状态码: %d", resp.Code))
			return
		}

		// 将字符串Body转换为字节切片
		imgBytes := []byte(resp.Body)

		// 使用base64发送图片
		imgBase64 := base64.StdEncoding.EncodeToString(imgBytes)
		go ctx.Send(message.Image("base64://" + imgBase64))
	})
}

type seasonStatusResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Day  string `json:"day"`
		Days string `json:"days"`
		Time string `json:"time"`
		Data struct {
			Seasonal    string `json:"seasonal"`
			NotSeasonal string `json:"notSeasonal"`
			Reward      string `json:"reward"`
			OrderReward string `json:"orderReward"`
			Grad        string `json:"grad"`
		} `json:"data"`
	} `json:"data"`
}

func getSeasonStatus() {
	zero.OnCommand("季节状态").Handle(func(ctx *zero.Ctx) {
		api := "https://api.t1qq.com/api/sky/sc/gf/djs"
		headers := map[string]string{}
		params := map[string]string{"key": API_KEY}

		resp, err := gofer.Get(api, headers, params)
		if err != nil {
			go ctx.Send("请求失败: " + err.Error())
			return
		}

		// 检查响应状态
		if resp.Code != 200 {
			go ctx.Send(fmt.Sprintf("API返回错误状态码: %d", resp.Code))
			return
		}

		// 解析JSON响应
		var status seasonStatusResponse
		err = json.Unmarshal([]byte(resp.Body), &status)
		if err != nil {
			go ctx.Send("解析数据失败: " + err.Error())
			return
		}

		// 检查API业务状态码
		if status.Code != 200 {
			go ctx.Send("API返回错误: " + status.Msg)
			return
		}

		// 构建消息内容
		message := strings.Join([]string{
			"✨ 季节状态信息 ✨",
			"",
			"📅 " + status.Data.Time,
			"⏳ " + status.Data.Days,
			"",
			"🎫 季卡信息:",
			"  - " + status.Data.Data.Seasonal,
			"  - " + status.Data.Data.NotSeasonal,
			"",
			"🕯️ 季蜡获取:",
			"  - " + status.Data.Data.Reward,
			"  - " + status.Data.Data.OrderReward,
			"",
			"🎓 " + status.Data.Data.Grad,
		}, "\n")

		go ctx.Send(message)
	})
}

func getEventSchedule() {
	zero.OnCommand("活动日历").Handle(func(ctx *zero.Ctx) {
		api := "https://api.t1qq.com/api/sky/sc/hdrl"
		headers := map[string]string{}
		params := map[string]string{"key": API_KEY, "type": "img"}

		resp, err := gofer.Get(api, headers, params)
		if err != nil {
			go ctx.Send("请求失败: " + err.Error())
			return
		}

		// 检查响应状态
		if resp.Code != 200 {
			go ctx.Send(fmt.Sprintf("API返回错误状态码: %d", resp.Code))
			return
		}

		// 将字符串Body转换为字节切片
		imgBytes := []byte(resp.Body)

		// 使用base64发送图片
		imgBase64 := base64.StdEncoding.EncodeToString(imgBytes)
		go ctx.Send(message.Image("base64://" + imgBase64))
	})
}

func getChineseServerTravelingSpirit() {
	zero.OnCommand("国服复刻").Handle(func(ctx *zero.Ctx) {
		api := "https://api.t1qq.com/api/sky/sc/scfk"
		headers := map[string]string{}
		params := map[string]string{"key": API_KEY}

		resp, err := gofer.Get(api, headers, params)
		if err != nil {
			go ctx.Send("请求失败: " + err.Error())
			return
		}

		// 检查响应状态
		if resp.Code != 200 {
			go ctx.Send(fmt.Sprintf("API返回错误状态码: %d", resp.Code))
			return
		}

		// 将字符串Body转换为字节切片
		imgBytes := []byte(resp.Body)

		// 使用base64发送图片
		imgBase64 := base64.StdEncoding.EncodeToString(imgBytes)
		go ctx.Send(message.Image("base64://" + imgBase64))
	})
}

// 定义API响应结构体
type WeatherResponse struct {
	Code int      `json:"code"`
	Msg  string   `json:"msg"`
	Data []string `json:"data"`
}

func getSkyWeatherForecast() {
	zero.OnCommand("天气预报").Handle(func(ctx *zero.Ctx) {
		api := "https://api.t1qq.com/api/sky/gytq"
		headers := map[string]string{}
		params := map[string]string{"key": API_KEY}

		resp, err := gofer.Get(api, headers, params)
		if err != nil {
			go ctx.Send("请求失败: " + err.Error())
			return
		}

		// 检查HTTP响应状态
		if resp.Code != 200 {
			go ctx.Send(fmt.Sprintf("API返回错误状态码: %d", resp.Code))
			return
		}

		// 解析JSON响应
		var weatherResp WeatherResponse
		err = json.Unmarshal([]byte(resp.Body), &weatherResp)
		if err != nil {
			go ctx.Send("解析数据失败: " + err.Error())
			return
		}

		// 检查API业务状态码
		if weatherResp.Code != 200 {
			go ctx.Send("API返回错误: " + weatherResp.Msg)
			return
		}

		// 检查是否有图片数据
		if len(weatherResp.Data) == 0 {
			go ctx.Send("未获取到天气预报图片")
			return
		}
		imgURL := weatherResp.Data[0]
		// 直接使用图片URL发送
		go ctx.Send(message.Image(imgURL))
	})
}

func getSkyMenu() {
	zero.OnCommand("光遇菜单").Handle(func(ctx *zero.Ctx) {
		message := strings.Join([]string{
			"✨ 可用指令列表-skygo ✨",
			"----------",
			".今日国服",
			".国服复刻",
			".季节蜡烛",
			".季节状态",
			".活动日历",
			".天气预报",
			"----------",
			"基于go开发的光遇攻略助手",
			"当前版本：v" + VERSION,
		}, "\n")

		go ctx.Send(message)
	})
}
