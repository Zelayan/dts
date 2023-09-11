package span

import (
	"context"
	"github.com/Zelayan/dts/cmd/colletcor/config"
	"github.com/Zelayan/dts/pkg/tenancy"
	pb "github.com/Zelayan/dts/proto-gen/v1/dts"
	"sync"
)

type MemoryStore struct {
	sync.RWMutex
	config    config.Config
	preTenant map[string]*Tenant
}

type Tenant struct {
	sync.RWMutex
	// traceId
	ids    []string
	traces map[string]*pb.Trace
	spans  map[string][]*pb.Span
	// 服务列表
	services map[string]struct{}
}

func (s *MemoryStore) ListServices(ctx context.Context) ([]string, error) {
	tenant := s.getTenant(tenancy.GetTenancy(ctx))
	s.RLock()
	defer s.RUnlock()
	var retMe []string
	for key := range tenant.services {
		retMe = append(retMe, key)
	}
	return retMe, nil
}

func (s *MemoryStore) ListSpan(ctx context.Context, serviceName string) []*pb.Span {
	//TODO implement me
	panic("implement me")
}

func (s *MemoryStore) WriteSpan(ctx context.Context, span *pb.Span) error {
	tenant := s.getTenant(tenancy.GetTenancy(ctx))
	tenant.Lock()
	defer tenant.Unlock()
	if _, ok := tenant.traces[span.TraceId]; !ok {
		// TODO:淘汰
		// 如果没有traces 则新建
		tenant.traces[span.TraceId] = &pb.Trace{}
		// 保存同一个traces的所有span
		tenant.spans[span.TraceId] = make([]*pb.Span, 0)
	}
	tenant.spans[span.SpanId] = append(tenant.spans[span.SpanId], span)
	tenant.services[span.ServiceName] = struct{}{}
	tenant.ids = append(tenant.ids, span.TraceId)
	return nil
}

func NewMemoryStorage() Storage {
	defaultOptions := config.CollectorOptions{MaxTraces: 0}
	return WithConfiguration(config.Config{Default: config.DefaultOptions{Collector: defaultOptions}})
}

func WithConfiguration(config config.Config) Storage {
	return &MemoryStore{
		config:    config,
		preTenant: make(map[string]*Tenant),
	}
}

func (s *MemoryStore) getTenant(tenantID string) *Tenant {
	s.RLock()
	tenant, ok := s.preTenant[tenantID]
	s.RUnlock()
	if !ok {
		s.Lock()
		defer s.Unlock()
		tenant, ok = s.preTenant[tenantID]
		if !ok {
			tenant = newTenant(s.config)
			s.preTenant[tenantID] = tenant
		}
	}
	return tenant

}

func newTenant(cfg config.Config) *Tenant {
	return &Tenant{
		ids:      make([]string, 0, cfg.Default.Collector.MaxTraces),
		traces:   make(map[string]*pb.Trace),
		services: make(map[string]struct{}),
		spans:    make(map[string][]*pb.Span),
	}
}
