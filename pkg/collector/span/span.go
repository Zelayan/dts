package span

import (
	"context"
	"github.com/Zelayan/dts/cmd/collector/config"
	"github.com/Zelayan/dts/pkg/store"
	"github.com/Zelayan/dts/proto-gen/v1/dts"
)

type SpanGetter interface {
	Span() Interface
}

type Interface interface {
	Create(ctx context.Context, sp *dts.Span) error
	BetchCreate(ctx context.Context, sps []*dts.Span) error
	ListSpan(ctx context.Context, service string) ([]*dts.Span, error)
	ListService(ctx context.Context) ([]string, error)
	GetTraces(ctx context.Context, id string) ([]*dts.Span, error)
}

type span struct {
	cc      config.Config
	factory store.ShareDaoFactory
}

func (s *span) GetTraces(ctx context.Context, traceId string) ([]*dts.Span, error) {
	spans := s.factory.SpanStore().ListSpanByTraceId(ctx, traceId)
	return spans, nil
}

func (s *span) ListService(ctx context.Context) ([]string, error) {
	services, err := s.factory.SpanStore().ListServices(ctx)
	return services, err
}

func (s *span) ListSpan(ctx context.Context, service string) ([]*dts.Span, error) {
	//TODO implement me
	panic("implement me")
}

func (s *span) Create(ctx context.Context, sp *dts.Span) error {
	err := s.factory.SpanStore().WriteSpan(ctx, sp)
	if err != nil {
		return err
	}
	return nil
}

func (s *span) BetchCreate(ctx context.Context, sps []*dts.Span) error {
	//TODO implement me
	panic("implement me")
}

func NewSpan(cfg config.Config, f store.ShareDaoFactory) Interface {
	return &span{
		cc:      cfg,
		factory: f,
	}
}
