package command

import "fmt"

func Next(token string) (newToken string){
  _, newToken, err := createRequest(token, "POST", "https://api.spotify.com/v1/me/player/next", nil)

  if err != nil {
    fmt.Println("Error: ", err)
    return
  }

  GetPlayStatus(token)

  return
}
