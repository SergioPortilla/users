package users

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/ceiba-meli-demo/users/domain/model"
	"github.com/ceiba-meli-demo/users/infrastructure/adapters/repository/models"
	"github.com/ceiba-meli-demo/users/infrastructure/database_client"
	"github.com/ceiba-meli-demo/users/infrastructure/mappers/users_mapper"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

var (
	userSqlRepository UserSqlRepository
)

func TestMain(m *testing.M) {
	fmt.Println("about to start users tests")
	containerMockServer, ctx := load()
	code := m.Run()
	beforeAll(ctx, containerMockServer)
	os.Exit(code)
}

func load() (testcontainers.Container, context.Context) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "postgres:latest",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_USER":     "root",
			"POSTGRES_PASSWORD": "secret",
			"POSTGRES_DB":       "users_db",
		},
		WaitingFor: wait.ForLog("Listening on 0.0.0.0"),
	}
	postgresC, _ := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	host, _ := postgresC.Host(ctx)
	p, _ := postgresC.MappedPort(ctx, "5432/tcp")
	port := p.Port()
	_ = os.Setenv("POSTGRES_HOST", host)
	_ = os.Setenv("POSTGRES_PORT", port)
	_ = os.Setenv("POSTGRES_DB", "users_db")
	_ = os.Setenv("POSTGRES_USER", "root")
	_ = os.Setenv("POSTGRES_PASSWORD", "secret")

	userSqlRepository = UserSqlRepository{
		Db: database_client.GetDatabaseInstance(),
	}
	return postgresC, ctx
}

func beforeAll(ctx context.Context, container testcontainers.Container) {
	_ = container.Terminate(ctx)
}

func TestUserSqlRepository_Update(t *testing.T) {
	tx := userSqlRepository.Db.Begin()
	defer tx.Rollback()
	var user model.User
	
	user, _ = user.CreateUser(7, "usuario3", "Test", 8)
	var userDb models.UserDb
	userDb = users_mapper.UserToUserDb(user)
	if err := userSqlRepository.Db.Create(&userDb).Error; err != nil {
		assert.Fail(t, err.Error())
	}

	msg, err := userSqlRepository.UpdateQuantityMovies(userDb.DNI)
	assert.Nil(t, err)
	assert.EqualValues(t, msg.Done, true)
}

func TestUpdateQuantityMovies(t *testing.T) {
	tx := userSqlRepository.Db.Begin()
	defer tx.Rollback()
	var user model.User
	
	user, _ = user.CreateUser(77, "usuario3", "Test", 8)
	var userDb models.UserDb
	userDb = users_mapper.UserToUserDb(user)
	if err := userSqlRepository.Db.Create(&userDb).Error; err != nil {
		assert.Fail(t, err.Error())
	}

	msg, err := userSqlRepository.UpdateQuantityMovies(userDb.DNI)
	assert.Nil(t, err)
	assert.EqualValues(t, msg.Done, true)
}

func TestUpdateQuantityMoviesNotFound(t *testing.T) {
	tx := userSqlRepository.Db.Begin()
	defer tx.Rollback()
	_, err := userSqlRepository.UpdateQuantityMovies(-100)
	assert.NotNil(t, err)
}

func TestUserSqlRepository_Get(t *testing.T) { 

	tx := userSqlRepository.Db.Begin()
	defer tx.Rollback()
	var user model.User
	user, _ = user.CreateUser(1, "usuario3", "Test", 8)
	var userDb models.UserDb
	userDb = users_mapper.UserToUserDb(user)
	if err := userSqlRepository.Db.Create(&userDb).Error; err != nil {
		assert.Fail(t, err.Error())
	}
	userRes, err := userSqlRepository.GetByDNI(userDb.DNI)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, userDb.DNI, userRes.DNI)
	assert.EqualValues(t, "Test", userRes.LastName)
}
