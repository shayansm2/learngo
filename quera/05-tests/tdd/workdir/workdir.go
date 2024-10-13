package workdir

import (
	"fmt"
	cp "github.com/otiai10/copy"
	"os"
	"path/filepath"
	"vc/commands"
)

// you can use this library freely: "github.com/otiai10/copy"

type WorkDir struct {
	whereami     string
	subscribedVC *commands.VC
}

func (wd *WorkDir) CreateFile(name string) error {
	path := fmt.Sprintf("%s/%s", wd.whereami, name)
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}

func (wd *WorkDir) CreateDir(name string) error {
	path := fmt.Sprintf("%s/%s", wd.whereami, name)
	err := os.Mkdir(path, 0755)
	if err != nil {
		return err
	}
	return nil
}

func (wd *WorkDir) WriteToFile(name string, content string) error {
	path := fmt.Sprintf("%s/%s", wd.whereami, name)
	file, err := os.OpenFile(path, os.O_WRONLY, 0775)
	defer file.Close()
	if err != nil {
		return err
	}

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}
	wd.PublishChangedFiles(name)
	return nil
}

func (wd *WorkDir) AppendToFile(name string, content string) error {
	path := fmt.Sprintf("%s/%s", wd.whereami, name)
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0775)
	defer file.Close()
	if err != nil {
		return err
	}

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}
	wd.PublishChangedFiles(name)
	return nil
}

func (wd *WorkDir) Clone() *WorkDir {
	newDir := fmt.Sprintf("%s_copy", wd.whereami)
	cp.Copy(wd.whereami, newDir)
	return &WorkDir{
		whereami: newDir,
	}
}

func (wd *WorkDir) ListFilesRoot() []string {
	result, err := wd.ListFilesIn("")
	if err != nil {
		return []string{}
	}
	return result
}

func (wd *WorkDir) ListFilesIn(path string) ([]string, error) {
	var files []string

	err := filepath.Walk(fmt.Sprintf("%s/%s", wd.whereami, path), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			relPath, err := filepath.Rel(wd.whereami, path)
			if err != nil {
				return err
			}
			files = append(files, relPath)
		}

		return nil
	})

	if err != nil {
		return []string{}, err
	}

	return files, nil
}

func (wd *WorkDir) CatFile(fileName string) (string, error) {
	path := fmt.Sprintf("%s/%s", wd.whereami, fileName)
	bytes, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func InitEmptyWorkDir() *WorkDir {
	testDir := "test_dir"
	os.Mkdir(testDir, 0755)
	wd := &WorkDir{
		whereami: fmt.Sprintf("./%s", testDir),
	}
	return wd
}

func (wd *WorkDir) SubscribeVC(vc *commands.VC) {
	wd.subscribedVC = vc
}

func (wd *WorkDir) PublishChangedFiles(file string) {
	if wd.subscribedVC == nil {
		return
	}

	wd.subscribedVC.AddChangedFiles(file)
}
