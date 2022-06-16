package speedtest

import (
	"encoding/json"
	"os/exec"
)

type SpeedTestResult struct {
	Type      string
	Timestamp string
	Ping      struct {
		Jitter  float64
		Latency float64
	}
	Download struct {
		Bandwidth uint64
		Bytes     uint64
		Elapsed   uint64
	}
	Upload struct {
		Bandwidth uint64
		Bytes     uint64
		Elapsed   uint64
	}
	PacketLoss float64
	Isp        string
	Interface  struct {
		InternalIp string
		Name       string
		MacAddr    string
		IsVpn      bool
		ExternalIp string
	}
	Server struct {
		Id       uint64
		Host     string
		Port     uint64
		Name     string
		Location string
		Country  string
		Ip       string
	}
	Result struct {
		Id        string
		Url       string
		Persisted bool
	}
}

func Run(result *SpeedTestResult) error {
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
