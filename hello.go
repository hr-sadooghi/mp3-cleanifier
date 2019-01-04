package main

import (
    "fmt"
    "io/ioutil"
     "log"
     "strings"
     "strconv"
)
import (
    id3 "github.com/mikkyang/id3-go"
)

func main() {

    files, err := ioutil.ReadDir("./")
    if err != nil {
        log.Fatal(err)
    }
    
    for _, f := range files {
        fileName:=f.Name()
        fileExt:=fileName[len(fileName)-4:len(fileName)]
        // fmt.Println(fileExt)
        if ( fileExt == ".mp3" ) {
            fmt.Println("FileName:\t" + fileName)
            mp3File, err := id3.Open(fileName)
            if ( err != nil ) {
                fmt.Println("Error open file!")
            } else {
                defer mp3File.Close()

                Title, TitleHasDiff := cleanify(mp3File.Title())
                if(TitleHasDiff){
                    mp3File.SetTitle(Title)
                }

                Album, AlbumHasDiff := cleanify(mp3File.Album())
                if(AlbumHasDiff){
                    mp3File.SetAlbum(Album)
                }
                
                Artist, ArtistHasDiff := cleanify(mp3File.Artist())
                if(ArtistHasDiff){
                    mp3File.SetArtist(Artist)
                }

                fmt.Println("Title:\t" + Title + strconv.FormatBool(TitleHasDiff) )
                fmt.Println("Album:\t" + Album + strconv.FormatBool(AlbumHasDiff))
                fmt.Println("Artist:\t"+Artist+strconv.FormatBool(ArtistHasDiff))
                fmt.Println("Genre:\t"+mp3File.Genre())
                fmt.Println("Year:\t"+mp3File.Year())
                // fmt.Println("Composer:\t"+mp3File.Composer())
                fmt.Println("Comments:")
                fmt.Println(mp3File.Comments())

                
            }
        }
        fmt.Println("")
    }
}

func cleanify(uncleanstring string) (string, bool) {
    str := uncleanstring
    str = strings.Replace(str, ".:: PardisMusic.net ::.", "", -1)
    str = strings.Replace(str, "پردیس موزیک", "", -1)
    str = strings.TrimSpace(str)
    return str, uncleanstring != str
}
