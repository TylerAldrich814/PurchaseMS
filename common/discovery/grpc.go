package discovery

import (
	"context"
	"math/rand"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ServiceConnection(
  ctx         context.Context,
  serviceName string,
  registry    Registry,
)( *grpc.ClientConn, error ) {
  address, err := registry.Discover(ctx, serviceName)
  if err != nil {
    return nil, err
  }

  return grpc.Dial(
    address[rand.Intn(len(address))],
    grpc.WithTransportCredentials(
      insecure.NewCredentials(),
    ),
  )
}
