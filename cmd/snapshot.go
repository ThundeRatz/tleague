package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.thunderatz.org/tleague/pkg/do"
)

func init() {
	snapshotCmd.AddCommand(snapshotListCmd)
	rootCmd.AddCommand(snapshotCmd)
}

var snapshotCmd = &cobra.Command{
	Use:     "snapshot",
	Short:   "snapshot utilities",
	Aliases: []string{"snap"},

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		c = do.NewClient(viper.GetString("token"))
	},
}

var snapshotListCmd = &cobra.Command{
	Use:   "list",
	Short: "list account snapshots",

	Run: func(cmd *cobra.Command, args []string) {
		ifVerbose("Authenticating and finding droplets...")
		snapshots, err := c.SnapshotList()
		cobra.CheckErr(err)
		ifVerbose("Done...")

		fmt.Println("Snapshots:")

		for _, d := range snapshots {
			fmt.Println(d.ID, d.Name)
		}
	},
}
