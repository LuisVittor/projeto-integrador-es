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
    db = fazConexaoComBanco() // Estabelece a conexão com o banco de dados
    templates = template.Must(template.ParseGlob("templates/*.html")) // Carrega todos os templates HTML

    // Configuração do servidor para servir arquivos estáticos a partir da pasta "static"
    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    // Rotas da aplicação
    http.HandleFunc("/tela1", tela1)
    http.HandleFunc("/tela2", tela2)
    http.HandleFunc("/tela3", tela3)
    http.HandleFunc("/login", login)
    http.HandleFunc("/inicio", inicio)
    http.HandleFunc("/index", index)
    http.HandleFunc("/index2", index2)
    http.HandleFunc("/index3", index3)
    http.HandleFunc("/index4", index4)
    http.HandleFunc("/index5", index5)
    http.HandleFunc("/index7", index7)
    http.HandleFunc("/index8", index8)
    http.HandleFunc("/index10", index10)
    http.HandleFunc("/index11", index11)
    http.HandleFunc("/index20", index20)
    http.HandleFunc("/index21", index21)
    http.HandleFunc("/index22", index22)
    http.HandleFunc("/index23", index23)
    http.HandleFunc("/index24", index24)
    http.HandleFunc("/index25", index25)
    http.HandleFunc("/index26", index26)
    http.HandleFunc("/menu", menu)
    http.HandleFunc("/pacientes", pacientes) // Rota para listar pacientes
    http.HandleFunc("/adicionar-paciente", adicionarPaciente) // Rota para adicionar paciente
    http.HandleFunc("/cadastro-paciente", renderCadastroPaciente) // Rota para renderizar a página de cadastro
    http.HandleFunc("/deletar-paciente", deletarPaciente) // Rota para deletar paciente
    http.HandleFunc("/editar-paciente", renderEditarPaciente) // Rota para renderizar a página de edição
    http.HandleFunc("/atualizar-paciente", editarPaciente) // Rota para atualizar paciente

    // Alimentar o banco de dados com dados iniciais
    alimentaBancoDeDados()

    // Iniciar o servidor
    log.Println("Server rodando na porta 8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal(err)
    }
}

// Função para listar pacientes
func pacientes(w http.ResponseWriter, r *http.Request) {
    err := r.ParseForm()
    if err != nil {
        http.Error(w, "Erro ao processar formulário", http.StatusBadRequest)
        return
    }
    busca := strings.TrimSpace(r.Form.Get("busca"))
    filtroFumante := r.Form.Get("filtro_fumante") == "on"
    filtroBebe := r.Form.Get("filtro_bebe") == "on"
    filtroRua := r.Form.Get("filtro_rua") == "on"

    log.Println("Buscando por:", busca, "Fumante:", filtroFumante, "Bebe:", filtroBebe, "Rua:", filtroRua)

    pacientes := buscaPacientePorNome(busca, filtroFumante, filtroBebe, filtroRua)

    err = templates.ExecuteTemplate(w, "pacientes.html", pacientes)
    if err != nil {
        http.Error(w, "Erro ao renderizar template", http.StatusInternalServerError)
        log.Println("Erro ao renderizar template:", err)
    }
}

// Função para renderizar a tela de menu
func inicio(w http.ResponseWriter, r *http.Request) {
    err := templates.ExecuteTemplate(w, "inicio.html", nil)
    if err != nil {
        http.Error(w, "Erro ao renderizar template", http.StatusInternalServerError)
        log.Println("Erro ao renderizar template:", err)
    }
}

func login(w http.ResponseWriter, r *http.Request) {
    err := templates.ExecuteTemplate(w, "login.html", nil)
    if err != nil {
        http.Error(w, "Erro ao renderizar template", http.StatusInternalServerError)
        log.Println("Erro ao renderizar template:", err)
    }
}

func tela1(w http.ResponseWriter, r *http.Request) {
    err := templates.ExecuteTemplate(w, "tela1.html", nil)
    if err != nil {
        http.Error(w, "Erro ao renderizar template", http.StatusInternalServerError)
        log.Println("Erro ao renderizar template:", err)
    }
}

func tela2(w http.ResponseWriter, r *http.Request) {
    err := templates.ExecuteTemplate(w, "tela2.html", nil)
    if err != nil {
        http.Error(w, "Erro ao renderizar template", http.StatusInternalServerError)
        log.Println("Erro ao renderizar template:", err)
    }
}

func tela3(w http.ResponseWriter, r *http.Request) {
    err := templates.ExecuteTemplate(w, "tela3.html", nil)
    if err != nil {
        http.Error(w, "Erro ao renderizar template", http.StatusInternalServerError)
        log.Println("Erro ao renderizar template:", err)
    }
}


// Função para renderizar a tela de menu
func menu(w http.ResponseWriter, r *http.Request) {
    err := templates.ExecuteTemplate(w, "menu.html", nil)
    if err != nil {
        http.Error(w, "Erro ao renderizar template", http.StatusInternalServerError)
        log.Println("Erro ao renderizar template:", err)
    }
}

func index(w http.ResponseWriter, r *http.Request) {
    err := templates.ExecuteTemplate(w, "index.html", nil)
    if err != nil {
        http.Error(w, "Erro ao renderizar template", http.StatusInternalServerError)
        log.Println("Erro ao renderizar template:", err)
    }
}

    func index2(w http.ResponseWriter, r *http.Request) {
        err := templates.ExecuteTemplate(w, "index2.html", nil)
        if err != nil {
            http.Error(w, "Erro ao renderizar template", http.StatusInternalServerError)
            log.Println("Erro ao renderizar template:", err)
        }
}

func index3(w http.ResponseWriter, r *http.Request) {
    err := templates.ExecuteTemplate(w, "index3.html", nil)
    if err != nil {
        http.Error(w, "Erro ao renderizar template", http.StatusInternalServerError)
        log.Println("Erro ao renderizar template:", err)
    }
}

func index4(w http.ResponseWriter, r *http.Request) {
    err := templates.ExecuteTemplate(w, "index4.html", nil)
    if err != nil {
        http.Error(w, "Erro ao renderizar template", http.StatusInternalServerError)
        log.Println("Erro ao renderizar template:", err)
    }
}

func index5(w http.ResponseWriter, r *http.Request) {
    err := templates.ExecuteTemplate(w, "index5.html", nil)
    if err != nil {
        http.Error(w, "Erro ao renderizar template", http.StatusInternalServerError)
        log.Println("Erro ao renderizar template:", err)
    }
}

func index7(w http.ResponseWriter, r *http.Request) {
    err := templates.ExecuteTemplate(w, "index7.html", nil)
    if err != nil {
        http.Error(w, "Erro ao renderizar template", http.StatusInternalServerError)
        log.Println("Erro ao renderizar template:", err)
    }
}

func index8(w http.ResponseWriter, r *http.Request) {
    err := templates.ExecuteTemplate(w, "index8.html", nil)
    if err != nil {
        http.Error(w, "Erro ao renderizar template", http.StatusInternalServerError)
        log.Println("Erro ao renderizar template:", err)
    }
}

func index10(w http.ResponseWriter, r *http.Request) {
    err := templates.ExecuteTemplate(w, "index10.html", nil)
    if err != nil {
        http.Error(w, "Erro ao renderizar template", http.StatusInternalServerError)
        log.Println("Erro ao renderizar template:", err)
    }
}

func index11(w http.ResponseWriter, r *http.Request) {
    err := templates.ExecuteTemplate(w, "index11.html", nil)
    if err != nil {
        http.Error(w, "Erro ao renderizar template", http.StatusInternalServerError)
        log.Println("Erro ao renderizar template:", err)
    }
}


func index20(w http.ResponseWriter, r *http.Request) {
    err := templates.ExecuteTemplate(w, "index20.html", nil)
    if err != nil {
        http.Error(w, "Erro ao renderizar template", http.StatusInternalServerError)
        log.Println("Erro ao renderizar template:", err)
    }
}

func index21(w http.ResponseWriter, r *http.Request) {
    err := templates.ExecuteTemplate(w, "index21.html", nil)
    if err != nil {
        http.Error(w, "Erro ao renderizar template", http.StatusInternalServerError)
        log.Println("Erro ao renderizar template:", err)
    }
}

func index22(w http.ResponseWriter, r *http.Request) {
    err := templates.ExecuteTemplate(w, "index22.html", nil)
    if err != nil {
        http.Error(w, "Erro ao renderizar template", http.StatusInternalServerError)
        log.Println("Erro ao renderizar template:", err)
    }
}

func index23(w http.ResponseWriter, r *http.Request) {
    err := templates.ExecuteTemplate(w, "index23.html", nil)
    if err != nil {
        http.Error(w, "Erro ao renderizar template", http.StatusInternalServerError)
        log.Println("Erro ao renderizar template:", err)
    }
}

func index24(w http.ResponseWriter, r *http.Request) {
    err := templates.ExecuteTemplate(w, "index24.html", nil)
    if err != nil {
        http.Error(w, "Erro ao renderizar template", http.StatusInternalServerError)
        log.Println("Erro ao renderizar template:", err)
    }
}

func index25(w http.ResponseWriter, r *http.Request) {
    err := templates.ExecuteTemplate(w, "index25.html", nil)
    if err != nil {
        http.Error(w, "Erro ao renderizar template", http.StatusInternalServerError)
        log.Println("Erro ao renderizar template:", err)
    }
}

func index26(w http.ResponseWriter, r *http.Request) {
    err := templates.ExecuteTemplate(w, "index26.html", nil)
    if err != nil {
        http.Error(w, "Erro ao renderizar template", http.StatusInternalServerError)
        log.Println("Erro ao renderizar template:", err)
    }
}
// Função para adicionar um paciente
func adicionarPaciente(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
        return
    }

    nome := r.FormValue("nome")
    cpf := r.FormValue("cpf")
    dataNascimento := r.FormValue("data_nasc")
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
    http.Redirect(w, r, "/pacientes", http.StatusSeeOther) // Redireciona para a lista de pacientes após adicionar
}

// Função para renderizar a página de cadastro de paciente
func renderCadastroPaciente(w http.ResponseWriter, r *http.Request) {
    err := templates.ExecuteTemplate(w, "cadastro_paciente.html", nil)
    if err != nil {
        http.Error(w, "Erro ao renderizar template", http.StatusInternalServerError)
        log.Println("Erro ao renderizar template:", err)
    }
}

// Função para deletar um paciente
func deletarPaciente(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
        return
    }

    cpf := r.FormValue("cpf")
    if cpf == "" {
        http.Error(w, "CPF não fornecido", http.StatusBadRequest)
        return
    }

    _, err := db.Exec("DELETE FROM paciente WHERE cpf = $1", cpf)
    if err != nil {
        http.Error(w, "Erro ao deletar paciente", http.StatusInternalServerError)
        log.Println("Erro ao deletar paciente:", err)
        return
    }

    http.Redirect(w, r, "/pacientes", http.StatusSeeOther) // Redireciona para a lista de pacientes após deletar
}

// Função para editar um paciente
func editarPaciente(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
        return
    }

    cpf := r.FormValue("cpf")
    if cpf == "" {
        http.Error(w, "CPF não fornecido", http.StatusBadRequest)
        return
    }

    nome := r.FormValue("nome")
    dataNascimento := r.FormValue("data_nasc")
    telefone := r.FormValue("telefone")
    sexo := r.FormValue("sexo")
    estaFumante := r.FormValue("esta_fumante") == "on"
    fazUsoAlcool := r.FormValue("faz_uso_alcool") == "on"
    estaSituacaoRua := r.FormValue("esta_situacao_rua") == "on"
    endereco := r.FormValue("endereco")
    nomeMae := r.FormValue("nome_mae")

    _, err := db.Exec(`
        UPDATE paciente 
        SET nome = $1, data_nascimento = $2, telefone = $3, sexo = $4, esta_fumante = $5, faz_uso_alcool = $6, esta_situacao_rua = $7, endereco = $8, nome_mae = $9
        WHERE cpf = $10
    `, nome, dataNascimento, telefone, sexo, estaFumante, fazUsoAlcool, estaSituacaoRua, endereco, nomeMae, cpf)
    if err != nil {
        http.Error(w, "Erro ao atualizar paciente", http.StatusInternalServerError)
        log.Println("Erro ao atualizar paciente:", err)
        return
    }

    http.Redirect(w, r, "/pacientes", http.StatusSeeOther) // Redireciona para a lista de pacientes após editar
}

// Função para renderizar a página de edição de paciente
func renderEditarPaciente(w http.ResponseWriter, r *http.Request) {
    cpf := r.URL.Query().Get("cpf")
    if cpf == "" {
        http.Error(w, "CPF não fornecido", http.StatusBadRequest)
        return
    }

    var paciente Paciente
    err := db.QueryRow(`
        SELECT id, cpf, nome, data_nascimento, telefone, sexo, esta_fumante, faz_uso_alcool, esta_situacao_rua, endereco, nome_mae
        FROM paciente
        WHERE cpf = $1
    `, cpf).Scan(&paciente.Id, &paciente.Cpf, &paciente.Nome, &paciente.DataNascimento, &paciente.Telefone, &paciente.Sexo, &paciente.EstaFumante, &paciente.FazUsoAlcool, &paciente.EstaSituacaoRua, &paciente.Endereco, &paciente.NomeMae)
    if err != nil {
        http.Error(w, "Paciente não encontrado", http.StatusNotFound)
        log.Println("Erro ao buscar paciente:", err)
        return
    }

    err = templates.ExecuteTemplate(w, "editar_paciente.html", paciente)
    if err != nil {
        http.Error(w, "Erro ao renderizar template", http.StatusInternalServerError)
        log.Println("Erro ao renderizar template:", err)
    }
}

// Função para conectar ao banco de dados
func fazConexaoComBanco() *sql.DB {
    err := godotenv.Load() // Carrega variáveis de ambiente do arquivo .env
    if err != nil {
        log.Fatalf("Erro ao carregar arquivo .env")
    }

    usuarioBancoDeDados := os.Getenv("USUARIO")
    senhaDoUsuario := os.Getenv("SENHA")
    nomeDoBancoDeDados := os.Getenv("NOME_BANCO_DE_DADOS")
    dadosParaConexao := fmt.Sprintf("user=%s dbname=%s password=%s host=localhost port=5432 sslmode=disable", usuarioBancoDeDados, nomeDoBancoDeDados, senhaDoUsuario)

    database, err := sql.Open("postgres", dadosParaConexao) // Abre a conexão com o banco de dados
    if err != nil {
        log.Fatal("Erro ao conectar com o banco de dados:", err)
    }

    err = database.Ping() // Verifica se a conexão foi bem-sucedida
    if err != nil {
        log.Fatal("Erro ao pingar o banco de dados:", err)
    }

    // Criação da tabela de pacientes, se não existir
    _, err = database.Exec(`
        CREATE TABLE IF NOT EXISTS paciente (
            id SERIAL PRIMARY KEY,
            nome VARCHAR(255) NOT NULL,
            cpf VARCHAR(15) UNIQUE NOT NULL,
            data_nascimento VARCHAR(12),
            telefone VARCHAR(20),
            sexo VARCHAR(30),
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

// Função para cadastrar um paciente no banco de dados
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

// Função para buscar pacientes pelo nome com filtros
func buscaPacientePorNome(nome string, filtroFumante, filtroBebe, filtroRua bool) Pacientes {
    query := `
        SELECT id, cpf, nome, data_nascimento, telefone, sexo, esta_fumante, faz_uso_alcool, esta_situacao_rua, endereco, nome_mae
        FROM paciente 
        WHERE (nome ILIKE '%' || $1 || '%' OR cpf = $1)
    `
    args := []interface{}{nome}
    index := 2

    if filtroFumante {
        query += " AND esta_fumante = $"+fmt.Sprint(index)
        args = append(args, true)
        index++
    }
    if filtroBebe {
        query += " AND faz_uso_alcool = $"+fmt.Sprint(index)
        args = append(args, true)
        index++
    }
    if filtroRua {
        query += " AND esta_situacao_rua = $"+fmt.Sprint(index)
        args = append(args, true)
        index++
    }

    busca, err := db.Query(query, args...)
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

// Função para alimentar o banco de dados com dados iniciais a partir de um arquivo JSON
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

// Estrutura que representa um paciente
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

// Estrutura que representa uma lista de pacientes
type Pacientes struct {
    Pacientes []Paciente `json:"pacientes"`
}
