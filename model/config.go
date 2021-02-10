package model

type Config struct {
	Queues           int         `json:"queues"`
	DefaultBandwidth int         `json:"default_bandwidth"`
	Interfaces       []Interface `json:"interfaces"`
	Protocols        []Protocol  `json:"protocols"`
	Subnets          []Subnet    `json:"subnets"`
}

type Interface struct {
	Id        string   `json:"id"`
	BackupId  string   `json:"backup_id"`
	Bandwidth int      `json:"bandwidth"`
	IPs       []string `json:"ips"`
}

type Protocol struct {
	Name             string `json:"name"`
	Protocol         string `json:"protocol"`
	Ports            []int  `json:"ports"`
	DefaultBandwidth int    `json:"default_bandwidth"`
}

type Subnet struct {
	Subnet    string           `json:"subnet"`
	Bandwidth int              `json:"bandwidth"`
	Protocols []SubnetProtocol `json:"protocols"`
}

type SubnetProtocol struct {
	Name      string `json:"name"`
	Bandwidth int    `json:"bandwidth"`
}
