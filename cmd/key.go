package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.thunderatz.org/tleague/pkg/do"
)

func init() {
	keyCmd.AddCommand(keyListCmd)
	rootCmd.AddCommand(keyCmd)
}

var keyCmd = &cobra.Command{
	Use:   "key",
	Short: "key utilities",

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		checkTokenPresent()

		c = do.NewClient(viper.GetString("token"))
	},
}

var keyListCmd = &cobra.Command{
	Use:   "list",
	Short: "list created keys",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("> Authenticating and finding keys...")
		keys, err := c.KeyList()
		cobra.CheckErr(err)

		fmt.Println("keys:")

		for _, k := range keys {
			fmt.Println(k.ID, k.Name)
		}
	},
}
