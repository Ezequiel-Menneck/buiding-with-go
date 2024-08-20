
- Testes devem:
	- ficar num arquivo que termine com \_test.go
	- ficar no mesmo package que terá o código a ser testado
	- ficar em funções com o nome "func TestNome(t \*testing.T)"
- Para rodar os testes:
	- go test
	- go test -v
- Para falhas, utilizamos o t.Error(), onde a maneira idiomática é algo como: "expected: x. got: y"