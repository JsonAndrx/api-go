package types

type ClientExcel struct {
	Nombre      string `json:"nombre"`
	Apellido    string `json:"apellido"`
	Telefono    string `json:"telefono"`
	Fecha       string `json:"fecha"`
	Descripcion string `json:"descripcion"`
}

type ClientCreateRequest struct {
	UsuarioID   int    `json:"usuario_id" binding:"required"`
	Nombre      string `json:"nombre" binding:"required"`
	Apellido    string `json:"apellido" binding:"required"`
	Telefono    string `json:"telefono" binding:"required"`
	Fecha       string `json:"fecha" binding:"required"`
	Hora		string `json:"hora" binding:"required"`
	Descripcion string `json:"descripcion" binding:"required"`
}
