package internal

import (
	"bytes"
	"embed"
	"fmt"
	"io"
	"io/fs"
	"path/filepath"
	"strings"
)

var IgnoreFiles = []string{
	".gitignore",
	".bzrignore",
	".chefignore",
	".cfignore",
	".cvsignore",
	".boringignore",
	".deployignore",
	".dockerignore",
	".ebignore",
	".eleventyignore",
	".eslintignore",
	".flooignore",
	".gcloudignore",
	".helmignore",
	".jpmignore",
	".jshintignore",
	".hgignore",
	".mtn-ignore",
	".nodemonignore",
	".npmignore",
	".nuxtignore",
	".openapi-generator-ignore",
	".p4ignore",
	".prettierignore",
	".stylelintignore",
	".stylintignore",
	".swagger-codegen-ignore",
	".terraformignore",
	".tfignore",
	".tokeignore",
	".upignore",
	".vercelignore",
	".yarnignore",
}

//go:embed templates/*
var templates embed.FS

type TemplateRegistry struct {
	templates fs.FS
}

func NewTemplateRegistry() *TemplateRegistry {
	return &TemplateRegistry{templates: templates}
}

func (tr *TemplateRegistry) HasTemplate(name string) bool {
	_, err := fs.Stat(tr.templates, fmt.Sprintf("templates/%s.gitignore", name))
	return err == nil
}

func (tr *TemplateRegistry) List() []string {
	var templates []string
	fs.WalkDir(tr.templates, ".", func(path string, d fs.DirEntry, _ error) error {
		if d.IsDir() {
			return nil
		}
		template := strings.TrimSuffix(filepath.Base(path), ".gitignore")
		templates = append(templates, template)
		return nil
	})
	return templates
}

func (tr *TemplateRegistry) CopyTemplate(name string, dst io.Writer) error {
	defer tr.writeFooter(dst)

	b, err := fs.ReadFile(tr.templates, fmt.Sprintf("templates/%s.gitignore", name))
	if err != nil {
		return err
	}
	io.Copy(dst, bytes.NewReader(b))
	return nil
}

func (tr *TemplateRegistry) writeHeader(name string, dst io.Writer) {
	const header = `# Auto-generated %s .gitignore by ignore: https://github.com/neptship/ignore

`

	fmt.Fprintf(dst, header, name)
}

func (tr *TemplateRegistry) writeFooter(dst io.Writer) {
	const footer = `
# End of ignore: https://github.com/neptship/ignore`

	fmt.Fprintln(dst, footer)
}

func (tr *TemplateRegistry) WriteTemplate(name string, dst io.Writer) error {
	tr.writeHeader(name, dst)
	return tr.CopyTemplate(name, dst)
}
