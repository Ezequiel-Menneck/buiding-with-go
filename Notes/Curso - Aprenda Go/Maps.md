Mesmo esquema de sempre, aquele padrão de Keys e Values
Declaramos um map com: `meumap := map[T key]T val{}`

Introducão: https://go.dev/play/p/24cNmpfnjaV

### Range e Deletando

Range em map o primeiro valor diferente de um slice ou array n é o index, e sim a key (bem normal, até mesmo porque eles são desordenados)

Para deletar um valor do map usamos a função delete(map, T key)

Range e Delete no map: https://go.dev/play/p/8mORm3g1jla
