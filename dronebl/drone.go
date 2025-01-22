package dronebl

import (
	"errors"
	"fmt"
	"net"
	"strings"
)

const (
	droneblDNS = "dnsbl.dronebl.org."
)

func Probe(ip string) error {

	revip := reverseOctets(ip)

	queryIP := fmt.Sprintf("%s.%s", revip, droneblDNS)

	_, err := net.LookupIP(queryIP)
	if err != nil {
		var errDNS *net.DNSError
		if errors.As(err, &errDNS) {
			if errDNS.IsNotFound {
				return nil
			}
		}

		return err
	}

	return fmt.Errorf("IP %s informed in %s", ip, droneblDNS)
}

func reverseOctets(ip string) string {

	octets := strings.SplitN(ip, ".", 4)

	return fmt.Sprintf("%s.%s.%s.%s", octets[3], octets[2], octets[1], octets[0])

}
