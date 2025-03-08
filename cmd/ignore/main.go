package main

import (
	"log"
	"os"

	"github.com/neptship/ignore/internal"
	"github.com/spf13/cobra"
)

const usageMessage = `Usage:
ignore [language]

Examples:
	ignore go
	ignore python
	ignore java
`

func handlePanic() {
	if err := recover(); err != nil {
		log.Fatal("crashed", "err", err)
		os.Exit(1)
	}
}

func main() {
	defer handlePanic()

	var rootCmd = &cobra.Command{
		Use:   "ignore",
		Short: "Create .gitignore files quickly and simply",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				cmd.Help()
				os.Exit(1)
			}

			internal.AddIgnoreTemplate(".gitignore", args[0])
		},
	}

	rootCmd.SetUsageFunc(func(cmd *cobra.Command) error {
		cmd.Print(usageMessage)
		return nil
	})

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
		internal.CallClear()
	}
}
