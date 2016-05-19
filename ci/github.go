package ci

//package main

import (
	"fmt"
	//"net/http"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var org = "wothing"

func Githubmain() {
	//client := github.NewClient(nil)
	//orgs, _, err := client.Organizations.List("wothing", nil)

	//orgs, _, err := client.Organizations.GetTeam(1)

	//if err != nil {
	//fmt.Println(err)
	//}
	//fmt.Println(orgs)
	//GetComment()
	client := connection()
	//wathchUser(client)
	//getHooks(client)
	//CreateComment(client, 2, nil)
	//GetComment()
	//PullRequestsList()
	//RepoGet(client)
	ListComments(client)

}
func RepoGet(client *github.Client) {
	gitRepo, _, _ := client.Repositories.Get("izgnod", "test")
	fmt.Println(gitRepo)
}
func ListComments(client *github.Client) {
	//t1, t2, t3 := client.Repositories.ListComments("izgnod", "test", nil)
	t1, t2, t3 := client.PullRequests.GetComment("FuckAll", "test", 3)
	//t1, t2, t3 := client.Repositories.ListComments("rod6", "17mei", nil)
	fmt.Println(t1, t2, t3)

}

func CheckAuth() {

}

func connection() *github.Client {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: ""})
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	return github.NewClient(tc)

}

func GetCommentA() {
	client := connection()
	//fmt.Println(client.Activity.ListEventsForOrganization("wothing", nil))
	//fmt.Println(client.Activity.ListIssueEventsForRepository("FuckAll", "17mei", nil))
	fmt.Println(client.Activity.ListUserEventsForOrganization("woting", "FuckAll", nil))
}

func test() {
	client := connection()
	//client := github.NewClient(nil)
	//fmt.Println(client.Authorizations.List(nil))
	//fmt.Println(client.Issues.GetEvent("FuckAll", "17mei", 1))
	//fmt.Println(client.Issues.List(true, nil))
	//fmt.Println(client.PullRequests.List("rod6", "17mei", nil))
	//fmt.Println(client.Organizations.IsMember("wothing", "FuckAll"))
	//fmt.Println(client.Organizations.IsPublicMember("wothing", "FuckAll"))
	//fmt.Println(client.Organizations.ListMembers("wothing", nil))
	//fmt.Println(client.Organizations.GetHook("wothing", 70160592))
	//fmt.Println(client.Organizations.ListHooks("wothing", nil))
	//hooks, _, _ := client.Repositories.ListHooks("rod6", "wothing/17mei", nil)
	//hooks, _, _ := client.Repositories.ListServiceHooks() //("rod6", "wothing/17mei", nil)
	hooks, _, _ := client.ListServiceHooks()
	fmt.Println(hooks)

}

func PullRequestEvent() {

}
func OrgUser(client *github.Client) ([]github.User, error) {
	members, _, err := client.Organizations.ListMembers(org, nil)
	if err != nil {
		return members, err
	}
	fmt.Println(members)
	return members, nil
}

func getHooks(client *github.Client) error {
	//hooks, _, err := client.Repositories.ListServiceHooks()
	//hooks, _, err := client.Repositories.ListHooks("ord6""wothing", nil)
	//hooks, _, err := client.Organizations.ListHooks("o", nil)
	//hooks, _, err := client.Organizations.Get("wothing")
	hooks, _, err := client.Repositories.ListHooks("rod6", "17mei", nil)
	if err != nil {
		return err
	}
	for _, hook := range hooks {
		fmt.Println(hook)
	}
	return err
}

func CreateComment(client *github.Client, number int, comment *github.PullRequestComment) {
	t1, t2, t3 := client.PullRequests.EditComment("izgond", "test", number, comment)
	fmt.Println(t1, t2, t3)
}

func GetComment() {
	client := connection()
	t1, t2, t3 := client.PullRequests.GetComment("izgnod", "test", 2)
	fmt.Println(t1, t2, t3)
}

func PullRequestsList() {
	client := connection()
	t1, t2, t3 := client.PullRequests.List("izgnod", "test", nil)
	fmt.Println(t1, t2, t3)
	fmt.Println("----------------------")
	fmt.Println(&t1[0].Body)
}
