package main

import (
    "log"
    "net/http"
    "os"

    "github.com/go-echarts/go-echarts/charts"
)

var (
    unixbenchItems = []string {
        "Dhrystone",
        "Double-Precision",
        "Execl Throughput",
        "File Copy 256",
        "File Copy 1024",
        "File Copy 4096",
        "Pipe Throughput",
        "Context Switching",
        "Process Creation ",
        "Shell Scripts 1",
        "Shell Scripts 8",
        "SysCall Overhead",
    }
)

func unixbenchRandInt() []int {
    cnt := len(unixbenchItems)
    r := make([]int, 0)

    for i := 0; i < cnt; i++ {
        r = append(r, int(seed.Int63())%maxNum)
    }

    return r
}

func unixbenchBarReverse() *charts.Bar {
    bar := charts.NewBar()
    bar.SetGlobalOptions(charts.TitleOpts{Title: "UnixBench Analyse View", Subtitle: "UnixBench 5.1.3"}, 
        charts.ToolboxOpts{Show: false}, 
        charts.InitOpts{Width: "1500px", Height: "768px", 
        PageTitle: "Tainshu View System"})

    bar.AddXAxis(unixbenchItems).
    AddYAxis("FT-2000plus/64", unixbenchRandInt()).
    AddYAxis("FT-S2500/64", unixbenchRandInt()).
    AddYAxis("Zebu-S2500/64", unixbenchRandInt())

    bar.XYReversal()
    return bar
}

func unixbenchHandler(w http.ResponseWriter, _ *http.Request) {
    page := charts.NewPage()
    page.Add(unixbenchBarReverse())

    f, err := os.Create(getRenderPath("unixbench.html"))
    if err != nil {
        log.Println(err)
    }

    page.Render(w, f)
}
