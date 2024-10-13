package commands

import "vc/workdir"

type Files struct {
	ModifiedFiles []string
	StagedFiles   []string
}

type VC struct {
	wd *workdir.WorkDir
	Files
}

func Init(wd *workdir.WorkDir) *VC {
	return &VC{
		Files: Files{
			ModifiedFiles: make([]string, 0),
			StagedFiles:   make([]string, 0),
		},
		wd: wd,
	}
}

func (vs *VC) Status() Files {
	return vs.Files
}

func (vs *VC) Add(files ...string) {
	for _, file := range files {
		for i := 0; i < len(vs.ModifiedFiles); i++ {
			if vs.ModifiedFiles[i] != file {
				continue
			}
			vs.ModifiedFiles = append(vs.ModifiedFiles[:i], vs.ModifiedFiles[i+1:]...)
			vs.StagedFiles = append(vs.StagedFiles, file)
			break
		}
	}
	vs.StagedFiles = removeDuplicates(vs.StagedFiles)
}

func (vs *VC) AddAll() {
	vs.StagedFiles = append(vs.StagedFiles, vs.ModifiedFiles...)
	vs.StagedFiles = removeDuplicates(vs.StagedFiles)
	vs.ModifiedFiles = []string{}
}

func (vs *VC) Commit(message string) {
}

func (vs *VC) GetWorkDir() *workdir.WorkDir {
	return &workdir.WorkDir{}
}

func (vs *VC) Checkout(tag string) (*workdir.WorkDir, error) {
	return &workdir.WorkDir{}, nil
}

func (vs *VC) Log() []string {
	return []string{}
}

func (vs *VC) AddChangedFiles(file string) {
	vs.ModifiedFiles = removeDuplicates(append(vs.ModifiedFiles, file))
}

func removeDuplicates(arr []string) []string {
	hashmap := make(map[string]bool)
	for _, item := range arr {
		hashmap[item] = true
	}
	result := make([]string, len(hashmap))
	i := 0
	for item := range hashmap {
		result[i] = item
		i++
	}
	return result
}
