# Notes

#### Rotas
Criação de rotas com o Chi pode ser feito com chi.NewRouter() \
Chi é usado mais como um criador de rotas, o resto temos default em go

#### Handler x HandlerFunc
Handler é mais parrudão o HandlerFunc já é mais roubusto. Não precisamos criar uma struct c ele, apenas passamos uma
função com a escrita e a leitura (ResposeWrites e Request)

#### Middlewares - Chi
Pontos programaveis p filtar as requests. \
Podemos usar com chi.Use(middleware) \
Podemos criar um middleware com:
```go
func myMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("before")
		next.ServeHTTP(w, r)
		fmt.Println("after")
	})
}
```
Chi tem alguns middlewares defaults, como:
- RealIP - RealIP é um middleware que define o RemoteAddr de um http.Request para os resultados da análise do cabeçalho X-Real-IP ou do cabeçalho X-Forwarded-For.
- Logger - Loga o começo e final da request
- Recoverer - Da um recover de um panic, loga o panic e retorna um 500 Internal Server Error se possível.

