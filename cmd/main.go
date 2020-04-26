package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	tra "traffic-bot/pkg/apis/tra/v1alpha1"
)

var (
	//tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	//certFile   = flag.String("cert_file", "", "The TLS cert file")
	//keyFile    = flag.String("key_file", "", "The TLS key file")
	port = flag.Int("port", 10000, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	tra.RegisterSearchServer(grpcServer, tra.NewSearchService())

	grpcServer.Serve(lis)
}

//func tlsCheck(tls string) {
//	if *tls {
//		if *certFile == "" {
//			*certFile = testdata.Path("server1.pem")
//		}
//		if *keyFile == "" {
//			*keyFile = testdata.Path("server1.key")
//		}
//		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
//		if err != nil {
//			log.Fatalf("Failed to generate credentials %v", err)
//		}
//		opts = []grpc.ServerOption{grpc.Creds(creds)}
//	}
//}
