# Spotify\_CLI

## OverView
Spotify Client in command line.

## GetStarted
1. Create Spotify ClientID/ClientSercretID
2. Create .env
3. Write ID and SecretID to .env with the format `SPOTIFY_ID=<ClientID>` and `SPOTIFY_SECRET=<ClientSercretID>`
4. Execute `go build main.go`
5. Execute `./main`

## commands
### Play
command: `play`

Start playing playlist with playlist url.

### Save
command: `save`

Save playlist to `playlist.json`.

### Load
command: `load`

Load and play playlist that saved in `playlist.json`.

### Resume
command: `resume`

Resume pausing truck.

### Pause
command: `pause`

Pause playing truck.

### Show
command: `show`

Show all saved playlists.

### Status
command: `status`

Show playing truck infomation.

### Refresh
command: `refresh`

Regenerate accesstoken.

### Exit
command: `exit`

Exit from this client.
