package main

import (
	_ "embed"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/robfig/cron/v3"
	"log"
	"math/rand"
	"time"
	"water/zh"
)

var w fyne.Window

// 提示喝水频率 默认三十分钟提示一次
var frequency = "30m"

var Tip = []string{
	"今天你多喝水了么?",
	"别光喝水，起来走动走动吧~",
	"每天八杯水，开心的源泉~",
	"忙一天了，去摸鱼办看看热搜新鲜事儿",
	"起来扭扭腰，动动脖吧",
	"辛苦撸代码的同时，也要多喝水哟~",
}

var timeSelect = []string{"半小时", "一小时", "一分钟"}

////go:embed1 wx1.png
//var icon []byte

func main() {
	a := app.New()
	a.Settings().SetTheme(&zh.MyTheme{})

	w = a.NewWindow("喝水提示小工具")

	//fyne.CurrentApp().SetIcon(&fyne.StaticResource{
	//	StaticName:    "w2.svg",
	//	StaticContent: icon,
	//})
	//
	selectEntry := widget.NewSelect(timeSelect, SelectTime())
	selectEntry.PlaceHolder = "时长,默认半小时"

	w.SetContent(container.NewVBox(
		selectEntry,
		widget.NewButton("开始", Start()),
	))

	w.Resize(fyne.NewSize(219, 66))

	w.ShowAndRun()

}

func Start() func() {
	return func() {
		c := cron.New()
		_, err := c.AddFunc("@every "+frequency, func() {
			nowTime := time.Now().Format("2006-01-02 15:04:05")
			log.Println(frequency, "he"+nowTime)
			SendMsgFunc()
		})
		if err != nil {
			dialog.ShowError(err, w)
		}
		c.Start()
		// 第一次按钮也提示一次方便查看样式
		SendMsgFunc()
	}
}

func SendMsgFunc() {
	fyne.CurrentApp().SendNotification(&fyne.Notification{
		Title:   "亲，该喝水了",
		Content: Tip[rand.Intn(len(Tip))],
	})
}

func SelectTime() func(s string) {
	return func(s string) {
		switch s {
		case "一小时":
			frequency = "1h"
			break
		case "一分钟":
			frequency = "1m"
			break
		default:
			frequency = "30m"
			break

		}
		log.Println(s, frequency)
	}
}
