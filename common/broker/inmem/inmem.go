package inmem

import (
	"context"
	"errors"
	"sync"
	"time"
)

var (
  ErrorInstanceIDNotFound = errors.New("Instance ID Not Found in Registry")
  ErrorServiceNameNotFound  = errors.New("Service Name Not Found in Registry")
)

type serviceInstance struct {
  hostPort   string
  lastActive time.Time
}

type Registry struct {
  sync.Mutex
  addrs map[string]map[string]*serviceInstance
}

func NewRegistry() *Registry {
  return &Registry{
    addrs: map[string]map[string]*serviceInstance{},
  }
}

func(r *Registry) Register(
  ctx         context.Context,
  instanceID  string,
  serviceName string,
  hostPort    string,
) error {
  r.Lock()
  defer r.Unlock()

  if _, ok := r.addrs[serviceName]; !ok {
    r.addrs[serviceName] = map[string]*serviceInstance{}
  }

  r.addrs[serviceName][instanceID] = &serviceInstance{
    hostPort   : hostPort,
    lastActive : time.Now(),
  }

  return nil
}

func(r *Registry) Deregister(
  ctx         context.Context,
  instanceID  string,
  serviceName string,
) error {
  r.Lock()
  defer r.Unlock()

  if _, ok := r.addrs[serviceName]; !ok {
    return ErrorServiceNameNotFound
  }

  if _, ok := r.addrs[serviceName][instanceID]; !ok {
    return ErrorInstanceIDNotFound
  }

  delete(r.addrs[serviceName], instanceID)

  return nil
}

func(r *Registry) HealthCheck(
  instanceID  string,
  serviceName string,
) error {
  r.Lock()
  defer r.Unlock()

  if _, ok := r.addrs[serviceName]; !ok {
    return ErrorServiceNameNotFound
  }

  if _, ok := r.addrs[serviceName][instanceID]; !ok {
    return ErrorInstanceIDNotFound
  }

  r.addrs[serviceName][instanceID].lastActive = time.Now()

  return nil
}

func(r *Registry) Discover(
  ctx         context.Context,
  serviceName string,
)( []string,error ){
  r.Lock()
  defer r.Unlock()

  if len(r.addrs[serviceName]) == 0 {
    return nil, ErrorServiceNameNotFound
  }

  var res []string
  for _, i := range r.addrs[serviceName] {
    res = append(res, i.hostPort)
  }

  return res,nil
}

func(r *Registry) ServiceAddresses(
  ctx         context.Context,
  serviceName string,
)( []string, error ){
  r.Lock()
  defer r.Unlock()

  if len(r.addrs[serviceName]) == 0 {
    return nil, ErrorServiceNameNotFound
  }

  var res []string
  for _, i := range r.addrs[serviceName]{
    if i.lastActive.Before(time.Now().Add(-5 * time.Second)){
      continue
    }
    res = append(res, i.hostPort)
  }


  return res,nil
}
