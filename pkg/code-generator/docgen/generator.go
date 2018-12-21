package docgen

import (
	"bytes"
	"text/template"

	"github.com/iancoleman/strcase"
	"github.com/pseudomuto/protokit"
	"github.com/solo-io/solo-kit/pkg/code-generator"
	"github.com/solo-io/solo-kit/pkg/code-generator/docgen/templates"
	"github.com/solo-io/solo-kit/pkg/code-generator/model"
)

const fileHeader = `<!-- Code generated by solo-kit. DO NOT EDIT. -->
`

// must ignore validate.proto from lyft
// may need to add more here
var ignoredFiles = []string{
	"validate/validate.proto",
}

func generateFilesForProtoFiles(project *model.Project, protoFiles []*protokit.FileDescriptor) (code_generator.Files, error) {
	var v code_generator.Files
	for suffix, tmpl := range map[string]*template.Template{
		".sk.md": templates.ProtoFileTemplate(project),
	} {
		for _, protoFile := range protoFiles {
			var ignore bool
			for _, ignoredFile := range ignoredFiles {
				if protoFile.GetName() == ignoredFile {
					ignore = true
					break
				}
			}
			if ignore {
				continue
			}
			content, err := generateProtoFileFile(protoFile, tmpl)
			if err != nil {
				return nil, err
			}
			fileName := protoFile.GetName() + suffix
			v = append(v, code_generator.File{
				Filename: fileName,
				Content:  content,
			})
		}
	}

	return v, nil
}

func GenerateFiles(project *model.Project) (code_generator.Files, error) {
	protoFiles := protokit.ParseCodeGenRequest(project.Request)

	files, err := generateFilesForProject(project)
	if err != nil {
		return nil, err
	}
	messageFiles, err := generateFilesForProtoFiles(project, protoFiles)
	if err != nil {
		return nil, err
	}
	files = append(files, messageFiles...)

	for i := range files {
		files[i].Content = fileHeader + files[i].Content
	}
	return files, nil
}

func generateFilesForProject(project *model.Project) (code_generator.Files, error) {
	var v code_generator.Files
	for suffix, tmpl := range map[string]*template.Template{
		".project.sk.md": templates.ProjectDocsRootTemplate(project),
	} {
		content, err := generateProjectFile(project, tmpl)
		if err != nil {
			return nil, err
		}
		v = append(v, code_generator.File{
			Filename: strcase.ToSnake(project.Name) + suffix,
			Content:  content,
		})
	}
	return v, nil
}

func generateProjectFile(project *model.Project, tmpl *template.Template) (string, error) {
	buf := &bytes.Buffer{}
	if err := tmpl.Execute(buf, project); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func generateProtoFileFile(protoFile *protokit.FileDescriptor, tmpl *template.Template) (string, error) {
	buf := &bytes.Buffer{}
	if err := tmpl.Execute(buf, protoFile); err != nil {
		return "", err
	}
	return buf.String(), nil
}
