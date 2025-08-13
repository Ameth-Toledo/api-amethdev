package application

import "AmethToledo/src/likes/domain"

type GetLikesByUserWithModuleInfoUseCase struct {
	repository domain.ILike
}

func NewGetLikesByUserWithModuleInfoUseCase(repository domain.ILike) *GetLikesByUserWithModuleInfoUseCase {
	return &GetLikesByUserWithModuleInfoUseCase{
		repository: repository,
	}
}

func (uc *GetLikesByUserWithModuleInfoUseCase) Execute(usuarioID int) ([]domain.LikeWithUserInfo, error) {
	return uc.repository.GetLikesByUserWithModuleInfo(usuarioID)
}