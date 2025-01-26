package core

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

var (
	once     sync.Once
	rowMutex sync.Mutex
)

func StartListening(options ListenArgs) {
	// TODO device should depends on OS
	handle, err := pcap.OpenLive("en0", 65535, true, pcap.BlockForever)
	if err != nil {
		log.Fatal("Error opening device:", err)
	}
	defer handle.Close()

	// Create a packet source to process packets
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	fmt.Println("Listening for outgoing requests... Press Ctrl + C to quit")

	for packet := range packetSource.Packets() {
		// Decode packet layers
		networkLayer := packet.NetworkLayer()
		if networkLayer == nil {
			continue
		}

		var destinationIP string

		// Handle IPv4 and IPv6
		switch ipLayer := networkLayer.(type) {
		case *layers.IPv4:
			{
				if options.IpLayer == "v6" {
					continue
				}
				destinationIP = ipLayer.DstIP.String()
			}
		case *layers.IPv6:
			{
				if options.IpLayer == "v4" {
					continue
				}
				destinationIP = ipLayer.DstIP.String()
			}
		default:
			continue // Skip non-IP packets
		}

		// Extract transport layer (TCP/UDP)
		transportLayer := packet.TransportLayer()
		if transportLayer == nil {
			continue
		}

		switch transport := transportLayer.(type) {
		case *layers.TCP:
			{
				if options.Protocol == "upd" {
					continue
				}

				if options.HttpOnly == true {
					filterAndPrintHttpRequests(transport, destinationIP)
				} else {
					printConnectionInfo(destinationIP, int(transport.DstPort), "TCP")
				}
			}
		case *layers.UDP:
			if options.Protocol == "tcp" || options.HttpOnly == true {
				continue
			}
			printConnectionInfo(destinationIP, int(transport.DstPort), "UDP")
		}
	}
}

func filterAndPrintHttpRequests(transport *layers.TCP, destinationIP string) {
	if transport.DstPort == 80 || transport.DstPort == 443 {
		protocol := "HTTP"
		if transport.DstPort == 443 {
			protocol = "HTTPS"
		}
		printConnectionInfo(destinationIP, int(transport.DstPort), protocol)
	}
}

var timestampFormat = "2006-01-02 15:04:05" // Default timestamp format TODO: should be possible to modify

func printConnectionInfo(destinationIP string, port int, protocol string) {
	// Ensure headers are printed only once
	once.Do(func() {
		fmt.Printf("%-20s %-30s %-10s %-20s\n", "TIMESTAMP", "DESTINATION", "PORT", "PROTOCOL")
		fmt.Printf("%-20s %-30s %-10s %-20s\n", strings.Repeat("-", 20), strings.Repeat("-", 30), strings.Repeat("-", 10), strings.Repeat("-", 20))
	})

	destinationHostname := resolveHostname(destinationIP)
	if destinationHostname == "" {
		destinationHostname = destinationIP
	}

	currentTime := time.Now().Format(timestampFormat)

	// Print connection info as a table row
	rowMutex.Lock()
	defer rowMutex.Unlock()
	fmt.Printf("%-20s %-30s %-10d %-10s\n", currentTime, destinationHostname, port, protocol)
}
