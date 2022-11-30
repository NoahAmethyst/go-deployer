package grpc_client

import (
	"alliance-widget/entity/entity_pb/metadata_pb"
	"alliance-widget/utils/log"
	"google.golang.org/grpc"
)

var metadataConn *grpc.ClientConn

func InitMetadataClient(address string) {
	if conn, err := StartConnection(address); err != nil {
		log.Error().Fields(map[string]interface{}{
			"action":  "init metadata client",
			"address": address,
			"error":   err,
		}).Send()
	} else {
		log.Info().Fields(map[string]interface{}{
			"action":  "init metadata client",
			"address": address,
		}).Send()
		metadataConn = conn
	}
}

func GetMetadataClient() metadata_pb.MetadataServiceClient {
	return metadata_pb.NewMetadataServiceClient(metadataConn)
}
