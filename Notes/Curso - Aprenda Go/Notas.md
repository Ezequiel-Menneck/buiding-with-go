Tudo que começa com letra maiúscula é exportado, visível p outros packages quando importamos o nosso, em minuscula fica invisível (public, private do java)

### Compilação cruzada
GOOS
GOARCH
`GOOS=windows GOARCH=base64 go build main.do` 
gera um .exe do windows
O mesmo funcionaria p mac com:
`GOOS=darwin GOARCH=base64 go build main.do` 
