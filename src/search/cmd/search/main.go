package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"

	"time"

	"strconv"

	"cloud.google.com/go/trace"
	"github.com/disiqueira/frango/lib"
	"github.com/disiqueira/frango/src/search/proto/search"
	"github.com/kelseyhightower/envconfig"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type config struct {
	Port string `default:"8080" envconfig:"PORT"`
}

var (
	cfg config
)

func init() {
	envconfig.MustProcess("", &cfg)
}

type searchServer struct {
	traceClient *trace.Client
}

// Nearby returns all hotels within a given distance.
func (s *searchServer) Asa(ctx context.Context, req *search.AsaFilter) (*search.AsaList, error) {
	// add some artificial time so traces display nicely
	time.Sleep(time.Duration(rand.Int31n(5)) * time.Millisecond)

	res := &search.AsaList{}
	res.AsaIds = append(res.AsaIds, strconv.Itoa(rand.Intn(5)))

	return res, nil
}

func main() {
	// tcp listener
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	tc := lib.NewTraceClient(
		os.Getenv("TRACE_PROJECT_ID"),
		os.Getenv("TRACE_JSON_CONFIG"),
	)

	// grpc server
	srv := grpc.NewServer(
		grpc.UnaryInterceptor(trace.GRPCServerInterceptor(tc)),
	)
	search.RegisterSearchServer(srv, &searchServer{
		traceClient: tc,
	})
	srv.Serve(lis)
}
