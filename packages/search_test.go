package packages

import (
	"testing"
	"log"
)

func TestGetArch(t *testing.T) {
	log.Println(GetArch())
}

func TestGetRepository(t *testing.T) {
	log.Println(GetRepository())
}

func TestGetMaintainer(t *testing.T) {
	log.Println(GetMaintainer())
}

func TestGetFlagged(t *testing.T) {
	log.Println(GetFlagged())
}