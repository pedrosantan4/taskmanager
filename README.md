🚀 Task Manager API - Uma Jornada em Go, PostgreSQL e Docker
Go
PostgreSQL
Docker

🎯 Objetivo do Projeto
Desenvolver uma API RESTful robusta para gerenciamento de tarefas, implementando boas práticas modernas de desenvolvimento backend com Go. O projeto serve como laboratório para explorar:

Conexão segura com PostgreSQL via Docker

Operações CRUD com ORM (GORM)

Arquitetura limpa e escalável

Tratamento profissional de erros

Documentação de endpoints

🧩 O Que Esta API Faz?
Um CRUD completo para gerenciar tasks com os seguintes campos:

go
type Task struct {
    gorm.Model
    Title       string `json:"title"`
    Description string `json:"description"`
    Completed   bool   `json:"completed"`
}
🛠️ Tecnologias Utilizadas
Tecnologia	Função	Destaque
Go 1.20+	Linguagem principal	Performance nativa e concorrência
PostgreSQL	Banco de dados	Relacional e confiável
Docker	Containerização	Ambiente isolado e reproduzível
GORM	ORM	Mapeamento objeto-relacional
Chi Router	Roteamento	Leve e eficiente
Git	Controle de versão	Gestão de código
🏗️ Arquitetura do Projeto
taskmanager/
├── cmd/
│   └── main.go          # Ponto de entrada
├── database/
│   ├── db.go            # Conexão e migração
│   └── seed.go          # Dados iniciais
├── handlers/
│   └── task.go          # Lógica dos endpoints
├── models/
│   └── task.go          # Estrutura de dados
├── router/
│   └── router.go        # Configuração de rotas
├── .env                 # Variáveis sensíveis
├── go.mod               # Dependências
└── README.md            # Você está aqui!
💡 Práticas Aprendidas e Implementadas
1. Dockerização do PostgreSQL
Configuração profissional de container com persistência de dados:

bash
docker run --name pg-taskdb \
  -e POSTGRES_PASSWORD=senhasecreta \
  -p 5432:5432 \
  -v pg_data:/var/lib/postgresql/data \
  -d postgres
2. Conexão Robusta com Banco de Dados
Código resiliente que trata falhas de conexão:

go
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
3. Migrações Automatizadas
Sistema que atualiza o schema do banco:

go
func Migrate() {
    if err := DB.AutoMigrate(&models.Task{}); err != nil {
        log.Fatalf("❌ Database migration failed: %v", err)
    }
    log.Println("✅ Database migrated successfully!")
}
4. Seed de Dados Iniciais
População inteligente do banco para desenvolvimento:

go
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
5. Handlers Profissionais
Padrões de tratamento de requisições e respostas:

go
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

// Helper para respostas JSON
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    json.NewEncoder(w).Encode(payload)
}
6. Roteamento Eficiente
Configuração clara de endpoints com Chi:

go
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
7. Tratamento de Erros Resiliente
Solução para o famoso "Task not found":

go
// No handler GetTaskByID
if err := database.DB.First(&task, id).Error; err != nil {
    if errors.Is(err, gorm.ErrRecordNotFound) {
        respondWithError(w, http.StatusNotFound, "Task não encontrada")
    } else {
        respondWithError(w, http.StatusInternalServerError, "Erro no servidor")
    }
    return
}
8. Documentação de Endpoints
Exemplos práticos usando curl:

bash
# Listar todas tasks
curl http://localhost:8080/tasks

# Criar nova task
curl -X POST -H "Content-Type: application/json" \
  -d '{"title":"Reunião","description":"Com equipe às 15h"}' \
  http://localhost:8080/tasks

# Buscar task específica
curl http://localhost:8080/tasks/1
🚀 Como Executar o Projeto
Pré-requisitos
Go 1.20+

Docker

Git

Passo a Passo
bash
# 1. Clonar repositório
git clone https://github.com/seu-usuario/taskmanager.git
cd taskmanager

# 2. Iniciar PostgreSQL
docker run --name pg-taskdb -e POSTGRES_PASSWORD=senha -p 5432:5432 -d postgres

# 3. Configurar ambiente (Linux/Mac)
cp .env.example .env
nano .env  # Ajuste as variáveis se necessário

# 4. Instalar dependências
go mod tidy

# 5. Executar aplicação
go run cmd/main.go

# 6. Acessar endpoints
curl http://localhost:8080/tasks
🌟 Lições Aprendidas
Docker é essencial para desenvolvimento consistente

GORM acelera operações de banco mas requer entendimento

Separação de conceitos é fundamental para código sustentável

Tratamento de erros deve ser priorizado desde o início

Chi Router oferece equilíbrio perfeito entre simplicidade e poder

📈 Próximos Passos
Implementar autenticação com JWT

Adicionar testes automatizados

Criar sistema de paginação

Implementar cache com Redis

Configurar CI/CD com GitHub Actions

🤝 Contribuição
Contribuições são bem-vindas! Siga o processo:

Fork o repositório

Crie uma branch para sua feature (git checkout -b feature/incrivel)

Commit suas mudanças (git commit -am 'Adiciona feature incrível')

Push para a branch (git push origin feature/incrivel)

Abra um Pull Request

Feito com ❤️ e go build
"O conhecimento compartilhado é o único que realmente cresce" - Provérbio hacker
