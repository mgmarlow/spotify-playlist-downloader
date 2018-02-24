# Playlist Downloader

Download a Spotify Playlist to a csv file.

## How to Use

1. Create a `config.json` file at the root of the project with the following contents:
    ```
    {
        "clientId": "<your app client ID>",
        "clientSecret": "<your app client secret>"
    }
    ```
2. Run `go build`
3. Run `./spotify-playlist-downloader.exe`
4. Open up the Spotify client and right-click your playlist -> "Share" -> "Copy Spotify URI"
5. Paste the playlist URI into the command prompt
6. A `playlist.csv` file will appear in the directory where the program is run
