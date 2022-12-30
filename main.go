package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"

	"oss.terrastruct.com/d2/d2layouts/d2dagrelayout"
	"oss.terrastruct.com/d2/d2lib"
	"oss.terrastruct.com/d2/d2renderers/d2svg"
	"oss.terrastruct.com/d2/d2themes/d2themescatalog"
	"oss.terrastruct.com/d2/lib/textmeasure"
)

func getRoot(w http.ResponseWriter, r *http.Request) {

	ruler, _ := textmeasure.NewRuler()
	diagram, _, _ := d2lib.Compile(context.Background(), "x -> y", &d2lib.CompileOptions{
		Layout:  d2dagrelayout.Layout,
		Ruler:   ruler,
		ThemeID: d2themescatalog.GrapeSoda.ID,
	})

	out, err := d2svg.Render(diagram, &d2svg.RenderOpts{
		Pad: d2svg.DEFAULT_PADDING,
	})

	if err != nil {
		fmt.Printf("not able render graph: %s \n", err)
	}

	_, err = w.Write(out)

	if err != nil {
		fmt.Printf("not able to show graph: %s \n", err)
	}

}

func main() {
	http.HandleFunc("/", getRoot)
	err := http.ListenAndServe(":3333", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")

	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
