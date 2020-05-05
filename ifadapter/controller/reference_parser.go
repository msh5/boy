package controller

import (
	"strings"
)

type referenceType int

const (
	gistURLSegmentCount                = 3
	gitHubRLSegmentCount               = 4
	unknownReferenceType referenceType = iota
	gistFileReferenceType
	gitHubBlobReferenceType
)

func detectReferenceType(ref string) referenceType {
	if strings.HasPrefix(ref, "gist.github.com") {
		return gistFileReferenceType
	}

	if strings.HasPrefix(ref, "github.com") {
		return gitHubBlobReferenceType
	}

	return unknownReferenceType
}

type gistFileReference struct {
	UserID   string
	GistName string
}

type urlSegmentError struct{}

func (e *urlSegmentError) Error() string {
	return "reference is unexpected syntax"
}

func parseGistFileReference(ref string) (*gistFileReference, error) {
	segments := strings.Split(ref, "/")

	if len(segments) != gistURLSegmentCount {
		return nil, &urlSegmentError{}
	}

	refObj := gistFileReference{
		UserID:   segments[1],
		GistName: segments[2],
	}

	return &refObj, nil
}

type gitHubBlobReference struct {
	RepositoryOwner string
	RepositoryName  string
	BlobPath        string
}

func parseGitHubBlobReference(ref string) (*gitHubBlobReference, error) {
	segments := strings.SplitN(ref, "/", 4)

	if len(segments) != gitHubRLSegmentCount {
		//return nil, xerrors.Errorf("reference is unexpected syntax")
		//return nil, errors.New("reference is unexpected syntax")
		return nil, &urlSegmentError{}
	}

	refObj := gitHubBlobReference{
		RepositoryOwner: segments[1],
		RepositoryName:  segments[2],
		BlobPath:        segments[3],
	}

	return &refObj, nil
}
