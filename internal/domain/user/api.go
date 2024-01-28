package user

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/happsie/go-webserver-template/internal/architecture"
	"github.com/labstack/echo/v4"
)

type api struct {
	Container  *architecture.Container
	Repository Repository
}

func InitAPI(container *architecture.Container) *api {
	return &api{
		Container:  container,
		Repository: Repository{Container: container},
	}
}

func (a api) Register(e *echo.Echo) {
	g := e.Group("/api/users/user-v1")
	g.POST("", a.create)
	g.GET("/:id", a.read)
	g.PUT("", a.update)
	g.DELETE("/:id", a.delete)
}

func (a api) create(c echo.Context) error {
	req := CreateUserRequest{}
	err := c.Bind(&req)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	creationTime := time.Now()
	user := User{
		ID:          uuid.New(),
		DisplayName: req.DisplayName,
		Email:       req.Email,
		CreatedAt:   creationTime,
		UpdatedAt:   creationTime,
		Version:     1,
	}
	err = a.Repository.Create(user)
	if err != nil {
		a.Container.L.Error("error storing user", "error", err)
		return c.NoContent(http.StatusBadRequest)
	}
	return c.JSON(http.StatusCreated, user)
}

func (a api) read(c echo.Context) error {
	ID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	user, err := a.Repository.Read(ID)
	if err != nil {
		a.Container.L.Warn("Could not read user", "id", ID, "error", err)
		return c.NoContent(http.StatusNotFound)
	}
	return c.JSON(http.StatusOK, user)
}

func (a api) update(c echo.Context) error {
	req := UpdateUserRequest{}
	err := c.Bind(&req)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	user, err := a.Repository.Read(req.ID)
	if err != nil {
		a.Container.L.Warn("Could not read user", "id", req.ID, "error", err)
		return c.NoContent(http.StatusNotFound)
	}
	user.DisplayName = req.DisplayName
	user.Email = req.Email
	err = a.Repository.Update(user)
	if err != nil {
		a.Container.L.Error("error updating user", "id", req.ID, "error", err)
		return c.NoContent(http.StatusBadRequest)
	}
	return c.JSON(http.StatusOK, user)
}

func (a api) delete(c echo.Context) error {
	ID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	err = a.Repository.Delete(ID)
	if err != nil {
		a.Container.L.Error("error deleting user", "id", ID, "error", err)
		return c.NoContent(http.StatusBadRequest)
	}
	return c.NoContent(http.StatusNoContent)
}
