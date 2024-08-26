Preferir io.Copy p escrever uma resposta direto em um arquivo ao invés de usar io.ReadAll() e depois uma função Write.


#### Alocar variável na Heap X Stack (Go Compiler)
Essa escolha não tem relação com declarar a variável com **new** ou não, ela se faz devido ao escopo que a variável fica acessível.

```go
var global *int
func f() {
	var x int
	x = 1
	global = &x
}
```
Nesse exemplo x será alocada na heap, já que continua sendo acessivel a partir da variável global depois que a função f() retornar, apesar de ser declarada como uma variável local, dizemos que x *escapa* de **f**

```go
func g() {
	y := new(int)
	*y = 1
}
```
Nesse caso quando a função **g()** retorna a variável torna-se inacessível e pode ser reciclada. Como \*y não *escapa* de **g()** é seguro o compilador alocar \*y na stack, mesmo sendo declarada com o **new**.  