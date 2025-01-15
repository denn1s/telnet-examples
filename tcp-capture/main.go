package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

type Filters struct {
	sourceIP   string
	destIP     string
	sourcePort int
	destPort   int
	showSYN    bool
	showACK    bool
	showFIN    bool
	showPSH    bool
	showRST    bool
	showURG    bool
}

func main() {
	// Define command line flags
	filters := Filters{}
	flag.StringVar(&filters.sourceIP, "from", "", "Filter source IP address")
	flag.StringVar(&filters.destIP, "to", "", "Filter destination IP address")
	flag.IntVar(&filters.sourcePort, "sport", 0, "Filter source port")
	flag.IntVar(&filters.destPort, "dport", 0, "Filter destination port")
	flag.BoolVar(&filters.showSYN, "syn", false, "Show only SYN packets")
	flag.BoolVar(&filters.showACK, "ack", false, "Show only ACK packets")
	flag.BoolVar(&filters.showFIN, "fin", false, "Show only FIN packets")
	flag.BoolVar(&filters.showPSH, "psh", false, "Show only PSH packets")
	flag.BoolVar(&filters.showRST, "rst", false, "Show only RST packets")
	flag.BoolVar(&filters.showURG, "urg", false, "Show only URG packets")

	flag.Parse()

	// Check if any flag filters are active
	flagFiltersActive := filters.showSYN || filters.showACK || filters.showFIN ||
		filters.showPSH || filters.showRST || filters.showURG

	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}

	// for lan testing:
	handle, err := pcap.OpenLive(devices[0].Name, 1600, true, pcap.BlockForever)
	// for localhost testing:
	// handle, err := pcap.OpenLive("lo", 1600, true, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	err = handle.SetBPFFilter("tcp")
	if err != nil {
		log.Fatal(err)
	}

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	fmt.Println("Starting packet capture...")
	if filters.sourceIP != "" {
		fmt.Printf("Filtering source IP: %s\n", filters.sourceIP)
	}
	if filters.destIP != "" {
		fmt.Printf("Filtering destination IP: %s\n", filters.destIP)
	}
	if filters.sourcePort != 0 {
		fmt.Printf("Filtering source port: %d\n", filters.sourcePort)
	}
	if filters.destPort != 0 {
		fmt.Printf("Filtering destination port: %d\n", filters.destPort)
	}

	for packet := range packetSource.Packets() {
		tcpLayer := packet.Layer(layers.LayerTypeTCP)
		if tcpLayer != nil {
			tcp, _ := tcpLayer.(*layers.TCP)

			ipLayer := packet.Layer(layers.LayerTypeIPv4)
			if ipLayer != nil {
				ip, _ := ipLayer.(*layers.IPv4)

				// Apply filters
				if !shouldShowPacket(ip, tcp, filters, flagFiltersActive) {
					continue
				}

				// Print packet separator
				fmt.Println("\n" + strings.Repeat("=", 50))

				// fmt.Printf("Time: %v\n", time.Now().Format("15:04:05"))
				fmt.Printf("%s ", tcpFlagsToString(tcp))
				fmt.Printf("%s:%d -> %s:%d\n",
					ip.SrcIP, tcp.SrcPort,
					ip.DstIP, tcp.DstPort)
				// fmt.Printf("Sequence: %d\n", tcp.Seq)

				applicationLayer := packet.ApplicationLayer()
				if applicationLayer != nil {
					payload := applicationLayer.Payload()
					if len(payload) > 0 {
						fmt.Println("Payload:")
						fmt.Printf("\n%s\n\n", string(payload))
						// fmt.Printf("Hex: % X\n", payload)
					}
				}

				fmt.Println(strings.Repeat("=", 50))
			}
		}
	}
}

func shouldShowPacket(ip *layers.IPv4, tcp *layers.TCP, filters Filters, flagFiltersActive bool) bool {
	// Check IP filters
	if filters.sourceIP != "" && ip.SrcIP.String() != filters.sourceIP {
		return false
	}
	if filters.destIP != "" && ip.DstIP.String() != filters.destIP {
		return false
	}

	// Check port filters - convert TCPPort to uint16 for comparison
	if filters.sourcePort != 0 && uint16(filters.sourcePort) != uint16(tcp.SrcPort) {
		return false
	}
	if filters.destPort != 0 && uint16(filters.destPort) != uint16(tcp.DstPort) {
		return false
	}

	// Check flag filters only if any flag filter is active
	if flagFiltersActive {
		return (filters.showSYN && tcp.SYN) ||
			(filters.showACK && tcp.ACK) ||
			(filters.showFIN && tcp.FIN) ||
			(filters.showPSH && tcp.PSH) ||
			(filters.showRST && tcp.RST) ||
			(filters.showURG && tcp.URG)
	}

	return true
}

func tcpFlagsToString(tcp *layers.TCP) string {
	var flags []string
	if tcp.FIN {
		flags = append(flags, "FIN")
	}
	if tcp.SYN {
		flags = append(flags, "SYN")
	}
	if tcp.RST {
		flags = append(flags, "RST")
	}
	if tcp.PSH {
		flags = append(flags, "PSH")
	}
	if tcp.ACK {
		flags = append(flags, "ACK")
	}
	if tcp.URG {
		flags = append(flags, "URG")
	}
	return fmt.Sprintf("[%s]", strings.Join(flags, " "))
}
