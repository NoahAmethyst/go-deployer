package deployer_service

import (
	"context"
	"go-deployer/entity/entity_pb/deployer_pb"
)

type Server struct{}

func (s Server) DeployCat(ctx context.Context, req *deployer_pb.DeployCatReq) (resp *deployer_pb.DeployResp, err error) {

	return
}
