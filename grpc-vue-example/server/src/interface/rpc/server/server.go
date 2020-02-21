package server

import (
	"log"
	"[[.ModName]]/src/infrastructure/datasource/database"
	"[[.ModName]]/src/interface/rpc"
	"net/http"

	"github.com/improbable-eng/grpc-web/go/grpcweb"

	"google.golang.org/grpc"
)

func GrpcRun(dbConfigPath, env, addr string) {

	//dbm init
	dbm := database.GetDBM()
	if err := dbm.InitDB(dbConfigPath, env); err != nil {
		panic(err)
	}

	grpcSrv := grpc.NewServer()
	rpc.RegisterRoomHandlerServer(grpcSrv, NewGrpcRoomServer(dbm))

	//for client(wrap)
	wrappedGrpc := grpcweb.WrapServer(
		grpcSrv,
		grpcweb.WithOriginFunc(func(origin string) bool { return true }),
	)

	handler := func(res http.ResponseWriter, req *http.Request) {
		wrappedGrpc.ServeHTTP(res, req)
	}

	httpServer := http.Server{
		Addr:    addr,
		Handler: http.HandlerFunc(handler),
	}
	log.Printf("[%s]starting http server port: %s", env, addr)

	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatalf("failed to start http server:%v", err)
	}
}
