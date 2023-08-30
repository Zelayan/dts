package service

import (
	"context"
	"fmt"
	"time"

	pb "github.com/Zelayan/dts/proto-gen/v1/dts"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewSender(ctx context.Context, collectorAddress string, ch <-chan *pb.Span) (*Sender, error) {
	conn, err := grpc.Dial(collectorAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := pb.NewCollectorServiceClient(conn)
	buf := make(chan *pb.Span, 10)
	return &Sender{
		ctx:           ctx,
		ch:            ch,
		buf:           buf,
		fullCh:        make(chan struct{}),
		collectClient: client,
	}, nil
}

type Sender struct {
	ctx           context.Context
	ch            <-chan *pb.Span
	buf           chan *pb.Span
	fullCh        chan struct{}
	collectClient pb.CollectorServiceClient
}

func (s *Sender) Run() {
	for ch := range s.ch {
		s.buf <- ch
		if (len(s.buf)) == cap(s.buf) {
			s.fullCh <- struct{}{}
		}
	}
}

func (s *Sender) BatchSendG() {
	ticker := time.NewTicker(time.Millisecond * 100)
	defer ticker.Stop()
	errChan := make(chan error)
	go func() {
		for err := range errChan {
			fmt.Println(err)
		}
	}()

	for {
		select {
		case <-s.fullCh:
			err := s.BatchSend()
			if err != nil {
				errChan <- err
			}
		case <-ticker.C:
			err := s.BatchSend()
			if err != nil {
				errChan <- err
			}
		case <-s.ctx.Done():
			close(errChan)
			fmt.Println("done")
			return
		}
	}
}

func (s *Sender) BatchSend() error {
	spans := s.readBuf()
	if len(spans) == 0 {
		return nil
	}
	return s.RequestCollector(spans)
}

func (s *Sender) RequestCollector(spans []*pb.Span) error {
	_, err := s.collectClient.PostSpans(context.Background(), &pb.PostSpansRequest{Batch: &pb.Batch{Spans: spans}})
	if err != nil {
		return err
	}
	return nil
}

func (s *Sender) readBuf() []*pb.Span {
	ret := make([]*pb.Span, 0, cap(s.buf))
	done := false
	for i := 0; i < cap(s.buf); i++ {
		select {
		case span := <-s.buf:
			ret = append(ret, span)
		default:
			done = true
		}
		if done {
			break
		}
	}
	return ret
}
