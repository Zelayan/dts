package collector

import (
	"context"
	"github.com/Zelayan/dts/cmd/colletcor/options"
	"github.com/Zelayan/dts/pkg/collector"
	pb "github.com/Zelayan/dts/proto-gen/v1/dts"
	"k8s.io/klog/v2"
)

func NewCollectorServer(ops *options.Options) pb.CollectorServiceServer {
	return &collectorServer{
		cc: ops.Collector,
	}
}

type collectorServer struct {
	pb.UnimplementedCollectorServiceServer
	cc collector.CollectorInterface
}

func (c *collectorServer) PostSpans(ctx context.Context, request *pb.PostSpansRequest) (*pb.PostSpansResponse, error) {
	spans := request.Batch.GetSpans()
	for _, x := range spans {
		err := c.cc.Span().Create(ctx, x)
		klog.Info(x)
		if err != nil {
			return &pb.PostSpansResponse{}, err
		}
	}
	return &pb.PostSpansResponse{}, nil
}
