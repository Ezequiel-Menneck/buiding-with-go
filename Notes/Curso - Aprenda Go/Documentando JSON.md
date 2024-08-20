
### json.Marshal
- Usamos para converter uma struct nossa em um objeto Json valido -  https://go.dev/play/p/03qTTz9slCn

### json.Unmarshal
- Usamos para converter um JSON em uma struct nossa para utilizarmos em nosso código - https://go.dev/play/p/pAnNbGPnfvz

### Tags
- Serve p mapear o valor que virá do Json para a nossa variável, permite dizer que nosso código tenha um struct c valor Robson e recebe do Json Carlos e os 2 deem match, a tag faz esse mapeamento, exp:
```go
type Carrodesom struct {
	Nome          string  `json:"Nome"`
	Sobrenome     string  `json:"Sobrenome"`
	Idade         int     `json:"Idade"`
	Profissao     string  `json:"Profissao"`
	Contabancaria float64 `json:"Contabancaria"`
}
```
- Os valores ente \`\` são quem fazem o mapeamento pra gente, dizem que no json vai vir como N e p gente quando fizermos Unmarshal vira J, ou vice versa, teremos N e faremos virar J no Marshal

### Marshal/Unmarshal VS Encode/Decode
- Quando executamos Marshal o valor vai p uma váriavel, já no encode ele vai direto pra interface que colocamos no NewEncoder, sem precisar salvar em uma variável antes 