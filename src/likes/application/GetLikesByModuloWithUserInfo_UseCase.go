package application

import "AmethToledo/src/likes/domain"

type GetLikesByModuloWithUserInfoUseCase struct {
	repository domain.ILike
}

func NewGetLikesByModuloWithUserInfoUseCase(repository domain.ILike) *GetLikesByModuloWithUserInfoUseCase {
	return &GetLikesByModuloWithUserInfoUseCase{
		repository: repository,
	}
}

func (uc *GetLikesByModuloWithUserInfoUseCase) Execute(moduloID int) ([]domain.LikeWithUserInfo, error) {
	return uc.repository.GetByModuloWithUserInfo(moduloID)
}