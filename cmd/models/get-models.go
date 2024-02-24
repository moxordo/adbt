// Demo code for a timer based update
package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"

	"github.com/rivo/tview"
)

var (
	view *tview.Modal
	app  *tview.Application
)

//
func listModels() {
	cmd := exec.Command("dbt", "list", "--resource-type", "model")

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Command failed: %s", err)
	}

	fmt.Printf("Output:\n%s", out.String())
}



func main() {
	app = tview.NewApplication()

	inputField := tview.NewInputField().
		SetLabel("🚬").
		SetFieldWidth(20)

	// The overall layout of the application
	flex := tview.NewFlex().
		AddItem(inputField, 0, 1, false).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("Contents"), 0, 2, false)

	if err := app.SetRoot(flex, false).Run(); err != nil {
		panic(err)
	}
}
