package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"text/template"
)

type Question struct {
	Text    string
	Options []Option
}

type Option struct {
	Text     string
	IsAnswer bool
}

var questions []Question

func main() {
	loadQuestions("src/questions.txt")
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("src/img"))))
	http.HandleFunc("/", menuHandler)
	http.HandleFunc("/jeux", startQuizHandler) // Nouvelle route pour accéder au jeu
	http.HandleFunc("/question", quizHandler)
	http.HandleFunc("/information", informationHandler)
	http.HandleFunc("/end", endHandler)
	http.HandleFunc("/restart", restartHandler)
	http.ListenAndServe(":9999", nil)
}

func loadQuestions(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var q Question
	var optionsStarted bool
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Question") {
			if q.Text != "" {
				questions = append(questions, q)
			}
			q = Question{}
			q.Text = strings.TrimSpace(strings.SplitN(line, ":", 2)[1])
			optionsStarted = false
		} else if strings.HasPrefix(line, "Réponse :") {
			optionsStarted = false
			answer := strings.TrimSpace(strings.SplitN(line, ":", 2)[1])
			for i := range q.Options {
				if q.Options[i].Text == answer {
					q.Options[i].IsAnswer = true
					break
				}
			}
		} else if strings.HasPrefix(line, "A.") || strings.HasPrefix(line, "B.") || strings.HasPrefix(line, "C.") || strings.HasPrefix(line, "D.") {
			if !optionsStarted {
				q.Options = []Option{}
				optionsStarted = true
			}
			q.Options = append(q.Options, Option{Text: strings.TrimSpace(line)})
		}
	}
	if q.Text != "" {
		questions = append(questions, q)
	}
}

func menuHandler(w http.ResponseWriter, r *http.Request) {
	// Charger et afficher le contenu de menu.html
	tmpl := template.Must(template.ParseFiles("src/templates/menu.html"))
	tmpl.Execute(w, nil)
}

func startQuizHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/question?id=0", http.StatusSeeOther)
}

func informationHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("src/templates/information.html"))
	tmpl.Execute(w, nil)
}

func quizHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.NotFound(w, r)
		return
	}

	questionID, err := strconv.Atoi(id)
	if err != nil || questionID < 0 || questionID >= len(questions) {
		http.NotFound(w, r)
		return
	}

	if r.Method == "POST" {
		r.ParseForm()
		selectedOptionIndex, err := strconv.Atoi(r.Form.Get("option"))
		if err != nil || selectedOptionIndex < 0 || selectedOptionIndex >= len(questions[questionID].Options) {
			http.Error(w, "Invalid option selected", http.StatusBadRequest)
			return
		}

		if questions[questionID].Options[selectedOptionIndex].IsAnswer {
			nextQuestionID := questionID + 1
			if nextQuestionID < len(questions) {
				http.Redirect(w, r, "/question?id="+strconv.Itoa(nextQuestionID), http.StatusSeeOther)
				return
			} else {
				http.Redirect(w, r, "/end", http.StatusSeeOther)
				return
			}
		}
	}

	tmpl := template.Must(template.ParseFiles("src/templates/quiz.html"))
	tmpl.Execute(w, struct {
		Text    string
		Options []Option
		ID      int
	}{
		Text:    questions[questionID].Text,
		Options: questions[questionID].Options,
		ID:      questionID,
	})
}

func endHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("src/templates/end.html"))
	tmpl.Execute(w, nil)
}

func restartHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
