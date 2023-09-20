package span

import (
	"context"
	"github.com/Zelayan/dts/cmd/colletcor/config"
	pb "github.com/Zelayan/dts/proto-gen/v1/dts"
	"github.com/olivere/elastic/v7"
)

const (
	trace = "trace"
)

type EsStore struct {
	config config.Config
	client *elastic.Client
}

func (e *EsStore) WriteSpan(ctx context.Context, span *pb.Span) error {
	_, err := e.client.Index().Index(trace).BodyJson(span).Do(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (e *EsStore) ListSpan(ctx context.Context, serviceName string) []*pb.Span {
	//TODO implement me
	panic("implement me")
}

func (e *EsStore) ListServices(ctx context.Context) ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (e *EsStore) ListSpanByTraceId(ctx context.Context, traceId string) []*pb.Span {
	//TODO implement me
	panic("implement me")
}

func NewEsStore() (Storage, error) {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}
	return &EsStore{
		config: config.Config{},
		client: client,
	}, nil
}
