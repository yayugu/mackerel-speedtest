package speedtest

import (
	"encoding/json"
	"os/exec"
	"time"

	"github.com/showwin/speedtest-go/speedtest"
)

type SpeedTestResult struct {
	Latency time.Duration
	DLSpeed float64
	ULSpeed float64
}

func Run_(result *SpeedTestResult) error {
	speedtestCmd := exec.Command("speedtest", "--server-id=21569", "--format=json")
	out, err := speedtestCmd.Output()
	if err != nil {
		return err
	}
	err2 := json.Unmarshal(out, &result)
	if err2 != nil {
		return err2
	}
	return nil
}

func Run(result *SpeedTestResult) error {
	user, err := speedtest.FetchUserInfo()
	if err != nil {
		return err
	}

	serverList, _ := speedtest.FetchServers(user)
	if err != nil {
		return err
	}

	targets, err := serverList.FindServer([]int{21569})
	if err != nil {
		return err
	}

	for _, s := range targets {
		s.PingTest()
		s.DownloadTest(false)
		s.UploadTest(false)

		result.Latency = s.Latency
		result.DLSpeed = s.DLSpeed
		result.ULSpeed = s.ULSpeed
	}

	return nil
}
