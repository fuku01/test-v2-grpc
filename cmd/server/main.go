package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/fuku01/test-v2-api/db/config"
	"github.com/fuku01/test-v2-api/pkg/grpc/pb"
	handler "github.com/fuku01/test-v2-api/pkg/handler/grpc"
	repository "github.com/fuku01/test-v2-api/pkg/infrastructure/mysql"
	"github.com/fuku01/test-v2-api/pkg/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// 1. 6000番portのlistenerを作成
	port := os.Getenv("PORT")
	debug := true

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err)
	}

	db, err := config.NewDatabase()
	if err != nil {
		panic(err)
	}

	// 依存性の注入
	todoRepository := repository.NewTodoRepository(db)
	todoUsecase := usecase.NewTodoUsecase(todoRepository)
	todoHandler := handler.NewTodoHandler(todoUsecase)

	srv := handler.GRPCServiceServer{
		TodoHandler: todoHandler,
	}

	// 2. gRPCサーバーを作成
	s := grpc.NewServer()

	// 3. Register***ServiceServerを呼び出し、gRPCサーバーにサービスを登録
	pb.RegisterGRPCServiceServer(s, srv)

	// 4. gRPCサーバーにリフレクションを登録（APIの構造を外部に公開する。※gRPC UIでデバッグする用）
	if debug {
		reflection.Register(s)
	}

	// 5. 作成したgRPCサーバーを、8080番ポートで稼働させる
	go func() {
		log.Printf("start gRPC server port: %v", port)
		err := s.Serve(listener)
		if err != nil {
			panic(err)
		}
	}()

	// 6. Ctrl+Cで終了するように設定
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()
}
