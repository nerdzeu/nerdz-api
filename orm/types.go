import (
    "database/sql"
    "time"
)

//TODO: tables that don't respect gorm naming conventions
// posts_no_notify, comments_no_notify, comments_notify, ban
// blacklist, whitelist, groups_notify, groups_posts_no_notify,
// groups_comments_no_notify, groups_comments_notify

type User struct {
    Id          int64 //PRIMARY KEY: counter
    Last        time.Time
    NotifyStory string `sql:"type:json"`
    Private     bool
    Lang        string `sql:"type:varchar(2)"`
    Username    string `sql:"type:varchar(90)"`
    Password    string `sql:"type:varchar(40)"`
    Name        string `sql:"type:varchar(60)"`
    Surname     string `sql:"tyoe:varchar(60)"`
    Email       string `sql:"type:varchar(350)"`
    Gender      bool
    BirthDate   time.Time
    BoardLang   string `sql:"type:varchar(2)"`
    Timezone    string `sql:"type:varchar(35)"`
    Viewonline  bool
}

type Profile struct {
    Id             int64 //PRIMARY KEY: counter
    RemoteAddr     string `sql:"type:inet"`
    HttpUserAgent  string `sql:"type:text"`
    Website        string `sql:"type:varchar(350)"`
    Quotes         string `sql:"type:text"`
    Biography      string `sql:"type:text"`
    Interests      string `sql:"type:text"`
    Github         string `sql:"type:varchar(350)"`
    Skype          string `sql:"type:varchar(350)"`
    Jabber         string `sql:"type:varchar(350)"`
    Yahoo          string `sql:"type:varchar(350)"`
    Userscript     string `sql:"type:varchar(128)"`
    Template       int16
    MobileTemplate int16
    Dateformat     string `sql:"type:varchar(25)"`
    Facebook       string `sql:"type:varchar(350)"`
    Twitter        string `sql:"type:varchar(350)"`
    Steam          string `sql:"type:varchar(350)"`
    Push           bool
    Pushregtime    time.Time
}

type ClosedProfile struct {
    Id     int64 //PRIMARY KEY: counter
}

type Post struct {
    Id      int64 //PRIMARY KEY: hpid
    From    int64
    To      int64
    Pid     int64
    Message string `sql:"type:text"`
    Notify  bool
    Time    time.Time
}

//TODO: *(no_)?_notify - and all other tables with singular name in db defintion

type Thumb struct {
    Hpid  int64
    User  int64
    Vote  int16
}

type Lurker struct {
    User int64
    Post int64
    Time time.Time
}

type Comment struct {
    Id      int64 //PRIMARY KEY: hcid
    Hpid    int64
    From    int64
    To      int64
    Message string `sql:"type:text"`
    Time    time.Time
}

type Bookmark struct {
    Hpid int64
    From int64
    Time time.Time
}

type Pm struct {
    Id       int64 //PRIMARY KEY:  pmid
    Pmid     int64
    From     int64
    To       int64
    Pid      int64
    Message  string `sql:"type:text"`
    Read     bool
    Time    time.Time
}

type Group struct {
    Id          int64 //PRIMARY KEY: counter
    Description string `sql:"type:text"`
    Owner       int64
    Name        string `sql:"type:varchar(30)"`
    Private     bool
    Photo       string `sql:"type:varchar(350)"`
    Website     string `sql:"type:varchar(350)"`
    Goal        string `sql:"type:text"`
    Visible     bool
    Open        bool
}

type GroupsMember {
    Group int64
    User  int64
}

type GroupsPost struct {
    Id      int64 //PRIMARY KEY: hpid
    From    int64
    To      int64
    Pid     int64
    Message string `sql:"type:text"`
    News    bool
    Time    time.Time
}

type GroupsThumb struct {
    Hpid   int64
    User   int64
    Vote   int16
}

type GroupsLurker struct {
    User int64
    Post int64
    Time time.Time
}

type GroupsComment struct {
    Id      int64 //PRIMARY KEY: hcid
    Hpid    int64
    From    int64
    To      int64
    Message string `sql:"type:text"`
    Time    time.Time
}

type GroupsBookmark struct {
    Hpid int64
    From int64
    Time time.Time
}

type GroupsFollower struct {
    Group int64
    User  int64
}

type CommentThumb struct {
    Hcid int64
    User int64
    Vote int16
}

type GroupsCommentThumb struct {
    Hcid int64
    User int64
    Vote int16
}