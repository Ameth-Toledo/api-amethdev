package application

import (
	"AmethToledo/src/likes/domain"
	"AmethToledo/src/likes/domain/entities"
	"errors"
)

type GetLikesByUser struct {
	db domain.ILike
}

func NewGetLikesByUser(db domain.ILike) *GetLikesByUser {
	return &GetLikesByUser{db: db}
}

type UserLikesResponse struct {
	UsuarioID   *int            `json:"usuario_id,omitempty"`
	Fingerprint *string         `json:"fingerprint_hash,omitempty"`
	Likes       []entities.Like `json:"likes"`
	TotalLikes  int             `json:"total_likes"`
}

func (glbu *GetLikesByUser) Execute(usuarioID *int, fingerprintHash *string) (*UserLikesResponse, error) {
	if usuarioID == nil && fingerprintHash == nil {
		return nil, errors.New("debe proporcionar usuario_id o fingerprint_hash")
	}

	if usuarioID != nil && fingerprintHash != nil {
		return nil, errors.New("no se pueden proporcionar tanto usuario_id como fingerprint_hash")
	}

	var likes []entities.Like
	var err error

	if usuarioID != nil {
		likes, err = glbu.db.GetLikesByUser(*usuarioID)
	} else {
		likes, err = glbu.db.GetLikesByFingerprint(*fingerprintHash)
	}

	if err != nil {
		return nil, err
	}

	// Inicializar slice vacío si es nil
	if likes == nil {
		likes = []entities.Like{}
	}

	return &UserLikesResponse{
		UsuarioID:   usuarioID,
		Fingerprint: fingerprintHash,
		Likes:       likes,
		TotalLikes:  len(likes),
	}, nil
}
