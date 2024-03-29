# go-api-boilerplate

Projeto base para criar APIs usando Go e [Chi] como router.

Bibliotecas externas:

- Logger: [Zerolog]
- Gerador de ids únicos: [xid]

Funcionalidades:

- [x] Log de dados de requests automático
- [x] Geração automática de correlation ID
- [x] Graceful shutdown do servidor HTTP
- [x] Módulo de settings isolado da API
- [x] Módulo de settings com carregamento de valores de variáveis de ambiente
- [x] Endpoint de healtcheck
- [x] Compressão de response (deflate e gzip)
- [x] Timeouts do servidor pré-configurados
- [x] Implementar handlers de 404 e 405
- [x] Funções helper para criar respostas HTTP em JSON
- [x] Padronização de estrutura JSON de erros e erros de validação
- [ ] Carregamento de arquivos `.env`
- [ ] "Pretty logs" quando em desenvolvimento local
- [ ] Documentação OpenAPI automática (ou o mais automático possível)
- [ ] OpenTelemetry



[Chi]: https://github.com/go-chi/chi
[Zerolog]: https://github.com/rs/zerolog
[xid]: https://github.com/rs/xid
