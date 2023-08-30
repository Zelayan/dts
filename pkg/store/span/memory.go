package span

import (
	"context"
	"errors"
	pb "github.com/Zelayan/dts/proto-gen/v1/dts"
)

type MemorySpan struct {
	spans []*pb.Span
}

func (s *MemorySpan) ListSpan(ctx context.Context, serviceName string) []*pb.Span {
	//TODO implement me
	panic("implement me")
}

func (s *MemorySpan) WriteSpan(ctx context.Context, span *pb.Span) error {
	if span == nil {
		return errors.New("span must not be nil")
	}
	s.spans = append(s.spans, span)
	return nil
}

func NewMemoryStorage() Storage {
	return &MemorySpan{}
}
