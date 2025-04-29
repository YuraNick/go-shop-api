package main

import (
	"3-validation-api/configs"
	"3-validation-api/inernal/verify"
	"3-validation-api/pkg/jsonfile"
	"fmt"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	router := http.NewServeMux()
	jsonFile := jsonfile.NewJsonFile("./data/hash_emails.json")

	verifyRepo := verify.NewVerifyRepository(jsonFile)

	verify.NewVerifyHandler(router, verify.VerifyHandlerDeps{
		Config:           conf,
		VerifyRepository: verifyRepo,
	})

	port := 81
	fmt.Printf("server is running http://localhost:%d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}
