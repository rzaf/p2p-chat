package main

import (
	_ "github.com/glebarez/go-sqlite"
	"github.com/rzaf/p2p-chat/gui/config"
	"github.com/rzaf/p2p-chat/gui/room"
)

func main() {
	config.Init()
	config.ChatApp.Lifecycle().SetOnEnteredForeground(func() {
		config.IsInForground = true
	})
	config.ChatApp.Lifecycle().SetOnExitedForeground(func() {
		config.IsInForground = false
	})
	config.ChatApp.Lifecycle().SetOnStopped(func() {
		room.StopServer()
		config.DisconnectFromSql()
	})
	config.Window = config.ChatApp.NewWindow(config.WindowTitle)
	config.Window.Resize(config.DefaultWindowSize)
	config.Window.CenterOnScreen()

	room.Rooms = make(map[string]*room.Room)
	room.RoomsContainer = room.MakeRoomList()
	room.LoadRooms()

	go room.StartServer()

	config.Window.SetContent(room.RoomsContainer)
	config.Window.ShowAndRun()
}

// export FYNE_FONT=./assets/noto.ttf
