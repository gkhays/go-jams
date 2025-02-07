package fips

import (
	"os"
	"strings"
)

// Determine whether the underlying operating system is running in FIPS mode
func IsFIPSModeEnabled() bool {
	// Check environment variable typically used on Linux systems
	if os.Getenv("OPENSSL_FIPS") == "1" {
		return true
	}

	// Check for FIPS mode flag in system configuration files
	fipsFiles := []string{
		"/proc/sys/crypto/fips_enabled",
		"/etc/system-fips",
	}

	for _, file := range fipsFiles {
		content, err := os.ReadFile(file)
		if err == nil && strings.TrimSpace(string(content)) == "1" {
			return true
		}
	}

	return false
}
