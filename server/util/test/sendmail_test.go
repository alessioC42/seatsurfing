package test

import (
	"path/filepath"
	"testing"

	. "github.com/seatsurfing/seatsurfing/server/testutil"
	. "github.com/seatsurfing/seatsurfing/server/util"
)

func TestGetEmailTemplatePathExists(t *testing.T) {
	res, err := GetEmailTemplatePath(GetEmailTemplatePathResetpassword(), "de")
	CheckStringNotEmpty(t, res)
	CheckTestBool(t, true, err == nil)
}

func TestGetEmailTemplatePathFallback(t *testing.T) {
	res, err := GetEmailTemplatePath(GetEmailTemplatePathResetpassword(), "notexists")
	CheckStringNotEmpty(t, res)
	CheckTestBool(t, true, err == nil)
}

func TestGetEmailTemplatePathNotExists(t *testing.T) {
	path, _ := filepath.Abs("./res/notexisting.txt")
	res, err := GetEmailTemplatePath(path, "en")
	CheckTestString(t, "", res)
	CheckTestBool(t, true, err != nil)
}

func TestGetLocalPartFromEmailAddress(t *testing.T) {
	CheckTestString(t, "test", GetLocalPartFromEmailAddress("test@domain.com"))
	CheckTestString(t, "test", GetLocalPartFromEmailAddress("test@domain"))
	CheckTestString(t, "test", GetLocalPartFromEmailAddress("test"))
	CheckTestString(t, "\"a@b\"", GetLocalPartFromEmailAddress("\"a@b\"@example.com"))
}
