- Maneira de transmitir dados entre as go routines (a main é uma go routine)
- Modo certo de fazer sincronização e código concorrente.
- Servem p coordenar, sincronizar, orquestrar e buffering.
- P criar um canal é `make(chan typo)`
- Numa mesma goroutine não podemos colocar e tirar informações em um canal (analogia a um corredor de bastão, só pode passar o bastão p outro corredor) ex: mostrando dando errado: https://go.dev/play/p/abOysHdtB4_e
- Exemplo funcionando colocando um valor pro canal com go routine e tirando c outra: https://go.dev/play/p/fAanVs0jeAC (a go routine anonima coloca a info no canal, a go routine main tira a info do canal (sim a main é uma go routine))
- **Buffer** para adicionarmos buffer no canal a sintaxe muda, fica `make(chan typo, <qtd>)`
	- qtd significada a quantidade de dados q podemos colocar ali numa paulada sem precisar de outras go routines - https://go.dev/play/p/SL8SEMeFpWE
- Via de regra n vamos utilizar Buffers no nosso código
	- Existem casos úteis para serem utilizados, porém, via de regra não usamos


### Canais Direcionais 
- Podem ser direcionais
	- Podemos ter canais que apenas enviam informações como canais que apenas recebem informações
- Um send channel e um receive channel tem typos diferentes. Isso faz com que o compilador entenda que não é possível escrever num canal de leitura 
- Canais bidirecionais
	- send chan<-
	- receive <-chan
	- Seta sempre aponta pra esquerda
		- Exemplos:
			- https://go.dev/play/p/8yZ9OqN8FHB
			- https://go.dev/play/p/9DC0-Fk0oEL
	- Posso pegar um canal bidirecional e transformalo em um receiver ou um sender

### Range e Close

- Depois de enviar todas as informações para um Chanel sender precisamos dar um close(<nosso_canal>) pra q ele saiba que não irá receber mais nada e não tomemos o erro de deadklock 
- range em canais é assim:
```go
c := make(chan int)

go meuloop(10, c)

for v := range c {
	fmt.Println(v)
}
```
https://go.dev/play/p/jpImwjsZMGd


### Select

- É como um switch só que pra canais e não é sequencial
- Se eu tiver mais de um caso valido ele escolhe um aleatório

- Exemplo 1: https://go.dev/play/p/47D1X0sHbxM
- Exemplo 2: https://go.dev/play/p/APaeoDhC3Cd
- Exemplo 3: https://go.dev/play/p/Z9QIAFQcNKa
- Exemplo 4: https://go.dev/play/p/oeB1iRcoQfV


### Expressão comma ok
https://go.dev/play/p/FANmM-tiz3Q
Nesse caso no segundo print se eu o comma OK serve p gente saber que o 0 não é um valor em si, ele é o zero-value, já que nosso ok está como false.


### Convergência
- Observamos convergência quando informação de vários canais é enviada a um número menor de canais.

- Exemplo bem básico de convergência: https://go.dev/play/p/QPAVMw0ybh6
	- Tiramos info de 2 canais e colocamos em 1 só

- Exemplo mais complexo: https://go.dev/play/p/UEkLiwGEQHN
	- Temos uma função que cria um canal e atocha de dado dentro dele com uma string q passamos e o retornamos
	- Temos outra função que recebe 2 canais, cria um novo canal, coloca todas as infos dos 2 canais que passamos p ela dentro de um canal e converge tudo e retorna o canal novo 

### Divergência
- O contrário de convergência, separamos uma coisa em várias.
- https://go.dev/play/p/xUiTUazKbGt
	- Exercico de convergencia e dps divergencia
- https://go.dev/play/p/H1bJgLdbsmT
	- Nesse caso ficou mais top, deixamos um número maximo de go routines que podem ser criadas então cada execução é chamada de bloco em bloco, que seria o número máximo de go routines.

### Context
- Usamos esse cara quando queremos matar uma go routine no meio da execução
- Destaques:
    - ctx := context.Background
    - ctx, cancel = context.WithCancel(context.Background)
    - goroutine: select case <-ctx.Done(): return; default: continua trabalhando.
    - check ctx.Err();
    - Tambem tem WithDeadline/Timeout
- Exemplos (Todd):
    - Analisando:
        - Background: https://play.golang.org/p/cByXyrxXUf 
        - WithCancel: https://play.golang.org/p/XOknf0aSpx
        - Função Cancel: https://play.golang.org/p/UzQxxhn_fm 
    - Exemplos práticos:
        - func WithCancel: https://play.golang.org/p/Lmbyn7bO7e
        - func WithCancel: https://play.golang.org/p/wvGmvMzIMW 
        - func WithDeadline: https://play.golang.org/p/Q6mVdQqYTt 
        - func WithTimeout: https://play.golang.org/p/OuES9sP_yX 
        - func WithValue: https://play.golang.org/p/8JDCGk1K4P 