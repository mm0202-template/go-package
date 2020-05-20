package main

import (
  "testing"
  "strings"
)

func TestMain(t *testing.T) {
	if strings.EqualFold("Hoge!!", HelloWorld()) {
		t.Errorf("error!") 
	}

	if strings.EqualFold("Fuga!!", HelloWorld()) {
		t.Errorf("error!") 
	}

  	if !strings.EqualFold("Hello World!!", HelloWorld()) {
    	t.Errorf("error!") 
  	}
}