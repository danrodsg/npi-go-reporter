# NPI-Go-Reporter

> Uma ferramenta de automação para Engenharia de Introdução de Novos Produtos (NPI), desenvolvida em **Go (Golang)**.

![Go Version](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go)
![Status](https://img.shields.io/badge/Status-Functional-success)

## 📋 Sobre o Projeto

O **NPI-Go-Reporter** foi desenvolvido para atender às necessidades de um escritório de projetos (PMO) focado em engenharia de produtos. O objetivo da ferramenta é automatizar a consolidação de dados técnicos e a geração de relatórios de status para reuniões de governança.

O sistema simula a coleta de dados de múltiplos setores (Engenharia, Qualidade, Financeiro), processa regras de negócio para KPIs e gera um *dashboard* visual para tomada de decisão.

### 🎯 Funcionalidades Principais

* **Coleta de Dados:** Leitura e estruturação de dados brutos (JSON) simulando integração com sistemas de gestão.
* **Análise Inteligente de KPIs (Business Logic):**
    * O sistema não apenas exibe dados, mas **interpreta** se o resultado é positivo ou negativo.
    * *Exemplo:* Para "Orçamento", valores acima da meta geram alertas. Para "Qualidade", valores abaixo da meta geram alertas.
* **Geração de Relatórios (Templating):** Renderização automática de relatórios em HTML limpo e responsivo.
* **Servidor Web (Live Reload):** O projeto roda em um servidor HTTP local, permitindo atualização em tempo real dos dados sem reiniciar a aplicação.

---

## 🛠️ Tecnologias Utilizadas

* **Linguagem:** Go (Golang)
* **Bibliotecas Padrão:** `net/http` (Servidor Web), `html/template` (Motor de Renderização), `encoding/json` (Parsing de Dados).
* **Frontend:** HTML5 & CSS3 (Design responsivo para relatórios).
* **Dados:** JSON (Simulação de banco de dados NoSQL/API Response).

---

## 📂 Estrutura do Projeto

```text
npi-go-reporter/
├── main.go           # O cérebro da aplicação (Servidor, Lógica de KPIs e Renderização)
├── dados.json        # Fonte de dados (Simula o input das áreas de engenharia)
├── template.html     # O modelo visual do relatório
└── README.md         # Documentação do projeto


