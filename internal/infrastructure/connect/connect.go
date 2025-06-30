package connect

import "google.golang.org/grpc"

type ServiceConnector interface {
	Connect() (*grpc.ClientConn, error)
}
