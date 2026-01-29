<div align="center">

# AbacatePay Go REST 🥑

Cliente HTTP oficial em Go para integração com a API da AbacatePay.  
Rápido, seguro, idiomático e pronto para produção.

<img src="https://res.cloudinary.com/dkok1obj5/image/upload/v1767631413/avo_clhmaf.png" width="100%" alt="AbacatePay Open Source"/>

## Instalação

</div>

```bash
go get github.com/AbacatePay/go-rest@latest
```

<div align="center">

## Configuração básica

</div>

```go
import "github.com/AbacatePay/go-rest"

client := abacatepay.New(abacatepay.Options{
	Secret: "ABACATEPAY_SECRET",
})
```

Ou via variável de ambiente

```bash
export ABACATEPAY_SECRET=...
```

<div align="center">

## Conceitos importantes

### Versão da API

Por padrão, o SDK utiliza a versão v1 da API.

Você pode sobrescrever:

</div>

```go
client := abacatepay.New(abacatepay.Options{
	Version: 2,
})
```

<div align="center">

### Retry automático

Por padrão:

</div>

- **3 tentativas**
- Retry em **408**, **429**, **5xx**
- Backoff exponencial

<div align="center">

**Customização**

</div>

```go
client := abacatepay.New(abacatepay.Options{
	Retry: abacatepay.RetryOptions{
		Max: 5,
		OnRetry: func(attempt int) {
			log.Println("retry", attempt)
		},
	},
})
```

<div align="center">

### Timeout
</div>

```go
client := abacatepay.New(abacatepay.Options{
	Timeout: 3 * time.Second,
})
```

<div align="center">

## Quickstart

**Crie um novo cupom**
</div>

```go
package main

import (
	"fmt"

    types "github.com/AbacatePay/go-types/v2"
	abacatepay "github.com/AbacatePay/go-rest"
)

func main() {
	client := abacatepay.New(abacatepay.Options{
		Secret: "YOUR_SECRET",
	})

	body := types.RESTPostCreateCouponBody{
        Value: 10,
		Code: "PROMO10",
		Type: "PERCENTAGE",
	}

	res, err := client.Post[types.RESTPostCreateCouponData](
		v2.RouteCreateCoupon,
		body,
	)

	if err != nil {
		panic(err)
	}

	fmt.Println("Coupon ID:", res.ID)
}
```

<div align="center">

## Tratamento de erros

O SDK expõe erros tipados:

</div>

```go
if err != nil {
	switch e := err.(type) {
	case *abacatepay.AbacatePayError:
		fmt.Println("API error:", e.Message)
	case *abacatepay.HTTPError:
		fmt.Println("HTTP error:", e.Status)
	default:
		fmt.Println("Unknown error:", err)
	}
}
```

<p align="center"> Feito com 🥑 pela equipe AbacatePay <br/> Open source, de verdade. </p>
