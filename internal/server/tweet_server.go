package server

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/AI1411/go-grpc-graphql/grpc"
	"github.com/AI1411/go-grpc-graphql/internal/infra/db"
	tweetRepo "github.com/AI1411/go-grpc-graphql/internal/infra/repository/tweet"
	userRepo "github.com/AI1411/go-grpc-graphql/internal/infra/repository/user"
	"github.com/AI1411/go-grpc-graphql/internal/server/form"
	tweetForm "github.com/AI1411/go-grpc-graphql/internal/server/form/tweet"
	"github.com/AI1411/go-grpc-graphql/internal/usecase/tweet"
)

type TweetServer struct {
	grpc.UnimplementedTweetServiceServer
	dbClient  *db.Client
	zapLogger *zap.Logger
	userRepo  userRepo.UserRepository
	tweetRepo tweetRepo.TweetRepository
}

func NewTweetServer(
	dbClient *db.Client,
	zapLogger *zap.Logger,
	userRepo userRepo.UserRepository,
	tweetRepo tweetRepo.TweetRepository,
) *TweetServer {
	return &TweetServer{
		dbClient:  dbClient,
		zapLogger: zapLogger,
		userRepo:  userRepo,
		tweetRepo: tweetRepo,
	}
}

func (s *TweetServer) ListTweet(ctx context.Context, _ *emptypb.Empty) (*grpc.ListTweetResponse, error) {
	usecase := tweet.NewListTweetUsecaseImpl(s.userRepo, s.tweetRepo)
	res, err := usecase.Exec(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *TweetServer) CreateTweet(ctx context.Context, in *grpc.CreateTweetRequest) (*grpc.CreateTweetResponse, error) {
	validator := form.NewFormValidator(tweetForm.NewCreateTweetForm(in))
	if err := validator.Validate(); err != nil {
		return nil, err
	}

	usecase := tweet.NewCreateTweetUsecaseImpl(s.tweetRepo)

	return usecase.Exec(ctx, in)
}
