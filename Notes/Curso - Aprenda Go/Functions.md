- Toda função segue esse padrão: func (receiver) nome(parameters) (returns) {}
- Tudo em GO é *pass by value*, com uma exceção
- https://go.dev/play/p/lmiNBsZ1T8I
- Numa função que recebe um argumento variadico (aquele ...tipo) não precisamos passar nenhum valor se não quiseremos https://go.dev/play/p/U6Dlz_00RlU


### Defer
- é uma keyword que diz que algo será executado após todas as instruções da função serem executadas
- Defer não pode vir depois do return, precisa ser antes
- Serve p deixar o código modular (pode ser usado p fechar um arquivo q nós abrimos em nossa func)
- https://go.dev/play/p/G8DpPJnWmXl
- https://go.dev/play/p/WuUaO_D9yRN

### Métodos
- Pará ter métodos criamos uma função com o receiver, esse receiver a transforma em um método
- É usado para determinar um contexto especifico para esse cara, como adicionar um método a um tipo especifico que nós criamos (fica parecido com Java quando a gente chama o .método, mesma ideia) https://go.dev/play/p/ozaNaWgh8EH
- É uma função especifica para o Tipo que definimos, para usá-la precisamos de um valor do Tipo que declaramos e assim podemos chama lá a partir desse tipo
- Para tipos básicos não podemos fazer esse esquema de métodos, precisa ser tipos nossos que criamos

### Interface e Polimorfismo
#### Interface
- Interface é um conjunto de métodos (mesmo esquema do java)
- Não é preciso implementar explicitamente, se teu Tipo pessoal criado tiver todos os métodos da interface pessoal criada ele implementa ela automagicamente
- https://go.dev/play/p/vqiX8zdOamx
#### Polimorfismo 
- Métodos que vistos na interface são iguais mas podem ter diversas implementações apenas necessário respeitar a assinatura

### Funções anonimas
Criamos uma função sem nome que pode ou não receber parâmetros e chamamos a mesma loga em seguida, caso ela receba algum parâmetro precisamos declara-lo antes da criação da mesma: https://go.dev/play/p/px6NqfVo4kA

### Funções como expressão
É quando damos de valor á uma variável uma função, nesse caso ela não precisará de nome por que assumirá o nome da variável, já que a variável será a função. https://go.dev/play/p/AeOygIGVyqf

### Retornando uma função
É quando uma função nos retorna uma outra função: https://go.dev/play/p/glee81nYuyM
Similar ao callback hell in js -> https://go.dev/play/p/vmpFkYD3X62

### Callback
É passar uma função como argumento: https://go.dev/play/p/OyhcstFPOkw

### Closure
É quando a gente captura um escopo. https://go.dev/play/p/yCvIFHJZaku

### Recursividade