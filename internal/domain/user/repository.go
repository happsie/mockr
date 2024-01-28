package user

import (
	"errors"

	"github.com/google/uuid"
	"github.com/happsie/go-webserver-template/internal/architecture"
)

type Repository struct {
	Container *architecture.Container
}

func (r Repository) Create(user User) error {
	res, err := r.Container.DB.NamedExec(`INSERT INTO users (id, email, display_name, created_at, updated_at, version) 
									VALUES (:id, :email, :display_name, :created_at, :updated_at, 1)`, user)
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return errors.New("minumum affected rows not reached")
	}
	return nil
}

func (r Repository) Read(ID uuid.UUID) (User, error) {
	user := User{}
	err := r.Container.DB.Get(&user, "SELECT * FROM users WHERE id = $1", ID.String())
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (r Repository) Update(user User) error {
	res, err := r.Container.DB.NamedExec(`UPDATE users
										SET id = :id, display_name = :display_name, email = :email, created_at = :created_at, updated_at = :updated_at, version = :version + 1
										WHERE id = :id AND version = :version`, user)
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return errors.New("minumum affected rows not reached")
	}
	return nil
}

func (r Repository) Delete(ID uuid.UUID) error {
	res, err := r.Container.DB.Exec(`DELETE FROM users where id = $1`, ID)
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return errors.New("minumum affected rows not reached")
	}
	return nil
}
