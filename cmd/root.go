package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mytool",
	Short: "MyTool is a CLI utility with multiple functions",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
