package q2

func AverageLettersPerWord(text string) (float64, error) {
words := strings.Fields(text)
	numWords := len(words)

	if numWords == 0 {
		return 0, errors.New("texto vazio")
	}

	totalLetters := 0
	for _, word := range words {
		totalLetters += len(word)
	}

	average := float64(totalLetters) / float64(numWords)
	return average, nil
}

func main() {
	text := "O rato roeu a roupa do rei de Roma"
	average, err := AverageLettersPerWord(text)
	if err != nil {
		panic(err)
	}
	println("Média de letras por palavra:", average)
}
