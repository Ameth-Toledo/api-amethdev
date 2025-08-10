package entities

type Contenido struct {
	ID                int    `json:"id" db:"id"`
	IdModulo          int    `json:"id_modulo" db:"id_modulo"`
	ImagenPortada     string `json:"imagen_portada" db:"imagen_portada"`
	Titulo            string `json:"titulo" db:"titulo"`
	Descripcion       string `json:"descripcion" db:"descripcion"`
	VideoURL          string `json:"video_url" db:"video_url"`
	DescripcionModule string `json:"descripcion_module" db:"descripcion_module"`
	Repositorio       string `json:"repositorio" db:"repositorio"`
}
