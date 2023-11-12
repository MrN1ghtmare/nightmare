/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Starts a new project",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var packageName string
		var err error
		if len(args) == 1 {
			packageName = args[0]
		} else {
			packageName, err = os.Getwd()
			if err != nil {
				log.Fatal(err)
			}

			packageName = filepath.Base(packageName)
		}

		if err := exec.Command("go", "mod", "init", packageName).Run(); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Go module created [%s]!\n", packageName)

		const defaultPerm = 0755
		os.Mkdir("build", defaultPerm)
		os.Mkdir("cmd", defaultPerm)
		os.Mkdir("configs", defaultPerm)
		os.Mkdir("deployments", defaultPerm)
		os.Mkdir("docs", defaultPerm)
		os.Mkdir("examples", defaultPerm)
		os.Mkdir("internal", defaultPerm)
		os.MkdirAll("pkg", defaultPerm)
		os.Mkdir("test", defaultPerm)
		os.MkdirAll("internal/controller", defaultPerm)
		os.MkdirAll("internal/domain", defaultPerm)
		os.MkdirAll("internal/domain/entity", defaultPerm)
		os.MkdirAll("internal/domain/port", defaultPerm)
		os.MkdirAll("internal/domain/service", defaultPerm)

		_, err = os.Create("./configs/nightmare.yml")

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Default directories created!")
		fmt.Println("Project initialized!")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
