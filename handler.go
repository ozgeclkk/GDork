package main

import (
	"fmt"
	"net/url"
)

func SorguHazirla(hedef string) []Dork {
	var sonuclar []Dork

	for _, dork := range SorguListesi {
		tamSorgu := fmt.Sprintf(dork.SorguMetni, hedef)
		googleLink := "https://www.google.com/search?q=" + url.QueryEscape(tamSorgu)

		sonuclar = append(sonuclar, Dork{
			Baslik:   dork.Baslik,
			Kategori: dork.Kategori,
			Sorgu:    tamSorgu,
			Link:     googleLink,
		})

	}
	return sonuclar

}
