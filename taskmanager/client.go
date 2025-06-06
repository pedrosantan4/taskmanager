package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Task struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func Main() {
	task := Task{
		Title:     "Aprender Go Client",
		Completed: false,
	}

	// Serializa a task para JSON
	taskJSON, err := json.Marshal(task)
	if err != nil {
		fmt.Println("Erro ao serializar JSON:", err)
		return
	}

	// Faz a requisição POST
	resp, err := http.Post("http://localhost:8080/tasks", "application/json", bytes.NewBuffer(taskJSON))
	if err != nil {
		fmt.Println("Erro ao fazer requisição:", err)
		return
	}
	defer resp.Body.Close()

	// Lê a resposta
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro ao ler resposta:", err)
		return
	}

	fmt.Println("Status:", resp.Status)
	fmt.Println("Resposta:", string(body))
}
