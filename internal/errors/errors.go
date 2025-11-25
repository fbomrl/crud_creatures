package errors

import "errors"

var (
	ErrDatabaseQuery      = errors.New("erro ao executar query")
	ErrDatabaseConnection = errors.New("erro de conexão com banco de dados")
	ErrNotFound           = errors.New("registro não encontrado")
	ErrMandatoryName      = errors.New("nome obrigatório")
)

/* MOVE ERRORS */
var (
	ErrMandatoryType     = errors.New("tipo do movimento é obrigatório")
	ErrMandatoryCategory = errors.New("categoria do movimento é obrigatório")
	ErrMoveAlreadyExists = errors.New("movimento já existe")
)

/* CREATURES ERRORS */
var (
	ErrCreatureAlreadyExists = errors.New("criatura já existe")
)

/* REGION ERRORS */
var (
	ErrRegionbAlreadyExists = errors.New("região já existe")
)
