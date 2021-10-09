package cmd

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.thunderatz.org/tleague/pkg/do"
)

var c *do.Client

func init() {
	dropletCreateCmd.Flags().String("snapshot", "thunderleague", "snapshot name")
	dropletCreateCmd.Flags().String("name", "thunderleague", "droplet name")
	dropletCreateCmd.Flags().String("key", "tleague", "key name")

	viper.BindPFlag("snapshot", dropletCreateCmd.Flags().Lookup("snapshot"))
	viper.BindPFlag("droplet", dropletCreateCmd.Flags().Lookup("name"))
	viper.BindPFlag("key", dropletCreateCmd.Flags().Lookup("key"))

	dropletCmd.AddCommand(dropletListCmd)
	dropletCmd.AddCommand(dropletDestroyCmd)
	dropletCmd.AddCommand(dropletCreateCmd)

	rootCmd.AddCommand(dropletCmd)
}

var dropletCmd = &cobra.Command{
	Use:   "droplet",
	Short: "droplet utilities",

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		checkTokenPresent()

		c = do.NewClient(viper.GetString("token"))
	},
}

var dropletListCmd = &cobra.Command{
	Use:   "list",
	Short: "list created droplets",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("> Authenticating and finding droplets...")
		droplets, err := c.DropletList()
		cobra.CheckErr(err)

		if len(droplets) == 0 {
			fmt.Println("Nenhum droplet encontrado")
			return
		}

		fmt.Println("Droplets:")

		for _, d := range droplets {
			fmt.Println(d.ID, d.Name, do.DropletGetPublicIP(d))
		}
	},
}

var dropletCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create a droplet from a snaphot",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("> Finding snapshot", viper.GetString("snapshot"))
		snap, err := c.GetSnapshotByName(viper.GetString("snapshot"))
		cobra.CheckErr(err)
		fmt.Println("> Found snapshot with ID", snap.ID)

		fmt.Println("> Finding SSH Key", viper.GetString("key"))
		key, err := c.KeyGetByName(viper.GetString("key"))

		if err != nil {
			prompt := promptui.Select{
				Label: "Error finding key, do you want to create a new key?",
				Items: []string{"Yes", "No"},
			}
			_, result, perr := prompt.Run()
			cobra.CheckErr(perr)

			if result == "Yes" {
				fmt.Println("> Creating SSH Key", viper.GetString("key"))
				key, err = c.KeyCreateDefault(viper.GetString("key"))
			}
		}

		cobra.CheckErr(err)
		fmt.Println("> Found key with ID", key.ID)

		fmt.Println("> Creating droplet", viper.GetString("droplet"))
		drop, err := c.DropletCreateC32(viper.GetString("droplet"), snap.ID, key.ID)
		cobra.CheckErr(err)
		fmt.Println("> Droplet", drop.Name, "created with ID", drop.ID)
	},
}

var dropletDestroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "destroy a droplet by name",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("> Getting droplet list")
		droplets, err := c.DropletList()
		cobra.CheckErr(err)

		for _, droplet := range droplets {
			if droplet.Name != viper.GetString("droplet") {
				continue
			}

			prompt := promptui.Select{
				Label: fmt.Sprintf("Destroy droplet %d?", droplet.ID),
				Items: []string{"Yes", "No"},
			}
			_, result, perr := prompt.Run()
			cobra.CheckErr(perr)

			if result == "Yes" {
				fmt.Println("> Destroying droplet", droplet.ID)
				err := c.DropletDestroy(droplet.ID)
				if err != nil {
					fmt.Println("Failed to delete droplet", err)
				}
				fmt.Println("> Done")
			}
		}
	},
}
