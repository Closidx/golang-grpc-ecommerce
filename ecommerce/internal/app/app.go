package app

import (
	grpcapp "github/closidx/golang-grpc-ecommerce/ecommerce/internal/app/grpc"
	"log/slog"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
) *App {
	grpcApp := grpcapp.New(log, grpcPort)

	return &App{
		GRPCSrv: grpcApp,
	}
}