package demo

import (
	"encoding/json"
	"fmt"
)

func ConvertStringToStruct(srcStr string, destStruct interface{}) error {
	if srcStr == "" {
		msg := fmt.Sprintf("ConvertStringToStruct failed, invalid input: src: %v", srcStr)
		print(msg)
		return nil
	}

	if err := json.Unmarshal([]byte(srcStr), &destStruct); err != nil {
		msg := fmt.Sprintf("ConvertStringToStruct failed, err: %v", err)
		print(msg)
		return err
	}

	return nil
}
type PodDescription struct {
	ExpectedConfiguration
}
type ExpectedConfiguration struct {
	ModuleConfMap       map[string]Module      `json:"moduleConfMap"`
	ShardPodCounts      []int                  `json:"shard_pod_counts"`
	MaxPortInclude      int                    `json:"maxPortInclude,omitempty"`
	MinPortInclude      int                    `json:"minPortInclude,omitempty"`
	HostTagConstraints  []string               `json:"host_tag_constraints"`
	ExtraInfo           map[string]interface{} `json:"extra_info"`
}
type Module struct {
	DeployConf         DeployConf        `json:"deployConf"`
	KeepAliveOnHost    KeepAliveOnHost   `json:"keepAliveOnHost"`
	MigratePodIfError  bool              `json:"migratePodIfError"`
	MonitorConf        MonitorConf       `json:"monitorConf"`
	PackageConf        PackageConf       `json:"packageConf"`
	Port               PortForSolaria    `json:"port"`
	UpstreamModuleList []string          `json:"upstreamModuleList"`
	BNSTags            map[string]string `json:"bns_tags"`
	Deployed           bool              `json:"deployed"`
}

type DeployConf struct {
	DeployDir        string `json:"deployDir"`
	DeployTimeoutSec int    `json:"deployTimeoutSec"`
	User             string `json:"user"`
}

type KeepAliveOnHost struct {
	MaxKeepAliveCount int `json:"maxKeepAliveCount"`
}

type PortForSolaria struct {
	Range               Range       `json:"range"`
	DynamicPortNames    []string               `json:"dynamicPortName"` // 与Opera DB中t_deployment.expected_configuration使用的变量格式一致
	AnonymousPortsCount int                    `json:"anonymousPortsCount"`
	StaticPort          map[string]interface{} `json:"staticPort"`
}

type Range struct {
	Min int `json:"min"`
	Max int `json:"max"`
}

type PackageConf struct {
	PackageSource  string `json:"packageSource"`
	PackageType    string `json:"packageType"`
	PackageVersion string `json:"packageVersion"`
}
type MonitorConf struct {
	HealthCheckTimeoutSec    int               `json:"healthCheckTimeoutSec"`
	Process                  map[string]string `json:"process"`
	UseCustomHealthCheckOnly bool              `json:"useCustomHealthCheckOnly"`
}


