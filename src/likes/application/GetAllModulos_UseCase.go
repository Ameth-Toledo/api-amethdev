package application

import "AmethToledo/src/likes/domain"

type GetAllModulosUseCase struct {
	repository domain.ILike
}

func NewGetAllModulosUseCase(repository domain.ILike) *GetAllModulosUseCase {
	return &GetAllModulosUseCase{
		repository: repository,
	}
}

func (uc *GetAllModulosUseCase) Execute() ([]domain.ModuloOption, error) {
	return uc.repository.GetAllModulos()
}