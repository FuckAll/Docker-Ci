package ci

import (
	//	"fmt"
	"testing"
)

func TestCiRun(t *testing.T) {
	//CiRun("OnlyBuild", "663d2166")
	CiRun("Push", "663d2166")
}
