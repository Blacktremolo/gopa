package gopa

import "fmt"
import "os"
import "strconv"

const Version = "v0.2"

// Message being reported when an error occurs
var Failmsg string = "Parsing error"
// Exit code being returned when an error occurs
var Retnum int = 0

func IsValidArgpos(argpos int) bool {
	return argpos < len(os.Args) && argpos >= 0
}

func fail() {
	fmt.Println(Failmsg)
	os.Exit(Retnum)
}

func GetInt(argpos int) (val int, ok bool) {
    val, ok = 0, false
    if IsValidArgpos(argpos) {
        var err error
        val, err = strconv.Atoi(os.Args[argpos])
        ok = (err == nil)
    }
    return
}

func RequireInt(argpos int) (val int) {
	val, ok := GetInt(argpos)

    if !ok {
        fail()
    }

	return
}

func GetString(argpos int) (val string, ok bool) {
    val, ok = "", false
    if IsValidArgpos(argpos) {
        val, ok = os.Args[argpos], true
    }
    return
}

func RequireString(argpos int) (val string) {
    val, ok := GetString(argpos)

    if !ok {
        fail()
    }

    return
}
