package cidsdk

import (
	"github.com/cidverse/cidverseutils/pkg/archive/tar"
	"github.com/cidverse/cidverseutils/pkg/archive/zip"
)

// ZIPCreate creates a zip archive of the directory at the given path.
func (sdk SDK) ZIPCreate(inputDirectory string, outputFile string) error {
	return zip.Create(inputDirectory, outputFile)
}

// ZIPExtract unzips the zip archive at the given path into the given directory.
func (sdk SDK) ZIPExtract(archiveFile string, outputDirectory string) error {
	return zip.Extract(archiveFile, outputDirectory)
}

// TARCreate creates a tar archive of the directory at the given path.
func (sdk SDK) TARCreate(inputDirectory string, outputFile string) error {
	return tar.Create(inputDirectory, outputFile)
}

// TARExtract extracts a tar archive at the given path into the given directory.
func (sdk SDK) TARExtract(archiveFile string, outputDirectory string) error {
	return tar.Extract(archiveFile, outputDirectory)
}
