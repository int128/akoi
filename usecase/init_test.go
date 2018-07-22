package usecase

import (
	"fmt"
	"testing"

	"github.com/suzuki-shunsuke/akoi/domain"
	"github.com/suzuki-shunsuke/akoi/testutil"
)

func TestInitConfigFile(t *testing.T) {
	methods := &domain.InitMethods{
		MkdirAll: testutil.NewFakeMkdirAll(nil),
		Exist:    testutil.NewFakeExistFile(true),
		Write:    testutil.NewFakeWrite(nil),
	}
	params := &domain.InitParams{Dest: "dest"}
	if err := InitConfigFile(params, methods); err != nil {
		t.Fatal(err)
	}
	methods.Exist = testutil.NewFakeExistFile(false)
	if err := InitConfigFile(params, methods); err != nil {
		t.Fatal(err)
	}
	methods.MkdirAll = testutil.NewFakeMkdirAll(fmt.Errorf("failed to create a directory"))
	if err := InitConfigFile(params, methods); err == nil {
		t.Fatal("it should be failed to create a directory")
	}
}
