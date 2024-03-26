package grpchelper

import (
	log "github.com/rs/zerolog/log"
	structpb "google.golang.org/protobuf/types/known/structpb"
)

func ConverMapStringInterfaceToStructpb(d map[string]interface{}) *structpb.Struct {
	res, err := structpb.NewStruct(d)
	if err != nil {
		log.Error().Stack().Str("func", "ConverMapStringInterfaceToStructpb").Err(err).Send()
	}
	return res
}
