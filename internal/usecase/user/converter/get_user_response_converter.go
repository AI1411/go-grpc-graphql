package converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/AI1411/go-grpc-graphql/grpc"
	commonEntity "github.com/AI1411/go-grpc-graphql/internal/domain/common/entity"
	"github.com/AI1411/go-grpc-graphql/internal/domain/user/entity"
	grpcTool "github.com/AI1411/go-grpc-graphql/internal/infra/grpc"
	"github.com/AI1411/go-grpc-graphql/internal/util"
)

func UserEntityToGRPC(user *entity.User) *grpc.GetUserResponse {
	status := grpcTool.ToGRPCEnum(grpc.Status_value, entity.UserStatusValue[user.Status.String()])
	prefecture := grpcTool.ToGRPCEnum(grpc.Prefecture_value, commonEntity.PrefectureValue[user.Prefecture.String()])
	bloodType := grpcTool.ToGRPCEnum(grpc.BloodType_value, commonEntity.BloodTypeValue[user.BloodType.String()])

	return &grpc.GetUserResponse{
		User: &grpc.User{
			Id:           util.NullUUIDToString(user.ID),
			Email:        user.Email,
			Username:     user.Username,
			Password:     string(user.Password),
			Status:       grpc.Status(status),
			Prefecture:   grpc.Prefecture(prefecture),
			BloodType:    grpc.BloodType(bloodType),
			Introduction: user.Introduction,
			CreatedAt:    timestamppb.New(user.CreatedAt),
			UpdatedAt:    timestamppb.New(user.UpdatedAt),
		},
	}
}
