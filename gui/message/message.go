package message

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var (
	BackGround1 color.Color = color.RGBA{54, 92, 181, 255}
	BackGround2 color.Color = color.RGBA{30, 138, 52, 255}
)

type Message struct {
	Text        string
	Username    string
	Container   *fyne.Container
	IsMe        bool
	time        time.Time
	backgroundR *canvas.Rectangle
	Tick1       *widget.Icon
	Tick2       *widget.Icon
}

func NewMessage(username string, text string, isMe bool, time time.Time) *Message {
	m := &Message{
		Username: username,
		Text:     text,
		IsMe:     isMe,
		time:     time,
	}
	m.MakeContainer()
	return m
}

func (m *Message) MakeContainer() {
	m.backgroundR = canvas.NewRectangle(BackGround1)
	m.backgroundR.CornerRadius = 12

	textLabel := widget.NewLabel(m.Text)
	textLabel.Alignment = fyne.TextAlignLeading
	textLabel.Wrapping = fyne.TextWrapWord

	var border *fyne.Container
	m.Tick1 = widget.NewIcon(theme.ConfirmIcon())
	m.Tick2 = widget.NewIcon(theme.ConfirmIcon())
	timeLable := container.NewHBox(
		widget.NewLabel(m.time.Format(time.Kitchen)),
	)
	if m.IsMe {
		timeLable.Add(container.NewStack(container.NewWithoutLayout(m.Tick2), container.NewWithoutLayout(m.Tick1)))
		m.Tick1.Resize(fyne.NewSize(20, 20))
		m.Tick2.Resize(fyne.NewSize(25, 20))
		m.Tick1.Move(fyne.NewPos(-10, 6))
		m.Tick2.Move(fyne.NewPos(-6, 6))
		m.Tick1.Hide()
		m.Tick2.Hide()
		textLabel.Alignment = fyne.TextAlignLeading
		timeBox := container.NewVBox(layout.NewSpacer(), timeLable)
		border = container.NewBorder(nil, nil, nil, timeBox, container.NewPadded(container.NewStack(m.backgroundR, textLabel)))
	} else {
		usernmaLabel := widget.NewLabel(m.Username)
		usernmaLabel.Alignment = fyne.TextAlignTrailing
		usernmaLabel.TextStyle = fyne.TextStyle{Bold: true}

		m.backgroundR.FillColor = BackGround2
		textLabel.Alignment = fyne.TextAlignTrailing
		timeBox := container.NewVBox(layout.NewSpacer(), timeLable)
		border = container.NewBorder(usernmaLabel, nil, timeBox, nil, container.NewPadded(container.NewStack(m.backgroundR, textLabel)))
	}
	m.Container = border
}

func (m *Message) UpdateBgColor() {
	if m.IsMe {
		m.backgroundR.FillColor = BackGround1
	} else {
		m.backgroundR.FillColor = BackGround2
	}
	m.backgroundR.Refresh()
}
