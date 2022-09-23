package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/mpreath/netcalc/network"
	"github.com/spf13/cobra"
)

var vlsmCmd = &cobra.Command{
	Use:   "vlsm",
	Short: "Given a network and comma-separated list of subnet lengths break it into smaller networks",
	Long: `
This command subnets a network based on a comma-separated list of subnet lengths.
Usage: netcalc vlsm <vlsm list> <ip_address> <subnet_mask>.`,
	Args: cobra.MinimumNArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		net, err := network.GenerateNetwork(args[1], args[2])
		if err != nil {
			log.Fatal(err)
		}
		// generate network from args
		node := network.NetworkNode{
			Network: net,
		}

		vlsm_args := strings.Split(args[0], ",")
		var vlsm_list = make([]int, len(vlsm_args))
		for idx, val := range vlsm_args {
			vlsm_list[idx], err = strconv.Atoi(val)
			if err != nil {
				log.Fatal(err)
			}
		}
		sort.Slice(vlsm_list, func(i, j int) bool {
			return vlsm_list[i] < vlsm_list[j]
		})

		for _, vlsm := range vlsm_list {
			err = network.SplitToVlsmCount(&node, vlsm)

			if err != nil {
				log.Fatal(err)
			}
		}

		if JSON_FLAG {
			// json output
			s, _ := json.MarshalIndent(node, "", "  ")
			fmt.Println(string(s))
		} else {
			// std output
			printNetworkTree(&node)
		}

	},
}

func init() {
	// subnetCmd.Flags().IntVar(&HOST_COUNT, "hosts", 0, "Specifies the number of hosts to include each subnet.")
	// subnetCmd.Flags().IntVar(&NET_COUNT, "networks", 0, "Specifies the number of subnets to create.")
	rootCmd.AddCommand(vlsmCmd)
}
