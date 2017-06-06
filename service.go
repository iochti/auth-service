package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io/ioutil"

	pb "github.com/iochti/auth-service/proto"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

// AuthSvc represents the auth service
type AuthSvc struct {
	Conf *oauth2.Config
}

//HandleAuth handles the exchange code to initiate transport
func (a *AuthSvc) HandleAuth(ctx context.Context, in *pb.AuthRequest) (*pb.AuthResponse, error) {
	tok, err := a.Conf.Exchange(oauth2.NoContext, in.GetCode())
	if err != nil {
		return nil, err
	}
	client := a.Conf.Client(oauth2.NoContext, tok)
	user, err := client.Get("https://api.github.com/user")
	if err != nil {
		return nil, err
	}
	res := pb.AuthResponse{}
	defer user.Body.Close()
	data, _ := ioutil.ReadAll(user.Body)
	res.User = string(data)
	return &res, nil
}

// Generates a state random token
func randToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

// GetLoginURL sends a login url for a given user
func (a *AuthSvc) GetLoginURL(ctx context.Context, in *pb.LoginURLRequest) (*pb.LoginURLResponse, error) {
	state := in.GetState()
	if state == "" {
		return nil, fmt.Errorf("Error: missing state in context")
	}
	url := a.Conf.AuthCodeURL(state)
	return &pb.LoginURLResponse{Url: url}, nil
}
