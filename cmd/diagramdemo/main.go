package main

import (
	"fmt"
	"image/color"
	"time"

	"fyne.io/x/fyne/widget/diagramwidget"
	"fyne.io/x/fyne/widget/diagramwidget/arrowhead"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var forceticks int = 0

func forceanim() {

	// XXX: very naughty -- accesses shared memory in potentially unsafe
	// ways, this almost certainly has race conditions... don't do this!

	for {
		if forceticks > 0 {
			diagramwidget.Globaldiagram.StepForceLayout(300)
			diagramwidget.Globaldiagram.Refresh()
			forceticks--
			fmt.Printf("forceticks=%d\n", forceticks)
		}

		time.Sleep(time.Millisecond * (1000 / 30))
	}
}

func main() {
	app := app.New()
	w := app.NewWindow("Diagram Demo")

	w.SetMaster()

	diagramWidget := diagramwidget.NewDiagramWidget()
	diagramwidget.Globaldiagram = diagramWidget

	go forceanim()

	// Node 0
	node0Label := widget.NewLabel("Node0")
	node0 := diagramwidget.NewDiagramNode(diagramWidget, node0Label)
	diagramWidget.Nodes["node0"] = node0

	// Node 1
	node1Button := widget.NewButton("Node1 Button", func() { fmt.Printf("tapped Node1!\n") })
	node1 := diagramwidget.NewDiagramNode(diagramWidget, node1Button)
	node1.Move(fyne.Position{X: 200, Y: 200})
	diagramWidget.Nodes["node1"] = node1

	// Node 2
	node2 := diagramwidget.NewDiagramNode(diagramWidget, nil)
	node2Container := container.NewVBox(
		widget.NewLabel("Node2 - with structure"),
		widget.NewButton("Up", func() {
			node2.Displace(fyne.Position{X: 0, Y: -10})
			node2.Refresh()
		}),
		widget.NewButton("Down", func() {
			node2.Displace(fyne.Position{X: 0, Y: 10})
			node2.Refresh()
		}),
		container.NewHBox(
			widget.NewButton("Left", func() {
				node2.Displace(fyne.Position{X: -10, Y: 0})
				node2.Refresh()
			}),
			widget.NewButton("Right", func() {
				node2.Displace(fyne.Position{X: 10, Y: 0})
				node2.Refresh()
			}),
		),
	)
	node2.InnerObject = node2Container
	node2.Move(fyne.Position{X: 300, Y: 300})
	diagramWidget.Nodes["node2"] = node2

	// Node 3
	node3 := diagramwidget.NewDiagramNode(diagramWidget, widget.NewButton("Node3: Force layout step", func() {
		diagramWidget.StepForceLayout(300)
		diagramWidget.Refresh()
	}))
	node3.Move(fyne.Position{X: 400, Y: 200})
	diagramWidget.Nodes["node3"] = node3

	// Node 4
	node4 := diagramwidget.NewDiagramNode(diagramWidget, widget.NewButton("Node4: auto layout", func() {
		forceticks += 100
		diagramWidget.Refresh()
	}))
	node4.Move(fyne.Position{X: 400, Y: 500})
	diagramWidget.Nodes["node4"] = node4

	link0 := diagramwidget.NewDiagramLink(diagramWidget, node0, node1)
	diagramWidget.Links["link0"] = link0
	link0.AddSourceAnchoredText("sourceRole", "sourceRole")

	link1 := diagramwidget.NewDiagramLink(diagramWidget, node2, node1)
	diagramWidget.Links["link1"] = link1
	link1.LinkColor = color.RGBA{255, 64, 64, 255}
	link1.TargetDecorations = append(diagramWidget.Links["link1"].TargetDecorations, arrowhead.NewArrowhead())
	link1.TargetDecorations = append(diagramWidget.Links["link1"].TargetDecorations, arrowhead.NewArrowhead())
	link1.MidpointDecorations = append(diagramWidget.Links["link1"].MidpointDecorations, arrowhead.NewArrowhead())
	link1.MidpointDecorations = append(diagramWidget.Links["link1"].MidpointDecorations, arrowhead.NewArrowhead())
	link1.SourceDecorations = append(diagramWidget.Links["link1"].SourceDecorations, arrowhead.NewArrowhead())
	link1.SourceDecorations = append(diagramWidget.Links["link1"].SourceDecorations, arrowhead.NewArrowhead())

	diagramWidget.Links["link2"] = diagramwidget.NewDiagramLink(diagramWidget, node0, node3)

	link3 := diagramwidget.NewDiagramLink(diagramWidget, node2, node3)
	link3.AddSourceAnchoredText("sourceRole", "sourceRole")
	link3.AddMidpointAnchoredText("linkName", "Link 3")
	link3.AddTargetAnchoredText("targetRole", "targetRole")
	diagramWidget.Links["link3"] = link3

	diagramWidget.Links["link4"] = diagramwidget.NewDiagramLink(diagramWidget, node4, node3)

	w.SetContent(diagramWidget)

	w.ShowAndRun()
}
