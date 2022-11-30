package grpc_client

import (
	"alliance-widget/constant"
	"alliance-widget/entity/entity_pb/taskmonitor_pb"

	"alliance-widget/utils/log"
	"context"
	"google.golang.org/grpc"
)

var taskMonitorConn *grpc.ClientConn

func InitTaskMonitorClient(address string) {
	if conn, err := StartConnection(address); err != nil {
		log.Error().Fields(map[string]interface{}{
			"action":  "init task-monitor client",
			"address": address,
			"error":   err,
		}).Send()
	} else {
		log.Info().Fields(map[string]interface{}{
			"action":  "init task-monitor client",
			"address": address,
		}).Send()
		taskMonitorConn = conn
	}
}

func GetTaskMonitorClient() taskmonitor_pb.TaskMonitorServiceClient {
	return taskmonitor_pb.NewTaskMonitorServiceClient(taskMonitorConn)
}

func TaskMonitorTriggerRetweet(tweet *taskmonitor_pb.TaskMonitorReTweet) (*taskmonitor_pb.TaskMonitorCommonResp, error) {
	client := GetTaskMonitorClient()
	c, cancel := context.WithTimeout(context.Background(), constant.TimeOut)
	defer cancel()
	return client.TaskReTweet(c, tweet)
}

func TaskMonitorTriggerTwitterFollow(follow *taskmonitor_pb.TaskMonitorTwitterFollow) (*taskmonitor_pb.TaskMonitorCommonResp, error) {

	client := GetTaskMonitorClient()
	c, cancel := context.WithTimeout(context.Background(), constant.TimeOut)
	defer cancel()
	return client.TaskTwitterFollow(c, follow)

}
