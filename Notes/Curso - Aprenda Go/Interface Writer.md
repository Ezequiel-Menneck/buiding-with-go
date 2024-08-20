Interface Writer tem um método Write(p []byte) (n int, err error)
O Typo File implementa a interface Writer, ele tem o método Write
A função NewEncoder do package json recebe um Writer - NewEncoder(w io.Writer) \*Encoder
Ou seja, qualquer coisa que implementa a interface Writer pode ser usada como parametro para o NewEncoder.

Quando eu tenho uma função (ou qualquer coisa) que recebe uma interface do typo N, posso passar de argumento para ela qualquer outra coisa que implemente essa interface do typo N. Por exemplo:
![[Pasted image 20240527160814.png]]![[Pasted image 20240527160826.png]]
Nesse caso quando eu faço o Typo Celta implementar a interface carro eu posso passalo como argumento pra minha função carreta, caso ele não implemente eu não posso, eles não se conversam. - https://go.dev/play/p/v0dVjw8l7Zf
