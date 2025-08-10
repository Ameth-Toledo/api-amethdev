package domain

import "AmethToledo/src/comentarios/domain/entities"

type IComentario interface {
	Save(comentario entities.Comentario) (entities.Comentario, error)
	GetAll() ([]entities.Comentario, error)
	GetById(id int) (*entities.Comentario, error)
	GetByModuloId(moduloId int) ([]entities.Comentario, error)
	GetByUsuarioId(usuarioId int) ([]entities.Comentario, error)
	GetTotal() (int, error)
	Update(comentario entities.Comentario) error
	Delete(id int) error
	GetByModuloIdWithUser(moduloId int) ([]entities.ComentarioConUsuario, error)
	GetAllWithUser() ([]entities.ComentarioConUsuario, error)
}
