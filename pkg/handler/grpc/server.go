package handler

import "github.com/fuku01/test-v2-api/pkg/grpc/pb"

// 複数のHandlerを一つの構造体にまとめる（埋め込みフィールド）
type GRPCServiceServer struct {
	TodoHandler
}

// pbの型を満たしているかチェックする
var _ pb.GRPCServiceServer = (*GRPCServiceServer)(nil)
