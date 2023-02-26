package repository

import "github.com/sk10az/go_anime_crud/internal/app/model"

type Interface interface {
	GetAnimeCharacter(id model.IdAC) (*model.AnimeCharacter, error)
	GetAllAnimeCharacters() []*model.AnimeCharacter
	CreateAnimeCharacter(a *model.AnimeCharacter) model.IdAC
	UpdateAnimeCharacter(id model.IdAC, a *model.AnimeCharacter) (*model.AnimeCharacter, error)
	DeleteAnimeCharacter(id model.IdAC) error
}

type MapAnimeCharacters map[model.IdAC]*model.AnimeCharacter
