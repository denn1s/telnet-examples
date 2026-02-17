package main

import (
	"log"

	"github.com/miekg/dns"
)

// Map of domain names to IPs
var domainMap = map[string]string{
	"robertillo.com.": "192.168.1.130",
	"bran.neto.":      "192.168.1.145",
	"victor.net.":     "192.168.1.132",
}

func handleDNS(w dns.ResponseWriter, r *dns.Msg) {
	m := new(dns.Msg)
	m.SetReply(r)
	m.Authoritative = true

	// Process each question in the DNS query
	for _, q := range r.Question {
		ip, found := domainMap[q.Name] // Look up the IP for the domain name
		if found && q.Qtype == dns.TypeA {
			rr, err := dns.NewRR(q.Name + " 3600 IN A " + ip)
			if err != nil {
				log.Printf("Error creating A record: %v", err)
				continue
			}
			m.Answer = append(m.Answer, rr)
		}
	}

	// Write the response back to the client
	w.WriteMsg(m)
}

func main() {
	// Create a new DNS server
	server := &dns.Server{Addr: ":53", Net: "udp"}
	dns.HandleFunc(".", handleDNS) // Handle all DNS queries with handleDNS

	log.Println("Starting DNS server on :53")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
