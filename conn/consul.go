package conn

import (
	"log"
	"os"

	"github.com/hashicorp/consul/api"
)

func InitConsul() *api.Client {
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		log.Printf("Error connecting on Consul, %v", err)
		return nil
	}

	registerThis(client)

	return client
}

func registerThis(client *api.Client) {
	hostname, _ := os.Hostname()
	client.Agent().ServiceRegister(&api.AgentServiceRegistration{
		ID:      "aetherius-" + hostname,
		Name:    "aetherius",
		Port:    4000,
		Address: "http://" + hostname,
	})
}
