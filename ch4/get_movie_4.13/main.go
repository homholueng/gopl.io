package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/metal3d/go-slugify"
)

const (
	APIURL = "http://www.omdbapi.com/?"
	APIKEY = "ad1b8d7f"
)

type Movie struct {
	Title  string
	Year   string
	Poster string
}

func (m Movie) posterFileName() string {
	ext := filepath.Ext(m.Poster)
	title := slugify.Marshal(m.Title)
	return fmt.Sprintf("%s_(%s)%s", title, m.Year, ext)
}

func getMovie(title string) (movie Movie, err error) {
	url := fmt.Sprintf("%sapikey=%s&t=%s", APIURL, url.QueryEscape(APIKEY), url.QueryEscape(title))
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("%d response from %s", resp.StatusCode, url)
		return
	}
	err = json.NewDecoder(resp.Body).Decode(&movie)
	if err != nil {
		return
	}
	return
}

func (m Movie) writePoster() error {
	postURL := m.Poster
	resp, err := http.Get(postURL)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("%d response from %s", resp.StatusCode, postURL)
		return err
	}
	defer resp.Body.Close()
	file, err := os.Create(m.posterFileName())
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	_, err = writer.ReadFrom(resp.Body)
	if err != nil {
		return err
	}
	err = writer.Flush()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: poster MOVIE_TITLE")
		os.Exit(1)
	}
	title := os.Args[1]
	movie, err := getMovie(title)
	if err != nil {
		log.Fatal(err)
	}

	if zero := new(Movie); movie == *zero {
		fmt.Fprintf(os.Stderr, "No results for '%s'\n", title)
		os.Exit(2)
	}

	err = movie.writePoster()
	if err != nil {
		log.Fatal(err)
	}
}
