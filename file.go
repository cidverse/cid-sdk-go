package cidsdk

import (
	"os"
	"path/filepath"
	"strings"

	cp "github.com/otiai10/copy"
)

type FileRequest struct {
	Directory  string   `json:"dir"`
	Extensions []string `json:"ext"`
}

type File struct {
	Path      string `json:"path"`
	Directory string `json:"dir"`
	Name      string `json:"name"`
	NameShort string `json:"name_short"`
	Extension string `json:"ext"`
}

func NewFile(path string) File {
	split := strings.SplitN(filepath.Base(path), ".", 2)
	fileName := split[0]
	fileExt := ""
	if len(split) > 1 && split[1] != "" {
		fileExt = "." + split[1]
	}

	return File{
		Path:      path,
		Directory: filepath.Dir(path),
		Name:      filepath.Base(path),
		NameShort: fileName,
		Extension: fileExt,
	}
}

// FileRead command
func (sdk SDK) FileRead(file string) (string, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

// FileList command
func (sdk SDK) FileList(req FileRequest) (files []File, err error) {
	err = filepath.Walk(req.Directory, func(path string, info os.FileInfo, err error) error {
		if info != nil && !info.IsDir() {
			if len(req.Extensions) > 0 {
				for _, ext := range req.Extensions {
					if strings.HasSuffix(path, ext) {
						files = append(files, NewFile(path))
						break
					}
				}
			} else {
				files = append(files, NewFile(path))
			}
		}

		return nil
	})

	return
}

// FileRename command
func (sdk SDK) FileRename(old string, new string) error {
	err := os.Rename(old, new)
	if err != nil {
		return err
	}

	return nil
}

// FileCopy command
func (sdk SDK) FileCopy(old string, new string) error {
	err := cp.Copy(old, new)

	return err
}

// FileRemove command
func (sdk SDK) FileRemove(file string) error {
	err := os.Remove(file)
	if err != nil {
		return err
	}

	return nil
}
