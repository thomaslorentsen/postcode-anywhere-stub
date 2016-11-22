package main

import (
	"net/http"
	"encoding/json"
	"log"
)

type AddressList struct {
	Items []Address
}

type Address struct {
	Id string
	DomesticId string
	Language string
	LanguageAlternatives string
	Department string
	Company string
	SubBuilding string
	BuildingNumber string
	BuildingName string
	SecondaryStreet string
	Street string
	Block string
	Neighbourhood string
	District string
	City string
	Line1 string
	Line2 string
	Line3 string
	Line4 string
	Line5 string
	AdminAreaName string
	AdminAreaCode string
	Province string
	ProvinceName string
	ProvinceCode string
	PostalCode string
	CountryName string
	CountryIso2 string
	CountryIso3 string
	CountryIsoNumber string
	SortingNumber1 string
	SortingNumber2 string
	Barcode string
	POBoxNumber string
	Label string
	Type string
	DataLevel string
}

type ItemList struct {
	Items []Item
}

type Item struct {
	Id string
	Type string
	Text string
	Highlight string
	Description string
}

func findHandler(w http.ResponseWriter, r *http.Request) {

	items := ItemList{
		Items: make([]Item, 0),
	}

	container := r.URL.Query().Get("Container")
	if (container == "") {
		item := Item{
			Id: "GB|RM|A|P-SW1A-1RB",
			Type: "Postcode",
			Text: "SW1A 1RB",
			Highlight: "0-1,1-2,2-3,3-4,5-6,6-7,7-8",
			Description: "Arlington Street, London",
		}
		items.Items = append(items.Items, item);
	} else {
		item := Item{
			Id: "GB|RM|A|1000676865",
			Type: "Address",
			Text: "Ritz Hotel, Arlington Street",
			Highlight: "",
			Description: "London, SW1A 1RB",
		}
		items.Items = append(items.Items, item);

		item = Item{
			Id: "GB|RM|A|2",
			Type: "Address",
			Text: "The long test",
			Highlight: "",
			Description: "A long address",
		}
		items.Items = append(items.Items, item);
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "pca-source")
	json.NewEncoder(w).Encode(items)
}

func retrieveHandler(w http.ResponseWriter, r *http.Request) {

	items := AddressList{
		Items: make([]Address, 0),
	}
	id := r.URL.Query().Get("Id")
	if (id == "GB|RM|A|2") {
		address := Address{
			Id: "GB|RM|A|2",
			DomesticId: "2",
			Language: "ENG",
			LanguageAlternatives: "ENG",
			Company: "Super Long Widget Company",
			SubBuilding: "Widget Research Department",
			BuildingName: "The Big Blue Building",
			Street: "The Longest Street In The World",
			Line1: "This is a line that should exceed the limit of fifty characters",
			Line2: "Line 2",
			District: "Central District",
			City: "London",
			PostalCode: "SW1 1AA",
			CountryName: "United Kingdom",
			Label: "This is a really long address",
			Type: "Commercial",
			DataLevel: "Premise",
		}
		items.Items = append(items.Items, address)
	} else {
		address := Address{
			Id: "GB|RM|A|1000676865",
			DomesticId: "1000676865",
			Language: "ENG",
			LanguageAlternatives: "ENG",
			Company: "Ritz Hotel",
			Street: "Arlington Street",
			Line1: "Arlington Street",
			City: "London",
			PostalCode: "SW1A 1RB",
			CountryName: "United Kingdom",
			Label: "Ritz Hotel\nArlington Street\nLONDON\nSW1A 1RB\nUNITED KINGDOM",
			Type: "Commercial",
			DataLevel: "Premise",
		}
		items.Items = append(items.Items, address)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "pca-source")
	json.NewEncoder(w).Encode(items)
}

func main() {
	http.HandleFunc("/Capture/Interactive/Find/v1.00/json3ex.ws", findHandler)
	http.HandleFunc("/Capture/Interactive/Retrieve/v1.00/json3ex.ws", retrieveHandler)
	err := http.ListenAndServeTLS(":9040", "server.pem", "server.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
