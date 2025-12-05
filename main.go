package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
)

// --- MANTEMOS AS MESMAS STRUCTS E LÓGICA ---

type KPI struct {
	Nome            string  `json:"nome"`
	ValorAtual      float64 `json:"valor_atual"`
	Meta            float64 `json:"meta"`
	Unidade         string  `json:"unidade"`
	StatusTexto     string  `json:"-"`
	StatusClasseCSS string  `json:"-"`
}

func (k *KPI) CalcularStatus() {
	if strings.Contains(k.Nome, "Gasto") || strings.Contains(k.Nome, "Orçamento") {
		if k.ValorAtual > k.Meta {
			k.StatusTexto = "⚠️ Estourou o Orçamento!"
			k.StatusClasseCSS = "status-alert"
		} else {
			k.StatusTexto = "✅ Dentro do Orçamento"
			k.StatusClasseCSS = "status-ok"
		}
	} else {
		if k.ValorAtual < k.Meta {
			k.StatusTexto = "⚠️ Abaixo da Meta"
			k.StatusClasseCSS = "status-alert"
		} else {
			k.StatusTexto = "✅ Meta Batida"
			k.StatusClasseCSS = "status-ok"
		}
	}
}

type Entregavel struct {
	Nome   string `json:"nome"`
	Status string `json:"status"`
}

type DadosProjeto struct {
	NomeProjeto   string       `json:"nome_projeto"`
	Gerente       string       `json:"gerente"`
	DataRelatorio string       `json:"data_relatorio"`
	Kpis          []KPI        `json:"kpis"`
	Entregaveis   []Entregavel `json:"entregaveis"`
}

// --- AQUI MUDA TUDO: O HANDLER DO SERVIDOR ---

func renderizarRelatorio(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Recebi uma requisição! Processando dados atualizados...")

	// 1. Ler o JSON (Toda vez que alguém acessa a página)
	arquivoJSON, err := os.ReadFile("dados.json")
	if err != nil {
		http.Error(w, "Erro ao ler dados.json", http.StatusInternalServerError)
		return
	}

	var projeto DadosProjeto
	json.Unmarshal(arquivoJSON, &projeto)

	// 2. Processar Lógica
	for i := range projeto.Kpis {
		projeto.Kpis[i].CalcularStatus()
	}

	// 3. Renderizar direto no navegador (ResponseWriter)
	tpl, err := template.ParseFiles("template.html")
	if err != nil {
		http.Error(w, "Erro ao ler template", http.StatusInternalServerError)
		return
	}

	tpl.Execute(w, projeto)
}

func main() {
	// Define a rota "/" para chamar nossa função
	http.HandleFunc("/", renderizarRelatorio)

	fmt.Println("🌐 Servidor rodando em http://localhost:8080")
	fmt.Println("Pressione Ctrl+C para parar.")
	
	// Inicia o servidor
	http.ListenAndServe(":8080", nil)
}