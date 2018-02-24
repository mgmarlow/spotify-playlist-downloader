# Playlist Downloader

Download a Spotify Playlist to a csv file.

# How to Use

## From Release

1. Grab a [release](https://github.com/mgmarlow/spotify-playlist-downloader/releases/tag/1.0.0) and save it to a folder
   (e.g. `spotify-playlist-downloader/spotify-playlist-downloader.exe`)
2. Register a [Spotify App](https://beta.developer.spotify.com/documentation/general/guides/app-settings/)
3. Create a `config.json` file in the same folder of the executable (`spotify-playlist-downloader/config.json`) with the following:
    ```
    {
        "clientId": "<your app client ID>",
        "clientSecret": "<your app client secret>"
    }
    ```
4. Run `spotify-playlist-downloader.exe`
5. Open up the Spotify client and right-click your playlist -> "Share" -> "Copy Spotify URI"
6. Paste the playlist URI into the command prompt
7. A `playlist.csv` file will appear in the directory where the program is run

## From Source

1. `go get github.com/mgmarlow/spotify-playlist-downloader`
2. Register a [Spotify App](https://beta.developer.spotify.com/documentation/general/guides/app-settings/)
3. Create a `config.json` file at the root of the project with the following contents:
    ```
    {
        "clientId": "<your app client ID>",
        "clientSecret": "<your app client secret>"
    }
    ```
4. Run `go build`
5. Run `./spotify-playlist-downloader.exe`
6. Open up the Spotify client and right-click your playlist -> "Share" -> "Copy Spotify URI"
7. Paste the playlist URI into the command prompt
8. A `playlist.csv` file will appear in the directory where the program is run
