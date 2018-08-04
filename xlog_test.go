package xlog

import (
	"fmt"
	"testing"
)

func TestXlog(t *testing.T) {
	fmt.Println(Blue("Blue"))
	fmt.Println(LightGreen("LightGreen"))
	fmt.Println(Purple("Purple"))
	fmt.Println(White("White"))
	fmt.Println(Gray("Gray"))
	fmt.Println(Red("Red"))
	fmt.Println(Green("Green"))
	fmt.Println(Yellow("Yellow"))

	Default.Level = DebugLevel
	Infof("infof test ...")
	Debugf("debug test ...")
	Warnf("warnf test ...")
	Errorf("error test ...")
	Fatalf("fatalf test ...")
}
