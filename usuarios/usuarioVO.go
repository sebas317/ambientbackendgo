package usuarios

type usuarioVO struct {
	usuaID            int
	usuaUsuario       string
	usuaContrasena    string
	usuaUltimoAcceso  string
	usuaUltimaID      string
	estaID            int
	usuaRegistradoPor int
	usuaFechaCambio   string
	usuaImage         string
	pegeID            int
	funcionalidadVO
}
