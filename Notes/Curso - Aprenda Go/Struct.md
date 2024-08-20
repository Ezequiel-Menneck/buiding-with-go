É um tipo que criamos que contem vários valores com diversos tipos como:
```go
type cliente struct {
	nome      string
	sobrenome string
	fumante   bool
}
```

Podemos definir de 2 maneiras:
```go
cliente1 := cliente{
	nome:      "jorge",
	sobrenome: "augusto",
	fumante:   true,
}

cliente2 := cliente{"Joana", "Silva", false}
```
https://go.dev/play/p/MJeq0GwQyjX
### Struct embutido

É basicamente um Struct dentro de outro Struct
Podemos ter quantos structs dentro de outros structs
https://go.dev/play/p/2znGcQ45iKf

### Acessando valores
- Para acessar os valores é só pegar a variavel do type N struct q criamos e colocar . atributo a ser acessado, algo como cliente1.nome, que o resultado seria jorge (parecido com o java).
- Isso pode ir ficando sempre mais deep, depende quantas structs embutidas teremos
- Se formos ter uma valor dentro de uma struct que será outra struct e esse valor tiver o mesmo nome da struct não precisamos definir o typo dele, será automagicamente definido, é chamado de embedded field. Com embedded fields podemos acessar propriedades direto dessa struct caso o ela tenha um atributo que não dê conflito com o da struct de nível acima, algo como:
```go
type pessoa struct {
	nome  string
	idade int
}

type profissional struct {
	pessoa
	titulo  string
	salario int
}

pessoa2 := profissional{
	pessoa: pessoa{
		nome:  "maricota",
		idade: 31,
	},
	titulo:  "pizzaiolo",
	salario: 10000,
}

fmt.Println(pessoa2.nome)
```
Nesse caso no fmt.Println(pessoa2.nome) podemos acessar nome direito por que ele é único no contexto do typo profissional.

### Struct anonimos
- É quando criamos uma struct sem criar ele no scopo do módulo, declaramos ela dentro de uma função ou algo assim, fazemos com :=, sendo algo descartável, usou joga fora.
```go
x := struct {
	nome  string
	idade int
}{
	nome:  "Mario",
	idade: 33,
}

fmt.Println(x)
```