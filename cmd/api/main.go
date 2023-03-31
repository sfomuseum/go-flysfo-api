package main

import (
	_ "gocloud.dev/runtimevar/constantvar"
	_ "gocloud.dev/runtimevar/filevar"
)

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"

	"github.com/sfomuseum/go-flags/multi"
	"github.com/sfomuseum/go-flysfo-api"
	"github.com/sfomuseum/runtimevar"
)

func main() {

	apikey_uri := flag.String("apikey", "", "A valid developers.flysfo.com API key encoded as a gocloud.dev/runtimevar URI. Supported schemes are: constant://, file://")
	method := flag.String("method", "", "The relative URL of the method to be executed.")

	var args multi.KeyValueString
	flag.Var(&args, "param", "Zero or more query parameters to append to the method being executed.")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Execute a method against the developers.flysfo.com API.\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t %s [options]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Valid options are:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	ctx := context.Background()

	apikey, err := runtimevar.StringVar(ctx, *apikey_uri)

	if err != nil {
		log.Fatalf("Failed to derive API key, %v", err)
	}

	cl, err := api.NewClient(ctx, apikey)

	if err != nil {
		log.Fatalf("Failed to create new API client, %v", err)
	}

	var params *url.Values

	if len(args) > 0 {

		params = &url.Values{}

		for _, p := range args {
			params.Set(p.Key(), p.Value().(string))
		}
	}

	rsp, err := cl.ExecuteMethod(ctx, *method, params)

	if err != nil {
		log.Fatalf("Failed to execute method '%s', %v", *method, err)
	}

	defer rsp.Close()

	_, err = io.Copy(os.Stdout, rsp)

	if err != nil {
		log.Fatalf("Failed to copy results to STDOUT, %v", err)
	}
}
