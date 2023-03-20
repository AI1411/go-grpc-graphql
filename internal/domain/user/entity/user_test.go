package entity_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	commonEntity "github.com/AI1411/go-grpc-graphql/internal/domain/common/entity"
	"github.com/AI1411/go-grpc-graphql/internal/domain/user/entity"
)

func TestNewUser(t *testing.T) {
	type args struct {
		username     string
		email        string
		password     string
		status       entity.UserStatus
		prefecture   commonEntity.Prefecture
		introduction string
		bloodType    commonEntity.BloodType
	}
	tests := []struct {
		name string
		args args
		want *entity.User
	}{
		{
			name: "success",
			args: args{
				username:     "username",
				email:        "test@gmail.com",
				password:     "$2a$10$Ig2ubFhcRtxTswDOZ95ymOfpnhRjm4DhmTPwlp1VtC.3NoCO4y2aC",
				status:       "通常会員",
				prefecture:   "岡山県",
				introduction: "自己紹介",
				bloodType:    "A型",
			},
			want: &entity.User{
				Username:     "username",
				Email:        "test@gmail.com",
				Password:     "$2a$10$Ig2ubFhcRtxTswDOZ95ymOfpnhRjm4DhmTPwlp1VtC.3NoCO4y2aC",
				Status:       "通常会員",
				Prefecture:   "岡山県",
				Introduction: "自己紹介",
				BloodType:    "A型",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := entity.NewUser(tt.args.username, tt.args.email, tt.args.password, tt.args.status, tt.args.prefecture, tt.args.introduction, tt.args.bloodType)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("diff %s", cmp.Diff(got, tt.want))
			}
		})
	}
}
