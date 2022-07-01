package mackerel

import (
	"mackerel-speedtest/internal/speedtest"
	"time"

	"github.com/mackerelio/mackerel-client-go"
)

type MackerelClient struct {
	Client      *mackerel.Client
	ServiceName string
}

func NewMackerelClient(apiKey string, serviceName string) *MackerelClient {
	m := new(MackerelClient)
	m.Client = mackerel.NewClient(apiKey)
	m.ServiceName = serviceName
	return m
}

func (m *MackerelClient) PostSpeedtestMetric(result speedtest.Result) error {
	t, err := time.Parse("2006-01-02T15:04:05Z", result.Timestamp)
	if err != nil {
		return err
	}

	unixTimestamp := t.Unix()
	err2 := m.Client.PostServiceMetricValues(m.ServiceName, []*mackerel.MetricValue{
		{
			Name:  "speedtest.ping.latency",
			Time:  unixTimestamp,
			Value: result.Ping.Latency / 1000, // ms -> s
		},
		{
			Name:  "speedtest.ping.jitter",
			Time:  unixTimestamp,
			Value: result.Ping.Jitter / 1000, // ms -> s
		},
		{
			Name:  "speedtest.bandwidth.download",
			Time:  unixTimestamp,
			Value: result.Download.Bandwidth * 8, // bytes/s -> bps
		},
		{
			Name:  "speedtest.bandwidth.upload",
			Time:  unixTimestamp,
			Value: result.Upload.Bandwidth * 8, // bytes/s -> bps
		},
	})
	return err2
}

func (m *MackerelClient) CreateGraphDefs() error {
	var payloads = []*mackerel.GraphDefsParam{
		{
			Name: "speedtest.ping",
			Unit: "seconds",
			Metrics: []*mackerel.GraphDefsMetric{
				{
					Name:      "speedtest.ping.latency",
					IsStacked: false,
				},
				{
					Name:      "speedtest.ping.jitter",
					IsStacked: false,
				},
			},
		},
		{
			Name: "speedtest.bandwidth",
			Unit: "bits/sec",
			Metrics: []*mackerel.GraphDefsMetric{
				{
					Name:      "speedtest.bandwidth.download",
					IsStacked: false,
				},
				{
					Name:      "speedtest.bandwidth.upload",
					IsStacked: false,
				},
			},
		},
	}
	err := m.Client.CreateGraphDefs(payloads)
	return err
}
