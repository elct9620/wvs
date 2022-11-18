package command

import (
	"fmt"
	"os"
	"reflect"
	"regexp"

	"github.com/elct9620/wvs/internal/engine"
	"github.com/elct9620/wvs/internal/repository"
	"github.com/elct9620/wvs/internal/service"
	"github.com/elct9620/wvs/pkg/rpc"
)

type RPCService struct {
	rpc.RPC
	engine           *engine.Engine
	matchRepo        *repository.MatchRepository
	broadcastService *service.BroadcastService
	gameLoopService  *service.GameLoopService
}

func NewRPCService(engine *engine.Engine, matchRepo *repository.MatchRepository, broadcastService *service.BroadcastService, gameLoopService *service.GameLoopService) *RPCService {
	service := &RPCService{
		RPC:              *rpc.NewRPC(),
		engine:           engine,
		matchRepo:        matchRepo,
		broadcastService: broadcastService,
		gameLoopService:  gameLoopService,
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
