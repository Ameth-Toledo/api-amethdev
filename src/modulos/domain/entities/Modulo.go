package entities

type Modulo struct {
	ID            int    `json:"id" db:"id"`
	IdCurso       int    `json:"id_curso" db:"id_curso"`
	ImagenPortada string `json:"imagen_portada" db:"imagen_portada"`
	Titulo        string `json:"titulo" db:"titulo"`
	Descripcion   string `json:"descripcion" db:"descripcion"`
}
