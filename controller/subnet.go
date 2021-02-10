package controller

import (
	"bytes"
	"net"
	"sort"
	"strconv"
	"strings"
	"testCDN/model"
)

func (c *Controller) SortSubnet(config model.Config) (model.Config, error) {
	sorted := sorkByIP(config.Subnets)
	config.Subnets = sorkByMask(sorted)
	return config, nil
}

func sorkByMask(subnet []model.Subnet) []model.Subnet {
	sort.SliceStable(subnet, func(i, j int) bool {
		return getMask(subnet[i].Subnet) > getMask(subnet[j].Subnet)
	})
	return subnet
}

func sorkByIP(subnet []model.Subnet) []model.Subnet {
	sort.SliceStable(subnet, func(i, j int) bool {
		return bytes.Compare(net.ParseIP(getIP(subnet[i].Subnet)), net.ParseIP(getIP(subnet[j].Subnet))) < 0
	})
	return subnet
}

func getMask(subnet string) int {
	tmp := strings.Split(subnet, "/")
	mask, err := strconv.Atoi(tmp[1])
	if err != nil {
		return 0
	}
	return mask
}

func getIP(subnet string) string {
	return strings.Split(subnet, "/")[0]
}
