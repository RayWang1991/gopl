package main

import (
	"time"
	"text/tabwriter"
	"os"
	"fmt"
	"sort"
)

type Track struct {
	Title, Artist, Album string
	Year                 int
	Length               time.Duration
}

func length(s string) time.Duration {
	t, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return t
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Marthin Solveig", "Smash", 2011, length("4m24s")},
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()
}

func main() {
	printTracks(tracks)
	sort.Sort(byArtist(tracks))
	printTracks(tracks)
	sort.Sort(sort.Reverse(byArtist(tracks)))
	printTracks(tracks)
	fmt.Println()
	// by length
	c:=customSort{tracks,func(s,t *Track)bool{
		return s.Length < t.Length
	}}
	sort.Sort(c)
	printTracks(tracks)
}

type byArtist []*Track

func (b byArtist) Len() int {
	return len(b)
}

func (b byArtist) Less(i, j int) bool {
	return b[i].Artist < b[j].Artist
}

func (b byArtist) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

type customSort struct {
	t    []*Track
	less func(s, t *Track) bool
}

func (c customSort) Len() int {
	return len(c.t)
}

func (c customSort) Less(i, j int) bool {
	return c.less(c.t[i], c.t[j])
}

func (c customSort) Swap(i, j int) {
	c.t[i], c.t[j] = c.t[j], c.t[i]
}
