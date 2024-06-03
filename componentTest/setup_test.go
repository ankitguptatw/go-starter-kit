package componenttests

import (
	"context"
	app "myapp/app/server"
	"myapp/app/server/config"
	"myapp/componentTest/db"
	"myapp/crossCutting/logger"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"gopkg.in/h2non/gock.v1"
)

var TestClient *testClient
var TestServer *httptest.Server
var Conf config.ServerConfig

func TestMain(m *testing.M) {
	// update these variable accordingly to point your local path of your docker engine
	// os.Setenv("DOCKER_HOST", "unix:///Users/username/.colima/docker.sock")
	// os.Setenv("TESTCONTAINERS_DOCKER_SOCKET_OVERRIDE", "/var/run/docker.sock")

	ctx := context.Background()
	Conf = config.NewServerConfig(filepath.Join(".", "conf", "test.yaml"))
	c := setUpDBContainer(ctx)
	Conf.Database.Port = c.MappedPort.Int()
	initDB()
	setUpMockServer()
	TestServer = httptest.NewServer(app.InitWithConf(Conf))

	exitCode := m.Run()

	if c.Container != nil {
		TestServer.Close()
		err := c.Container.Terminate(ctx)
		if err != nil {
			logger.GetLogger(ctx).Error("Termination of container failed. Err: %s", err.Error())
		}
	}
	os.Exit(exitCode)
}

func setUpDBContainer(ctx context.Context) *db.PostgresContainer {
	c, err := db.InitPostgresContainer(Conf.Database, ctx)
	if err != nil {
		logger.GetLogger(ctx).Error("Failed to create postgres container. %s", err.Error())
		return c
	}
	return c
}

func setUpMockServer() {
	os.Setenv("RUNNING_COMPONENT_TESTS", "true")
	gock.EnableNetworking()
	gock.NetworkingFilter(func(request *http.Request) bool {
		//nolint:staticcheck
		if request.URL.Host == strings.Trim(Conf.ServiceProviders.BankProvider.BaseURL, "http://") ||
			request.URL.Host == strings.Trim(Conf.ServiceProviders.FraudProvider.BaseURL, "http://") {
			return false
		}
		return true
	})
}

func initDB() {
	db.InitDB(Conf.Database)
	db.GetDB().Seed()
}
