package service

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type sourceFileService struct {
	elementName *string
}

const pascalCaseRegex = `^[A-Z][a-z0-9]*$`

func (sfs *sourceFileService) validateName(name *string) error {
	if !regexp.MustCompile(pascalCaseRegex).MatchString(*name) {
		return fmt.Errorf("write %s name following PascalCase pattern", *sfs.elementName)
	}

	return nil
}

func (sfs *sourceFileService) genFileName(name *string) *string {
	pattern := regexp.MustCompile(pascalCaseRegex)
	fileName := pattern.ReplaceAllString(*name, "${1}_${2}")
	if fileName == "_" {
		fileName = strings.ToLower(*name)
	} else {
		fileName = strings.ToLower(fileName)
	}

	fileName = fileName + ".go"

	return &fileName
}

func (sfs *sourceFileService) createSourceFile(targetDir *string, fileName *string, content *string) error {
	srcFile, err := os.Create(fmt.Sprintf("%s/%s", *targetDir, *fileName))
	if err != nil {
		return err
	}

	if _, err := srcFile.WriteString(*content); err != nil {
		return err
	}

	return nil
}

func newSourceFileService(elementName string) *sourceFileService {
	return &sourceFileService{&elementName}
}
