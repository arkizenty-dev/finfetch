package main


import (
	"fmt"
	"time"
	. "github.com/logrusorgru/aurora"
)

func main() {
	then := time.Now()
	fmt.Println(Cyan("                *     * "))
	fmt.Println(Cyan("              * * * * * *"))
	fmt.Println(Cyan("            *     * *     * "))
	fmt.Println(Cyan("                  * * "))
	fmt.Println(Cyan("                  * * "))
	fmt.Println(Blue("            # # # # # # # # #       "))
	fmt.Println(Blue("#   #     # # # # # # # # # # #     "), Blue("Date:"), Cyan(then.Month()), Cyan(then.Day()), "," , Cyan(then.Year()))
	fmt.Println(Blue("  #     # # # # # # # # # #"), Black("@"), Blue("# #   "), Blue("Weekday:"), Cyan(then.Weekday()))
	fmt.Println(Blue("  #   # # # # # # # # # # #"), Black("@"), Blue("# # # "), Blue("Time:"), Cyan(then.Hour()), ":", Cyan(then.Minute()), ":", Cyan(then.Second()))
	fmt.Println(Blue("  # # # # # # # # # # # # # # # # # "), Blue("Timezone:"), Cyan(then.Location()))
	fmt.Println(Blue("    # # # # # # # # # # # # #"), Black("@"), Blue("# #"))
	fmt.Println(Blue("      # # # # # # # # # # # # #"), Black("@ @"))
	fmt.Println(Blue("        # # # # # # # # # # # # # #"))
	fmt.Println(Cyan("          @ @"), Blue("# # # # # # # #"), Cyan("@ @ @"))
	fmt.Println(Cyan("            @ @ @ @ @ @ @ @ @ @"))
	fmt.Println("\n\n")
}
