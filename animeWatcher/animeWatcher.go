package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func errorCheck(err error) {

	if err != nil {
		log.Fatal(err)
	}

}

func write(writer http.ResponseWriter, msg string) {

	_, err := writer.Write([]byte(msg))

	errorCheck(err)
}

func indexHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Yes, the app is working.")
}

type animeList struct {
	AnimeCount   int
	AnimeToWatch []string
}

func getStrings(fileName string) []string {

	var lines []string

	file, err := os.Open(fileName)

	if os.IsNotExist(err) {
		return nil
	}

	errorCheck(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	errorCheck(scanner.Err())
	return lines

}

func interactHandler(writer http.ResponseWriter, request *http.Request) {

	animeVals := getStrings("animeList.txt")

	fmt.Printf("%v\n", animeVals)

	tmpl, err := template.ParseFiles("view.html")

	errorCheck(err)

	animes := animeList{
		AnimeCount:   len(animeVals),
		AnimeToWatch: animeVals,
	}

	err = tmpl.Execute(writer, animes)

}

func newHandler(writer http.ResponseWriter, request *http.Request) {
	tmpl, err := template.ParseFiles("newAnime.html")

	errorCheck(err)

	err = tmpl.Execute(writer, nil)
}

func createHandler(writer http.ResponseWriter, request *http.Request) {

	anime := request.FormValue("anime")

	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE

	file, err := os.OpenFile("animeList.txt", options, os.FileMode(0600))

	errorCheck(err)

	_, err = fmt.Fprintln(file, anime)

	errorCheck(err)

	err = file.Close()

	errorCheck(err)

	http.Redirect(writer, request, "/interact", http.StatusFound)

}

func main() {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/interact", interactHandler)
	http.HandleFunc("/new", newHandler)
	http.HandleFunc("/create", createHandler)
	err := http.ListenAndServe("localhost:4444", nil)
	errorCheck(err)

}
