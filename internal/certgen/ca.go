package certgen

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"time"
)

type CertificateAuthority struct {
	privateKey    *rsa.PrivateKey
	privateKeyPem *bytes.Buffer

	certBytes []byte
	certPem   *bytes.Buffer
}

func NewCaWithSubject(subject pkix.Name) (*CertificateAuthority, error) {
	caRsaPrivKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return nil, fmt.Errorf("failed to generate rsa private key for ca: %w", err)
	}

	crtTemplate := &x509.Certificate{
		SerialNumber:          big.NewInt(2019),
		Subject:               subject,
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0),
		IsCA:                  true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
	}

	caBytes, err := x509.CreateCertificate(rand.Reader, crtTemplate, crtTemplate, &caRsaPrivKey.PublicKey, caRsaPrivKey)
	if err != nil {
		return nil, fmt.Errorf("failed to create x509 certificate for ca: %w", err)
	}

	caPEM := new(bytes.Buffer)
	pem.Encode(caPEM, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: caBytes,
	})

	caPrivKeyPEM := new(bytes.Buffer)
	pem.Encode(caPrivKeyPEM, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(caRsaPrivKey),
	})

	return &CertificateAuthority{
		privateKey:    caRsaPrivKey,
		privateKeyPem: caPrivKeyPEM,
		certBytes:     caBytes,
		certPem:       caPEM,
	}, nil
}

// func NewCA(subject pkix.Name) (*x509.Certificate, error) {
// 	caPrivKey, err := rsa.GenerateKey(rand.Reader, 4096)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to generate rsa private key for ca: %w", err)
// 	}

// 	crtTemplate := &x509.Certificate{
// 		SerialNumber:          big.NewInt(2019),
// 		Subject:               subject,
// 		NotBefore:             time.Now(),
// 		NotAfter:              time.Now().AddDate(10, 0, 0),
// 		IsCA:                  true,
// 		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
// 		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
// 		BasicConstraintsValid: true,
// 	}

// 	caBytes, err := x509.CreateCertificate(rand.Reader, crtTemplate, crtTemplate, &caPrivKey.PublicKey, caPrivKey)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to create x509 certificate for ca: %w", err)
// 	}

// 	return caBytes, nil
// }
