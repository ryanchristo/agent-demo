package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/a2aproject/a2a-go/a2a"
	"github.com/a2aproject/a2a-go/a2agrpc"
	"github.com/a2aproject/a2a-go/a2asrv"
	"github.com/a2aproject/a2a-go/a2asrv/eventqueue"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

// agentExecutor implements [a2asrv.AgentExecutor], which is a required [a2asrv.RequestHandler] dependency.
// It is responsible for invoking an agent, translating its outputs to a2a.Event object and writing them to the provided [eventqueue.Queue].
type agentExecutor struct{}

// Execute is called when an agent is invoked.
func (*agentExecutor) Execute(ctx context.Context, reqCtx *a2asrv.RequestContext, q eventqueue.Queue) error {
	response := a2a.NewMessage(a2a.MessageRoleAgent, a2a.TextPart{Text: "Hello, world!"})
	return q.Write(ctx, response)
}

// Cancel is called when an agent invocation is cancelled.
func (*agentExecutor) Cancel(ctx context.Context, reqCtx *a2asrv.RequestContext, q eventqueue.Queue) error {
	return nil
}

// startGRPCServer starts a gRPC server on the given port.
func startGRPCServer(port int, card *a2a.AgentCard) error {

	// Create a listener for the server.
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}

	// Log the server starting.
	log.Printf("Starting a gRPC server on 127.0.0.1:%d", port)

	// Create a transport-agnostic A2A request handler.
	requestHandler := a2asrv.NewHandler(&agentExecutor{}, a2asrv.WithExtendedAgentCard(card))

	// Wrap the handler into a transport implementation.
	grpcHandler := a2agrpc.NewHandler(requestHandler)

	// Create the gRPC server.
	s := grpc.NewServer()

	// Register the handler with the server.
	grpcHandler.RegisterWith(s)

	// Start serving requests on the listener.
	return s.Serve(listener)
}

// servePublicCard starts a public A2A AgentCard server on the given port.
func servePublicCard(port int, card *a2a.AgentCard) error {

	// Create a listener for the server.
	listener, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", port))
	if err != nil {
		return err
	}

	// Log the server starting.
	log.Printf("Starting a public AgentCard server on 127.0.0.1:%d", port)

	// Create a multiplexer for the server.
	mux := http.NewServeMux()

	// Register the AgentCard handler with the multiplexer.
	mux.Handle(a2asrv.WellKnownAgentCardPath, a2asrv.NewStaticAgentCardHandler(card))

	// Start serving requests on the listener.
	return http.Serve(listener, mux)
}

var (
	grpcPort = flag.Int("grpc-port", 9000, "Port for a gGRPC A2A server to listen on.")
	cardPort = flag.Int("card-port", 9001, "Port for a public A2A AgentCard server to listen on.")
)

func main() {
	flag.Parse()

	agentCard := &a2a.AgentCard{
		Name:               "Hello World Agent",
		Description:        "Just a hello world agent",
		URL:                fmt.Sprintf("127.0.0.1:%d", *grpcPort),
		PreferredTransport: a2a.TransportProtocolGRPC,
		DefaultInputModes:  []string{"text"},
		DefaultOutputModes: []string{"text"},
		Capabilities:       a2a.AgentCapabilities{Streaming: true},
		Skills: []a2a.AgentSkill{
			{
				ID:          "hello_world",
				Name:        "Hello, world!",
				Description: "Returns a 'Hello, world!'",
				Tags:        []string{"hello world"},
				Examples:    []string{"hi", "hello"},
			},
		},
	}

	var group errgroup.Group
	group.Go(func() error {
		return startGRPCServer(*grpcPort, agentCard)
	})
	group.Go(func() error {
		return servePublicCard(*cardPort, agentCard)
	})
	if err := group.Wait(); err != nil {
		log.Fatalf("Server shutdown: %v", err)
	}
}
