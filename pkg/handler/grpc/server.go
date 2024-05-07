package grpc

import "github.com/fuku01/test-v2-grpc/pkg/grpc/pb"

// 複数のHandlerを一つの構造体にまとめる（埋め込みフィールドにして、GRPCServiceServer構造体自体が全てのメソッドを持つようにする必要がある）
type GRPCServiceServer struct {
	TodoHandler
}

// pbの型を満たしているかチェックする
var _ pb.GRPCServiceServer = (*GRPCServiceServer)(nil)
