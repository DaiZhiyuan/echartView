package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-echarts/go-echarts/charts"
)

var (

    speccpu2006Items = []string {
        "410.bwaves",
        "416.gamess",
        "433.milc",
        "434.zeusmp",
        "435.gromacs",
        "436.cactusADM",
        "437.leslie3d",
        "444.namd",
        "447.dealII",
        "450.soplex",
        "453.povray",
        "454.calculix",
        "459.GemsFDTD",
        "465.tonto",
        "470.lbm",
        "481.wrf",
        "482.sphinx3",
        "SPECfp_rate_base2006",
    }
)


func ft2500StaticInit() []int {
    r := []int {332, 1960, 558, 972, 1460, 972, 417, 1610, 1300, 561, 2400, 1010, 384, 1230, 437, 854, 799, 864}
    return r
}

func zebe2500StaticInit() []int {
    r := []int {332, 1960, 558, 972, 417, 1610, 1300, 561, 2400, 1010, 384, 1230, 437, 854, 799, 864, 1000, 1300}
    return r
}

func speccpu2006BarReverse() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.TitleOpts{Title: "SPEC cpu2006 Analyse View", Subtitle: "SPEC CPU2006 v1.2"}, 
        charts.ToolboxOpts{Show: false}, 
        charts.InitOpts{Width: "1500px", Height: "768px", 
        PageTitle: "Tainshu View System"})

	bar.AddXAxis(speccpu2006Items).
		AddYAxis("FT-S2500/64", ft2500StaticInit()).
        AddYAxis("Zebu-S2500/64", zebe2500StaticInit())

	bar.XYReversal()
	return bar
}

func speccpu2006Handler(w http.ResponseWriter, _ *http.Request) {
	page := charts.NewPage()
	page.Add(speccpu2006BarReverse())

	f, err := os.Create(getRenderPath("speccpu2006.html"))
	if err != nil {
		log.Println(err)
	}

	page.Render(w, f)
}
