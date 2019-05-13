package jGollery

import (
	"net/http"
	"os"
	"strings"
)

var picEndings = []string{".gif", ".bmp", ".jpg", ".jpeg", ".png", ".webp", ".ico"}

type File struct {
	*os.File
}

func Open(filename string) (*File, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	return &File{file}, nil
}

// Get all (supported) pictures in this folder (NOT in subfolders)
// Supported filetypes are ".gif", ".bmp", ".jpg", ".jpeg", ".png", ".webp", ".ico"
func (f *File) Pictures() (*[]os.FileInfo, error) {
	if files, err := f.Readdir(0); err == nil {
		filtered := make([]os.FileInfo, 0)

		for _, fileInfo := range files {
			if !fileInfo.IsDir() {
				name := strings.ToLower(fileInfo.Name())

				// First check file ending
				for _, ending := range picEndings {
					if strings.HasSuffix(name, ending) {

						// Then check also mime-type for "image/..." just to be sure
						if file, err := os.Open(f.Name() + "/" + fileInfo.Name()); err == nil {
							if fileType, err := GetFileContentType(file); err == nil {
								if strings.HasPrefix(fileType, "image/") {
									// TODO: Maybe check if ending fits to mimetype
									filtered = append(filtered, fileInfo)
									break
								}
							} else {
								return nil, err
							}
						} else {
							return nil, err
						}
					}
				}

			}

		}

		return &filtered, nil
	} else {
		return nil, err
	}
}

func GetFileContentType(out *os.File) (string, error) {

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}

	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}
