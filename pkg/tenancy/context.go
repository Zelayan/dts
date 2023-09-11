package tenancy

import "context"

type tenancyKeyType string

const (
	tenancyKey = tenancyKeyType("tenant")
)

func WithTenancy(ctx context.Context, tenancy string) context.Context {
	return context.WithValue(ctx, tenancyKey, tenancy)
}

func GetTenancy(ctx context.Context) string {
	value := ctx.Value(tenancyKey)
	if value == nil {
		return ""
	}
	if s, ok := value.(string); ok {
		return s
	}
	return ""
}
