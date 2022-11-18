package command

import (
	"fmt"
	"os"
	"reflect"
	"regexp"

	"github.com/elct9620/wvs/internal/infrastructure/container"
	"github.com/elct9620/wvs/internal/infrastructure/rpc"
	"github.com/elct9620/wvs/internal/repository"
)

type RPCService struct {
	rpc.RPC
	container *container.Container
	matchRepo *repository.MatchRepository
}

func NewRPCService(container *container.Container, matchRepo *repository.MatchRepository) *RPCService {
	service := &RPCService{
		RPC:       *rpc.NewRPC(),
		container: container,
		matchRepo: matchRepo,
	}

	service.setup()

	return service
}

func (s *RPCService) setup() {
	methodFinder := reflect.TypeOf(s)

	for i := 0; i < methodFinder.NumMethod(); i++ {
		method := methodFinder.Method(i)

		ok, err := regexp.MatchString("^Setup", method.Name)
		if err != nil {
			fmt.Fprintf(os.Stderr, "invalid regexp for: %s\n", err)
			os.Exit(1)
		}

		if !ok {
			continue
		}

		method.Func.Call([]reflect.Value{reflect.ValueOf(s)})
	}
}
