package room

import (
	"github.com/rzaf/p2p-chat/gui/config"
	"github.com/rzaf/p2p-chat/gui/message"
	"github.com/rzaf/p2p-chat/models"

	"errors"
	"fmt"
	"image/color"
	"math/rand"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var (
	RoomsContainer *fyne.Container
	RoomsList      *fyne.Container
	Rooms          map[string]*Room

	addRoomContainer *fyne.Container
	settingContainer *fyne.Container
)

func LoadRooms() {
	if config.Incognito || !config.SaveRooms {
		return
	}
	fmt.Println("loading rooms")
	rooms, _ := models.GetAllRooms()
	fmt.Println("rooms:", rooms)
	for _, r := range rooms {
		AddRoom(r)
	}
}

func NewRoom(name, addr, port, userUuid string) *Room {
	r := &Room{
		Secret:    config.RandStringBytes(32),
		Uuid:      string(config.RandStringBytes(8)),
		UserId:    userUuid,
		Name:      name,
		Addr:      addr,
		Port:      port,
		seperator: widget.NewSeparator(),
	}
	r.createContainer()
	r.createChatRoom()
	r.add()
	if config.SaveRooms {
		newRoom, err := models.InsertRoom(models.Room{
			Secrete: string(r.Secret),
			Uuid:    r.Uuid,
			Name:    r.Name,
			Addr:    r.Addr,
			Port:    r.Port,
		})
		if err == nil {
			r.Id = newRoom.Id
		}
	} else {
		r.Id = rand.Int63()
	}
	return r
}

func MakeRoomList() *fyne.Container {
	roomsLabel := widget.NewLabel("Rooms")
	roomsLabel.TextStyle.Bold = true
	roomsLabel.Alignment = fyne.TextAlignCenter

	RoomsList = container.NewVBox()
	roomsListScroll := container.NewVScroll(RoomsList)

	roomToolBar := widget.NewToolbar(
		makeSettingToolbarAction(),
		makeNightModeToolBarAction(),
		widget.NewToolbarSpacer(),
		makeAddRoomToolbarAction(),
	)

	return container.NewBorder(roomsLabel, roomToolBar, nil, nil, roomsListScroll)
}

func makeAddRoomToolbarAction() *widget.ToolbarAction {
	// add room toolbar button
	addRoomNameInput := widget.NewEntry()
	addRoomNameInput.Validator = roomNameValidator
	addRoomNameForm := widget.NewFormItem("Name:", addRoomNameInput)
	addRoomNameForm.HintText = "chat name"

	addRoomIpInput := widget.NewEntry()
	addRoomIpInput.Validator = config.Ipv4Validator()
	addRoomIpForm := widget.NewFormItem("Ip:", addRoomIpInput)
	addRoomIpForm.HintText = "ipv4 ex: 127.0.0.1"

	addRoomPortInput := widget.NewEntry()
	addRoomPortInput.Validator = config.PortValidator()
	addRoomPortForm := widget.NewFormItem("port:", addRoomPortInput)

	addRoomUserIdInput := widget.NewEntry()
	addRoomUserIdInput.Validator = func(s string) error {
		if s == "" {
			return fmt.Errorf("empty")
		}
		return nil
	}
	addRoomUserIdForm := widget.NewFormItem("user id:", addRoomUserIdInput)
	addRoomUserIdForm.HintText = "user uuid of person that can send messages in this room"

	addRoomLabel := widget.NewLabel("Add Room")
	addRoomLabel.TextStyle.Bold = true
	addRoomLabel.Alignment = fyne.TextAlignCenter

	addRoomForm := widget.NewForm(addRoomNameForm, addRoomIpForm, addRoomPortForm, addRoomUserIdForm)

	addRoomContainer = container.NewBorder(
		addRoomLabel,
		container.NewCenter(
			container.NewHBox(
				widget.NewButtonWithIcon("Cancel", theme.CancelIcon(), func() {
					config.Window.SetContent(RoomsContainer)
				}),
				widget.NewButtonWithIcon("Add", theme.ConfirmIcon(), func() {
					addRoomNameInput.Text = strings.TrimSpace(addRoomNameInput.Text)
					err := addRoomForm.Validate()
					if err != nil {
						return
					}
					NewRoom(addRoomNameInput.Text, addRoomIpInput.Text, addRoomPortInput.Text, addRoomUserIdInput.Text)
					addRoomNameInput.Text = ""
					addRoomIpInput.Text = ""
					addRoomPortInput.Text = ""
					addRoomUserIdInput.Text = ""
					config.Window.SetContent(RoomsContainer)
				}),
			),
		),
		nil,
		nil,
		addRoomForm,
	)
	addRoomToolBarAction := widget.NewToolbarAction(theme.ContentAddIcon(), func() {
		config.Window.SetContent(addRoomContainer)
	})
	return addRoomToolBarAction
}

func makeNightModeToolBarAction() *widget.ToolbarAction {
	toolBarAction := widget.NewToolbarAction(nil, nil)
	if config.IsNightMode {
		toolBarAction.SetIcon(config.DayIcon)
	} else {
		toolBarAction.SetIcon(config.NightIcon)
	}
	toolBarAction.OnActivated = func() {
		if config.IsNightMode {
			config.ChatApp.Settings().SetTheme(theme.LightTheme())
			toolBarAction.SetIcon(config.NightIcon)
		} else {
			config.ChatApp.Settings().SetTheme(theme.DarkTheme())
			toolBarAction.SetIcon(config.DayIcon)
		}
		config.IsNightMode = !config.IsNightMode

		addRoomContainer.Refresh()
		settingContainer.Refresh()
	}
	return toolBarAction
}

func makeSettingToolbarAction() *widget.ToolbarAction {
	// setting toolbar button

	settingPortInput := widget.NewEntry()
	settingPortInput.Text = config.Port
	settingPortInput.Validator = config.PortValidator()
	settingPortForm := widget.NewFormItem("port:", settingPortInput)
	settingPortForm.HintText = "server port"

	settingUserUuidInput := widget.NewEntry()
	settingUserUuidInput.Text = config.UserUuid
	settingUserUuidInput.Disable()
	settingUserUuidForm := widget.NewFormItem("id:", settingUserUuidInput)
	settingUserUuidForm.HintText = "user uuid"

	settingUsernameInput := widget.NewEntry()
	settingUsernameInput.Text = config.Username
	settingUsernameInput.Validator = func(s string) error {
		if s == "" {
			return fmt.Errorf("empty")
		}
		return nil
	}
	settingUsernameForm := widget.NewFormItem("username:", settingUsernameInput)
	settingUsernameForm.HintText = "public name"

	settingSaveRoomsCheck := widget.NewCheck("", nil)
	settingSaveRoomsFrom := widget.NewFormItem("store rooms:", settingSaveRoomsCheck)

	size := 2 * theme.IconInlineSize()
	r1 := canvas.NewRectangle(message.BackGround1)
	r1.SetMinSize(fyne.NewSize(size, size))
	r2 := canvas.NewRectangle(message.BackGround2)
	r2.SetMinSize(fyne.NewSize(size, size))

	var color1, color2 color.Color

	cp1 := dialog.NewColorPicker("Pick color", "", func(c color.Color) {
		color1 = c
		r1.FillColor = c
		r1.Refresh()
	}, config.Window)
	cp1.Advanced = true
	bg1Form := widget.NewFormItem("color 1",
		container.NewStack(
			widget.NewButton("pick color", func() {
				// cp1.SetColor(message.BackGround1)
				cp1.Show()
			}),
			r1,
		),
	)
	bg1Form.HintText = "background color 1"

	cp2 := dialog.NewColorPicker("Pick color", "", func(c color.Color) {
		color2 = c
		r2.FillColor = c
		r2.Refresh()
	}, config.Window)

	cp2.Advanced = true
	bg2Form := widget.NewFormItem("color 2",
		container.NewStack(
			widget.NewButton("pick color", func() {
				// cp2.SetColor(message.BackGround2)
				cp2.Show()
			}),
			r2,
		),
	)
	bg2Form.HintText = "background color 2"

	settingLabel := widget.NewLabel("Setting")
	settingLabel.TextStyle.Bold = true

	settingForm := widget.NewForm(settingPortForm, settingUserUuidForm, settingUsernameForm, settingSaveRoomsFrom, bg1Form, bg2Form)
	settingForm.SubmitText = "Change"
	settingForm.OnSubmit = func() {
		err := settingForm.Validate()
		if err != nil {
			return
		}
		config.Username = settingUsernameInput.Text
		config.SaveRooms = settingSaveRoomsCheck.Checked

		message.BackGround1 = color1
		message.BackGround2 = color2
		UpdateBgColors()
		if config.Port != settingPortInput.Text {
			config.Port = settingPortInput.Text
			StopServer()
			go StartServer()
			// config.Stop()
			// go config.ServeAt(config.Addr + ":" + config.Port)
		}
		config.StoreSetting()
		config.Window.SetContent(RoomsContainer)
	}

	settingContainer = container.NewBorder(
		container.NewHBox(
			container.NewBorder(
				nil, nil, nil, settingLabel,
				widget.NewButtonWithIcon(
					"",
					theme.NavigateBackIcon(), func() {
						config.Window.SetContent(RoomsContainer)
					}),
			),
		),
		nil,
		nil,
		nil,
		settingForm,
	)
	settingToolBarAction := widget.NewToolbarAction(theme.SettingsIcon(), func() {
		settingUsernameInput.SetText(config.Username)
		settingPortInput.SetText(config.Port)
		settingSaveRoomsCheck.SetChecked(config.SaveRooms)
		color1 = message.BackGround1
		color2 = message.BackGround2
		r1.FillColor = color1
		r1.Refresh()
		r2.FillColor = color2
		r2.Refresh()
		config.Window.SetContent(settingContainer)
	})
	return settingToolBarAction
}

func roomNameValidator(s string) error {
	s = strings.TrimSpace(s)
	if s == "" {
		return errors.New("empty")
	}
	_, found := Rooms[s]
	if found {
		return errors.New("`" + s + "` already exists")
	}
	return nil
}

func UpdateBgColors() {
	for _, r := range Rooms {
		for _, m := range r.Messages {
			m.UpdateBgColor()
		}
	}
}
