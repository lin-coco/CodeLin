package lin

import (
	"math"
	"math/rand"
	"time"
)

// 选择的地址要均衡、响应快
// 健康检查
type LoadBalance struct {
	ips         []string
	roundFacter int
	conns       []int
}

func NewLoadBalance(ips []string) *LoadBalance {
	if len(ips) <= 0 {
		return nil
	}
	return &LoadBalance{ips: ips, roundFacter: -1}
}

// 随机
func (o *LoadBalance) Random() string {
	rand.Seed(time.Now().UnixNano())
	return o.ips[rand.Intn(len(o.ips))]
}

// 轮训
func (o *LoadBalance) Round() string {
	if o.roundFacter == len(o.ips)-1 {
		o.roundFacter = -1
	}
	o.roundFacter++
	return o.ips[o.roundFacter]
}

// 一致性hash
func (o *LoadBalance) Hash(clientIp string) string {
	// 计算Hash值
	hash := func(clientIp string) uint {
		var hash uint
		for i := 0; i < len(clientIp); i++ {
			hash += uint(clientIp[i] - '0')
		}
		hash += hash % 10 * (hash%10 + hash%100)
		return hash
	}(clientIp)
	return o.ips[int(hash%uint(len(clientIp)))]
}

// 最少连接
func (o *LoadBalance) MinConn() string {
	min := math.MaxInt
	for i := 0; i < len(o.conns); i++ {
		if o.conns[i] < min {
			min = i
		}
	}
	return o.ips[min]
}

// 最小响应
func (o *LoadBalance) MinResponse() string {
	return func(ips []string) string {
		// TODO ping ip and get time
		rand.Seed(time.Now().UnixNano())
		//pingTimes := make([]time.Duration, len(ips))
		min := math.MaxInt
		for i := 0; i < len(ips); i++ {
			pingTimes := rand.Intn(100)
			if pingTimes < min {
				min = i
			}
		}
		return ips[min]
	}(o.ips)
}
