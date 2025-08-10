package application

import (
	"AmethToledo/src/likes/domain"
	"AmethToledo/src/likes/domain/entities"
	"errors"
)

type ToggleLike struct {
	db domain.ILike
}

func NewToggleLike(db domain.ILike) *ToggleLike {
	return &ToggleLike{db: db}
}

type ToggleLikeRequest struct {
	ModuloID        int     `json:"modulo_id"`
	UsuarioID       *int    `json:"usuario_id,omitempty"`
	FingerprintHash *string `json:"fingerprint_hash,omitempty"`
}

type ToggleLikeResponse struct {
	Action    string `json:"action"` // "liked" o "unliked"
	LikeCount int    `json:"like_count"`
	Message   string `json:"message"`
}

func (tl *ToggleLike) Execute(req ToggleLikeRequest) (*ToggleLikeResponse, error) {
	if req.ModuloID <= 0 {
		return nil, errors.New("el id del módulo es obligatorio y debe ser válido")
	}

	// Validar que se proporcione al menos usuario_id o fingerprint_hash
	if req.UsuarioID == nil && req.FingerprintHash == nil {
		return nil, errors.New("debe proporcionar usuario_id o fingerprint_hash")
	}

	// Validar que no se proporcionen ambos
	if req.UsuarioID != nil && req.FingerprintHash != nil {
		return nil, errors.New("no se pueden proporcionar tanto usuario_id como fingerprint_hash")
	}

	var action string
	var err error

	// Manejar like por usuario autenticado
	if req.UsuarioID != nil {
		exists, checkErr := tl.db.CheckIfUserLiked(req.ModuloID, *req.UsuarioID)
		if checkErr != nil {
			return nil, checkErr
		}

		if exists {
			// Unlike - eliminar like existente
			err = tl.db.DeleteByModuloAndUser(req.ModuloID, *req.UsuarioID)
			if err != nil {
				return nil, err
			}
			action = "unliked"
		} else {
			// Like - crear nuevo like
			like := entities.Like{
				ModuloID:  req.ModuloID,
				UsuarioID: req.UsuarioID,
			}
			_, err = tl.db.Save(like)
			if err != nil {
				return nil, err
			}
			action = "liked"
		}
	} else {
		// Manejar like por fingerprint (usuario anónimo)
		exists, checkErr := tl.db.CheckIfFingerprintLiked(req.ModuloID, *req.FingerprintHash)
		if checkErr != nil {
			return nil, checkErr
		}

		if exists {
			// Unlike - eliminar like existente
			err = tl.db.DeleteByModuloAndFingerprint(req.ModuloID, *req.FingerprintHash)
			if err != nil {
				return nil, err
			}
			action = "unliked"
		} else {
			// Like - crear nuevo like
			like := entities.Like{
				ModuloID:        req.ModuloID,
				FingerprintHash: req.FingerprintHash,
			}
			_, err = tl.db.Save(like)
			if err != nil {
				return nil, err
			}
			action = "liked"
		}
	}

	// Obtener el conteo actualizado de likes
	likeCount, err := tl.db.GetLikeCount(req.ModuloID)
	if err != nil {
		return nil, err
	}

	var message string
	if action == "liked" {
		message = "Like agregado exitosamente"
	} else {
		message = "Like eliminado exitosamente"
	}

	return &ToggleLikeResponse{
		Action:    action,
		LikeCount: likeCount,
		Message:   message,
	}, nil
}
