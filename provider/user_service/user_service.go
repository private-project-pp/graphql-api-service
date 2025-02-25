package user_service

import (
	rpcUserService "github.com/private-project-pp/pos-grpc-contract/model/user_service"
)

type UserServiceRpc interface {
}

type userService struct {
	rpcClient rpcUserService.UserServiceClient
}

func SetupRpcClientUserService(rpcClient rpcUserService.UserServiceClient) UserServiceRpc {
	return &userService{
		rpcClient: rpcClient,
	}
}
