package config

import (
	"fmt"
	"github.com/rzaf/p2p-chat/gui/message"
	"image/color"
	"strconv"
)

func GetSetting() {
	if Incognito {
		UserUuid = string(RandStringBytes(10))
		return
	}
	fmt.Println("loading preferences ... ")
	p := ChatApp.Preferences().String("Port")
	if p == "" { // first time running
		fmt.Println("Preferences not found")
		UserUuid = string(RandStringBytes(10))
		StoreSetting()
		return
	}
	Port = p
	var err error
	saveChatsStr := ChatApp.Preferences().String("Save-chats")
	SaveRooms, err = strconv.ParseBool(saveChatsStr)
	if err != nil {
		SaveRooms = false
		fmt.Printf("failed to read `Save-chats:%s` from Preferences!\n", saveChatsStr)
	}
	UserUuid = ChatApp.Preferences().String("UserUuid")
	Username = ChatApp.Preferences().String("Username")
	var c color.RGBA
	b1 := ChatApp.Preferences().String("Background1")
	fmt.Sscanf(b1, "#%02x%02x%02x%02x", &c.R, &c.G, &c.B, &c.A)
	fmt.Println(c)
	message.BackGround1 = c
	b2 := ChatApp.Preferences().String("Background2")
	fmt.Sscanf(b2, "#%02x%02x%02x%02x", &c.R, &c.G, &c.B, &c.A)
	fmt.Println(c)
	message.BackGround2 = c
}

func StoreSetting() {
	if Incognito {
		return
	}
	ChatApp.Preferences().SetString("Save-chats", strconv.FormatBool(SaveRooms))
	ChatApp.Preferences().SetString("Port", Port)
	ChatApp.Preferences().SetString("UserUuid", UserUuid)
	ChatApp.Preferences().SetString("Username", Username)
	r, g, b, a := message.BackGround1.RGBA()
	ChatApp.Preferences().SetString("Background1", fmt.Sprintf("#%02x%02x%02x%02x", uint8(r), uint8(g), uint8(b), uint8(a)))
	r, g, b, a = message.BackGround2.RGBA()
	ChatApp.Preferences().SetString("Background2", fmt.Sprintf("#%02x%02x%02x%02x", uint8(r), uint8(g), uint8(b), uint8(a)))
}
