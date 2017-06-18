package main

import (
	"fmt"
	"log"

	"cloud.google.com/go/trace"
	"github.com/disiqueira/frango/src/api/lib"
	"github.com/disiqueira/frango/src/api/proto/search"
	"github.com/kelseyhightower/envconfig"
	"google.golang.org/grpc"
)

type config struct {
	TraceProjectID  string `envconfig:"TRACE_PROJECT_ID"`
	TraceJSONConfig string `envconfig:"TRACE_JSON_CONFIG"`
	Port            string `default:"8080" envconfig:"PORT"`
	SearchAddr      string `default:"search:8080" envconfig:"ASA_ADDR"`
}

type env struct {
	cfg config

	Tracer       *trace.Client
	SearchClient search.SearchClient
}

func newEnv() *env {
	var cfg config
	envconfig.MustProcess("", &cfg)

	tc := lib.NewTraceClient(
		cfg.TraceProjectID, cfg.TraceJSONConfig,
	)

	return &env{
		cfg:          cfg,
		Tracer:       tc,
		SearchClient: search.NewSearchClient(mustDial(cfg.SearchAddr, tc)),
	}
}

// mustDial ensures a tcp connection to specified address.
func mustDial(addr string, tc *trace.Client) *grpc.ClientConn {
	conn, err := grpc.Dial(
		addr,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(trace.GRPCClientInterceptor()),
	)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
		panic(err)
	}
	return conn
}

func (e *env) serviceAddr() string {
	return fmt.Sprintf(":%s", e.cfg.Port)
}
