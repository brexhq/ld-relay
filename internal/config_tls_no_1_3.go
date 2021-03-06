// +build !go1.12

package internal

import "crypto/tls"

func ParseTLSVersion(name string) (uint16, bool) {
	switch name {
	case "1.0":
		return tls.VersionTLS10, true
	case "1.1":
		return tls.VersionTLS11, true
	case "1.2":
		return tls.VersionTLS12, true
	default:
		return 0, false
	}
}
