package main

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"gopkg.in/go-playground/webhooks.v2"
	ghub "gopkg.in/go-playground/webhooks.v2/github"
)

const SECRET = "mutHupGhemopholOwdoocHuckmidBor4"

const BotToken = "760c6348f92cd0885737cf2e4a78dd0582915941"
const EmaxToken = "c08ab4c23359d4aa512bac979de246d7217bbf31"

func main() {
	fmt.Println("vim-go")
	WH()

}

func WH() {
	hook := ghub.New(&ghub.Config{Secret: ""})
	//hook.RegisterEvents(HandlePullRequest, ghub.PullRequestEvent)
	//hook.RegisterEvents(HandlePush, ghub.PushEvent)
	//hook.RegisterEvents(HandlePullRequestComments, ghub.PullRequestReviewCommentEvent)
	//hook.RegisterEvents(HandlePush, ghub.CommitCommentEvent)
	//
	hook.RegisterEvents(HandleAll, ghub.CommitCommentEvent, ghub.CreateEvent, ghub.PullRequestEvent, ghub.PullRequestReviewCommentEvent, ghub.PushEvent, ghub.IssueCommentEvent)

	err := webhooks.Run(hook, ":"+strconv.Itoa(8080), "/")
	if err != nil {
		fmt.Println(err)
	}
}

// HandlePullRequest handles GitHub pull_request events
func HandlePullRequest(payload interface{}, header webhooks.Header) {

	fmt.Println("Handling Pull Request")

	pl := payload.(ghub.PullRequestPayload)

	// Do whatever you want from here...
	fmt.Printf("%+v", pl)
}

// HandlePullRequest handles GitHub pull_request events
func HandleAll(payload interface{}, header webhooks.Header) {

	fmt.Println("Handling all")
	fmt.Printf("payload = %+v\n", payload)

	pl := payload.(ghub.PullRequestPayload)

	// Do whatever you want from here...
	fmt.Printf("%+v", pl)
}

// HandlePullRequest handles GitHub pull_request events
func HandlePullRequestComments(payload interface{}, header webhooks.Header) {

	fmt.Println("Handling Pull Request Comment")

	pl := payload.(ghub.PullRequestReviewCommentPayload)

	// Do whatever you want from here...
	fmt.Printf("%+v", pl)
}

// HandlePullRequest handles GitHub pull_request events
func HandlePush(payload interface{}, header webhooks.Header) {

	fmt.Println("Handling Push ")

	pl := payload.(ghub.PushPayload)

	// Do whatever you want from here...
	fmt.Printf("%+v", pl)
}

func Ver1() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		fmt.Printf("r.URL.Path = %+v\n", r.URL.Path)
		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Print(err)
		}
		fmt.Printf("body = %+v\n", string(body))
		event := map[string]interface{}{}
		err = json.Unmarshal(body, &event)
		if err != nil {
			log.Print(err)
		}
		res, err := json.MarshalIndent(event, "", "    ")
		if err != nil {
			log.Print(err)
		}
		fmt.Println(string(res))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func tag() {
	//ctx := context.Background()
	//ts := oauth2.StaticTokenSource(
	//&oauth2.Token{AccessToken: EmaxToken},
	//)
	//tc := oauth2.NewClient(ctx, ts)

	//client := github.NewClient(tc)
	//client.Repositories.ListTags()

	// list all repositories for the authenticated user
	//repos, _, err := client.Repositories.List(ctx, "", nil)
}
