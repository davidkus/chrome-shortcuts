package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {

	if err := godotenv.Load(".env"); err != nil {
		log.Print("Error loading .env file")
	}

	address := os.Getenv("ECHO_ADDRESS")
	configFile := os.Getenv("CONFIG_FILE")

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		query := strings.TrimSpace(c.QueryParam("q"))
		parts := strings.SplitN(query, " ", 2)

		if len(parts) == 0 || len(query) == 0 {
			return c.String(http.StatusBadRequest, "No query provided")
		}

		mappings, err := readConfigFileData(configFile)
		if err != nil {
			return c.String(http.StatusInternalServerError, fmt.Sprintf("Internal Server Error: %s", err.Error()))
		}

		var url string
		if len(parts) == 1 {
			url = mappings.Shortcuts[parts[0]]
		} else {
			url = mappings.ShortcutsWithParams[parts[0]]
			// Replace %s with the rest of the query
			if len(parts) > 1 && strings.Contains(url, "%s") {
				url = fmt.Sprintf(url, parts[1])
			}
		}

		if len(url) == 0 {
			// Redirect to Google if we can't find a match.
			return c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("https://www.google.com/search?q=%s", query))
		}

		return c.Redirect(http.StatusTemporaryRedirect, url)
	})
	e.GET("/config", func(c echo.Context) error {
		var result, err = readConfigFileData(configFile)
		if err != nil {
			return c.String(http.StatusInternalServerError, fmt.Sprintf("Internal Server Error: %s", err.Error()))
		}
		return c.JSONPretty(http.StatusOK, result, "  ")
	})
	e.Logger.Fatal(e.Start(address))
}

type config struct {
	Shortcuts           map[string]string `json:"shortcuts"`
	ShortcutsWithParams map[string]string `json:"shortcutsWithParams"`
}

func readConfigFileData(filename string) (*config, error) {
	configFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	rawConfig, _ := io.ReadAll(configFile)
	var result *config
	if err = json.Unmarshal(rawConfig, &result); err != nil {
		return nil, err
	}
	configFile.Close()
	return result, nil
}
