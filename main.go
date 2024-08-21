package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/hmcalister/TeaTimer/internal/tui"
)

func main() {
	mainInterface := tui.NewMainModel()
	if _, err := tea.NewProgram(mainInterface).Run(); err != nil {
		fmt.Printf("error encountered during bubble tea program: %v", err)
		os.Exit(1)
	}
}
