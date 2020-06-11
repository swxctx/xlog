package ozlog

import (
	"fmt"
	"testing"
)

func TestXlog(t *testing.T) {
	fmt.Println(blue("Blue"))
	fmt.Println(lightGreen("LightGreen"))
	fmt.Println(purple("Purple"))
	fmt.Println(red("Red"))
	fmt.Println(green("Green"))
	fmt.Println(yellow("Yellow"))
	fmt.Println(orange("Orange"))

	Default.Level = TraceLevel
	Tracef("Trace test->%d", TraceLevel)
	Debugf("Debug test->%d", DebugLevel)
	Infof("Infof test->%d", InfoLevel)
	Warnf("Warnf test->%d", WarnLevel)
	Errorf("Error test->%d", ErrorLevel)
	Fatalf("Fatalf test->%d", FatalLevel)
}
