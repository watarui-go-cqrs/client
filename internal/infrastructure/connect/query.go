package connect

import (
	"crypto/tls"
	"crypto/x509"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type queryConnector struct{}

func NewQueryConnector() ServiceConnector {
	return &queryConnector{}
}

func (c *queryConnector) Connect() (*grpc.ClientConn, error) {
	certFile := "certs/queryservice.pem"

	pemData, err := os.ReadFile(certFile)
	if err != nil {
		return nil, err
	}
	cp := x509.NewCertPool()
	if !cp.AppendCertsFromPEM(pemData) {
		return nil, err
	}
	config := &tls.Config{
		RootCAs: cp,
	}
	creds := credentials.NewTLS(config)
	server := "query-service:8083"
	if conn, err := grpc.Dial(
		server,
		grpc.WithTransportCredentials(creds),
	); err != nil {
		return nil, err
	} else {
		return conn, nil
	}
}
