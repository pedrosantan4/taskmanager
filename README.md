
# 🚀 Task Manager API - Uma Jornada em Go, PostgreSQL e Docker

**Tecnologias:**
`Go` • `PostgreSQL` • `Docker`

---

## 🎯 Objetivo do Projeto

Desenvolver uma **API RESTful robusta** para gerenciamento de tarefas, implementando boas práticas modernas de desenvolvimento backend com Go. Este projeto serve como um laboratório para explorar:

- ✅ Conexão segura com PostgreSQL via Docker  
- ✅ Operações CRUD utilizando ORM (GORM)  
- ✅ Arquitetura limpa e escalável  
- ✅ Tratamento profissional de erros  
- ✅ Documentação clara dos endpoints  

---

## 🧩 O Que Esta API Faz?

Um CRUD completo para gerenciar tarefas com os seguintes campos:

```go
type Task struct {
    gorm.Model
    Title       string `json:"title"`
    Description string `json:"description"`
    Completed   bool   `json:"completed"`
}
```

---

## 🛠️ Tecnologias Utilizadas

| Tecnologia   | Função                  | Destaque                                 |
|--------------|--------------------------|-------------------------------------------|
| **Go 1.20+** | Linguagem principal      | Performance nativa e concorrência leve    |
| **PostgreSQL** | Banco de dados         | Relacional, confiável e robusto           |
| **Docker**    | Containerização         | Ambiente isolado e reproduzível           |
| **GORM**      | ORM                     | Mapeamento objeto-relacional elegante     |
| **Chi Router**| Roteamento              | Leve, modular e altamente performático    |
| **Git**       | Controle de versão      | Gestão eficiente do código-fonte          |

---

## 🏗️ Arquitetura do Projeto

```
taskmanager/
├── cmd/
│   └── main.go          # Ponto de entrada da aplicação
├── database/
│   ├── db.go            # Conexão e migração do banco
│   └── seed.go          # População inicial de dados
├── handlers/
│   └── task.go          # Lógica dos endpoints da API
├── models/
│   └── task.go          # Estrutura de dados (Task)
├── router/
│   └── router.go        # Definição e organização de rotas
├── .env                 # Variáveis de ambiente sensíveis
├── go.mod               # Gerenciamento de dependências
└── README.md            # Documentação do projeto
```

---

## 💡 Práticas Aprendidas e Implementadas

### 1. Dockerização do PostgreSQL

```bash
docker run --name pg-taskdb \
  -e POSTGRES_PASSWORD=senhasecreta \
  -p 5432:5432 \
  -v pg_data:/var/lib/postgresql/data \
  -d postgres
```

### 2. Conexão Robusta com o Banco de Dados

```go
func Connect() {
    dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        os.Getenv("DB_HOST"), os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"),
        os.Getenv("DB_PORT"),
    )

    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("❌ Failed to connect to database: %v", err)
    }
    log.Println("✅ Database connected!")
}
```

### 3. Migrações Automatizadas

```go
func Migrate() {
    if err := DB.AutoMigrate(&models.Task{}); err != nil {
        log.Fatalf("❌ Database migration failed: %v", err)
    }
    log.Println("✅ Database migrated successfully!")
}
```

### 4. Seed de Dados Iniciais

```go
func Seed() {
    var count int64
    DB.Model(&models.Task{}).Count(&count)

    if count == 0 {
        tasks := []models.Task{
            {Title: "Configurar API", Description: "Criar endpoints REST", Completed: false},
            {Title: "Conectar Frontend", Description: "Integrar com React", Completed: false},
        }

        DB.Create(&tasks)
        log.Println("🌱 Database seeded with initial data!")
    }
}
```

### 5. Handlers Profissionais

```go
func GetTaskByID(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(chi.URLParam(r, "id"))
    if err != nil {
        respondWithError(w, http.StatusBadRequest, "ID inválido")
        return
    }

    var task models.Task
    if err := database.DB.First(&task, id).Error; err != nil {
        respondWithError(w, http.StatusNotFound, "Task não encontrada")
        return
    }

    respondWithJSON(w, http.StatusOK, task)
}
```

**Helper para respostas JSON:**

```go
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    json.NewEncoder(w).Encode(payload)
}
```

### 6. Roteamento Eficiente

```go
func SetupRouter() *chi.Mux {
    r := chi.NewRouter()

    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)

    r.Route("/tasks", func(r chi.Router) {
        r.Get("/", handlers.GetAllTasks)       // GET /tasks
        r.Post("/", handlers.CreateTask)       // POST /tasks
        r.Get("/{id}", handlers.GetTaskByID)   // GET /tasks/{id}
        r.Put("/{id}", handlers.UpdateTask)    // PUT /tasks/{id}
        r.Delete("/{id}", handlers.DeleteTask) // DELETE /tasks/{id}
    })

    return r
}
```

### 7. Tratamento de Erros Resiliente

```go
if err := database.DB.First(&task, id).Error; err != nil {
    if errors.Is(err, gorm.ErrRecordNotFound) {
        respondWithError(w, http.StatusNotFound, "Task não encontrada")
    } else {
        respondWithError(w, http.StatusInternalServerError, "Erro no servidor")
    }
    return
}
```

### 8. Documentação de Endpoints

```bash
# Listar todas as tarefas
curl http://localhost:8080/tasks

# Criar nova tarefa
curl -X POST -H "Content-Type: application/json" \
  -d '{"title":"Reunião","description":"Com equipe às 15h"}' \
  http://localhost:8080/tasks

# Buscar tarefa específica
curl http://localhost:8080/tasks/1
```

---

## 🚀 Como Executar o Projeto

### 🔧 Pré-requisitos

- Go 1.20+
- Docker
- Git

### 📋 Passo a Passo

```bash
git clone https://github.com/seu-usuario/taskmanager.git
cd taskmanager

docker run --name pg-taskdb -e POSTGRES_PASSWORD=senha -p 5432:5432 -d postgres

cp .env.example .env
nano .env

go mod tidy
go run cmd/main.go

curl http://localhost:8080/tasks
```

---

## 🌟 Lições Aprendidas

- Docker é essencial para desenvolvimento consistente  
- GORM acelera o desenvolvimento, mas exige bom entendimento  
- Separação de responsabilidades melhora a manutenção  
- Tratamento de erros deve ser prioridade desde o início  
- Chi Router oferece o equilíbrio ideal entre simplicidade e poder  

---

## 📈 Próximos Passos

- 🔐 Implementar autenticação com JWT  
- ✅ Adicionar testes automatizados  
- 📄 Criar sistema de paginação  
- ⚡ Implementar cache com Redis  
- 🔁 Configurar CI/CD com GitHub Actions  

---

## 🤝 Contribuição

Contribuições são **muito bem-vindas**!

```bash
# Etapas para contribuir:
1. Fork do repositório
2. Criar branch: git checkout -b feature/sua-feature
3. Commit: git commit -am 'Nova feature'
4. Push: git push origin feature/sua-feature
5. Pull Request
```

---

> Feito com ❤️ e `go build`  
>  
> **"O conhecimento compartilhado é o único que realmente cresce" – Provérbio hacker**
