package nerdz_test

import (
	"fmt"
	"github.com/nerdzeu/nerdz-api/nerdz"
	"testing"
	"time"
)

var me, other, blacklisted, withClosedProfile *nerdz.User

func init() {
	blacklisted, _ = nerdz.NewUser(5)
	withClosedProfile, _ = nerdz.NewUser(7)

	var err error

	me, err = nerdz.NewUser(1)
	if err != nil {
		panic(fmt.Sprintf("No error should happen when create existing user, but got: %+v", err))
	}

	other, err = nerdz.NewUser(2)
	if err != nil {
		panic(fmt.Sprintf("No error should happen when create existing user, but got: %+v", err))
	}
}

func TestContactInfo(t *testing.T) {
	info := me.ContactInfo()
	if info == nil {
		t.Error("null info")
	}
}

func TestPersonalInfo(t *testing.T) {
	info := me.PersonalInfo()
	if info == nil {
		t.Error("null info")
	}

	t.Logf("Struct: %+v\nINTERESTES:", *info)
	for i, elem := range info.Interests {
		t.Logf("%d) %s\n", i, elem)
	}

	t.Log("Quotes:\n")
	for i, elem := range info.Quotes {
		t.Logf("%d) %s\n", i, elem)
	}
}

func TestBoardInfo(t *testing.T) {
	info := me.BoardInfo()
	if info == nil {
		t.Error("null info")
	}

	// If whitelist is not empty, the output will be huge (if tested with -v flag)
	t.Logf("%+v\n", *info)
	t.Logf("Template: %+v", *info.Template)
}

func TestBlackList(t *testing.T) {
	bl := me.Blacklist()
	if len(bl) != 1 {
		t.Errorf("Expected 1 user in blacklist, but got: %v\n", len(bl))
	}
}

func TestHome(t *testing.T) {
	// At most the last 10 posts from italian users
	userHome := me.UserHome(&nerdz.PostlistOptions{Following: false, Language: "it", N: 10})
	if len(*userHome) != 10 {
		t.Errorf("Expected 10 posts, but got: %+v\n", len(*userHome))
	}

	t.Logf("%+v\n", *userHome)

	// At most the last 10 project posts from italian users
	projectHome := me.ProjectHome(&nerdz.PostlistOptions{Following: false, Language: "it", N: 10})
	if len(*projectHome) != 10 {
		t.Errorf("Expected 10 posts, but got: %+v\n", len(*projectHome))
	}

	t.Logf("%+v\n", *projectHome)

	// At most the last 10 posts from German users
	userHome = me.UserHome(&nerdz.PostlistOptions{Following: false, Language: "de", N: 10})
	if len(*userHome) != 0 {
		t.Errorf("Expected 0 posts, but got: %+v\n", len(*userHome))
	}

	t.Logf("%+v\n", *userHome)

	// At most the last 10 posts to English users from users that "user" is following
	userHome = me.UserHome(&nerdz.PostlistOptions{Following: true, Language: "en", N: 10})

	if len(*userHome) == 0 {
		t.Error("Expected at leat 1 post from an english user the 'user' is following. But 0 found")
	}

	t.Logf("%+v\n", *userHome)

	// The single post older (created before) the one with hpid 1000, from some user that 'user' follow and to an english speaking one
	userHome = me.UserHome(&nerdz.PostlistOptions{Following: true, Language: "en", N: 1, Older: 1000})

	if len(*userHome) != 1 {
		t.Errorf("Expeted 1 post, but got: %d", len(*userHome))
	}

	t.Logf("THE POST: %+v", (*userHome)[0])

	if (*userHome)[0].Hpid != 36 {
		t.Errorf("Post with hpid 36 expected, but got: %d", (*userHome)[0].Hpid)
	}

	// At most 2 posts in the Homepage formed by my posts and my friends posts
	userHome = me.UserHome(&nerdz.PostlistOptions{Following: true, Followers: true, N: 2})

	if len(*userHome) != 2 {
		t.Errorf("Expeted 2 posts, but got: %d", len(*userHome))
	}

	t.Logf("FRIENDZ: %v", *userHome)

	lastFriendPost := (*userHome)[0]

	// Get the (at max 20, in this case only 1) newer posts than the one with the "Newer" from friends
	userHome = me.UserHome(&nerdz.PostlistOptions{
		Following: true,
		Followers: true,
		Newer:     (*userHome)[1].Hpid})

	if len(*userHome) > 1 || (*userHome)[0].Hpid != lastFriendPost.Hpid {
		t.Errorf("Expected 1 post with hpid %d, but got %d posts and the first post has hpid = %d", lastFriendPost.Hpid, len(*userHome), (*userHome)[0].Hpid)
	}
}

func TestUserPostlist(t *testing.T) {
	postList := me.Postlist(nil).([]nerdz.UserPost)
	if len(postList) != 20 {
		t.Errorf("Expected 20  posts, but got: %+v\n", len(postList))
	}

	// Older than 1 (all) and newer than 8000 (no one) -> empty
	postList = me.Postlist(&nerdz.PostlistOptions{
		Older: 1,
		Newer: 80000}).([]nerdz.UserPost)

	if len(postList) != 0 {
		t.Errorf("Expected 0 posts. But got: %d", len(postList))
	}

	// Find posts between 103 and 97 inclusive, in user profile, from everybody.
	postList = me.Postlist(&nerdz.PostlistOptions{
		Older: 103,
		Newer: 97,
	}).([]nerdz.UserPost)

	if len(postList) != 4 {
		t.Errorf("Expected 4 posts. But got: %d", len(postList))
	}
}

func TestAddEditDeleteUserPost(t *testing.T) {
	var post nerdz.UserPost

	// New post on my board (To = 0)
	post.Message = "All right"
	post.Lang = "en"
	if err := me.Add(&post); err != nil {
		t.Errorf("Add user post should work but, got: %v", err)
	}

	if err := me.Delete(&post); err != nil {
		t.Errorf("Delete with hpid %v shoud work, but got error: %v", post.Hpid, err)
	}

	post.Message = "All right2"
	post.Lang = "en"

	if err := me.Add(&post); err != nil {
		t.Errorf("Add with ID should work but, got: %v", err)
	}

	post.Message = "Post updated -> :D\nwow JA JA JA"
	post.Lang = "fu"
	// Language "fu" does not exists, this edit should fail
	if err := me.Edit(&post); err == nil {
		t.Errorf("Edit post language and message not failed! - %v", err)
	}

	post.Lang = "de"
	if err := me.Edit(&post); err != nil {
		t.Errorf("This edit shold work but got %s", err)
	}

	oldHpid := post.Hpid
	post.Hpid = 0 //default value for uint64
	if err := me.Delete(&post); err == nil {
		t.Errorf("Delete with hpid 0 should fail")
	}

	post.Hpid = oldHpid
	if err := me.Delete(&post); err != nil {
		t.Errorf("Delete a valid post should work")
	}

}

func TestAddEditDeleteUserPostComment(t *testing.T) {
	existingPost := me.Postlist(&nerdz.PostlistOptions{N: 1}).([]nerdz.UserPost)[0]

	var comment nerdz.UserPostComment
	comment.Message = "Nice <html>"
	comment.Hpid = existingPost.Hpid

	if err := me.Add(&comment); err != nil {
		t.Errorf("Add failed: %s", err)
	}

	comment.Message = "LOL EDIT"
	// Wait 5 second to avoid flood limit (db side)
	time.Sleep(5000 * time.Millisecond)
	if err := me.Edit(&comment); err != nil {
		t.Errorf("Edit comment failed with error: %s", err)
	}

	if err := me.Delete(&comment); err != nil {
		t.Errorf("Delete comment with hcid %v shoud work, but got error: %v", comment.Hcid, err)
	}
}

func TestAddEditDeleteProjectPost(t *testing.T) {
	var post nerdz.ProjectPost

	myProject := me.Projects()[0]
	post.To = myProject.Counter
	post.Message = "BEST ADMIN EVER :>\nHello!"
	post.Lang = "en"

	if err := me.Add(&post); err != nil {
		t.Errorf("No errors should occur whie adding a post to a project of mine, but got: %v", err)
	}

	post.Message = "WORST ADMIN EVER :<\a <- some random character"
	if err := me.Edit(&post); err != nil {
		t.Errorf("Project Post edit should work, but failed with error: %s\n", err)
	}

	if err := me.Delete(&post); err != nil {
		t.Errorf("Delete failed with error: %s", err.Error())
	}
}

func TestAddEditDeleteProjectPostComment(t *testing.T) {
	myProject := me.Projects()[0]
	projectPost := myProject.Postlist(&nerdz.PostlistOptions{N: 1}).([]nerdz.ProjectPost)[0]

	var projectPostComment nerdz.ProjectPostComment
	projectPostComment.Hpid = projectPost.Hpid
	projectPostComment.Message = "lol k"

	if err := me.Add(&projectPostComment); err != nil {
		t.Errorf("Add comment on an existing project post sould work but failed with error: %s", err.Error())
	}

	projectPostComment.Message = "lol, k"
	// Wait 5 second to avoid flood limit (db side)
	time.Sleep(5000 * time.Millisecond)
	if err := me.Edit(&projectPostComment); err != nil {
		t.Errorf("Edit project post comment failed with error: %s", err)
	}

	if err := me.Delete(&projectPostComment); err != nil {
		t.Errorf("Delete with hcid %v shoud work, but got error: %v", projectPostComment.Hcid, err)
	}
}

func TestAddEditDeletePm(t *testing.T) {
	var pm nerdz.Pm

	pm.Message = "Hi bro. Join telegram now"
	pm.To = withClosedProfile.Counter

	if err := me.Add(&pm); err != nil {
		t.Errorf("No errors should occur while adding a new pm to a non blacklisted user, but got %v", err)
	}

	pm.Message = "Pm edit is impossible (since in IM messages are not editable)"
	if err := me.Edit(&pm); err == nil {
		t.Errorf("Pm edit shouldn't work")
	}

	if err := me.Delete(&pm); err != nil {
		t.Errorf("Pm delete failed with error: %s", err.Error())
	}
}

func TestFollowUser(t *testing.T) {
	other, _ = nerdz.NewUser(3)

	t.Logf("User(%d) follows User(%d)", me.Counter, other.Counter)

	oldNumFollowers := len(other.NumericFollowers())

	if err := me.Follow(other); err != nil {
		t.Log("The user should correctly follow the other user but: ")
		t.Error(err)
	}

	if len(other.NumericFollowers()) != oldNumFollowers+1 {
		t.Log("There isn't a new follower for the user!")
		t.Error("No new follower")
	}
}

func TestFollowProject(t *testing.T) {
	project, _ := nerdz.NewProject(1)

	t.Log("I want to follow a fantastic project whose name is: ", project.Name)
	oldNumFollowers := len(project.NumericFollowers())

	if err := me.Follow(project); err != nil {
		t.Log("The user should correctly follow the project but: ")
		t.Error(err)
	}

	if len(project.NumericFollowers()) != oldNumFollowers+1 {
		t.Log("There isn't a new follower for the project!")
		t.Error("No new follower")
	}
}

func TestUnfollowUser(t *testing.T) {
	other, _ = nerdz.NewUser(3)
	t.Logf("User(%d) unfollows User(%d)", me.Counter, other.Counter)

	oldNumFollowers := len(other.NumericFollowers())

	if err := me.Unfollow(other); err != nil {
		t.Error(err)
	}

	newNumFollowers := len(other.NumericFollowers())

	if newNumFollowers != oldNumFollowers-1 {
		t.Errorf("The follower isn't removed from the followers list! (old %d, new %d)", oldNumFollowers, newNumFollowers)
	}
}

func TestUnfollowProject(t *testing.T) {
	project, _ := nerdz.NewProject(2)

	t.Log("I want to unfollow a useless project whose name is: ", project.Name)
	oldNumFollowers := len(project.Followers())

	if err := me.Unfollow(project); err != nil {
		t.Error(err)
	}

	if len(project.Followers()) != oldNumFollowers-1 {
		t.Error("The follower isn't removed from the project's followers!")
	}
}

func TestUserPostBookmark(t *testing.T) {
	post, _ := nerdz.NewUserPost(13)

	t.Logf("User(%d) bookmarks the user's post(%d) ", me.Counter, post.Hpid)

	oldNumBookmarkers := len(post.NumericBookmarkers())

	if err := me.Bookmark(post); err != nil {
		t.Error(err)
	}

	if len(post.NumericBookmarkers()) != oldNumBookmarkers+1 {
		t.Error("There isn't a new bookmark for the user's post ", post.Hpid)
	}
}

func TestUserPostUnbookmark(t *testing.T) {
	post, _ := nerdz.NewUserPost(13)

	t.Logf("User(%d) unbookmarks the user's post(%d) ", me.Counter, post.Hpid)

	oldNumBookmarkers := len(post.NumericBookmarkers())

	if err := me.Unbookmark(post); err != nil {
		t.Error(err)
	}

	if len(post.NumericBookmarkers()) != oldNumBookmarkers-1 {
		t.Error("Bookmark isn't removed for the user's post ", post.Hpid)
	}
}

func TestProjectPostBookmark(t *testing.T) {
	post, _ := nerdz.NewProjectPost(2)

	t.Logf("User(%d) bookmarks the project's post(%d) ", me.Counter, post.Hpid)

	oldNumBookmarkers := len(post.NumericBookmarkers())

	if err := me.Bookmark(post); err != nil {
		t.Error(err)
	}

	if len(post.NumericBookmarkers()) != oldNumBookmarkers+1 {
		t.Error("There isn't a new bookmark for the project's post ", post.Hpid)
	}
}

func TestProjectPostUnbookmark(t *testing.T) {
	post, _ := nerdz.NewProjectPost(2)

	t.Logf("User(%d) unbookmarks the project's post(%d) ", me.Counter, post.Hpid)

	oldNumBookmarkers := len(post.NumericBookmarkers())

	if err := me.Unbookmark(post); err != nil {
		t.Error(err)
	}

	if len(post.NumericBookmarkers()) != oldNumBookmarkers-1 {
		t.Error("Bookmark isn't removed for the project ", post.Hpid)
	}
}

func TestPms(t *testing.T) {
	other, _ = nerdz.NewUser(2)
	t.Logf("User(%d) -pm-> User(%d)", me.Counter, other.Counter)
	// build a pm configuration in order to filter results
	pmConf := nerdz.NewPmConfig().WithDescOrder(true)

	pmList, err := me.Pms(other.Counter, pmConf)

	if err != nil {
		t.Errorf("Error trying to get pms between user(%s) and user(%s) - %v", me.Id(), other.Id(), err)
		return
	}

	t.Log("####### PMS  ########")

	for _, val := range *pmList {
		t.Logf("%+v", val)
	}

	t.Log("####################")

	pmConf = nerdz.NewPmConfig().WithOffset(2).WithLimit(4)

	pmListR, errR := me.Pms(other.Counter, pmConf)

	if errR != nil {
		t.Errorf("Error trying to get pms between user(%s) and user(%s) - %v", me.Id(), other.Id(), errR)
		return
	}

	t.Log("####### PMS between (2 - 4) ########")

	for _, val := range *pmListR {
		t.Logf("%+v", val)
	}

	t.Log("####################")

}

func TestConversation(t *testing.T) {
	t.Logf("Looking for conversation for user(%d)", me.Counter)

	convList, err := me.Conversations()

	if err != nil {
		t.Errorf("No private conversations available for user(%d)", me.Counter)
	}

	t.Logf("########## Conversations ###########")
	for _, val := range *convList {
		t.Log(val)

	}

	t.Logf("####################################")
}

func TestDoThumbs(t *testing.T) {
	userPost, _ := nerdz.NewUserPost(13)

	t.Logf("user(%d) likes user post(%d)", me.Counter, userPost.Hpid)

	if err := me.ThumbUp(userPost); err != nil {
		t.Errorf("User is unable to like user post - %v", err)
	}

	projPost, _ := nerdz.NewProjectPost(2)

	t.Logf("user(%d) likes project post(%d)", me.Counter, projPost.Hpid)

	if err := me.ThumbUp(projPost); err != nil {
		t.Errorf("User is unable to like project post - %v", err)
	}

}
