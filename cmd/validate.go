package main

import (
	"github.com/spf13/cobra"
)

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "validate your stuff or whatever",
	Long:  "shibbidy doobidy goobidy",
}

func init() {

	rootCmd.AddCommand(validateCmd)
}
