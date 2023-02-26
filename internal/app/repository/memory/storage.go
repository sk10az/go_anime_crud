package memory

import (
	"github.com/sk10az/go_anime_crud/internal/app/model"
	"github.com/sk10az/go_anime_crud/internal/app/repository"
	"github.com/sk10az/go_anime_crud/pkg/logger"
)

type ServiceInMemory struct {
	logger  logger.Interface
	mapAc   repository.MapAnimeCharacters
	IndexAC model.IdAC
}

func New(l logger.Interface) (*ServiceInMemory, error) {
	mapAC := make(repository.MapAnimeCharacters)
	s := &ServiceInMemory{
		logger:  l,
		mapAc:   mapAC,
		IndexAC: 1,
	}
	return s, nil
}

func validateNonZeroId(l logger.Interface, id model.IdAC) error {
	if id == 0 {
		err := &repository.ZeroIdError{}
		l.Error(err.Error())
		return err
	}
	return nil
}

func (s *ServiceInMemory) GetAnimeCharacter(id model.IdAC) (*model.AnimeCharacter, error) {
	err := validateNonZeroId(s.logger, id)
	if err != nil {
		return nil, err
	}

	AC, ok := s.mapAc[id]
	if !ok {
		err := &repository.ACNotFoundError{
			Id: id,
		}
		s.logger.Error(err.Error())
		return nil, err
	}
	s.logger.Info("Returning Anime Character by id: ", id)
	return AC, nil
}

func (s *ServiceInMemory) GetAllAnimeCharacters() []*model.AnimeCharacter {
	result := make([]*model.AnimeCharacter, 0, len(s.mapAc))

	s.logger.Info("Iterating in all Anime Characters in memory")

	for key := range s.mapAc {
		result = append(result, s.mapAc[key])
	}

	s.logger.Info("Returning all Anime Characters in memory")
	return result
}

func (s *ServiceInMemory) CreateAnimeCharacters(a *model.AnimeCharacter) model.IdAC {
	if a.Id != 0 {
		s.logger.Warn("Ignoring ID of Anime Character")
	}
	s.logger.Info("Adding Anime Character to memory with id: ", s.IndexAC)
	id := s.IndexAC
	a.Id = id
	s.mapAc[id] = a

	s.IndexAC++

	s.logger.Info("Returning id of Anime Character in memory: ", id)
	return id
}

func (s *ServiceInMemory) UpdateAnimeCharacter(id model.IdAC, a *model.AnimeCharacter) (*model.AnimeCharacter, error) {
	err := validateNonZeroId(s.logger, id)
	if err != nil {
		return nil, err
	}

	_, ok := s.mapAc[id]
	if !ok {
		err := &repository.ACNotFoundError{Id: id}
		s.logger.Error(err.Error())
		return nil, err
	}

	s.logger.Info("Updating Anime Character by id: ", id)
	s.mapAc[id] = a
	return a, nil
}

func (s *ServiceInMemory) DeleteAnimeCharacter(id model.IdAC) error {
	err := validateNonZeroId(s.logger, id)
	if err != nil {
		return err
	}

	_, ok := s.mapAc[id]
	if !ok {
		err := &repository.ACNotFoundError{Id: id}
		s.logger.Error(err.Error())
		return err
	}

	s.logger.Info("Deleting Anime Character by id: ", id)
	delete(s.mapAc, id)

	return nil
}