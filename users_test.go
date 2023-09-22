package main

import (
	"testing"
	structs "users/structs"
)

func TestUsers(t *testing.T) {
	t.Run("create holder struct", func(t *testing.T) {
		_, err := structs.NewUserManager()
		if err != nil {
			t.Errorf("got %q when trying to instantiate user manager", err)
		}
	})

}
