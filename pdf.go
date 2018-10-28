package main

type Answer struct {
	Answer  string `json:"answer"`
	Correct bool   `json:"correct"`
}

type Question struct {
	Question string   `json:"question"`
	Answers  []Answer `json:"answers"`
}

func NewAnswer(answer string, correct bool) Answer {
	return Answer{Answer: answer, Correct: correct}
}

func createPDFPage() {

}
