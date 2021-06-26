package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"go.thunderatz.org/tleague/pkg/docker"
)

var path string
var dockerfile string
var imageTag string

func init() {
	buildCmd.Flags().StringVarP(&path, "path", "p", ".", `caminho da pasta para buildar (default ".")`)
	buildCmd.Flags().StringVarP(&dockerfile, "dockerfile", "f", "tests.Dockerfile", `nome do dockerfile dentro da pasta (default "tests.Dockerfile")`)
	buildCmd.Flags().StringVarP(&imageTag, "tag", "t", "teste", `image tag (default "teste")`)

	rootCmd.AddCommand(buildCmd)
}

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "build ThunderLeague test docker image",

	Run: func(cmd *cobra.Command, args []string) {
		if err := docker.BuildImage(path, dockerfile, imageTag); err != nil {
			fmt.Println(err)
		}
	},
}
