package recordsrv

import (
	"os"
	"strings"
	"testing"
)

func TestCreate(t *testing.T) {
	recordSrv := New()
	recordSrv.Create("we'll see you tonigh")

	path := path()

	_, err := os.Open(path + "/we'll see you tonigh.mp3")
	if err != nil {
		t.Errorf("Create(we'll see you tonigh) = it wasn't created")
	}

	os.Remove(path + "/we'll see you tonigh.mp3")
}

func path() string {
	url, _ := os.Getwd()
	path := strings.Split(url, "internal")[0] + "records"

	return path
}
