package main

import (
	"fmt"
	"net"
	"os"

	"google.golang.org/grpc"

	"github.com/iochti/auth-service/models"
	pb "github.com/iochti/auth-service/proto"
	"github.com/namsral/flag"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"google.golang.org/grpc/credentials"
)

var conf *oauth2.Config
var ghubCreds models.Credentials

// RedirectURL used during oauth
var RedirectURL string

// Inits the service state & conf
func init() {
	flag.StringVar(&RedirectURL, "redirect_url", "http://127.0.0.1:3000/auth", "Redirect URL used during oauth")
	flag.Parse()
	// Gets Apps secrets & id for github
	ghubCreds.Init()
	fmt.Println(ghubCreds)
	conf = &oauth2.Config{
		ClientID:     ghubCreds.Cid,
		ClientSecret: ghubCreds.Csecret,
		RedirectURL:  RedirectURL,
		Scopes:       []string{"user:email"},
		Endpoint:     github.Endpoint,
	}
}

func dieIf(err error) {
	if err == nil {
		return
	}
	fmt.Fprintf(os.Stderr, "Error: %s. Try --help for help.\n", err)
	os.Exit(-1)
}

func main() {
	addr := flag.String("srv", ":5000", "TCP address to listen on (in host:port form)")
	certFile := flag.String("cert", "", "Path to PEM-encoded certificate")
	keyFile := flag.String("key", "", "Path to PEM-encoded secret key")
	flag.Parse()
	if flag.NArg() != 0 {
		dieIf(fmt.Errorf("expecting zero arguments but got %d", flag.NArg()))
	}
	svc := &AuthSvc{
		Conf: conf,
	}

	var server *grpc.Server
	if *keyFile == "" && *certFile == "" {
		server = grpc.NewServer()
	} else if *certFile == "" {
		dieIf(fmt.Errorf("key specified with no cert"))
	} else if *keyFile == "" {
		dieIf(fmt.Errorf("cert specified with no key"))
	} else {
		pair, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		dieIf(err)
		creds := grpc.Creds(pair)
		server = grpc.NewServer(creds)
	}
	lis, err := net.Listen("tcp", *addr)
	dieIf(err)
	pb.RegisterAuthSvcServer(server, svc)
	server.Serve(lis)
}
