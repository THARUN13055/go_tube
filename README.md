# GoTube - YouTube Downloader in Golang

GoTube is a command-line tool built with Golang that allows you to download YouTube videos and audio using the `github.com/kkdai/youtube/v2` package. This tool supports downloading individual videos, playlists, and extracting audio-only from videos.

## Features
- Download individual videos.
- Download entire playlists.
- Filter video quality (e.g., 1080p).
- Download audio only in MP3 format.

## Installation

1. **Clone the repository:**

    ```bash
    git clone https://github.com/yourusername/gotube.git
    cd go_tube
    ```

2. **Install Dependencies:**

    This project relies on the `github.com/kkdai/youtube/v2` package. You can install it by running:

    ```bash
    go get -u github.com/kkdai/youtube/v2
    ```

3. **Build the project:**

    Build the project by running:

    ```bash
    go build -o go_tube main.go
    ```

4. **Run the tool:**

    Use the following command structure to download videos, playlists, or audio only:

    ```bash
    ./go_tube <option> <URL>
    ```

## Usage

### 1. Download a Playlist
To download all videos in a playlist, run the following command:

```bash
./go_tube playlist <playlist-url> --all-videos --download --quality 1080p
    0       1          2               3            4         5       6

```