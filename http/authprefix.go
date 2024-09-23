package http

type AuthPrefix string

const (
	AuthPrefixBasic       AuthPrefix = "Basic"            // Basic Authentication
	AuthPrefixBearer      AuthPrefix = "Bearer"           // Bearer Token Authentication (commonly used with OAuth 2.0 and JWT)
	AuthPrefixDigest      AuthPrefix = "Digest"           // Digest Authentication
	AuthPrefixHOBA        AuthPrefix = "HOBA"             // HTTP Origin-Bound Authentication
	AuthPrefixMutual      AuthPrefix = "Mutual"           // Mutual Authentication (used with TLS)
	AuthPrefixAWS4HMAC    AuthPrefix = "AWS4-HMAC-SHA256" // AWS Signature Version 4
	AuthPrefixNTLM        AuthPrefix = "NTLM"             // NT LAN Manager (Microsoft's proprietary authentication protocol)
	AuthPrefixNegotiate   AuthPrefix = "Negotiate"        // SPNEGO-based Kerberos and NTLM Authentication
	AuthPrefixOAuth       AuthPrefix = "OAuth"            // OAuth 1.0 Authentication
	AuthPrefixSCRAMSHA256 AuthPrefix = "SCRAM-SHA-256"    // Salted Challenge Response Authentication Mechanism (SCRAM)
	AuthPrefixApiKey      AuthPrefix = "ApiKey"           // Custom API key-based authentication
)

func (p AuthPrefix) String() string {
	return string(p)
}

func (p AuthPrefix) Set(value string) string {
	return p.String() + " " + value
}
