package helpers

import (
	"errors"
	"fmt"
)

//flagsUtility is slice of flags the utility
var flagsUtility = []string{"-host", "-u", "-p", "-db", "-P", "-f", "-q"}

//VerifyFlags execute verification for flags
func VerifyFlags(flags map[string]string) error {
	for key, _ := range flags {
		if err := contains(flagsUtility, key); err != nil {
			return err
		}

	}
	return nil
}

//contains verify if flag exist.
func contains(flags []string, name string) error {
	existed := false
	var flg string
	for _, flag := range flags {
		if flag == name {
			existed = true
			flg = flag
			break
		}
	}
	if existed == false {
		return errors.New(fmt.Sprintf("flag: [%s] not permited", flg))
	}
	return nil
}

//empty verify if flag is empty.
func empty(flag string, value string) error {
	if len(value) == 0 {
		return errors.New(fmt.Sprintf("value: [%s] the flag -> [%s] not permited\n", value, flag))
	}
	return nil
}
