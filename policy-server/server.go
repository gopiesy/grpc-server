package server

import (
	"fmt"
	"time"

	"io"
	"log"

	"github.com/gopiesy/project-protos/policies"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type PolicyServiceServer struct {
}

func (p PolicyServiceServer) StreamSnapshots(stream policies.PolicyService_StreamSnapshotsServer) error {
	for {
		streamCpy := stream
		go func(stream policies.PolicyService_StreamSnapshotsServer) {
			status, err := stream.Recv()
			if err == io.EOF {
				log.Println("EOF received and exiting")
				return
			}
			if err != nil {
				log.Panic(err)
			}
			// received status
			log.Println("Status: ", status.GetSnapshotName())
		}(streamCpy)

		now := timestamppb.Now()
		name := fmt.Sprintf("NewSnapshot.%d", now.Seconds)
		if err := stream.Send(&policies.Snapshot{Name: name, Time: now, Data: nil}); err != nil {
			return err
		}
		<-time.After(time.Duration(time.Second * 2))
	}
}

func NewPolicyService() PolicyServiceServer {
	return PolicyServiceServer{}
}
