/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// glanceCmd represents the glance command
var glanceCmd = &cobra.Command{
	Use:   "glance",
	Short: "clug glance (<survey-name> | <path>) --(wcd | bathy | trackline)",
	Long: `Summarizes survey information including file count and files' size 
	for an equivalent get command`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("glance called")
	},
}

func init() {
	rootCmd.AddCommand(glanceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// glanceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// glanceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
