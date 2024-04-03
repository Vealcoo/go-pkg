package ginhttphelper

import (
	"github.com/Vealcoo/go-pkg/httphelper"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// SendErrorMessage SendErrorMessage
func SendErrorMessage(ctx *gin.Context, c codes.Code, mesg string, args ...interface{}) {
	var arg interface{}
	if len(args) > 0 {
		arg = args[0]
	}

	m := httphelper.CreateErrorMessage(c, mesg, arg)

	ctx.JSON(m.Code, m)
}

func SendGrpcError(ctx *gin.Context, err error) {
	if st, ok := status.FromError(err); !ok {
		log.Error().Stack().Err(errors.WithStack(err)).Msg("unexpected error")
		SendErrorMessage(ctx, codes.Internal, "")
	} else {
		if st.Code() != codes.Internal && st.Code() != codes.Unavailable {
			SendErrorMessage(ctx, st.Code(), st.Message())
		} else {
			SendErrorMessage(ctx, st.Code(), "")
		}
	}
}
