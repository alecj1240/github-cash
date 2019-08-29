package actiontests

import (
	"testing"

	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/suite"
	"github_cash/actions"
)

type ActionSuite struct {
	*suite.Action
}

func Test_ActionSuite(t *testing.T) {
	action, err := suite.NewActionWithFixtures(actions.App(), packr.New("Test_ActionSuite", "../fixtures"))
	if err != nil {
		t.Fatal(err)
	}

	as := &ActionSuite{
		Action: action,
	}
	suite.Run(t, as)
}
