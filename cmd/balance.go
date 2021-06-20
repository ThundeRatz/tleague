package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.thunderatz.org/tleague/pkg/do"
)

func init() {
	rootCmd.AddCommand(balanceCmd)
}

var balanceCmd = &cobra.Command{
	Use:     "balance",
	Short:   "gets current accout balance",
	Aliases: []string{"credits"},

	Run: func(cmd *cobra.Command, args []string) {
		checkTokenPresent()

		c = do.NewClient(viper.GetString("token"))
		b, err := c.GetCredits()

		cobra.CheckErr(err)

		fmt.Println(b)
	},
}
