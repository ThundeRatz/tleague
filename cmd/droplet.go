package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.thunderatz.org/tleague/pkg/do"
)

var c *do.Client

func init() {
	dropletCmd.AddCommand(dropletListCmd)
	rootCmd.AddCommand(dropletCmd)
}

var dropletCmd = &cobra.Command{
	Use:   "droplet",
	Short: "droplet utilities",

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		c = do.NewClient(viper.GetString("token"))
	},
}

var dropletListCmd = &cobra.Command{
	Use:   "list",
	Short: "list created droplets",

	Run: func(cmd *cobra.Command, args []string) {
		ifVerbose("Authenticating and finding snapshots...")
		droplets, err := c.DropletList()
		cobra.CheckErr(err)
		ifVerbose("Done...")

		fmt.Println("Droplets:")

		for _, d := range droplets {
			fmt.Println(d.ID, d.Name)
		}
	},
}
