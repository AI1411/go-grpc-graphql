package converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/AI1411/go-grpc-praphql/grpc"
	commonEntity "github.com/AI1411/go-grpc-praphql/internal/domain/common/entity"
	"github.com/AI1411/go-grpc-praphql/internal/domain/user/entity"
	grpcTool "github.com/AI1411/go-grpc-praphql/internal/infra/grpc"
	"github.com/AI1411/go-grpc-praphql/internal/util"
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
