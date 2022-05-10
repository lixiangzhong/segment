package main

import (
	"encoding/binary"
	"fmt"
	"net/netip"

	"github.com/lixiangzhong/segment"
)

func main() {
	s1 := segment.Must(1, 10, "value1")
	s2 := segment.Must(11, 20, "value1")
	s3 := segment.Must(21, 25, "value3")

	ss := segment.Merge(s1, s2, s3)
	fmt.Println(ss)
	//{1~20:value1}, {21~25:value3}

	fmt.Println(segment.Continuity(s1, s2, s3))
	//true

	fmt.Println(segment.Cover(ss, segment.Must(5, 15, "valueCover")))
	//{1~4:value1}, {5~15:valueCover}, {16~20:value1}, {21~25:value3}

	example_ip()
}

func example_ip() {
	ip1start := ipv4_to_uint32("1.0.0.0")
	ip1end := ipv4_to_uint32("1.0.0.255")
	ip1info := "中国上海"

	ips1 := segment.Must(int64(ip1start), int64(ip1end), ip1info)

	ip2start := ipv4_to_uint32("1.0.1.0")
	ip2end := ipv4_to_uint32("1.0.1.255")
	ip2info := "中国北京"

	ips2 := segment.Must(int64(ip2start), int64(ip2end), ip2info)

	ip3start := ipv4_to_uint32("1.0.0.10")
	ip3end := ipv4_to_uint32("1.0.0.20")
	ip3info := "中国广东"

	ss := segment.Cover(segment.Segments[string]{ips1, ips2}, segment.Must(int64(ip3start), int64(ip3end), ip3info))
	for _, v := range ss {
		fmt.Println(uint32_to_ipv4(uint32(v.Start())), uint32_to_ipv4(uint32(v.End())), v.Value())
	}
}

func ipv4_to_uint32(ip string) uint32 {
	addr, err := netip.ParseAddr(ip)
	if err != nil {
		return 0
	}
	if addr.Is4() {
		return binary.BigEndian.Uint32(addr.AsSlice())
	}
	return 0
}

func uint32_to_ipv4(i uint32) string {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, i)
	ip, ok := netip.AddrFromSlice(b)
	if !ok { //不应该发生
		panic("invalid ip")
	}
	return ip.String()
}
