package interceptor

import (
	"errors"
	"fmt"

	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ServerRecovery(log zerolog.Logger) grpc.UnaryServerInterceptor {
	return grpc_recovery.UnaryServerInterceptor(
		grpc_recovery.WithRecoveryHandler(
			func(p interface{}) (err error) {
				msg := fmt.Sprintf("panic: %v", p)

				if pe, ok := p.(error); ok {
					log.Error().Err(pe).Msg("server panic")
				} else {
					log.Error().Err(errors.New("unformat panic error")).Msg("server panic")
				}

				return status.New(codes.Unavailable, msg).Err()
			},
		),
	)
}
