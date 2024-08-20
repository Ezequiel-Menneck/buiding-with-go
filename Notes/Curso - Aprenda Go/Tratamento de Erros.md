- Em Go não temos exceções
- Não temos try-catch-finally
- Existe uma gama de built-in funtions para a gente se recuperar de um erro exepcional. Recovery mecanismo só é executado como parte do estado da função.
- Lidar com erros imediatamente é legal, boa prática.  Se a função retorna N e um Erro na próxima linha já tratemos o erro

#### Typo erro
- É uma interface convencional para representar uma condição de erro, sendo o valor nil representado por não ter erro. err == nil? não temos erro

### Print e Log
- fmt.Println: Normalzão, printa p gente no terminal o erro
- log.Println: Escreve para o standart error, mostra o erro junto de data + time
- log.SetOutput: Usamos esse cara p criar um arquivo de log p gente, ai nossos erros pegados com o log.Println que escreve para o standart error são escritos nesse arquivo de log que o log.SetOutput() criou para a gente
- log.Fatalln: Igual o Println() porém dps ele chama o os.Exit(1), (se o código for 0 deu sucesso, != 0 é erro) que faz com que o programa feche e mesmo funções defer não rodam. Usamos em erros catastróficos, onde o programa não pode continuar se esse erro ocorrer.
- log.Panicln(): Equivalente ao Println() seguido de uma chamada ao panic().
	- panic(): Para a execução da go routine atual, logo quando uma função F chama o panic() sua execução é parada instantaneamente, funções ou coisas que estiverem em um defer irão rodar.


### Recover
- Se tivermos uma função que chame um panic() e uma função anterior que chame essa função do panic tiver um recover() ele pegará oque passamos pro panic, e poderemos tratar como quisermos, o código não ira se encerrando até a main e tudo irá acabar, por que o recover faz a boa de falar que ele da conta e sabe oque fazer com o panic.


### Erros customizados
- Para que nossas funções retornem erros customizados podemos utilizar:
	- return errors.New(msg do erro que queremos)
	- return fmt.Errorf() -> tem um errors.New() embutido dentro dele
		- Nesse cara podemos utilizar de formatação com os % para string, p melhorar ainda mais a msg de erro
	- criar uma variavel com pare reutiliza-lá, pode ser errors.New ou fmt.Errorf
	- Criar um type que implemente a interface Error (só precisa ter o método Error q retorna uma string - satisfaz a interface e nosso typo pode ser usado onde tem parâmetros de typo erro)