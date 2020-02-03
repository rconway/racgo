package main

import "testing"

func checkEnvVar(t *testing.T, varname string, valExpected string) {
	valGot := ShowEnv(varname)
	if valExpected == valGot {
		t.Logf("Correct answer for EnvVar %v: expected %v, got %v", varname, valExpected, valGot)
	} else {
		t.Errorf("Wrong answer for EnvVar %v: expected %v, got %v", varname, valExpected, valGot)
	}
}

func TestMain(t *testing.T) {
	main()
	checkEnvVar(t, "FREDBOB", "Barry")
	checkEnvVar(t, "BURTLARRY", "Gordon")
	checkEnvVar(t, "SHELL", "/bin/bash")
}
