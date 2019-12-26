package test

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	t.Logf("%s [Error]", time.Now().Format("20060102150405")) //20060102150405.000 #这个相当于java中的yyyy-MM-ddHH
}

