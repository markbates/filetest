package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var configPath string
var failFast bool
var pwd string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "filetest",
	Short: "Runs tests against files/file system",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
	Run: func(cmd *cobra.Command, args []string) {
		err := Run(configPath)
		if err != nil {
			color.Red(err.Error())
			os.Exit(-1)
		}
		if len(errs) > 0 {
			color.Red("There were %d test failures:\n", len(errs))
			for _, e := range errs {
				color.Red("\t%s\n", e.Error())
			}
			os.Exit(-1)
		}
		color.Green("All of your tests have passed with flying colors!")
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func Run(path string) error {
	pwd, _ = os.Getwd()
	s, err := os.Stat(path)
	if err != nil {
		return err
	}
	if s.IsDir() {
		return RunDir(path)
	}
	return RunFile(path)
}

func RunDir(path string) error {
	return filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			return RunFile(p)
		}
		return nil
	})
}

func RunFile(path string) error {
	ext := strings.ToLower(filepath.Ext(path))
	if ext != ".json" {
		return errors.Errorf("%s is not a .json file", path)
	}
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	ff := Files{}
	err = json.Unmarshal(b, &ff)
	if err != nil {
		return err
	}
	return ff.Test()
}

func init() {
	RootCmd.Flags().StringVarP(&configPath, "config", "c", "filetest.json", "path to the filetest.json or directory containing many json files you want to run")
	RootCmd.Flags().BoolVarP(&failFast, "fail-fast", "f", false, "fail fast if there are any errors")
}
