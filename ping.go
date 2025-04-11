package icmp

import (
	probing "github.com/prometheus-community/pro-bing"
	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/icmp", new(IcmpModule))
}

type IcmpModule struct{}

func (m *IcmpModule) Ping(hostname string) (*probing.Statistics, error) {
	pinger, err := probing.NewPinger(hostname)
	if err != nil {
		panic(err)
	}
	pinger.Count = 3

	err = pinger.Run()
	if err != nil {
		panic(err)
	}

	stats := pinger.Statistics()
	return stats, nil
}
