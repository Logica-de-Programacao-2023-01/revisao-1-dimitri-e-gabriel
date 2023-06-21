package q1

func CalculateDiscount(currentPurchase float64, purchaseHistory []float64) (float64, error) {
if compraAtual <= 0 {
		return 0, errors.New("valor da compra inválido")
	}

	totalCompras := 0.0
	for _, valorCompra := range historicoCompras {
		totalCompras += valorCompra
	}

	desconto := 0.0

	if totalCompras > 1000 {
		desconto = 0.1 * compraAtual
	} else if totalCompras <= 1000 && totalCompras > 500 {
		desconto = 0.05 * compraAtual
	} else if totalCompras <= 500 {
		desconto = 0.02 * compraAtual
	}

	if len(historicoCompras) == 0 {
		desconto = 0.1 * compraAtual
	}

	mediaCompras := totalCompras / float64(len(historicoCompras))
	if mediaCompras > 1000 {
		desconto = 0.2 * compraAtual
	}

	return desconto, nil
}
