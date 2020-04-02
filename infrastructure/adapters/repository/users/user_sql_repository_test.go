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
			"POSTGRES_PASSWORD": "secret",
			"POSTGRES_DB":       "users_db",
		},
		WaitingFor: wait.ForLog("port: 5432  postgres Server"),
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
	fmt.Print("puerto test en uso : ", port)
	return postgresC, ctx
}

func beforeAll(ctx context.Context, container testcontainers.Container) {
	_ = container.Terminate(ctx)
}
/*
func TestUserSqlRepository_Update(t *testing.T) {
	tx := userSqlRepository.Db.Begin()
	defer tx.Rollback()
	var user model.User
	user, _ = user.CreateUser(1,"User1", "test", 7)
	dni, err := userSqlRepository.UpdateQuantityMovies(&user)

	assert.Nil(t, err)
	assert.EqualValues(t, user.FirstName, "User1", "user names are differences")
	assert.NotEqual(t, user.LastName, "test")
	assert.NotNil(t, user.Id, "user shouldn't be nil ")
}*/

func TestUserSqlRepository_Get(t *testing.T) {

	tx := userSqlRepository.Db.Begin()
	defer tx.Rollback()
	var user model.User
	user, _ = user.CreateUser(4,"usuario3", "Test", 8)
	var userDb models.UserDb
	userDb = users_mapper.UserToUserDb(user)
	if err := userSqlRepository.Db.Create(&userDb).Error; err != nil {
		assert.Fail(t, err.Error())
	}
	user, err := userSqlRepository.GetByDNI(userDb.DNI)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, userDb.DNI, user.DNI)
	assert.EqualValues(t, "Test", user.LastName)
}