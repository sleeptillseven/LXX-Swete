package main

import (
	"./types"
	"bufio"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	LXX_PATH = "src/First1KGreek-LXX-XML"
	BOOK_FOLDER = "src/First1KGreek-LXX-RAW"
)

func main() {
	allBooks := listAllFiles(LXX_PATH)
	log.Printf("found %d books in target folder", len(allBooks))

	for i := 0; i < len(allBooks); i++ {
		writingBook(allBooks[i])
	}
	log.Printf("Done processing %d files\n", len(allBooks))
}

func listAllFiles(root string) ([]string){
	var files []string

	err := filepath.Walk(LXX_PATH, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) != ".xml" || !strings.Contains(filepath.Base(path), "grc") {
			return nil
		}
		files = append(files, path)
		return nil
	})
	check(err)

	return files
}

func writingBook(book string) {
	xmlFile, err := os.Open(book)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(xmlFile)

	// we initialize our Users array
	var rawFormat types.Book

	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'users' which we defined above
	xml.Unmarshal(byteValue, &rawFormat)
	bookName := rawFormat.TeiHeader.FileDesc.TitleStmt.Title.Text
	if bookName == "" {
		log.Println("Book: ", bookName, book)
	}
	//log.Println("Book: ", bookName, book)
	text := rawFormat.Text.Body.Div

	f, err := os.Create(BOOK_FOLDER + "/" + bookName + ".txt")
	check(err)

	w := bufio.NewWriter(f)

	defer f.Close()

	for chapter := 0; chapter < len(text.Div); chapter++ {
		chapterNumber := text.Div[chapter].N

		for verse := 0; verse < len(text.Div[chapter].Div); verse++ {
			verseNumber := text.Div[chapter].Div[verse].N
			verseRawText := text.Div[chapter].Div[verse]
			verseText := ""
			if len(verseRawText.P.Text) < 1 {
				verseText = verseRawText.Text
				//log.Println(verseText)
			} else {
				verseText = verseRawText.P.Text
				//log.Println(verseText)
			}
			sanitizedText := strings.TrimSpace(strings.Replace(verseText, "\n", "", -1))

			_, errW := w.WriteString(bookName + "::" + chapterNumber + ":" + verseNumber + " " + sanitizedText + "\n")
			check(errW)
			w.Flush()
		}
	}
}

func check(e error) {
	if e != nil {
		log.Fatalln("ERROR: ", e)
	}
}
