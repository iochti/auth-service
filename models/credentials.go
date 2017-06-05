package models

import (
	"fmt"

	"github.com/namsral/flag" // We use namsral package so we can store flags as ENV vars
)

// Credentials represents the datas obtained by requesting oauth authorization
// to an application
type Credentials struct {
	Cid     string `json:"cid"`
	Csecret string `json:"csecret"`
}

// GHubCID => Github Credentials ID
var GHubCID string

// GHubCSecret => Github Secret ID
var GHubCSecret string

// Init a Credential instance by setting its IDS
func (c *Credentials) Init() error {
	flag.StringVar(&GHubCID, "ghub-cid", "", "Githbub Cid for Oauth")
	flag.StringVar(&GHubCSecret, "ghub-csecret", "", "Github Secret for Oauth")
	flag.Parse()

	if GHubCID == "" || GHubCSecret == "" {
		fmt.Println("Error, missing ghub credentials")
		return fmt.Errorf("Missing Github credential ID or Secret")
	}
	c.Cid = GHubCID
	c.Csecret = GHubCSecret
	return nil
}
