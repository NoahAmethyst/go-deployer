package grpc_client

import (
	"alliance-widget/constant"
	"alliance-widget/entity/entity_pb/alliance_bot_pb"
	"alliance-widget/utils/log"
	"context"
	"google.golang.org/grpc"
)

var botConn *grpc.ClientConn

func InitBotClient(address string) {
	if conn, err := StartConnection(address); err != nil {
		log.Error().Fields(map[string]interface{}{
			"action":  "init bot client",
			"address": address,
			"error":   err,
		}).Send()
	} else {
		log.Info().Fields(map[string]interface{}{
			"action":  "init bot client",
			"address": address,
		}).Send()
		botConn = conn
	}
}

func GetBotClient() alliance_bot_pb.BotServiceClient {
	return alliance_bot_pb.NewBotServiceClient(botConn)
}

func BotCreateQuiz(quiz *alliance_bot_pb.CreateQuizReq) (*alliance_bot_pb.CreateQuizResp, error) {
	initBotName()
	botClient := GetBotClient()
	c, cancel := context.WithTimeout(context.Background(), constant.TimeOut)
	defer cancel()
	resp, err := botClient.CreateQuiz(c, quiz)
	return resp, err
}

func GetQuizWinner(req *alliance_bot_pb.GetQuizWinnerReq) (*alliance_bot_pb.GetQuizWinnerResp, error) {
	initBotName()
	botClient := GetBotClient()
	c, cancel := context.WithTimeout(context.Background(), constant.TimeOut)
	defer cancel()
	return botClient.GetQuizWinner(c, req)
}

func BotSendMsg(req *alliance_bot_pb.SendMsgReq) (*alliance_bot_pb.SendMsgResp, error) {
	initBotName()
	botClient := GetBotClient()
	c, cancel := context.WithTimeout(context.Background(), constant.TimeOut)
	defer cancel()
	return botClient.SendMsg(c, req)
}

func BotSendGame(req *alliance_bot_pb.SendGameReq) (*alliance_bot_pb.CommonResp, error) {
	botClient := GetBotClient()
	c, cancel := context.WithTimeout(context.Background(), constant.TimeOut)
	defer cancel()
	return botClient.SendGame(c, req)
}

func BotDeleteMsg(req *alliance_bot_pb.DeleteMsgReq) (*alliance_bot_pb.CommonResp, error) {
	botClient := GetBotClient()
	c, cancel := context.WithTimeout(context.Background(), constant.TimeOut)
	defer cancel()

	return botClient.DelMsg(c, req)
}

func BotGetUsers(appId string, appType int32, groupId string, userIds []int64) (*alliance_bot_pb.GetUsersResp, error) {
	initBotName()
	botClient := GetBotClient()
	c, cancel := context.WithTimeout(context.Background(), constant.TimeOut)
	defer cancel()
	return botClient.GetUsers(c, &alliance_bot_pb.GetUsersReq{
		GroupId: groupId,
		UserId:  userIds,
		AppId:   appId,
		AppType: appType,
	})
}

func GenerateInviteLink(appId string, appType int32, groupId string, value string) (*alliance_bot_pb.InviteLinkResp, error) {
	initBotName()
	botClient := GetBotClient()
	c, cancel := context.WithTimeout(context.Background(), constant.TimeOut)
	defer cancel()
	req := &alliance_bot_pb.InviteLinkReq{
		GroupId: groupId,
		AppId:   appId,
		Value:   value,
		AppType: appType,
	}
	log.Info().Fields(map[string]interface{}{
		"action": "create invite link",
		"req":    req,
	}).Send()
	return botClient.GenerateInviteLink(c, req)
}

func initBotName() {
	//if len(BotName) > 0 {
	//	return
	//}
	//botClient := GetBotClient()
	//c, cancel := context.WithTimeout(context.Background(), constant.TimeOut)
	//defer cancel()
	//resp, err := botClient.GetBotInfo(c, &alliance_bot_pb.GetBotInfoReq{AppId: config.GetAppId()})
	//if err != nil {
	//	log.Error().Msgf("get bot info error:%s", err)
	//	return
	//}
	//if resp.Error != nil && resp.Error.Code != 0 {
	//	log.Error().Msgf("get bot info error:%s", resp.Error.Detail)
	//	return
	//}
	//BotName = resp.Data.BotName
	return
}
