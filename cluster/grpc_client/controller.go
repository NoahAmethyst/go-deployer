package grpc_client

import (
	"alliance-widget/constant"
	"alliance-widget/entity/entity_pb/controller_pb"
	"alliance-widget/utils/log"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"strconv"
)

var controllerConn *grpc.ClientConn

func InitControllerClient(address string) {
	if conn, err := StartConnection(address); err != nil {
		log.Error().Fields(map[string]interface{}{
			"action":  "init controller client",
			"address": address,
			"error":   err,
		}).Send()
	} else {
		log.Info().Fields(map[string]interface{}{
			"action":  "init controller client",
			"address": address,
		}).Send()
		controllerConn = conn
	}
}

func GetControllerClient() controller_pb.ControllerServiceClient {
	return controller_pb.NewControllerServiceClient(controllerConn)
}

func ControllerBatchGetUserByOpenIds(openIds []string, plat int32) ([]*controller_pb.User, error) {

	controllerClient := GetControllerClient()
	c, cancel := context.WithTimeout(context.Background(), constant.TimeOut)
	defer cancel()
	batchGetResp, err := controllerClient.BatchGetUser(c, &controller_pb.BatchGetUserReq{
		OpenIds:  openIds,
		OpenType: plat,
	})
	if err != nil {
		return nil, err
	} else if batchGetResp.CommonResponse.Code != constant.Success {
		return nil, fmt.Errorf("request wallet error %s", batchGetResp.GetCommonResponse().Inner)
	}
	return batchGetResp.Data, nil
}

func ControllerBatchGetUserByIds(ids []string) ([]*controller_pb.User, error) {

	controllerClient := GetControllerClient()
	c, cancel := context.WithTimeout(context.Background(), constant.TimeOut)
	defer cancel()
	batchGetResp, err := controllerClient.BatchGetUser(c, &controller_pb.BatchGetUserReq{
		UserNoList:    ids,
		IgnoreAddress: true,
	})
	if err != nil {
		return nil, err
	} else if batchGetResp.CommonResponse.Code != constant.Success {
		return nil, fmt.Errorf("request wallet error %s", batchGetResp.GetCommonResponse().Inner)
	}
	return batchGetResp.Data, nil
}

func ControllerGetUserByOpenId(openId int64, appType int32) (*controller_pb.GetUserResp, error) {
	controllerClient := GetControllerClient()
	c, cancel := context.WithTimeout(context.Background(), constant.TimeOut)
	requester := &controller_pb.Requester{
		RequesterOpenId:   strconv.FormatInt(openId, 10),
		RequesterOpenType: appType,
	}
	defer cancel()

	b, err := proto.Marshal(requester)
	if err != nil {
		log.Error().Msgf("proto mashal requester error:%s", err)
		return nil, nil
	}
	requestStr := base64.StdEncoding.EncodeToString(b)
	md := metadata.Pairs("requester", requestStr)
	c = metadata.NewOutgoingContext(c, md)
	return controllerClient.GetUser(c, &controller_pb.GetUserReq{
		OpenId:   strconv.FormatInt(openId, 10),
		OpenType: appType,
	})
}

func ControllerGetUserById(userId string, ignoreAddress bool) (*controller_pb.GetUserResp, error) {
	controllerClient := GetControllerClient()
	c, cancel := context.WithTimeout(context.Background(), constant.TimeOut)
	requester := &controller_pb.Requester{
		RequesterUserNo: userId,
	}
	defer cancel()

	b, err := proto.Marshal(requester)
	if err != nil {
		log.Error().Msgf("proto mashal requester error:%s", err)
		return nil, nil
	}
	requestStr := base64.StdEncoding.EncodeToString(b)
	md := metadata.Pairs("requester", requestStr)
	c = metadata.NewOutgoingContext(c, md)
	return controllerClient.GetUser(c, &controller_pb.GetUserReq{
		UserNo:        userId,
		IgnoreAddress: ignoreAddress,
	})
}

func ControllerAddAsset(addresses []string, token string, chainType int64, chainId int64) (*controller_pb.GetAssetResp, error) {
	controllerClient := GetControllerClient()
	c, cancel := context.WithTimeout(context.Background(), constant.TimeOut)
	requester := &controller_pb.Requester{}
	defer cancel()

	b, err := proto.Marshal(requester)
	if err != nil {
		log.Error().Msgf("proto mashal requester error:%s", err)
		return nil, nil
	}
	requestStr := base64.StdEncoding.EncodeToString(b)
	md := metadata.Pairs("requester", requestStr)
	c = metadata.NewOutgoingContext(c, md)

	return controllerClient.AddAsset(c, &controller_pb.AddAssetReq{
		ChainType:       uint32(chainType),
		ChainId:         uint64(chainId),
		AddressList:     addresses,
		ContractAddress: token,
		TokenType:       2,
	})
}

func ControllerGetTx(txId string) (*controller_pb.GetTxResp, error) {
	controllerClient := GetControllerClient()
	c, cancel := context.WithTimeout(context.Background(), constant.TimeOut)
	requester := &controller_pb.Requester{}
	defer cancel()
	b, err := proto.Marshal(requester)
	if err != nil {
		log.Error().Msgf("proto mashal requester error:%s", err)
		return nil, nil
	}
	requestStr := base64.StdEncoding.EncodeToString(b)
	md := metadata.Pairs("requester", requestStr)
	c = metadata.NewOutgoingContext(c, md)
	return controllerClient.GetTx(c, &controller_pb.GetTxReq{
		IsWait: false,
		TxId:   txId,
	})
}

func ControllerInitTemporaryToken(req *controller_pb.InitTemporaryTokenReq) (*controller_pb.InitAccessTokenResp, error) {
	controllerClient := GetControllerClient()
	c, cancel := context.WithTimeout(context.Background(), constant.TimeOut)
	defer cancel()

	return controllerClient.InitTemporaryToken(c, req)
}
