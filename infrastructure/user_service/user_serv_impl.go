package user_service

import (
	userService "github.com/private-project-pp/pos-grpc-contract/model/user_service"
	"google.golang.org/grpc"
)

type userServ struct {
	infra userService.UserServiceClient
}

func SetupUserService(clientConn *grpc.ClientConn) UserService {
	return &userServ{
		infra: userService.NewUserServiceClient(clientConn),
	}
}
