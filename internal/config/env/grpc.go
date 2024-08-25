package env

import (
	"errors"
	"net"
	"os"

	"github.com/Chuiko-GIT/auth/internal/config"
)

var _ config.GRPCConfig = (*GrpcConfig)(nil)

const (
	grpcHostEnvName = "GRPC_HOST"
	grpcPortEnvName = "GRPC_PORT"
)

type GrpcConfig struct {
	host string
	port string
}

func NewGRPCConfig() (*GrpcConfig, error) {
	host := os.Getenv(grpcHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("grpc host not found")
	}

	port := os.Getenv(grpcPortEnvName)
	if len(host) == 0 {
		return nil, errors.New("grpc port not found")
	}

	return &GrpcConfig{
		host: host,
		port: port,
	}, nil
}

func (g GrpcConfig) Address() string {
	return net.JoinHostPort(g.host, g.port)
}
