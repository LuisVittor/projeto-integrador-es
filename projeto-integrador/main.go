package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB
var templates *template.Template

func main() {
	// Configurações iniciais
	db = fazConexaoComBanco()
	templates = template.Must(template.ParseGlob("templates/*.html"))

	// Configuração do servidor para servir arquivos estáticos a partir da pasta "static"
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Rotas da aplicação
	http.HandleFunc("/pacientes", pacientes)
	http.HandleFunc("/adicionar-paciente", adicionarPaciente)
	http.HandleFunc("/cadastro-paciente", renderCadastroPaciente)

	// Alimentar o banco de dados com dados iniciais
	alimentaBancoDeDados()

	// Iniciar o servidor
	log.Println("Server rodando na porta 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func pacientes(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Erro ao processar formulário", http.StatusBadRequest)
		return
	}
	busca := strings.TrimSpace(r.Form.Get("busca"))
	log.Println("Buscando por:", busca)

	pacientes := buscaPacientePorNome(busca)

	err = templates.ExecuteTemplate(w, "pacientes.html", pacientes)
	if err != nil {
		http.Error(w, "Erro ao renderizar template", http.StatusInternalServerError)
		log.Println("Erro ao renderizar template:", err)
	}
}

func adicionarPaciente(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	nome := r.FormValue("nome")
	cpf := r.FormValue("cpf")
	dataNascimento := r.FormValue("data_nascimento")
	telefone := r.FormValue("telefone")
	sexo := r.FormValue("sexo")
	estaFumante := r.FormValue("esta_fumante") == "on"
	fazUsoAlcool := r.FormValue("faz_uso_alcool") == "on"
	estaSituacaoRua := r.FormValue("esta_situacao_rua") == "on"
	endereco := r.FormValue("endereco")
	nomeMae := r.FormValue("nome_mae")

	paciente := Paciente{
		Cpf:               cpf,
		Nome:              nome,
		DataNascimento:    dataNascimento,
		Telefone:          telefone,
		Sexo:              sexo,
		EstaFumante:       estaFumante,
		FazUsoAlcool:      fazUsoAlcool,
		EstaSituacaoRua:   estaSituacaoRua,
		Endereco:          endereco,
		NomeMae:           nomeMae,
	}

	cadastraPaciente(paciente)
	http.Redirect(w, r, "/pacientes", http.StatusSeeOther)
}

func renderCadastroPaciente(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "cadastro_paciente.html", nil)
	if err != nil {
		http.Error(w, "Erro ao renderizar template", http.StatusInternalServerError)
		log.Println("Erro ao renderizar template:", err)
	}
}

func fazConexaoComBanco() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar arquivo .env")
	}

	usuarioBancoDeDados := os.Getenv("USUARIO")
	senhaDoUsuario := os.Getenv("SENHA")
	nomeDoBancoDeDados := os.Getenv("NOME_BANCO_DE_DADOS")
	dadosParaConexao := fmt.Sprintf("user=%s dbname=%s password=%s host=localhost port=5432 sslmode=disable", usuarioBancoDeDados, nomeDoBancoDeDados, senhaDoUsuario)

	database, err := sql.Open("postgres", dadosParaConexao)
	if err != nil {
		log.Fatal("Erro ao conectar com o banco de dados:", err)
	}

	err = database.Ping()
	if err != nil {
		log.Fatal("Erro ao pingar o banco de dados:", err)
	}

	_, err = database.Exec(`
		CREATE TABLE IF NOT EXISTS paciente (
			id SERIAL PRIMARY KEY,
			nome VARCHAR(255) NOT NULL,
			cpf VARCHAR(15) UNIQUE NOT NULL,
			data_nascimento VARCHAR(12),
			telefone VARCHAR(20),
			sexo VARCHAR(10),
			esta_fumante BOOLEAN,
			faz_uso_alcool BOOLEAN,
			esta_situacao_rua BOOLEAN,
			endereco TEXT,
			nome_mae VARCHAR(255)
		)
	`)
	if err != nil {
		log.Fatal("Erro ao criar tabela:", err)
	}

	return database
}

func cadastraPaciente(paciente Paciente) {
	_, err := db.Exec(`
		INSERT INTO paciente (
			nome, cpf, data_nascimento, telefone, sexo, esta_fumante, faz_uso_alcool, esta_situacao_rua, endereco, nome_mae
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		ON CONFLICT (cpf) DO NOTHING
	`, paciente.Nome, paciente.Cpf, paciente.DataNascimento, paciente.Telefone, paciente.Sexo, paciente.EstaFumante, paciente.FazUsoAlcool, paciente.EstaSituacaoRua, paciente.Endereco, paciente.NomeMae)
	if err != nil {
		log.Println("Erro ao cadastrar paciente:", err)
	}
}

func buscaPacientePorNome(nome string) Pacientes {
	busca, err := db.Query(`
		SELECT id, cpf, nome, data_nascimento, telefone, sexo, esta_fumante, faz_uso_alcool, esta_situacao_rua, endereco, nome_mae
		FROM paciente 
		WHERE nome ILIKE '%' || $1 || '%'
		OR cpf = $1
	`, nome)
	if err != nil {
		log.Println("Erro ao buscar pacientes:", err)
		return Pacientes{}
	}
	defer busca.Close()

	var pacientes Pacientes

	for busca.Next() {
		var paciente Paciente
		err = busca.Scan(&paciente.Id, &paciente.Cpf, &paciente.Nome, &paciente.DataNascimento, &paciente.Telefone, &paciente.Sexo, &paciente.EstaFumante, &paciente.FazUsoAlcool, &paciente.EstaSituacaoRua, &paciente.Endereco, &paciente.NomeMae)
		if err != nil {
			log.Println("Erro ao scanear paciente:", err)
			continue
		}
		pacientes.Pacientes = append(pacientes.Pacientes, paciente)
	}
	if err = busca.Err(); err != nil {
		log.Println("Erro no loop de busca:", err)
	}

	return pacientes
}

func alimentaBancoDeDados() {
	var pacientes Pacientes

	jsonFile, err := os.Open("pacientes.json")
	if err != nil {
		log.Fatal("Erro ao abrir arquivo JSON:", err)
	}
	defer jsonFile.Close()

	byteJson, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatal("Erro ao ler arquivo JSON:", err)
	}

	err = json.Unmarshal(byteJson, &pacientes)
	if err != nil {
		log.Fatal("Erro ao fazer unmarshal do JSON:", err)
	}

	for _, paciente := range pacientes.Pacientes {
		cadastraPaciente(paciente)
	}
}

type Paciente struct {
	Id              uint64
	Nome            string `json:"nome"`
	Cpf             string `json:"cpf"`
	DataNascimento  string `json:"data_nasc"`
	Telefone        string `json:"celular"`
	Sexo            string `json:"sexo"`
	EstaFumante     bool   `json:"esta_fumante"`
	FazUsoAlcool    bool   `json:"faz_uso_alcool"`
	EstaSituacaoRua bool   `json:"esta_situacao_de_rua"`
	Endereco        string `json:"endereco"`
	NomeMae         string `json:"nome_mae"`
}

type Pacientes struct {
	Pacientes []Paciente `json:"pacientes"`
}
