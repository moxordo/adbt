package main

import (
	"fmt"
	"os/exec"

	"github.com/rivo/tview"
	"github.com/spf13/cobra"
)

func main() {
	app := tview.NewApplication()

	var rootCmd = &cobra.Command{
		Use:   "mycli",
		Short: "My CLI application",
	}

	var runCmd = &cobra.Command{
		Use:   "run",
		Short: "Run dbt",
		Run: func(cmd *cobra.Command, args []string) {
			output, err := exec.Command("dbt", "run").CombinedOutput()
			fmt.Printf("Output:\n%s", output)
			if err != nil {
				fmt.Printf("Error executing command: %v\n", err)
				return
			}
			showOutput(app, string(output))
		},
	}

	rootCmd.AddCommand(runCmd)
	rootCmd.Execute()
}

func showOutput(app *tview.Application, output string) {
	textView := tview.NewTextView().
		SetText(output).
		SetDynamicColors(true).
		SetRegions(true).
		SetWordWrap(true)

	if err := app.SetRoot(textView, true).Run(); err != nil {
		panic(err)
	}
}
