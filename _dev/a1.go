package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {
	fmt.Printf("Hello, %s\n", `World!`)
    
	color.Cyan("Hello, World!\n")
	fmt.Printf("Hello, %s\n", `World!`)
    
	color.Red("Hello, World!\n")
	fmt.Printf("Hello, %s\n", `World!`)

	color.Blue("Hello, World!\n")
	fmt.Printf("Hello, %s\n", `World!`)

	color.Yellow("Hello, World!\n")
	fmt.Printf("Hello, %s\n", `World!`)
    
	color.Magenta("Hello, World!\n")
	fmt.Printf("Hello, %s\n", `World!`)
    
	color.Green("Hello, World!\n")
	fmt.Printf("Hello, %s\n", `World!`)

	color.White("Hello, World!\n")
	fmt.Printf("Hello, %s\n", `World!`)
    
    
    yellow := color.New(color.FgYellow).SprintFunc()
    red := color.New(color.FgRed).SprintFunc()
    fmt.Printf("This is a %s and this is %s.\n", yellow("warning"), red("error"))

    info := color.New(color.FgWhite, color.BgGreen).SprintFunc()
    fmt.Printf("This %s rocks!\n", info("package"))


    info2 := color.New(color.FgWhite, color.BgCyan).SprintFunc()
    fmt.Printf("This %s rocks2!\n", info2("package2"))

    info3 := color.New(color.FgWhite, color.Bold).SprintFunc()
    fmt.Printf("This %s rocks3!\n", info3("package3"))
    
    info4 := color.New(color.FgMagenta, color.Bold).SprintFunc()
    fmt.Printf("This %s rocks4!\n", info4("package4"))

    info41 := color.New(color.FgRed, color.Bold).SprintFunc()
    fmt.Printf("This %s rocks4!\n", info41("package4"))

    info42 := color.New(color.FgBlue, color.Bold).SprintFunc()
    fmt.Printf("This %s rocks4!\n", info42("package4"))

    info43 := color.New(color.FgGreen, color.Bold).SprintFunc()
    fmt.Printf("This %s rocks4!\n", info43("package4"))

    info44 := color.New(color.FgYellow, color.Bold).SprintFunc()
    fmt.Printf("This %s rocks4!\n", info44("package4"))

    info45 := color.New(color.FgCyan, color.Bold).SprintFunc()
    fmt.Printf("This %s rocks4!\n", info45("package4"))

    info46 := color.New(color.FgRed, color.Bold).SprintFunc()
    fmt.Printf("This %s rocks4!\n", info46("package4"))

    info47 := color.New(color.FgWhite, color.Bold).SprintFunc()
    fmt.Printf("This %s rocks4!\n", info47("package4"))

    
    info5 := color.New(color.FgMagenta).SprintFunc()
    fmt.Printf("This %s rocks4!\n", info5("package4"))
    
    info6 := color.New(color.FgBlue, color.BgYellow, color.Bold).SprintFunc()
    fmt.Printf("This %s rocks4!\n", info6("package4"))
    
    
    // Use helper functions
    fmt.Printf("This", color.RedString("warning"), "should be not neglected.")
    fmt.Printf(color.GreenString("Info:"), "an important message." )

    // Windows supported too! Just don't forget to change the output to color.Output
    fmt.Fprintf(color.Output, "Windows support: %s", color.GreenString("PASS"))    
    
    
    
    
}
