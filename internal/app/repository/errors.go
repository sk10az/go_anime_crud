package repository

import (
	"fmt"
	"github.com/sk10az/go_anime_crud/internal/app/model"
)

type ACNotFoundError struct {
	Id model.IdAC
}

func (a *ACNotFoundError) Error() string {
	return fmt.Sprintf("not found AnimeCharacter by id '%id'", a.Id)
}

type ZeroIdError struct{}

func (z *ZeroIdError) Error() string {
	return "Id can't be zero (0)"
}
