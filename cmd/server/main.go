package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/fuku01/test-v2-api/db/config"
	"github.com/fuku01/test-v2-api/pkg/grpc/pb"
	grpc_handler "github.com/fuku01/test-v2-api/pkg/handler/grpc"
	repository "github.com/fuku01/test-v2-api/pkg/infrastructure/mysql"
	"github.com/fuku01/test-v2-api/pkg/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	debug := true

	// 1. 6000番ポートのlistenerを作成
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalf("環境変数 PORT が設定されていません")
	}
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err)
	}

	// 2. DB接続
	db, err := config.NewDatabase()
	if err != nil {
		panic(err)
	}

	// @ 3. 依存性の注入
	todoRepository := repository.NewTodoRepository(db)
	todoUsecase := usecase.NewTodoUsecase(todoRepository)
	todoHandler := grpc_handler.NewTodoHandler(todoUsecase)

	srv := grpc_handler.GRPCServiceServer{
		TodoHandler: todoHandler,
	}

	// 4. gRPC サーバーの作成
	s := grpc.NewServer()
	// 5. Register***ServiceServerを呼び出し、gRPCサーバーにサービス(実装)を登録
	pb.RegisterGRPCServiceServer(s, srv) // 引数srvは、実装したサービスのメソッド全てを持っている必要がある

	// 6. gRPCサーバーにリフレクションを登録（APIの構造を外部に公開する。※gRPC UIでデバッグする用）
	if debug {
		reflection.Register(s)
	}

	// 7. 作成したサーバーの起動
	go func() {
		log.Printf("start gRPC server port: %v", port)
		err := s.Serve(listener)
		if err != nil {
			panic(err)
		}
	}()

	// 8. Ctrl+Cでサーバーが終了するように設定
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()
}
