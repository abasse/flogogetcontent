package flogogetcontent

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// activityLog is the default logger for the exec Activity
var log = logger.GetLogger("activity-tibco-flogogetcontent")

type GetcontentActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new AppActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &GetcontentActivity{metadata: metadata}
}

// Metadata returns the activity's metadata
func (a *GetcontentActivity) Metadata() *activity.Metadata {
	return a.metadata
}

func (a *GetcontentActivity) Eval(context activity.Context) (done bool, err error) {

	basicAuthUser, _ := context.GetInput("basicAuthUser").(string)
	basicAuthPassword, _ := context.GetInput("basicAuthPassword").(string)
	url, _ := context.GetInput("URL").(string)

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("GET", url, nil)

	if basicAuthUser != "" {
		req.SetBasicAuth(basicAuthUser, basicAuthPassword)
	}

	// Fetch Request
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Failure : ", err)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)

	// Add Results
	context.SetOutput("result", string(respBody))
	context.SetOutput("status", resp.Status)
	context.SetOutput("header", resp.Header)

	return true, nil
}
