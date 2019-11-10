package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchevents"
	"log"
	"net/http"
)

const (
	EventDetailType = "FluxEvent"
	EventSource     = "flux"
	DetailTemplate  = "{}"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "Hello world!\n")

		cfg, err := external.LoadDefaultAWSConfig()
		if err != nil {
			log.Fatalln("failed to load config")

		}

		svc := cloudwatchevents.New(cfg)
		params := buildEventParams()
		req := svc.PutEventsRequest(&params)

		resp, err := req.Send(context.Background())
		if err == nil { // resp is now filled
			fmt.Println(resp)
		}
	})

	_ = http.ListenAndServe(":8080", nil)
}

func buildEventParams() cloudwatchevents.PutEventsInput {
	source := EventSource
	msg := EventDetailType
	detail := fmt.Sprintf(DetailTemplate)
	params := cloudwatchevents.PutEventsInput{
		Entries: []cloudwatchevents.PutEventsRequestEntry{
			{
				Detail:     &detail,
				DetailType: &msg,
				Source:     &source,
			},
		},
	}

	if err := params.Validate(); err != nil {
		panic("CloudWatchEvents request is not valid.  Error: " + err.Error())
	}

	return params
}
