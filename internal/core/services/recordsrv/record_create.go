package recordsrv

import (
	"log"
	"os"
	"strings"

	htgotts "github.com/hegedustibor/htgo-tts"
)

type RecordService struct {
}

func New() *RecordService {
	return &RecordService{}
}

func (r *RecordService) Create(sentence string) {
	path := buildPath()
	speech := htgotts.Speech{Folder: path, Language: "en"}
	speech.CreateSpeechFile(sentence, sentence)
}

func buildPath() string {
	url, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	path := strings.Split(url, "internal")[0] + "records"

	return path
}
