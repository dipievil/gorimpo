<div align="center">
  <img src="docs/gorimpo-logo.png" alt="GOrimpo Logo" width="200">
  <h1>⛏️ GOrimpo</h1>
  <p><strong>Garimpe as melhores ofertas de retro-gaming na OLX em tempo real.</strong></p>

  <p>
    <img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go">
    <img src="https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white" alt="Docker">
    <img src="https://img.shields.io/badge/SQLite-07405E?style=for-the-badge&logo=sqlite&logoColor=white" alt="SQLite">
    <img src="https://img.shields.io/badge/Playwright-2EAD33?style=for-the-badge&logo=playwright&logoColor=white" alt="Playwright">
   <img src="https://img.shields.io/github/v/release/LXSCA7/gorimpo?style=for-the-badge&color=00ADD8">
   <img src="https://img.shields.io/github/actions/workflow/status/LXSCA7/gorimpo/docker.yml?style=for-the-badge&label=BUILD" />
  </p>
</div>

---

## 📖 Sobre o Projeto

O **GOrimpo** é um motor de busca automatizado (scraper) desenvolvido em Go, focado em encontrar itens específicos no mercado de usados. Ele foi projetado para rodar 24/7 em uma VPS, monitorando termos de busca e notificando instantaneamente via Telegram sempre que uma nova oferta que se encaixe nos seus filtros de preço for detectada.

## ✨ Funcionalidades

* **Busca em Tempo Real:** Utiliza Playwright para renderizar a OLX e contornar proteções básicas.
* **Filtros Inteligentes:** Suporte a `min_price` e `max_price` por termo de busca.
* **Organização por Tópicos:** Cria automaticamente tópicos no Telegram para cada categoria (Console, Jogos, System).
* **Persistência Local:** Banco de dados SQLite garante que você nunca receba a mesma oferta duas vezes.
* **CI/CD nativo:** Deploy automatizado via GitHub Actions e auto-update com Watchtower.
* **Changelog no Bot:** Notifica no grupo sempre que o sistema é atualizado para uma nova versão.

## 🛠️ Arquitetura

O projeto segue os princípios da **Arquitetura Hexagonal**, separando a lógica de negócio (Core) dos detalhes de infraestrutura (Adapters como Telegram e OLX).

## 🚀 Como Rodar

### 1. Preparação
Crie um arquivo `.env` na raiz do projeto com suas credenciais:

```env
TELEGRAM_TOKEN=seu_token_aqui
TELEGRAM_CHAT_ID=seu_chat_id_aqui
```

### 2. Configuração (config.yaml)
Defina o que você quer garimpar:

```yaml
app:
  use_topics: true
  categories:
    - "🕹️ nintendo"
    - "🎮 playstation"

searches:
  - term: "Gameboy Color"
    category: "🕹️ nintendo"
    min_price: 200
    max_price: 500
```

### 3. Deploy com Docker
Para rodar em produção (VPS):

```bash
docker-compose up -d
```

Ou utilize o **Makefile** para testes locais:

```bash
make docker-build
make docker-run
```

## 📂 Estrutura de Pastas

* `/cmd`: Ponto de entrada da aplicação.
* `/internal/core`: Regras de negócio e interfaces (Ports).
* `/internal/adapters`: Implementações externas (Notifier, Scraper, Repo).
* `/internal/config`: Carregamento de YAML e variáveis de ambiente.

---
<p align="center">
Desenvolvido por <a href="https://github.com/LXSCA7">LXSCA</a> ⭐️ <br>
</p>