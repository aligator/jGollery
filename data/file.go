package data

import (
	"net/http"
	"os"
	"strings"
)

var picEndings = []string{".gif", ".bmp", ".jpg", ".jpeg", ".png", ".webp", ".ico"}

// File is a wrapper arround os.File. It provides some additional methods.
type File struct {
	*os.File
}

// Opens a file using os.Open()
func Open(filename string) (*File, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	return &File{file}, nil
}

// Get all (supported) pictures in this folder (NOT in subfolders)
// Supported filetypes are ".gif", ".bmp", ".jpg", ".jpeg", ".png", ".webp", ".ico"
// The mimetype is also checket.
// It returns a slice of all found picture names
func (f *File) Pictures() ([]string, error) {
	if files, err := f.Readdir(0); err == nil {
		filtered := make([]string, 0)

		for _, fileInfo := range files {
			if !fileInfo.IsDir() {
				// Then check also mime-type for "image/..." just to be sure
				if file, err := Open(f.Name() + "/" + fileInfo.Name()); err != nil {
					return nil, err
				} else if file.IsPicture() {
					filtered = append(filtered, fileInfo.Name())
					file.Close()
				}

			}

		}

		return filtered, nil
	} else {
		return nil, err
	}
}

// Checks if the File is a picture
func (f *File) IsPicture() bool {
	// First check file ending
	for _, ending := range picEndings {
		if strings.HasSuffix(strings.ToLower(f.Name()), ending) {
			// Then check MIME-type
			if fileType, err := f.getFileContentType(); err == nil {
				f.Close()
				return strings.HasPrefix(fileType, "image/")
				// TODO: Maybe check also if ending fits to mimetype
			}
		}
	}

	return false
}

// get the mimetype of the file
func (f *File) getFileContentType() (string, error) {

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, err := f.Read(buffer)
	if err != nil {
		return "", err
	}

	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}
