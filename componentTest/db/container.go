package db

import (
	"context"
	"fmt"
	"myapp/persistence/config"
	"time"

	"github.com/docker/go-connections/nat"
	_ "github.com/lib/pq"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type PostgresContainer struct {
	Container  testcontainers.Container
	MappedPort nat.Port
}

func InitPostgresContainer(cfg config.DBConfig, ctx context.Context) (*PostgresContainer, error) {
	// var host = cfg.Host
	var name = cfg.Name
	var user = cfg.User
	var password = cfg.Password
	var timeout = 5 * time.Second
	var port = "5432/tcp"

	var env = map[string]string{
		"POSTGRES_PASSWORD": password,
		"POSTGRES_USER":     user,
		"POSTGRES_DB":       name,
	}
	dbURL := func(host string, port nat.Port) string {
		return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			user,
			password,
			cfg.Host,
			port.Port(),
			name)
	}

	natPort := nat.Port(port)
	req := testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "postgres:latest",
			ExposedPorts: []string{port},
			Cmd:          []string{"postgres", "-c", "fsync=off"},
			Env:          env,
			WaitingFor:   wait.ForSQL(natPort, "postgres", dbURL).Timeout(timeout),
		},
		Started: true,
	}
	container, err := testcontainers.GenericContainer(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to start Container: %s", err)
	}

	mappedPort, err := container.MappedPort(ctx, natPort)
	if err != nil {
		return nil, fmt.Errorf("failed to get Container external port: %s", err)
	}
	fmt.Printf("postgres Container ready and running at port: %s\n", mappedPort)
	return &PostgresContainer{Container: container, MappedPort: mappedPort}, nil
}
