package main


import (
	"fmt"
	"time"
	c "github.com/logrusorgru/aurora"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	for {
		clear()
		then := time.Now()
		fmt.Println(c.Cyan("                *     * "))
		fmt.Println(c.Cyan("              * * * * * *"))
		fmt.Println(c.Cyan("            *     * *     * "))
		fmt.Println(c.Cyan("                  * * "))
		fmt.Println(c.Cyan("                  * * "))
		fmt.Println(c.Blue("            # # # # # # # # #       "))
		fmt.Println(c.Blue("#   #     # # # # # # # # # # #     "), c.Blue("Date:"), c.Cyan(then.Month()), c.Cyan(then.Day()), "," , c.Cyan(then.Year()))
		fmt.Println(c.Blue("  #     # # # # # # # # # #"), c.Black("@"), c.Blue("# #   "), c.Blue("Weekday:"), c.Cyan(then.Weekday()))
		fmt.Println(c.Blue("  #   # # # # # # # # # # #"), c.Black("@"), c.Blue("# # # "), c.Blue("Time:"), c.Cyan(then.Hour()), ":", c.Cyan(then.Minute()), ":", c.Cyan(then.Second()))
		fmt.Println(c.Blue("  # # # # # # # # # # # # # # # # # "), c.Blue("Timezone:"), c.Cyan(then.Location()))
		fmt.Println(c.Blue("    # # # # # # # # # # # # #"), c.Black("@"), c.Blue("# #"))
		fmt.Println(c.Blue("      # # # # # # # # # # # # #"), c.Black("@ @"))
		fmt.Println(c.Blue("        # # # # # # # # # # # # # #"))
		fmt.Println(c.Cyan("          @ @"), c.Blue("# # # # # # # #"), c.Cyan("@ @ @"))
		fmt.Println(c.Cyan("            @ @ @ @ @ @ @ @ @ @"))
		fmt.Println("")
		time.Sleep(1 * time.Second)
	}
}

func clear() {
	var cls string
	if runtime.GOOS == "windows" {
		cls = "cls"
	} else {
		cls = "clear"
	}
	c := exec.Command(cls)
	c.Stdout = os.Stdout
	c.Run()
}
