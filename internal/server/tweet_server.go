package server

import (
	"context"

	"github.com/bufbuild/connect-go"
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

func (s *TweetServer) ListTweet(ctx context.Context, _ *connect.Request[emptypb.Empty]) (*connect.Response[grpc.ListTweetResponse], error) {
	usecase := tweet.NewListTweetUsecaseImpl(s.userRepo, s.tweetRepo)
	res, err := usecase.Exec(ctx)
	if err != nil {
		return nil, err
	}

	resp := connect.NewResponse(&grpc.ListTweetResponse{
		Tweets: res.Tweets,
	})

	return resp, nil
}

func (s *TweetServer) CreateTweet(ctx context.Context, in *connect.Request[grpc.CreateTweetRequest]) (*connect.Response[grpc.CreateTweetResponse], error) {
	validator := form.NewFormValidator(tweetForm.NewCreateTweetForm(in.Msg))
	if err := validator.Validate(); err != nil {
		return nil, err
	}

	usecase := tweet.NewCreateTweetUsecaseImpl(s.tweetRepo)

	res, err := usecase.Exec(ctx, in.Msg)
	if err != nil {
		return nil, err
	}

	resp := connect.NewResponse(&grpc.CreateTweetResponse{
		Id: res.Id,
	})

	return resp, nil
}
