package application

import (
	"AmethToledo/src/likes/domain"
	"errors"
)

type GetLikeCount struct {
	db domain.ILike
}

func NewGetLikeCount(db domain.ILike) *GetLikeCount {
	return &GetLikeCount{db: db}
}

type LikeCountResponse struct {
	ModuloID  int  `json:"modulo_id"`
	LikeCount int  `json:"like_count"`
	UserLiked bool `json:"user_liked"`
}

func (glc *GetLikeCount) Execute(moduloID int, usuarioID *int, fingerprintHash *string) (*LikeCountResponse, error) {
	if moduloID <= 0 {
		return nil, errors.New("el id del módulo es obligatorio y debe ser válido")
	}

	likeCount, err := glc.db.GetLikeCount(moduloID)
	if err != nil {
		return nil, err
	}

	var userLiked bool
	if usuarioID != nil {
		userLiked, err = glc.db.CheckIfUserLiked(moduloID, *usuarioID)
		if err != nil {
			return nil, err
		}
	} else if fingerprintHash != nil {
		userLiked, err = glc.db.CheckIfFingerprintLiked(moduloID, *fingerprintHash)
		if err != nil {
			return nil, err
		}
	}

	return &LikeCountResponse{
		ModuloID:  moduloID,
		LikeCount: likeCount,
		UserLiked: userLiked,
	}, nil
}
