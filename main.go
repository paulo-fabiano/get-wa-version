package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type VersionInfo struct {
	App         string `json:"app"`
	Version     string `json:"version"`
	CollectedAt string `json:"collected_at"`
}

func normalizeVersion(version string) string {
	return strings.Split(version, "-")[0]
}

func main() {

	url := "https://wppconnect.io/whatsapp-versions/"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Erro ao acessar página:", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("Status inválido: %d", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal("Erro ao parsear HTML:", err)
	}

	// Seleciona exatamente o botão com a versão
	version := strings.TrimSpace(
		doc.Find("a.button--secondary").First().Text(),
	)

	if version == "" {
		log.Fatal("Versão não encontrada")
	}

	result := VersionInfo{
		App:         "whatsapp-web",
		Version:     normalizeVersion(version),
		CollectedAt: time.Now().UTC().Format(time.RFC3339),
	}

	output, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(output))
}
