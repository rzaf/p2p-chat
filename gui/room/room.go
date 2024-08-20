package room

import (
	"errors"
	"fmt"

	"github.com/rzaf/p2p-chat/gui/config"
	"github.com/rzaf/p2p-chat/gui/message"
	"github.com/rzaf/p2p-chat/models"
	"github.com/rzaf/p2p-chat/pb"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"golang.org/x/crypto/nacl/secretbox"
)

type Room struct {
	Secret            []byte
	Uuid              string
	UserId            string
	Name              string
	Addr              string
	Port              string
	Id                int64
	seperator         *widget.Separator
	container         *fyne.Container
	ChatRoom          *fyne.Container
	MessageList       *fyne.Container
	Messages          []*message.Message
	MessageListScroll *container.Scroll
}

func AddRoom(r2 *models.Room) *Room {
	r := &Room{
		Secret:    []byte(r2.Secrete),
		Uuid:      r2.Uuid,
		UserId:    r2.UserUuid,
		Name:      r2.Name,
		Addr:      r2.Addr,
		Port:      r2.Port,
		Id:        r2.Id,
		seperator: widget.NewSeparator(),
	}
	r.createContainer()
	r.createChatRoom()
	r.add()
	return r
}

func (r *Room) add() {
	fmt.Printf("room %v added \n", r)
	RoomsList.Add(r.seperator)
	RoomsList.Add(r.container)
	Rooms[r.Name] = r
}

func (r *Room) delete() {
	fmt.Printf("room %v deleted \n", r)
	if config.SaveRooms {
		models.DeleteRoom(r.Id)
	}
	RoomsList.Remove(r.container)
	if r.seperator != nil {
		RoomsList.Remove(r.seperator)
	}
	delete(Rooms, r.Name)
}

func (r *Room) createContainer() {
	roomName := widget.NewButton(r.Name, func() {
		config.Window.SetContent(r.ChatRoom)
		config.CurrentChatRoom = r.ChatRoom
	})
	roomIp := widget.NewLabel(r.Addr + ":" + r.Port)
	roomIp.Alignment = fyne.TextAlignCenter

	deleteButton := widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {
		deleteRoomDialog := dialog.NewConfirm("Delete Room", "`"+roomName.Text+"`", func(b bool) {
			if !b {
				return
			}
			r, found := Rooms[r.Name]
			if !found {
				return
			}
			r.delete()
		}, config.Window)
		deleteRoomDialog.Show()
	})

	infoButton := widget.NewButtonWithIcon("", theme.InfoIcon(), func() {
		addressEntry := widget.NewEntry()
		addressEntry.Text = r.Addr
		addressEntry.Validator = config.Ipv4Validator()
		addressItem := widget.NewFormItem("address:", addressEntry)

		portEntry := widget.NewEntry()
		portEntry.Text = r.Port
		portEntry.Validator = config.PortValidator()
		portItem := widget.NewFormItem("port:", portEntry)

		uuidEntry := widget.NewEntry()
		uuidEntry.Text = r.Uuid
		uuidEntry.Validator = func(s string) error {
			if s == "" {
				return fmt.Errorf("empty")
			}
			return nil
		}
		uuidItem := widget.NewFormItem("uuid:", uuidEntry)

		secretEntry := widget.NewEntry()
		secretEntry.Text = string(r.Secret)
		secretEntry.Validator = func(s string) error {
			if len(s) != 32 {
				return errors.New("length should be 32")
			}
			return nil
		}
		secretItem := widget.NewFormItem("secret:", secretEntry)

		userIdEntry := widget.NewEntry()
		userIdEntry.Text = r.UserId
		userIdEntry.Validator = func(s string) error {
			if s == "" {
				return fmt.Errorf("empty")
			}
			return nil
		}
		userIdItem := widget.NewFormItem("user id:", userIdEntry)
		userIdItem.HintText = "user uuid of person that can send messages in this room"

		infoRoomLabel := widget.NewLabel(r.Name + " Info")
		infoRoomLabel.Alignment = fyne.TextAlignCenter
		infoRoomLabel.TextStyle.Bold = true
		infoRoomForm := widget.NewForm(addressItem, portItem, uuidItem, secretItem, userIdItem)

		infoRoomsContainer := container.NewBorder(
			infoRoomLabel,
			container.NewCenter(
				container.NewHBox(
					widget.NewButtonWithIcon("Cancel", theme.CancelIcon(), func() {
						config.Window.SetContent(RoomsContainer)
					}),
					widget.NewButtonWithIcon("Change", theme.ConfirmIcon(), func() {
						if infoRoomForm.Validate() != nil {
							return
						}
						r.Port = portEntry.Text
						roomIp.SetText(r.Addr + ":" + r.Port)
						r.Addr = addressEntry.Text
						roomIp.SetText(r.Addr + ":" + r.Port)
						r.Uuid = uuidEntry.Text
						r.Secret = []byte(secretEntry.Text)
						r.UserId = userIdEntry.Text
						if config.SaveRooms {
							models.UpdateRoom(r.Id, models.Room{
								Secrete:  string(r.Secret),
								Uuid:     r.Uuid,
								UserUuid: r.UserId,
								Name:     r.Name,
								Addr:     r.Addr,
								Port:     r.Port,
							})
						}
						config.Window.SetContent(RoomsContainer)
					}),
				),
			),
			nil,
			nil,
			infoRoomForm,
		)

		config.Window.SetContent(infoRoomsContainer)
	})

	r.container = container.NewHBox(
		container.NewPadded(roomName), container.NewPadded(roomIp), layout.NewSpacer(), container.NewPadded(infoButton), container.NewPadded(deleteButton),
	)
}

func (r *Room) createChatRoom() {
	r.MessageList = container.NewVBox()
	r.MessageListScroll = container.NewVScroll(r.MessageList)

	input := widget.NewMultiLineEntry()
	sendBtn := widget.NewButtonWithIcon("", theme.MailSendIcon(), nil)
	sendBtn.OnTapped = func() {
		if input.Text != "" {
			r.MessageList.Add(widget.NewSeparator())
			m := message.NewMessage("", input.Text, true, time.Now())
			r.Messages = append(r.Messages, m)
			r.MessageList.Add(m.Container)
			r.MessageListScroll.ScrollToBottom()
			go func(secret []byte, roomUuid, text, addr string) {
				attempts := 0
				for {
					if attempts > 3 {
						m.Tick1.SetResource(theme.ErrorIcon())
						m.Tick1.Show()
						return
					}
					attempts++
					err := config.SendMessageTo(secret, roomUuid, config.UserUuid, text, addr, config.Username)
					if err != nil {
						fmt.Println(err)
						// ShowError(err)
						time.Sleep(time.Second * 2)
						continue
					}
					m.Tick1.Show()
					return
				}
			}(r.Secret, r.Uuid, input.Text, r.Addr+":"+r.Port)

			input.Text = ""
			input.Refresh()
		}
	}

	bottom := container.NewBorder(nil, nil, nil, sendBtn, input)
	center := container.NewBorder(nil, bottom, nil, nil, r.MessageListScroll)
	chatRoomName := widget.NewLabel(r.Name)
	chatRoomName.TextStyle.Bold = true
	backButton := widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() {
		config.Window.SetContent(RoomsContainer)
		config.CurrentChatRoom = nil
	})

	top := container.NewBorder(nil, nil, container.NewPadded(backButton), nil, container.NewPadded(chatRoomName))

	r.ChatRoom = container.NewBorder(top, nil, nil, nil, center)
}

func AddMessage(t *pb.Text) {
	fmt.Println(t)
	for _, r := range Rooms {
		fmt.Println(r)
		if r.Uuid == t.RoomUuid && r.UserId == t.UserUuid {
			encrypted := t.Message
			var decryptNonce [24]byte
			copy(decryptNonce[:], encrypted[:24])
			decrypted, ok := secretbox.Open(nil, encrypted[24:], &decryptNonce, (*[32]byte)(r.Secret))
			if ok {
				if config.CurrentChatRoom != r.ChatRoom || !config.IsInForground {
					config.SendNotification(
						fmt.Sprintf("New message in room:%s from:%s", r.Name, t.Username),
						fmt.Sprintf("%s: %s", t.Username, string(decrypted)),
					)
				}
				r.MessageList.Add(widget.NewSeparator())
				m := message.NewMessage(t.Username, string(decrypted), false, time.UnixMilli(t.UnixMilli))
				r.Messages = append(r.Messages, m)
				r.MessageList.Add(m.Container)
				r.MessageListScroll.ScrollToBottom()
			}
			return
		}
	}
}
