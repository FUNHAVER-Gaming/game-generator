package service

import (
	context "context"
	"github.com/FUNHAVER-Gaming/game-generator/pkg/proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"strconv"
)

type server struct {
	port int
}

func (s server) CreatePlayer(ctx context.Context, request *proto.CreatePlayerRequest) (*proto.CreatePlayerResponse, error) {
	return createPlayer(request)
}

func (s server) CreateGame(ctx context.Context, request *proto.CreateGameRequest) (*proto.CreateGameResponse, error) {
	return createGame(request)
}

func StartServer(port int) error {
	server := server{port: port}

	log.WithFields(log.Fields{
		"port": port,
	}).Info("Starting gRPC")

	s := grpc.NewServer()

	proto.RegisterLeagueServiceServer(s, &server)

	address := ":" + strconv.Itoa(port)

	lis, err := net.Listen("tcp", address)

	if err != nil {
		log.WithError(err).Error("failed to connect to grpc address")
		return err
	}

	log.Info("Serving gRPC")

	if err := s.Serve(lis); err != nil {
		log.WithError(err).Error("failed to start grpc server")
	}

	return nil
}
