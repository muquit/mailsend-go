package main

//
// Print certificate chain if ran with --info --printCerts
// With Assistance from Claude AI Sonnet 4
// Much better than my original cert info. My version was
// also wrong, it was printing a intermediate cert as server cert.
// Sep-26-2025 
//

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"strings"
	"time"
)

// CertInfo holds all collected certificate information
type CertInfo struct {
	ConnectionState tls.ConnectionState
	Certificates    []CertDetails
	Fingerprints    map[string]string // algorithm -> fingerprint
}

// CertDetails holds detailed information about a single certificate
type CertDetails struct {
	Certificate *x509.Certificate
	Subject     string
	Issuer      string
	NotBefore   time.Time
	NotAfter    time.Time
	DNSNames    []string
	IsCA        bool
	KeyUsage    x509.KeyUsage
	SerialNumber string
}

// collectTLSConnectionInfo gathers all TLS connection and certificate information
func collectTLSConnectionInfo(state tls.ConnectionState) *CertInfo {
	certInfo := &CertInfo{
		ConnectionState: state,
		Certificates:    make([]CertDetails, 0, len(state.PeerCertificates)),
		Fingerprints:    make(map[string]string),
	}

	// Collect details for each certificate in the chain
	for _, cert := range state.PeerCertificates {
		certDetails := collectCertificateDetails(cert)
		certInfo.Certificates = append(certInfo.Certificates, certDetails)
	}

	// Collect fingerprints for the leaf certificate (first in chain)
	if len(state.PeerCertificates) > 0 {
		leafCert := state.PeerCertificates[0]
		certInfo.Fingerprints["SHA-1"] = collectFingerprint(leafCert.Raw, "SHA-1")
		certInfo.Fingerprints["SHA-256"] = collectFingerprint(leafCert.Raw, "SHA-256")
	}

	return certInfo
}

// collectCertificateDetails extracts detailed information from a certificate
func collectCertificateDetails(cert *x509.Certificate) CertDetails {
	return CertDetails{
		Certificate:  cert,
		Subject:      cert.Subject.String(),
		Issuer:       cert.Issuer.String(),
		NotBefore:    cert.NotBefore,
		NotAfter:     cert.NotAfter,
		DNSNames:     cert.DNSNames,
		IsCA:         cert.IsCA,
		KeyUsage:     cert.KeyUsage,
		SerialNumber: cert.SerialNumber.String(),
	}
}

// collectFingerprint generates a fingerprint for the given certificate data
func collectFingerprint(certRaw []byte, algorithm string) string {
	var hash []byte
	
	switch strings.ToUpper(algorithm) {
	case "SHA-1":
		h := sha1.Sum(certRaw)
		hash = h[:]
	case "SHA-256":
		h := sha256.Sum256(certRaw)
		hash = h[:]
	default:
		return ""
	}

	// Format as colon-separated hex pairs
	var parts []string
	for _, b := range hash {
		parts = append(parts, fmt.Sprintf("%02X", b))
	}
	return strings.Join(parts, ":")
}

// printCollectedCertInfo outputs all collected certificate information
func printCollectedCertInfo(certInfo *CertInfo) {
	if certInfo == nil {
		return
	}

	fmt.Println("\n=== TLS Connection Information ===")
	
	// Connection state info
	fmt.Printf("TLS Version: %s\n", getTLSVersion(certInfo.ConnectionState.Version))
	fmt.Printf("Cipher Suite: %s\n", getCipherSuite(certInfo.ConnectionState.CipherSuite))
	fmt.Printf("Server Name: %s\n", certInfo.ConnectionState.ServerName)
	fmt.Printf("Negotiated Protocol: %s\n", certInfo.ConnectionState.NegotiatedProtocol)
	
	// Certificate chain information
	fmt.Printf("\nCertificate Chain (%d certificates):\n", len(certInfo.Certificates))
	
	for i, certDetails := range certInfo.Certificates {
		fmt.Printf("\n--- Certificate %d ---\n", i+1)
		fmt.Printf("Subject: %s\n", certDetails.Subject)
		fmt.Printf("Issuer: %s\n", certDetails.Issuer)
		fmt.Printf("Serial Number: %s\n", certDetails.SerialNumber)
		fmt.Printf("Not Before: %s\n", certDetails.NotBefore.Format(time.RFC3339))
		fmt.Printf("Not After: %s\n", certDetails.NotAfter.Format(time.RFC3339))
		fmt.Printf("Is CA: %v\n", certDetails.IsCA)
		
		if len(certDetails.DNSNames) > 0 {
			fmt.Printf("DNS Names: %s\n", strings.Join(certDetails.DNSNames, ", "))
		}
		
		// Show key usage
		if certDetails.KeyUsage != 0 {
			fmt.Printf("Key Usage: %s\n", formatKeyUsage(certDetails.KeyUsage))
		}
		
		// Show validity status
		now := time.Now()
		if now.Before(certDetails.NotBefore) {
			fmt.Printf("Status: Not yet valid\n")
		} else if now.After(certDetails.NotAfter) {
			fmt.Printf("Status: Expired\n")
		} else {
			fmt.Printf("Status: Valid\n")
		}
	}
	
	// Fingerprints for leaf certificate
	if len(certInfo.Fingerprints) > 0 {
		fmt.Println("\n--- Certificate Fingerprints (Leaf) ---")
		for algorithm, fingerprint := range certInfo.Fingerprints {
			fmt.Printf("%s: %s\n", algorithm, fingerprint)
		}
	}
	
	fmt.Println("=====================================\n")
}

// Helper function to get TLS version string
func getTLSVersion(version uint16) string {
	switch version {
	case tls.VersionTLS10:
		return "TLS 1.0"
	case tls.VersionTLS11:
		return "TLS 1.1"
	case tls.VersionTLS12:
		return "TLS 1.2"
	case tls.VersionTLS13:
		return "TLS 1.3"
	default:
		return fmt.Sprintf("Unknown (0x%04x)", version)
	}
}

// Helper function to get cipher suite name (simplified)
func getCipherSuite(suite uint16) string {
	// This would typically use tls.CipherSuiteName in Go 1.14+
	// or a lookup map for older versions
	return fmt.Sprintf("0x%04x", suite)
}

// Helper function to format key usage flags
func formatKeyUsage(usage x509.KeyUsage) string {
	var usages []string
	
	if usage&x509.KeyUsageDigitalSignature != 0 {
		usages = append(usages, "Digital Signature")
	}
	if usage&x509.KeyUsageKeyEncipherment != 0 {
		usages = append(usages, "Key Encipherment")
	}
	if usage&x509.KeyUsageDataEncipherment != 0 {
		usages = append(usages, "Data Encipherment")
	}
	if usage&x509.KeyUsageKeyAgreement != 0 {
		usages = append(usages, "Key Agreement")
	}
	if usage&x509.KeyUsageCertSign != 0 {
		usages = append(usages, "Certificate Sign")
	}
	if usage&x509.KeyUsageCRLSign != 0 {
		usages = append(usages, "CRL Sign")
	}
	
	if len(usages) == 0 {
		return "None"
	}
	
	return strings.Join(usages, ", ")
}
