package config

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/rzaf/p2p-chat/models"
	"math/rand"
	"os"
	"regexp"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/dialog"
)

var (
	Incognito  bool   = false
	SaveRooms  bool   = true
	AppId      string = "chat-app"
	Username   string = "New User"
	UserUuid   string
	Addr, Port string = "0.0.0.0", "56156"

	ChatApp           fyne.App
	DefaultWindowSize fyne.Size = fyne.NewSize(350, 500)
	WindowTitle       string    = "p2p chat app"
	Window            fyne.Window
	IsInForground     bool
	IsNightMode       bool

	CurrentChatRoom *fyne.Container

	DayIcon   *fyne.StaticResource
	NightIcon *fyne.StaticResource
	ChatIcon  *fyne.StaticResource
)

func parseArgs() {
	flag.StringVar(&Port, "p", Port, "port number of server")
	flag.StringVar(&Username, "u", Username, "public username")
	flag.BoolFunc("i", "incognito mode: to not save or load any data", func(s string) error {
		Incognito = true
		return nil
	})
	flag.Parse()
	if m, _ := regexp.MatchString(portValidatorRegex(), Port); !m {
		fmt.Printf("invalid port number:%s\n", Port)
		flag.Usage()
		os.Exit(1)
	}
}

func Init() {
	parseArgs()
	IsNightMode = true
	os.Setenv("FYNE_THEME", "dark")
	ChatApp = app.NewWithID(AppId)
	GetSetting()
	LoadResources()
	ConnectToSql()

	IsInForground = true
}

func InitWindow() {
	Window = ChatApp.NewWindow(WindowTitle)
	Window.Resize(DefaultWindowSize)
	Window.CenterOnScreen()
}

func ConnectToSql() {
	if Incognito || !SaveRooms {
		return
	}
	p := ChatApp.Storage().RootURI().Path() + "/sql.db"
	fmt.Println(p)
	db, err := sql.Open("sqlite", p)
	if err != nil {
		panic(err)
	}
	models.Db = db
	MigrateDb()
}

func LoadResources() {
	DayIcon = fyne.NewStaticResource("day.svg", []byte(`
		<svg viewBox="0 0 24 24" version="1.1"><g id="SVGRepo_bgCarrier" stroke-width="0"></g><g id="SVGRepo_tracerCarrier" stroke-linecap="round" stroke-linejoin="round"></g><g id="SVGRepo_iconCarrier"><title>ic_fluent_dark_theme_24_filled</title> <desc>Created with Sketch.</desc> <g id="🔍-Product-Icons" stroke="none" stroke-width="1" fill="none" fill-rule="evenodd"> <g id="ic_fluent_dark_theme_24_filled" fill="#ffffff" fill-rule="nonzero"> <path d="M12,22 C17.5228475,22 22,17.5228475 22,12 C22,6.4771525 17.5228475,2 12,2 C6.4771525,2 2,6.4771525 2,12 C2,17.5228475 6.4771525,22 12,22 Z M12,20 L12,4 C16.418278,4 20,7.581722 20,12 C20,16.418278 16.418278,20 12,20 Z" id="🎨-Color"> </path> </g> </g> </g></svg>
	`))
	NightIcon = fyne.NewStaticResource("night.svg", []byte(`
		<svg viewBox="0 0 24 24" version="1.1"><g id="SVGRepo_bgCarrier" stroke-width="0"></g><g id="SVGRepo_tracerCarrier" stroke-linecap="round" stroke-linejoin="round"></g><g id="SVGRepo_iconCarrier"><title>ic_fluent_dark_theme_24_filled</title> <desc>Created with Sketch.</desc> <g id="🔍-Product-Icons" stroke="none" stroke-width="1" fill="none" fill-rule="evenodd"> <g id="ic_fluent_dark_theme_24_filled" fill="#212121" fill-rule="nonzero"> <path d="M12,22 C17.5228475,22 22,17.5228475 22,12 C22,6.4771525 17.5228475,2 12,2 C6.4771525,2 2,6.4771525 2,12 C2,17.5228475 6.4771525,22 12,22 Z M12,20 L12,4 C16.418278,4 20,7.581722 20,12 C20,16.418278 16.418278,20 12,20 Z" id="🎨-Color"> </path> </g> </g> </g></svg>
	`))
	ChatApp.SetIcon(resourceChatIconPng)
}

func DisconnectFromSql() {
	if models.Db != nil {
		models.Db.Close()
	}
}

func MigrateDb() {
	err := models.MigrateRooms()
	if err != nil {
		panic(err)
	}
}

func ShowError(e error) {
	fmt.Printf("error:%v\n", e)
	d := dialog.NewError(e, Window)
	d.Show()
}

func SendNotification(title, content string) {
	ChatApp.SendNotification(fyne.NewNotification(title, content))
}

func Ipv4Validator() fyne.StringValidator {
	return validation.NewRegexp("^((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$", "invalid ip")
}

func portValidatorRegex() string {
	return "^([1-9][0-9]{0,3}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|6553[0-5])$"
}

func PortValidator() fyne.StringValidator {
	return validation.NewRegexp(portValidatorRegex(), "invalid port")
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandStringBytes(n int) []byte {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return b
}
