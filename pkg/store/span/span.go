package span

import (
	"context"
	pb "github.com/Zelayan/dts/proto-gen/v1/dts"
)

type Storage interface {
	SpanWriter
	SpanReader
}

type SpanWriter interface {
	WriteSpan(ctx context.Context, span *pb.Span) error
}

type SpanReader interface {
	ListSpan(ctx context.Context, serviceName string) []*pb.Span
	ListServices(ctx context.Context) ([]string, error)
	ListSpanByTraceId(ctx context.Context, traceId string) []*pb.Span
}
