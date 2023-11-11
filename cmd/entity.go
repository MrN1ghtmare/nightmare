/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

// entityCmd represents the entity command
var entityCmd = &cobra.Command{
	Use:   "entity",
	Short: "Create a new entity in your project",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		pascalCaseRegex := regexp.MustCompile(`^[A-Z][a-z0-9]*$`)
		entityName := args[0]
		if !pascalCaseRegex.MatchString(entityName) {
			log.Fatalln("write entity name following PascalCase pattern")
		}

		fileName := pascalCaseRegex.ReplaceAllString(entityName, "${1}_${2}")
		if fileName == "_" {
			fileName = strings.ToLower(entityName)
		} else {
			fileName = strings.ToLower(fileName)
		}

		fileName = fileName + ".go"

		srcFile, err := os.Create(fmt.Sprintf("./internal/domain/entity/%s", fileName))
		if err != nil {
			log.Fatalln(err)
		}

		_, err = srcFile.WriteString(
			fmt.Sprintf(
				"package entity\n\ntype %s struct {\n}\n\nfunc New%s() %s {\n\treturn %s{}\n}\n",
				entityName,
				entityName,
				entityName,
				entityName,
			),
		)

		if err != nil {
			log.Fatalln(err)
		}

		log.Printf("entity %s created!\n", entityName)
	},
}

func init() {
	createCmd.AddCommand(entityCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// entityCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// entityCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
