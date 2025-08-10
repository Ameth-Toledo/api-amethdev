package entities

type Curso struct {
	ID          int    `json:"id" db:"id"`
	Nombre      string `json:"nombre" db:"nombre"`
	Nivel       string `json:"nivel" db:"nivel"`
	Duracion    string `json:"duracion" db:"duracion"`
	Tecnologia  string `json:"tecnologia" db:"tecnologia"`
	Fecha       string `json:"fecha" db:"fecha"`
	Imagen      string `json:"imagen" db:"imagen"`
	Descripcion string `json:"descripcion" db:"descripcion"`
}
