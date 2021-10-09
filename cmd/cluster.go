package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.thunderatz.org/tleague/pkg/do"
)

var clusterName string

func init() {
	clusterCreateCmd.Flags().StringVar(&clusterName, "name", "tleague", `cluster base name (default "tleague")`)

	clusterCmd.AddCommand(clusterCreateCmd)

	rootCmd.AddCommand(clusterCmd)
}

var clusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "cluster utilities",

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		checkTokenPresent()

		c = do.NewClient(viper.GetString("token"))
	},
}

var clusterCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create a cluster",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("> Creating cluster", viper.GetString("cluster"))
		cluster, err := c.ClusterCreate(clusterName)
		cobra.CheckErr(err)
		fmt.Println("> cluster", cluster.Name, "created with ID", cluster.ID)
	},
}
