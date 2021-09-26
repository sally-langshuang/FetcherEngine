package demo

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T)  {
	str:="{\"minPortInclude\": 2000, \"maxPortInclude\": 40000, \"containerConf\": {\"priority\": \"STABLE\", \"baseEnv\": \"system\"}, \"resourceRequirement\": {\"network\": {\"outBandwidthMBPS\": {\"limit\": -1, \"quota\": 10}, \"inBandwidthMBPS\": {\"limit\": -1, \"quota\": 10}}, \"process\": {\"thread\": {\"limit\": -1, \"quota\": 1000}}, \"gpus\": [{\"exclusive\": true, \"numCores\": {\"limit\": -1, \"quota\": 386}, \"shareMode\": \"GPU_DEFAULT\", \"model\": \"Tesla P4\", \"videoMemoryMB\": {\"limit\": -1, \"quota\": 7415}, \"mixLowPriorityEnable\": false}], \"workspace\": {\"sizeMB\": {\"limit\": 10240, \"quota\": 10240}, \"exclusive\": false, \"type\": \"home\", \"inode\": {\"limit\": 500000, \"quota\": 500000}, \"sizeMBLimit\": 10240}, \"memory\": {\"sizeMB\": {\"limit\": -1, \"quota\": 1024}, \"ipcEnabled\": true}, \"volumes\": [], \"cpu\": {\"normalizedCore\": {\"limit\": -1, \"quota\": 10}}}, \"moduleConfMap\": {\"sandboxtestAPI\": {\"packageConf\": {\"packageSource\": \"http://xforce-data.baidu-int.com/gpu/online/deployments/archers/6-sandbox/testAPI/53252.testAPI-2-hn.tar.gz\", \"packageType\": \"ARCHER3\", \"packageVersion\": \"1.0\"}, \"deployConf\": {\"deployDir\": \"/home/work\", \"deployCmdOptions\": \"--timeout=12000\", \"deployTimeoutSec\": 12000, \"user\": \"work\"}, \"bns_tags\": {}, \"monitorConf\": {\"process\": {\"main\": \"sandboxtestAPI\"}, \"useCustomHealthCheckOnly\": true, \"healthCheckTimeoutSec\": 30}, \"migratePodIfError\": false, \"upstreamModuleList\": [], \"type\": \"PROCESS\", \"port\": {\"staticPort\": {}, \"dynamicPortName\": [\"main\"], \"AnonymousPortsCount\": 0}, \"ports\": {\"staticPort\": {}, \"dynamicPortName\": [\"main\"], \"AnonymousPortsCount\": 0}}}}"
	podDescStruct := PodDescription{}
	ConvertStringToStruct(str, &podDescStruct)
	fmt.Println(podDescStruct)
}
