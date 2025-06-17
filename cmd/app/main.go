package main

import (
    "fmt"
    "log"
    "os"
    "03_task/internal/generator"
)
func main() {
    // Task01()
    // Task02()
    Task03()
}
func Task03() {
    tmpl, err := os.ReadFile("assets/template/config_template.yml")
    if err != nil {
        log.Fatal(err)
    }

    if err := generator.ConfigGenerate(string(tmpl), "config.yml"); err != nil {
        log.Fatal(err)
    }
    
    fmt.Println("Config generated successfully!")
}
