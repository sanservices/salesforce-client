package salesforceclient

import (
	"crypto/tls"
	"fmt"

	"github.com/sanservices/salesforce-client/grpc/salesforce-proto/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type SettingsSfmc struct {
	Server    string `yaml:"Server"`
	Authority string `yaml:"Authority"`
}

type SfmcAgentClient struct {
	Client pb.SalesforceServiceClient
}

type GRPCConnectionManager struct {
	Conn *grpc.ClientConn
}

// New returns service instance

func newgRPCConn(target string, authority string, useTLS bool) (*GRPCConnectionManager, error) {
	var opts []grpc.DialOption
	if useTLS {
		creds := credentials.NewTLS(&tls.Config{
			InsecureSkipVerify: true, // Skips certificate verification
			ServerName:         authority,
		})
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	conn, err := grpc.NewClient(target, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to gRPC server at %s: %w", target, err)
	}
	return &GRPCConnectionManager{Conn: conn}, nil
}

func (m *GRPCConnectionManager) Close() error {
	return m.Conn.Close()
}

func SfmcNewClient(config *SettingsSfmc) (*SfmcAgentClient, error) {
	conn, err := newgRPCConn(config.Server, config.Authority, true)
	if err != nil {
		return nil, err
	}
	AgentClient := pb.NewSalesforceServiceClient(conn.Conn)
	return &SfmcAgentClient{Client: AgentClient}, nil
}
