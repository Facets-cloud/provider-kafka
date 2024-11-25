package kafka

import "strings"

// Config is a Kafka client configuration
type Config struct {
	Brokers []string `json:"brokers"`
	SASL    *SASL    `json:"sasl,omitempty"`
	TLS     *TLS     `json:"tls,omitempty"`
}

// SASL is an sasl option
type SASL struct {
	Mechanism string `json:"mechanism"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

// TLS is an option for enabling encryption in transit
type TLS struct {
	ClientCertificateSecretRef *ClientCertificateSecretRef `json:"clientCertificateSecretRef,omitempty"`
	InsecureSkipVerify         bool                        `json:"insecureSkipVerify"`
}

// ClientCertificateSecretRef is a TLS option for enable mTLS
type ClientCertificateSecretRef struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	KeyField  string `json:"keyField,omitempty"`
	CertField string `json:"certField,omitempty"`
}

// ParseBrokerURL parses a Confluent-style broker URL (e.g., sasl_ssl://host:port)
// and configures the appropriate SASL and TLS settings
func (c *Config) ParseBrokerURL(url string) {
	if strings.HasPrefix(url, "sasl_ssl://") {
		// Strip the protocol prefix
		broker := strings.TrimPrefix(url, "sasl_ssl://")
		
		// Update brokers list with the cleaned URL
		c.Brokers = []string{broker}
		
		// Enable TLS since it's a SASL_SSL URL
		if c.TLS == nil {
			c.TLS = &TLS{
				InsecureSkipVerify: false,
			}
		}
		
		// Ensure SASL is initialized
		if c.SASL == nil {
			c.SASL = &SASL{
				Mechanism: "PLAIN", // Default to PLAIN mechanism
			}
		}
	}
}
