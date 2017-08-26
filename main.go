package main

import (
	"os"
	"fmt"
	"encoding/json"
	"io/ioutil"

	"gopkg.in/alecthomas/kingpin.v2"
	"github.com/dogenzaka/go-iap/appstore"
	"github.com/mattn/go-isatty"
)

var (
	app = kingpin.New("iap-dump", "A command-line application dumping data from an InAppPurchase receipt")

	appStore = app.Command("appstore", "Dump a receipt from the AppStore (Apple)")
)

func main() {
	interactive := isatty.IsTerminal(os.Stdin.Fd())
	v := kingpin.MustParse(app.Parse(os.Args[1:]))
	if interactive {
		fmt.Sprintln(os.Stderr, "Print your full receipt:")
	}
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	in := string(data)
	if interactive {
		fmt.Sprintln(os.Stderr, "Got the receipt, checking...")
	}
	switch v {
	case appStore.FullCommand():
		dumpAppstore(in)
	}
}

func dumpAppstore(receipt string) {
	client := appstore.NewWithConfig(appstore.Config{
		IsProduction: true,
	})
	req := appstore.IAPRequest{
		ReceiptData: receipt,
	}
	resp := &appstore.IAPResponse{}
	err := client.Verify(req, resp)
	if err != nil {
		panic(err)
	}
	jsonVal, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonVal)
}

