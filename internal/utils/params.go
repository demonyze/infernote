package utils

import (
	"errors"
	"fmt"
	"strings"

	"github.com/demonyze/infernote/internal/constants"
	"github.com/demonyze/infernote/internal/model"
)

func GetParams(args []string) model.InfernoteParams {
	params := model.InfernoteParams{
		Language: constants.Language,
		FileName: constants.FileName,
		FilePath: constants.FilePath,
	}

	for _, arg := range args {
		namedParam, err := ExtractNamedParam(arg)
		if err != nil {
			continue
		}

		switch namedParam.Key {

		case "lang":
			params.Language = namedParam.Value

		case "name":
			params.FileName = namedParam.Value

		case "path":
			params.FilePath = namedParam.Value
		}

	}

	return params
}

func ExtractNamedParam(s string) (model.NamedParam, error) {
	namedParam := model.NamedParam{}
	splitDashes := strings.Split(s, "--")
	if len(splitDashes) < 2 {
		message := fmt.Sprintf("param '%s' does not follow -- syntax.", s)
		return namedParam, errors.New(message)
	}

	splitAssignment := strings.Split(splitDashes[1], "=")
	if len(splitAssignment) != 2 {
		message := fmt.Sprintf("param '%s' is not an assignment", s)
		return namedParam, errors.New(message)
	}

	namedParam.Key = splitAssignment[0]
	namedParam.Value = splitAssignment[1]

	return namedParam, nil
}
