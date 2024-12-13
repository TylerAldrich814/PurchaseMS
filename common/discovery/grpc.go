package discovery

import (
	"context"
	"log"
	"math/rand"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ServiceConnection(
  ctx         context.Context,
  serviceName string,
  registry    Registry,
)( *grpc.ClientConn, error ) {
  addrs, err := registry.Discover(ctx, serviceName)
  if err != nil {
    return nil, err
  }

	//   tracerProvider  := otel.GetTracerProvider()
	// otelInterceptor := otelgrpc.NewClientHandler(otelgrpc.WithTracerProvider(tracerProvider))

  log.Printf("Discovered %d Instances of %s", len(addrs), serviceName)

  return grpc.NewClient(
    addrs[rand.Intn(len(addrs))],
    grpc.WithTransportCredentials(insecure.NewCredentials()),
    grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
    grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
  )

}
