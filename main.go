package main

import (
	"FuriganaApiCall/request"
	"encoding/csv"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	err := godotenv.Load()
	failOnError(err)

	client, err := request.NewClient(os.Getenv("API_URL"))
	failOnError(err)

	output, err := os.OpenFile("output.csv", os.O_WRONLY|os.O_CREATE, 0600)
	failOnError(err)
	defer output.Close()
	err = output.Truncate(0)
	failOnError(err)

	target, err := os.Open("target.csv")
	failOnError(err)

	regx := regexp.MustCompile(`[[:space:]]`)
	r := csv.NewReader(target)
	writer := csv.NewWriter(output)
	for {
		line, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		req := request.FuriganaApiRequest{
			AppId:      os.Getenv("APP_ID"),
			Sentence:   line[0],
			OutputType: "hiragana",
		}
		res, err := client.CallApi(req)
		if err != nil {
			log.Fatal("Error fail call api")
			return
		}
		sentence := strings.Replace(req.Sentence, " ", "", -1)
		converted := regx.ReplaceAllString(res.Converted, "")
		writer.Write([]string{sentence, converted})
	}
	writer.Flush()
	fmt.Println("complete")
}

func failOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
