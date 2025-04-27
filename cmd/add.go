/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/upsaurav12/taskido-cli/storage"
)

// taskCmd represents the task command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "command for creating, updating , deleting your task.",
	Long:  `command for creating, updating , deleting your task.`,
	Run: func(cmd *cobra.Command, args []string) {
		// initialModel()

		title := args[0]

		tasks, err := storage.LoadTasks()
		if err != nil {
			fmt.Println("Error Loading Tasks", err)
			return
		}

		newTask := storage.Task{
			Title: title,
			Done:  true,
			ID:    len(tasks) + 1,
		}

		tasks = append(tasks, newTask)

		if err := storage.SaveTasks(tasks); err != nil {
			fmt.Println("Error saveing Task", err)
			return
		}

		fmt.Println("Task added:", title)

	},
}

// type (
// 	errMsg error
// )

// type model struct {
// 	textInpit textinput.Model
// 	err       error
// }

// func initialModel() model {
// 	fmt.Println("Hello world")
// 	ti := textinput.New()
// 	ti.Placeholder = "Coding..."
// 	ti.Focus()
// 	ti.CharLimit = 156
// 	ti.Width = 20

// 	return model{
// 		textInpit: ti,
// 		err:       nil,
// 	}
// }

// func (m model) Init() tea.Cmd {
// 	return textinput.Blink
// }

// func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	var cmd tea.Cmd

// 	switch msg := msg.(type) {
// 	case tea.KeyMsg:
// 		switch msg.Type {
// 		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
// 			return m, tea.Quit
// 		}
// 	case errMsg:
// 		m.err = msg
// 		return m, nil
// 	}

// 	m.textInpit, cmd = m.textInpit.Update(msg)
// 	return m, cmd
// }

// func (m model) View() string {
// 	return fmt.Sprintf(
// 		"Enter your Task?\n\n%s\n\n%s",
// 		m.textInpit.View(),
// 		"(esc to quit)",
// 	) + "\n"
// }

func init() {
	rootCmd.AddCommand(addCmd)
}
