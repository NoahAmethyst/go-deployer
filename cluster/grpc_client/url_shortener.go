package grpc_client

import (
	"alliance-widget/constant"
	urlshortener_pb "alliance-widget/entity/entity_pb/url_shortener_pb"
	"alliance-widget/utils/log"
	"context"
	"google.golang.org/grpc"
)

var urlShortenerConn *grpc.ClientConn

func InitUrlShortenerClient(address string) {
	if conn, err := StartConnection(address); err != nil {
		log.Error().Fields(map[string]interface{}{
			"action":  "init urlShortener client",
			"address": address,
			"error":   err,
		}).Send()
	} else {
		log.Info().Fields(map[string]interface{}{
			"action":  "init urlShortener client",
			"address": address,
		}).Send()
		urlShortenerConn = conn
	}
}

func GetUrlShortenerClient() urlshortener_pb.UrlShortenerServiceClient {
	return urlshortener_pb.NewUrlShortenerServiceClient(urlShortenerConn)
}

func UrlShortenerShortenUrl(url string) (*urlshortener_pb.ShortenUrlResp, error) {
	client := GetUrlShortenerClient()
	c, cancel := context.WithTimeout(context.Background(), constant.TimeOut)
	defer cancel()
	return client.ShortenUrl(c, &urlshortener_pb.ShortenUrlReq{Url: url})
}
