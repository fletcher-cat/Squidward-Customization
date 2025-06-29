package main

import (
	"image/color"
	"os/exec"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// generated using https://github.com/lusingander/fyne-theme-generator
type myTheme struct{}

func (myTheme) Color(c fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	switch c {
	case theme.ColorNameBackground:
		return color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
	case theme.ColorNameButton:
		return color.NRGBA{R: 0xf5, G: 0xf5, B: 0xf5, A: 0xff}
	case theme.ColorNameDisabledButton:
		return color.NRGBA{R: 0xf5, G: 0xf5, B: 0xf5, A: 0xff}
	case theme.ColorNameDisabled:
		return color.NRGBA{R: 0xe3, G: 0xe3, B: 0xe3, A: 0xff}
	case theme.ColorNameError:
		return color.NRGBA{R: 0xf4, G: 0x43, B: 0x36, A: 0xff}
	case theme.ColorNameFocus:
		return color.NRGBA{R: 0x0, G: 0x6c, B: 0xff, A: 0x2a}
	case theme.ColorNameForeground:
		return color.NRGBA{R: 0x56, G: 0x56, B: 0x56, A: 0xff}
	case theme.ColorNameHover:
		return color.NRGBA{R: 0x0, G: 0x0, B: 0x0, A: 0xf}
	case theme.ColorNameInputBackground:
		return color.NRGBA{R: 0xf3, G: 0xf3, B: 0xf3, A: 0xff}
	case theme.ColorNamePlaceHolder:
		return color.NRGBA{R: 0x88, G: 0x88, B: 0x88, A: 0xff}
	case theme.ColorNamePressed:
		return color.NRGBA{R: 0x0, G: 0x0, B: 0x0, A: 0x19}
	case theme.ColorNamePrimary:
		return color.NRGBA{R: 0x29, G: 0x6f, B: 0xf6, A: 0xff}
	case theme.ColorNameScrollBar:
		return color.NRGBA{R: 0x0, G: 0x0, B: 0x0, A: 0x99}
	case theme.ColorNameShadow:
		return color.NRGBA{R: 0x0, G: 0x0, B: 0x0, A: 0x33}
	default:
		return theme.DefaultTheme().Color(c, v)
	}
}

func (myTheme) Font(s fyne.TextStyle) fyne.Resource {
	if s.Monospace {
		return theme.DefaultTheme().Font(s)
	}
	if s.Bold {
		if s.Italic {
			return theme.DefaultTheme().Font(s)
		}
		return theme.DefaultTheme().Font(s)
	}
	if s.Italic {
		return theme.DefaultTheme().Font(s)
	}
	return theme.DefaultTheme().Font(s)
}

func (myTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (myTheme) Size(s fyne.ThemeSizeName) float32 {
	switch s {
	case theme.SizeNameCaptionText:
		return 11
	case theme.SizeNameInlineIcon:
		return 20
	case theme.SizeNamePadding:
		return 4
	case theme.SizeNameScrollBar:
		return 12
	case theme.SizeNameScrollBarSmall:
		return 3
	case theme.SizeNameSeparatorThickness:
		return 1
	case theme.SizeNameText:
		return 14
	case theme.SizeNameInputBorder:
		return 1
	default:
		return theme.DefaultTheme().Size(s)
	}
}

// end of theme
func main() {
	//initialize
	a := app.New()
	a.Settings().SetTheme(&myTheme{})
	w := a.NewWindow("Get Started (v1.0)")
	w.SetFixedSize(true) //stops window from resizing\
	//load the icon for window and taskbar
	icon, err := fyne.LoadResourceFromPath("assets/house.png")
	if err == nil {
		w.SetIcon(icon)
	}
	// Set the icon for the window. This is optional, but it can help to make
	//label and image. label is the text that will be displayed in the window, and image is the image that will be displayed in the window.

	label := widget.NewLabel("Hello! My name is Squidward Tentacles, and I will assist you into personalizing your computer.\nPersonalizing your computer is as easy as 1-2-3! So, sit back, relax, and get ready to personalize to your heart's content.\nTo get started, click one of the selections below.")
	//center the label text
	label.Alignment = fyne.TextAlignCenter
	image := canvas.NewImageFromFile("assets/squidward.jpg")
	//load resources for the button icons
	image2, err := fyne.LoadResourceFromPath("assets/monitor_brush.png")
	if err != nil { //placeholder
	}
	image3, err := fyne.LoadResourceFromPath("assets/colors.png")
	if err != nil { //placeholder
	}
	image4, err := fyne.LoadResourceFromPath("assets/windows-start.png")
	if err != nil { //placeholder
	}
	//fill mode
	image.FillMode = canvas.ImageFillContain
	//image size
	image.SetMinSize(fyne.NewSize(200, 500))
	//show the image and label
	//newhbox arranges the image and label horizontally, newvbox is for vertical

	w.SetContent(container.NewHBox(
		image,
		container.NewVBox(
			label,
			widget.NewButtonWithIcon("Change desktop background", image2, func() {
				exec.Command("cmd", "/C", "start ms-settings:personalization-background").Start()
			}),
			widget.NewButtonWithIcon("Change window glass colors", image3, func() {
				exec.Command("cmd", "/C", "start ms-settings:colors").Start()
			}),
			widget.NewButtonWithIcon("Customize the Start menu", image4, func() {
				exec.Command("cmd", "/C", "start ms-settings:personalization-start").Start()
			}),
		),
	))
	w.ShowAndRun()
}
