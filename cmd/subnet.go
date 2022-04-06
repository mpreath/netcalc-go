package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/mpreath/netcalc/network"
	"github.com/spf13/cobra"
)

var HOST_COUNT int
var NET_COUNT int

var subnetCmd = &cobra.Command{
	Use:   "subnet",
	Short: "Given a network break it into smaller networks",
	Long: `
This command subnets a network based on host count and network count parameters.
Usage: netcalc info <ip_address> <subnet_mask>.`,
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		net, err := network.GenerateNetwork(args[0], args[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		// generate network from args
		node := network.NetworkNode{
			Parent:  nil,
			Network: net,
		}

		if HOST_COUNT > 0 {
			err = network.SplitToHostCount(&node, HOST_COUNT)
			if err != nil {
				fmt.Println(err)
				return
			}

		} else if NET_COUNT > 0 {
			err = network.SplitToNetCount(&node, NET_COUNT)
			if err != nil {
				fmt.Println(err)
				return
			}
		}

		if JSON_FLAG {
			// json output
			s, _ := json.MarshalIndent(node, "", "  ")
			fmt.Println(string(s))
		} else {
			// std output
			if VERBOSE_FLAG {
				// verbose output
			}
		}

	},
}

func init() {
	subnetCmd.Flags().IntVar(&HOST_COUNT, "hosts", 0, "Specifies the number of hosts to include each subnet.")
	subnetCmd.Flags().IntVar(&NET_COUNT, "networks", 0, "Specifies the number of subnets to create.")
	rootCmd.AddCommand(subnetCmd)
}
