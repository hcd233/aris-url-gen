// Package cmd 命令行工具
//
//	update 2024-08-11 01:57:31
package cmd

import (
	"fmt"
	"os"

	"github.com/hcd233/Aris-url-gen/internal/logger"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "",
	Short: "Aris Blog API",
	Long:  `Aris Blog API`,
}

// Execute 执行命令行
//
//	author centonhuang
//	update 2024-12-05 16:05:32
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		msg := fmt.Sprintf("[Command] failed to execute command: %v", err)
		logger.Logger.Error(msg)
		os.Exit(1)
	}
}
