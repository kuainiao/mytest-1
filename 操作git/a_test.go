package 操作git

import (
	"github.com/libgit2/git2go"
	"log"
	"testing"
)

func TestA(t *testing.T) {
	a()
}

func credentialsCallback(url string, username string, allowedTypes git.CredType) (git.ErrorCode, *git.Cred) {
	ret, cred := git.NewCredSshKey("git", "/home/hero/.ssh/id_rsa.pub", "/home/hero/.ssh/id_rsa", "")
	return git.ErrorCode(ret), &cred
}

func certificateCheckCallback(cert *git.Certificate, valid bool, hostname string) git.ErrorCode {
	return 0
}

func a() {

	cloneOptions := &git.CloneOptions{}
	cloneOptions.FetchOptions = &git.FetchOptions{
		RemoteCallbacks: git.RemoteCallbacks{
			CredentialsCallback:      credentialsCallback,
			CertificateCheckCallback: certificateCheckCallback,
		},
	}
	repo, err := git.Clone("git@github.com:test/test.git", "/tmp/code", cloneOptions)
	if err != nil {
		log.Panic(err)
	}
	log.Print(repo)
}
