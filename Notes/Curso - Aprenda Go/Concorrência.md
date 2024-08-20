
### Concorrência VS Paralelismo
Go já foi pensando para trabalhar com Multi Cores, já que começou a ser desenvolvida depois do primeiro processador multi-core

Concorrência é um padrão de design que faz com que vários elementos sejam executados simultaneamente sem que esses caras dependam deles mesmo, Y n precisa esperar X rodar p ser executado.

### Race Condition 
Quando duas funções (ou qualquer coisa que seja) estão manipulando uma mesma variável. 
https://go.dev/play/p/cZglEGUFxN1

### Mutex
var mu sync.Mutex
mu.Lock()
mu.Unlock()
Usamos esses caras quando precisamos travar algum go routine para garantir que não teremos race condition nela. Parece com aquele semaforo do redis do Felipe

### Atomic
Mesmo esquema do Mutex, porém ao inves de usar Lock e Unlock a gente usa os métodos dele de adicionar/carregar os valores. Na surdina ele deixa tudo locado e n libera race conditions
https://go.dev/play/p/VbUUxF6D77N