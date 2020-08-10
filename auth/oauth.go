package auth

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/go-yaml/yaml"
	"github.com/joho/godotenv"
	"github.com/zmb3/spotify"
)

var (
	auth = spotify.NewAuthenticator(
		"http://localhost:8888/callback",
		spotify.ScopeUserModifyPlaybackState,
		spotify.ScopeUserReadPlaybackState,
		spotify.ScopeUserLibraryModify,
		spotify.ScopeUserReadRecentlyPlayed,
		spotify.ScopePlaylistModifyPrivate,
		spotify.ScopePlaylistModifyPublic,
	)
	state = "abc123"
	ch    = make(chan *spotify.Client)
)

func oauth() (*string, error) {
	clientID, secretID := getClientID()

	auth.SetAuthInfo(clientID, secretID)

	http.HandleFunc("/callback", handler)
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("Got requestTypes for:", request.URL.String())
	})

	go func() {
		err := http.ListenAndServe(":8888", nil)
		fmt.Println("Error: ", err)
	}()

	authURL := auth.AuthURL(state)
	fmt.Println("Please log in to Spotify by visiting the following page in your browser:", authURL)

	client := <-ch

	token, err := client.Token()
	if err != nil {
		return nil, err
	}

	err = createDotToken(token.RefreshToken)

	if err != nil {
		return nil, err
	}

	return &token.AccessToken, nil
}

func createDotToken(token string) error {
	file, err := os.Create(tokenFile)
	if err != nil {
		return err
	}
	err = file.Close()
	if err != nil {
		return err
	}

	bytes := []byte(token)

	return ioutil.WriteFile(tokenFile, bytes, 0666)
}

func handler(writer http.ResponseWriter, request *http.Request) {
	token, err := auth.Token(state, request)
	if err != nil {
		http.Error(writer, "Couldn't get token", http.StatusForbidden)
		log.Fatal(err)
	}
	if st := request.FormValue("state"); st != state {
		http.NotFound(writer, request)
		log.Fatalf("State mismatch: %s != %s\n", st, state)
	}

	client := auth.NewClient(token)
	_, err = fmt.Fprintf(writer, "<h1>Login!!!</h1><h3>You can start this app!!!</h3>")
	if err != nil {
		log.Fatal(err)
	}
	ch <- &client
}

type configs struct {
	ID     string `yaml:"SpotifyID"`
	Secret string `yaml:"SpotifySecret"`
}

func getClientID() (clientID, secretID string) {
	if err := godotenv.Load(); err == nil {
		clientID = os.Getenv("SPOTIFY_ID")
		secretID = os.Getenv("SPOTIFY_SECRET")
	}

	buf, err := ioutil.ReadFile(configFile)
	if err != nil {
		return
	}
	var ids configs
	err = yaml.Unmarshal(buf, &ids)
	if err != nil {
		return
	}
	clientID = ids.ID
	secretID = ids.Secret

	return
}
