package controllers

import (
    "context"
	"fmt"
	"os"
	"testing"
	"github.com/stretchr/testify/assert"
    "github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

var ip string
var port string

func TestMain(m *testing.M) {
	//ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "postgresSql",
		ExposedPorts: []string{"5433/tcp"},
	}
	postgC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer postgC.Terminate(ctx)
	ip, err = postgC.Host(ctx)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	postgPort, err := postgC.MappedPort(ctx, "5433/tcp")
	port = postgPort.Port()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
    }

	m.Run()

}

func TestFind(t *testing.T) {
	fmt.Printf("asdas")
	assert.Equal(t,1,1)
}