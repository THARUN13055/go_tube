package Download

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"

	"github.com/kkdai/youtube/v2"
)

type Downloader struct {
	youtube.Client
	OutputDir string
}

func (d *Downloader) Download(ctx context.Context, v *youtube.Video, format *youtube.Format, outputFile string) error {
	// Log downloading video information
	youtube.Logger.Info(
		"Downloading video",
		"id", v.ID,
		"title", v.Title,
		"quality", format.Quality,
		"mimetype", format.MimeType,
	)

	// Determine output file name and format
	if outputFile == "" {
		outputFile = SanitizeFilename(v.Title) + pickIdealFileExtension(format.MimeType)
	}
	if d.OutputDir != "" {
		if err := os.MkdirAll(d.OutputDir, 0o755); err != nil {
			return err
		}
		outputFile = filepath.Join(d.OutputDir, outputFile)
	}

	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	// Download the video or audio to the specified file.
	err = d.videoDLWorker(ctx, file, v, format)
	if err != nil {
		return err
	}

	// Convert WebM to MP4 using ffmpeg if necessary
	if format.MimeType == "video/webm" {
		err = ConvertWebMToMP4(outputFile)
		if err != nil {
			return err
		}
	}

	return nil
}

func (d *Downloader) videoDLWorker(ctx context.Context, file *os.File, v *youtube.Video, format *youtube.Format) error {
	// Use the youtube package to download the video/audio.
	if err := d.Client.G(ctx, v, format, file); err != nil {
		return err
	}
	return nil
}

// ConvertWebMToMP4 uses ffmpeg to convert WebM files to MP4.
func ConvertWebMToMP4(outputFile string) error {
	// Create the new output file name
	outputMP4File := outputFile[:len(outputFile)-len(filepath.Ext(outputFile))] + ".mp4"

	// Run the ffmpeg command to convert the file
	cmd := exec.Command("ffmpeg", "-i", outputFile, outputMP4File)
	err := cmd.Run()
	if err != nil {
		return err
	}

	// Optionally, you can remove the original WebM file after conversion
	err = os.Remove(outputFile)
	if err != nil {
		return err
	}

	return nil
}

// SanitizeFilename sanitizes a string to be used as a valid file name
func SanitizeFilename(fileName string) string {
	fileName = regexp.MustCompile(`[:/<>\:"\\|?*]`).ReplaceAllString(fileName, "")
	fileName = regexp.MustCompile(`\s+`).ReplaceAllString(fileName, " ")
	return fileName
}

// pickIdealFileExtension returns the ideal file extension based on the MIME type
func pickIdealFileExtension(mimeType string) string {
	switch mimeType {
	case "video/mp4":
		return ".mp4"
	case "video/webm":
		return ".webm"
	case "audio/mp3":
		return ".mp3"
	case "audio/mpeg":
		return ".mp3"
	case "audio/webm":
		return ".webm"
	default:
		return "" // No extension if the MIME type is unknown
	}
}
