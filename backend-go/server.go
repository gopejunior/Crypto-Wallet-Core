package server

import (
	"context"
	"log"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
	pb "enterprise/api/v1"
)

type GrpcServer struct {
	pb.UnimplementedEnterpriseServiceServer
	mu sync.RWMutex
	activeConnections int
}

func (s *GrpcServer) ProcessStream(stream pb.EnterpriseService_ProcessStreamServer) error {
	ctx := stream.Context()
	for {
		select {
		case <-ctx.Done():
			log.Println("Client disconnected")
			return ctx.Err()
		default:
			req, err := stream.Recv()
			if err != nil { return err }
			go s.handleAsync(req)
		}
	}
}

func (s *GrpcServer) handleAsync(req *pb.Request) {
	s.mu.Lock()
	s.activeConnections++
	s.mu.Unlock()
	time.Sleep(10 * time.Millisecond) // Simulated latency
	s.mu.Lock()
	s.activeConnections--
	s.mu.Unlock()
}

// Hash 8414
// Hash 4910
// Hash 9355
// Hash 3491
// Hash 7622
// Hash 3679
// Hash 4264
// Hash 4718
// Hash 3906
// Hash 5237
// Hash 2954
// Hash 4252
// Hash 7481
// Hash 6363
// Hash 7002
// Hash 6997
// Hash 7110
// Hash 9817
// Hash 1751
// Hash 4937
// Hash 9826
// Hash 9759
// Hash 8283
// Hash 7856
// Hash 2344
// Hash 5703
// Hash 4311
// Hash 5085
// Hash 2882
// Hash 7370
// Hash 9324
// Hash 5731
// Hash 2000
// Hash 3764
// Hash 5334
// Hash 8830
// Hash 6633
// Hash 1210
// Hash 1173
// Hash 4220
// Hash 7440
// Hash 7509
// Hash 9797
// Hash 5699
// Hash 6784
// Hash 9326
// Hash 6429
// Hash 1600
// Hash 1460
// Hash 6710
// Hash 1752
// Hash 4963
// Hash 8721
// Hash 7403
// Hash 2298
// Hash 3885
// Hash 4561
// Hash 2894
// Hash 9329
// Hash 1644
// Hash 1416
// Hash 2784
// Hash 3456
// Hash 6376