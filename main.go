package main

import (
	"strings"
	"time"

	"github.com/lxn/walk"
	dec "github.com/lxn/walk/declarative"
)

func main() {
	var inTE, outTE *walk.TextEdit

	dec.MainWindow{
		Title: "SCREAMO",
		MinSize: dec.Size{
			Width:  600,
			Height: 400},
		Layout: dec.VBox{},
		Children: []dec.Widget{
			dec.HSplitter{
				Children: []dec.Widget{
					dec.TextEdit{AssignTo: &inTE},
					dec.TextEdit{AssignTo: &outTE, ReadOnly: true},
				},
			},
			dec.PushButton{
				Text: "SCREAM",
				OnClicked: func() {
					outTE.SetText(strings.ToUpper(inTE.Text()))
				},
			},
		},
	}.Run()
	time.Sleep(time.Millisecond)
}
