# CRUD Creatures API

API REST para gerenciamento de criaturas, incluindo seus atributos, movimentos e regiões.

Projeto de estudos baseado em criaturas originais desenhadas pelo autor, com objetivo de catalogar e gerenciar informações detalhadas de cada criatura.

## 🎯 Sobre o Projeto

Sistema de CRUD desenvolvido em Go para armazenar e gerenciar:
- **Criaturas**: Informações completas incluindo tipos, habilidades, habitat, evolução
- **Movimentos**: Ataques e habilidades que as criaturas podem aprender
- **Regiões**: Localizações onde as criaturas habitam
- **Relações**: Vínculo entre criaturas e seus movimentos

## 🛠️ Tecnologias

- **Language**: Go 1.x
- **Database**: MySQL/MariaDB
- **Architecture**: Clean Architecture
- **Patterns**: Repository Pattern, Service Layer

## 📁 Estrutura do Projeto

```
crud_creatures/
├── cmd/
│   └── api/
│       └── main.go              # Entry point da aplicação
├── internal/
│   ├── sqlserver.go             # Conexão com SQL Server (ou MySQL)
│   ├── errors/
│   │   └── errors.go            # Erros customizados da aplicação
│   ├── handlers/                # HTTP handlers (controllers)
│   │   ├── interface/
│   │   │   ├── creatureInferface.go
│   │   │   └── moveInterface.go
│   ├── models/
│   │   ├── creatureModel.go     # Modelo de Criatura
│   │   ├── moveModel.go         # Modelo de Movimento
│   │   ├── regionModel.go       # Modelo de Região
│   │   ├── creatureMoveModel.go # Relação Criatura-Movimento
│   │   └── enums/
│   │       ├── typeEnum.go      # Enum de Tipos (Fire, Water, etc)
│   │       └── categoryEnum.go  # Enum de Categorias de Movimento
│   ├── repository/
│   │   ├── creatureRepository.go # Acesso a dados de Criaturas
│   │   └── moveRepository.go     # Acesso a dados de Movimentos
│   └── service/
│       ├── creatureService.go    # Regras de negócio - Criaturas
│       └── moveService.go        # Regras de negócio - Movimentos
├── migrations/                   # Scripts de migração do banco
├── pkg/                          # Pacotes reutilizáveis
├── config/                       # Configurações da aplicação
└── go.mod                        # Dependências Go

```

## 🚀 Instalação

### Pré-requisitos

- Go 1.18 ou superior
- MySQL 8.0 ou MariaDB 10.x
- Git

### Setup

1. **Clonar o repositório:**
```bash
git clone <seu-repositorio>
cd crud_creatures
```

2. **Instalar dependências:**
```bash
go mod download
```

3. **Configurar banco de dados:**
```sql
CREATE DATABASE creatures_db;
```

4. **Executar migrations:**
```bash
# (Instruções de migração serão adicionadas)
```

5. **Configurar variáveis de ambiente:**
```bash
# Criar arquivo .env
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=sua_senha
DB_NAME=creatures_db
```

## 🏃 Executando

### Desenvolvimento

```bash
go run cmd/api/main.go
```

### Build

```bash
go build -o bin/creatures-api cmd/api/main.go
```

### Executar build

```bash
./bin/creatures-api
```

## 📊 Modelos de Dados

### Creature (Criatura)
- Informações gerais (nome, espécie, descrição)
- Tipos primário e secundário (Fire, Water, Grass, Electric)
- Habilidades de ataque e defesa com descrições
- Atributos físicos (altura, peso)
- Habitat e alimentação
- Evolução e inspiração

### Move (Movimento)
- Nome e descrição
- Tipo elemental
- Categoria (Physical, Special, Status)
- Poder e precisão

### Region (Região)
- Nome e descrição da região

### CreatureMove (Relação)
- Vínculo entre criaturas e movimentos

## 🎮 Enums

### Types (Tipos Elementais)
- Fire (0)
- Water (1)
- Grass (2)
- Electric (3)

### Categories (Categorias de Movimento)
- Physical
- Special
- Status

## 🔄 Próximos Passos

- [ ] Implementar handlers HTTP
- [ ] Adicionar testes unitários
- [ ] Criar sistema de importação via Excel
- [ ] Implementar autenticação
- [ ] Adicionar documentação Swagger
- [ ] Deploy em produção

## 📝 Scripts Disponíveis

```bash
# Compilar o projeto
go build ./...

# Executar testes
go test ./...

# Verificar código
go vet ./...

# Formatar código
go fmt ./...
```

## 🏗️ Arquitetura

O projeto segue o padrão **Clean Architecture** com separação de responsabilidades:

- **Models**: Estruturas de dados
- **Repository**: Acesso ao banco de dados (SQL queries)
- **Service**: Regras de negócio e validações
- **Handlers**: Controllers HTTP (a implementar)

**Fluxo de dados:**
```
HTTP Request → Handler → Service → Repository → Database
                  ↓
            Validações
```

## 📄 Licença

© 2025 [metanikk.com.br](https://metanikk.com.br) - Todos os direitos reservados

---

**Projeto de Estudos** - Criado para fins educacionais e desenvolvimento de portfólio.

