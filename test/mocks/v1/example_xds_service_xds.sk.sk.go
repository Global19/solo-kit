// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"context"
	"errors"
	"fmt"

	discovery "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	"github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/solo-io/solo-kit/pkg/api/v1/control-plane/cache"
	"github.com/solo-io/solo-kit/pkg/api/v1/control-plane/client"
	"github.com/solo-io/solo-kit/pkg/api/v1/control-plane/server"
)

// Type Definitions:

const ExampleXDSResourceType = cache.TypePrefix + "/testing.solo.io.ExampleXDSResource"

/* Defined a resource - to be used by snapshot */
type ExampleXDSResourceXdsResourceWrapper struct {
	// TODO(yuval-k): This is public for mitchellh hashstructure to work properly. consider better alternatives.
	Resource *ExampleXDSResource
}

// Make sure the Resource interface is implemented
var _ cache.Resource = &ExampleXDSResourceXdsResourceWrapper{}

func NewExampleXDSResourceXdsResourceWrapper(resourceProto *ExampleXDSResource) *ExampleXDSResourceXdsResourceWrapper {
	return &ExampleXDSResourceXdsResourceWrapper{
		Resource: resourceProto,
	}
}

func (e *ExampleXDSResourceXdsResourceWrapper) Self() cache.XdsResourceReference {
	return cache.XdsResourceReference{Name: e.Resource.FavoriteMeme, Type: ExampleXDSResourceType}
}

func (e *ExampleXDSResourceXdsResourceWrapper) ResourceProto() cache.ResourceProto {
	return e.Resource
}
func (e *ExampleXDSResourceXdsResourceWrapper) References() []cache.XdsResourceReference {
	return nil
}

// Define a type record. This is used by the generic client library.
var ExampleXDSResourceTypeRecord = client.NewTypeRecord(
	ExampleXDSResourceType,

	// Return an empty message, that can be used to deserialize bytes into it.
	func() cache.ResourceProto { return &ExampleXDSResource{} },

	// Covert the message to a resource suitable for use for protobuf's Any.
	func(r cache.ResourceProto) cache.Resource {
		return &ExampleXDSResourceXdsResourceWrapper{Resource: r.(*ExampleXDSResource)}
	},
)

// Server Implementation:

// Wrap the generic server and implement the type sepcific methods:
type exampleXDSServiceServer struct {
	server.Server
}

func NewExampleXDSServiceServer(genericServer server.Server) ExampleXDSServiceServer {
	return &exampleXDSServiceServer{Server: genericServer}
}

func (s *exampleXDSServiceServer) StreamExampleXDSResource(stream ExampleXDSService_StreamExampleXDSResourceServer) error {
	return s.Server.Stream(stream, ExampleXDSResourceType)
}

func (s *exampleXDSServiceServer) FetchExampleXDSResource(ctx context.Context, req *discovery.DiscoveryRequest) (*discovery.DiscoveryResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.Unavailable, "empty request")
	}
	req.TypeUrl = ExampleXDSResourceType
	return s.Server.Fetch(ctx, req)
}

func (s *exampleXDSServiceServer) IncrementalExampleXDSResource(_ ExampleXDSService_IncrementalExampleXDSResourceServer) error {
	return errors.New("not implemented")
}

// Client Implementation: Generate a strongly typed client over the generic client

// The apply functions receives resources and returns an error if they were applied correctly.
// In theory the configuration can become valid in the future (i.e. eventually consistent), but I don't think we need to worry about that now
// As our current use cases only have one configuration resource, so no interactions are expected.
type ApplyExampleXDSResource func(version string, resources []*ExampleXDSResource) error

// Convert the strongly typed apply to a generic apply.
func applyExampleXDSResource(typedApply ApplyExampleXDSResource) func(cache.Resources) error {
	return func(resources cache.Resources) error {

		var configs []*ExampleXDSResource
		for _, r := range resources.Items {
			if proto, ok := r.ResourceProto().(*ExampleXDSResource); !ok {
				return fmt.Errorf("resource %s of type %s incorrect", r.Self().Name, r.Self().Type)
			} else {
				configs = append(configs, proto)
			}
		}

		return typedApply(resources.Version, configs)
	}
}

func NewExampleXDSResourceClient(nodeinfo *core.Node, typedApply ApplyExampleXDSResource) client.Client {
	return client.NewClient(nodeinfo, ExampleXDSResourceTypeRecord, applyExampleXDSResource(typedApply))
}
