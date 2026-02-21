package main

type State struct {
	LastMaxID int `json:"last_max_id"`
}

type Post struct {
	ID       int    `xml:"id,attr"`
	FileName string `xml:"file_name,attr"`
	FileURL  string `xml:"file_url,attr"`
	Date     string `xml:"date,attr"`
	Tags     string `xml:"tags,attr"`
	Author   string `xml:"author,attr"`
}

type Posts struct {
	Items []Post `xml:"tag"`
}
