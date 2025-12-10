package errors

import "errors"

var (
	ErrDatabaseQuery      = errors.New("erro ao executar query")
	ErrDatabaseConnection = errors.New("erro de conexão com banco de dados")
	ErrNotFound           = errors.New("registro não encontrado")
	ErrMandatoryName      = errors.New("nome obrigatório")
	ErrNameLenInvalid     = errors.New("quantidade de caracteres inválido (2 à 20)")
)

/* MOVE ERRORS */
var (
	ErrMandatoryType     = errors.New("tipo do movimento é obrigatório")
	ErrMandatoryCategory = errors.New("categoria do movimento é obrigatório")
	ErrMoveAlreadyExists = errors.New("movimento já existe")
	ErrMoveNoExists      = errors.New("movimento não existe no banco de dados")
	ErrMoveListEmpty     = errors.New("lista de movimentos está vazia")
	ErrMoveSameName      = errors.New("já existe um movimento com o mesmo nome")
)

/* CREATURES ERRORS */
var (
	ErrCreatureAlreadyExists      = errors.New("criatura já existe")
	ErrMandatoryIdGeneralOrRegion = errors.New("id geral/id da região é obrigatório")
	ErrNegativeId                 = errors.New("id geral/id da região não pode ser negativo")
	ErrCreatureSameName           = errors.New("já existe uma criatura com o mesmo nome")
	ErrNegativeHeightProhibited   = errors.New("altura não pode ser negativa")
	ErrNegativeWeightProhibited   = errors.New("peso não pode ser negativo")
	ErrCreatureNoExists           = errors.New("criatura não existe no banco de dados")
	ErrCreatureListEmpty          = errors.New("lista de criaturas está vazia")
	ErrMoveLimit                  = errors.New("criatura já tem 8 limites")
)

/* REGION ERRORS */
var (
	ErrRegionbAlreadyExists = errors.New("região já existe")
)
