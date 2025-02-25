package etc

import (
	"context"
	"time"

	"github.com/golanguzb71/livesphere-api-gateway/config"
	"google.golang.org/grpc/metadata"
)

func NewTimoutContext(ctx context.Context) (context.Context, context.CancelFunc) {
	md := metadata.Pairs()
	for _, key := range []string{} {
		if ctx.Value(key) != nil {
			val, ok := ctx.Value(key).(string)
			if ok {
				md.Set(key, val)
			}
		}
	}
	ctx = metadata.NewOutgoingContext(ctx, md)
	res, cancel := context.WithTimeout(ctx, time.Second*config.GrpcContextTimeout)
	return res, cancel
}
