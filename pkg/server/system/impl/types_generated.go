// Code generated by csi-proxy-api-gen. DO NOT EDIT.

package impl

import (
	"context"

	"github.com/kubernetes-csi/csi-proxy/client/apiversion"
	"google.golang.org/grpc"
)

type VersionedAPI interface {
	Register(grpcServer *grpc.Server)
}

// All the functions this group's server needs to define.
type ServerInterface interface {
	GetBIOSSerialNumber(context.Context, *GetBIOSSerialNumberRequest, apiversion.Version) (*GetBIOSSerialNumberResponse, error)
	GetService(context.Context, *GetServiceRequest, apiversion.Version) (*GetServiceResponse, error)
	StartService(context.Context, *StartServiceRequest, apiversion.Version) (*StartServiceResponse, error)
	StopService(context.Context, *StopServiceRequest, apiversion.Version) (*StopServiceResponse, error)
}