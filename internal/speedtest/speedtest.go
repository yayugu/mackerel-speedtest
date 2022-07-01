package speedtest

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

type SpeedTest struct {
	Path     string
	ServerId uint64
	Result
}

type Result struct {
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

func (s *SpeedTest) IsInstalled() error {
	speedtestCmd := exec.Command(s.Path, "--version")
	out, err := speedtestCmd.Output()
	if err != nil {
		return err
	}
	if !strings.Contains(string(out), "Speedtest by Ookla") {
		return fmt.Errorf("%s is not an official speedtest CLI provided by Ookla", s.Path)
	}
	return nil
}

func (s *SpeedTest) Run() error {
	speedtestCmd := exec.Command(s.Path, fmt.Sprintf("--server-id=%d", s.ServerId), "--format=json")
	out, err := speedtestCmd.Output()
	if err != nil {
		return err
	}
	err2 := json.Unmarshal(out, &s.Result)
	if err2 != nil {
		return err2
	}
	return nil
}
