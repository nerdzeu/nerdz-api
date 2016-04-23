package doc

type ProjectsIdMembers struct {
	Data []struct {
		Contacts struct {
			Facebook string `json:"facebook"`
			Github   string `json:"github"`
			Jabber   string `json:"jabber"`
			Skype    string `json:"skype"`
			Steam    string `json:"steam"`
			Twitter  string `json:"twitter"`
			Website  string `json:"website"`
			Yahoo    string `json:"yahoo"`
		} `json:"contacts"`
		Info struct {
			Board    string      `json:"board"`
			Closed   bool        `json:"closed"`
			ID       int         `json:"id"`
			Image    string      `json:"image"`
			Name     string      `json:"name"`
			Owner    interface{} `json:"owner"`
			Type     string      `json:"type"`
			Username string      `json:"username"`
			Website  string      `json:"website"`
		} `json:"info"`
		Personal struct {
			Biography string      `json:"biography"`
			Birthday  string      `json:"birthday"`
			Gender    bool        `json:"gender"`
			Gravatar  string      `json:"gravatar"`
			ID        int         `json:"id"`
			Interests interface{} `json:"interests"`
			Name      string      `json:"name"`
			Nation    string      `json:"nation"`
			Online    bool        `json:"online"`
			Quotes    []string    `json:"quotes"`
			Surname   string      `json:"surname"`
			Timezone  string      `json:"timezone"`
			Username  string      `json:"username"`
		} `json:"personal"`
	} `json:"data"`
	HumanMessage string `json:"humanMessage"`
	Message      string `json:"message"`
	Status       int    `json:"status"`
	Success      bool   `json:"success"`
}

type Me struct {
	Data struct {
		Contacts struct {
			Facebook string `json:"facebook"`
			Github   string `json:"github"`
			Jabber   string `json:"jabber"`
			Skype    string `json:"skype"`
			Steam    string `json:"steam"`
			Twitter  string `json:"twitter"`
			Website  string `json:"website"`
			Yahoo    string `json:"yahoo"`
		} `json:"contacts"`
		Info struct {
			Board    string      `json:"board"`
			Closed   bool        `json:"closed"`
			ID       int         `json:"id"`
			Image    string      `json:"image"`
			Name     string      `json:"name"`
			Owner    interface{} `json:"owner"`
			Type     string      `json:"type"`
			Username string      `json:"username"`
			Website  string      `json:"website"`
		} `json:"info"`
		Personal struct {
			Biography string   `json:"biography"`
			Birthday  string   `json:"birthday"`
			Gender    bool     `json:"gender"`
			Gravatar  string   `json:"gravatar"`
			ID        int      `json:"id"`
			Interests []string `json:"interests"`
			Name      string   `json:"name"`
			Nation    string   `json:"nation"`
			Online    bool     `json:"online"`
			Quotes    []string `json:"quotes"`
			Surname   string   `json:"surname"`
			Timezone  string   `json:"timezone"`
			Username  string   `json:"username"`
		} `json:"personal"`
	} `json:"data"`
	HumanMessage string `json:"humanMessage"`
	Message      string `json:"message"`
	Status       int    `json:"status"`
	Success      bool   `json:"success"`
}

type MeBlacklist struct {
	Data []struct {
		Contacts struct {
			Facebook string `json:"facebook"`
			Github   string `json:"github"`
			Jabber   string `json:"jabber"`
			Skype    string `json:"skype"`
			Steam    string `json:"steam"`
			Twitter  string `json:"twitter"`
			Website  string `json:"website"`
			Yahoo    string `json:"yahoo"`
		} `json:"contacts"`
		Info struct {
			Board    string      `json:"board"`
			Closed   bool        `json:"closed"`
			ID       int         `json:"id"`
			Image    string      `json:"image"`
			Name     string      `json:"name"`
			Owner    interface{} `json:"owner"`
			Type     string      `json:"type"`
			Username string      `json:"username"`
			Website  string      `json:"website"`
		} `json:"info"`
		Personal struct {
			Biography string      `json:"biography"`
			Birthday  string      `json:"birthday"`
			Gender    bool        `json:"gender"`
			Gravatar  string      `json:"gravatar"`
			ID        int         `json:"id"`
			Interests interface{} `json:"interests"`
			Name      string      `json:"name"`
			Nation    string      `json:"nation"`
			Online    bool        `json:"online"`
			Quotes    []string    `json:"quotes"`
			Surname   string      `json:"surname"`
			Timezone  string      `json:"timezone"`
			Username  string      `json:"username"`
		} `json:"personal"`
	} `json:"data"`
	HumanMessage string `json:"humanMessage"`
	Message      string `json:"message"`
	Status       int    `json:"status"`
	Success      bool   `json:"success"`
}

type MePostsPid struct {
	Data struct {
		Bookmarkers int  `json:"bookmarkers"`
		CanBookmark bool `json:"canBookmark"`
		CanComment  bool `json:"canComment"`
		CanDelete   bool `json:"canDelete"`
		CanEdit     bool `json:"canEdit"`
		CanLurk     bool `json:"canLurk"`
		Closed      bool `json:"closed"`
		Comments    int  `json:"comments"`
		From        struct {
			Board    string      `json:"board"`
			Closed   bool        `json:"closed"`
			ID       int         `json:"id"`
			Image    string      `json:"image"`
			Name     string      `json:"name"`
			Owner    interface{} `json:"owner"`
			Type     string      `json:"type"`
			Username string      `json:"username"`
			Website  string      `json:"website"`
		} `json:"from"`
		Hpid      int    `json:"hpid"`
		Lang      string `json:"lang"`
		Lurkers   int    `json:"lurkers"`
		Message   string `json:"message"`
		News      bool   `json:"news"`
		Pid       int    `json:"pid"`
		Rate      int    `json:"rate"`
		Revisions int    `json:"revisions"`
		Time      string `json:"time"`
		Timestamp int    `json:"timestamp"`
		To        struct {
			Board    string      `json:"board"`
			Closed   bool        `json:"closed"`
			ID       int         `json:"id"`
			Image    string      `json:"image"`
			Name     string      `json:"name"`
			Owner    interface{} `json:"owner"`
			Type     string      `json:"type"`
			Username string      `json:"username"`
			Website  string      `json:"website"`
		} `json:"to"`
		Type string `json:"type"`
		URL  string `json:"url"`
	} `json:"data"`
	HumanMessage string `json:"humanMessage"`
	Message      string `json:"message"`
	Status       int    `json:"status"`
	Success      bool   `json:"success"`
}

type ProjectsId struct {
	Data struct {
		Board  string `json:"board"`
		Closed bool   `json:"closed"`
		ID     int    `json:"id"`
		Image  string `json:"image"`
		Name   string `json:"name"`
		Owner  struct {
			Board    string      `json:"board"`
			Closed   bool        `json:"closed"`
			ID       int         `json:"id"`
			Image    string      `json:"image"`
			Name     string      `json:"name"`
			Owner    interface{} `json:"owner"`
			Type     string      `json:"type"`
			Username string      `json:"username"`
			Website  string      `json:"website"`
		} `json:"owner"`
		Type     string `json:"type"`
		Username string `json:"username"`
		Website  string `json:"website"`
	} `json:"data"`
	HumanMessage string `json:"humanMessage"`
	Message      string `json:"message"`
	Status       int    `json:"status"`
	Success      bool   `json:"success"`
}

type MeWhitelisting struct {
	Data []struct {
		Contacts struct {
			Facebook string `json:"facebook"`
			Github   string `json:"github"`
			Jabber   string `json:"jabber"`
			Skype    string `json:"skype"`
			Steam    string `json:"steam"`
			Twitter  string `json:"twitter"`
			Website  string `json:"website"`
			Yahoo    string `json:"yahoo"`
		} `json:"contacts"`
		Info struct {
			Board    string      `json:"board"`
			Closed   bool        `json:"closed"`
			ID       int         `json:"id"`
			Image    string      `json:"image"`
			Name     string      `json:"name"`
			Owner    interface{} `json:"owner"`
			Type     string      `json:"type"`
			Username string      `json:"username"`
			Website  string      `json:"website"`
		} `json:"info"`
		Personal struct {
			Biography string      `json:"biography"`
			Birthday  string      `json:"birthday"`
			Gender    bool        `json:"gender"`
			Gravatar  string      `json:"gravatar"`
			ID        int         `json:"id"`
			Interests interface{} `json:"interests"`
			Name      string      `json:"name"`
			Nation    string      `json:"nation"`
			Online    bool        `json:"online"`
			Quotes    []string    `json:"quotes"`
			Surname   string      `json:"surname"`
			Timezone  string      `json:"timezone"`
			Username  string      `json:"username"`
		} `json:"personal"`
	} `json:"data"`
	HumanMessage string `json:"humanMessage"`
	Message      string `json:"message"`
	Status       int    `json:"status"`
	Success      bool   `json:"success"`
}

type MeWhitelist struct {
	Data []struct {
		Contacts struct {
			Facebook string `json:"facebook"`
			Github   string `json:"github"`
			Jabber   string `json:"jabber"`
			Skype    string `json:"skype"`
			Steam    string `json:"steam"`
			Twitter  string `json:"twitter"`
			Website  string `json:"website"`
			Yahoo    string `json:"yahoo"`
		} `json:"contacts"`
		Info struct {
			Board    string      `json:"board"`
			Closed   bool        `json:"closed"`
			ID       int         `json:"id"`
			Image    string      `json:"image"`
			Name     string      `json:"name"`
			Owner    interface{} `json:"owner"`
			Type     string      `json:"type"`
			Username string      `json:"username"`
			Website  string      `json:"website"`
		} `json:"info"`
		Personal struct {
			Biography string      `json:"biography"`
			Birthday  string      `json:"birthday"`
			Gender    bool        `json:"gender"`
			Gravatar  string      `json:"gravatar"`
			ID        int         `json:"id"`
			Interests interface{} `json:"interests"`
			Name      string      `json:"name"`
			Nation    string      `json:"nation"`
			Online    bool        `json:"online"`
			Quotes    []string    `json:"quotes"`
			Surname   string      `json:"surname"`
			Timezone  string      `json:"timezone"`
			Username  string      `json:"username"`
		} `json:"personal"`
	} `json:"data"`
	HumanMessage string `json:"humanMessage"`
	Message      string `json:"message"`
	Status       int    `json:"status"`
	Success      bool   `json:"success"`
}

type MeBlacklisting struct {
	Data []struct {
		Contacts struct {
			Facebook string `json:"facebook"`
			Github   string `json:"github"`
			Jabber   string `json:"jabber"`
			Skype    string `json:"skype"`
			Steam    string `json:"steam"`
			Twitter  string `json:"twitter"`
			Website  string `json:"website"`
			Yahoo    string `json:"yahoo"`
		} `json:"contacts"`
		Info struct {
			Board    string      `json:"board"`
			Closed   bool        `json:"closed"`
			ID       int         `json:"id"`
			Image    string      `json:"image"`
			Name     string      `json:"name"`
			Owner    interface{} `json:"owner"`
			Type     string      `json:"type"`
			Username string      `json:"username"`
			Website  string      `json:"website"`
		} `json:"info"`
		Personal struct {
			Biography string      `json:"biography"`
			Birthday  string      `json:"birthday"`
			Gender    bool        `json:"gender"`
			Gravatar  string      `json:"gravatar"`
			ID        int         `json:"id"`
			Interests interface{} `json:"interests"`
			Name      string      `json:"name"`
			Nation    string      `json:"nation"`
			Online    bool        `json:"online"`
			Quotes    []string    `json:"quotes"`
			Surname   string      `json:"surname"`
			Timezone  string      `json:"timezone"`
			Username  string      `json:"username"`
		} `json:"personal"`
	} `json:"data"`
	HumanMessage string `json:"humanMessage"`
	Message      string `json:"message"`
	Status       int    `json:"status"`
	Success      bool   `json:"success"`
}

type UsersIdPostsPidCommentsCid struct {
	Data struct {
		CanDelete bool `json:"canDelete"`
		CanEdit   bool `json:"canEdit"`
		From      struct {
			Board    string      `json:"board"`
			Closed   bool        `json:"closed"`
			ID       int         `json:"id"`
			Image    string      `json:"image"`
			Name     string      `json:"name"`
			Owner    interface{} `json:"owner"`
			Type     string      `json:"type"`
			Username string      `json:"username"`
			Website  string      `json:"website"`
		} `json:"from"`
		Hcid    int    `json:"hcid"`
		Hpid    int    `json:"hpid"`
		Message string `json:"message"`
		Time    string `json:"time"`
		To      struct {
			Board    string      `json:"board"`
			Closed   bool        `json:"closed"`
			ID       int         `json:"id"`
			Image    string      `json:"image"`
			Name     string      `json:"name"`
			Owner    interface{} `json:"owner"`
			Type     string      `json:"type"`
			Username string      `json:"username"`
			Website  string      `json:"website"`
		} `json:"to"`
	} `json:"data"`
	HumanMessage string `json:"humanMessage"`
	Message      string `json:"message"`
	Status       int    `json:"status"`
	Success      bool   `json:"success"`
}

type UsersIdFollowing struct {
	Data []struct {
		Contacts struct {
			Facebook string `json:"facebook"`
			Github   string `json:"github"`
			Jabber   string `json:"jabber"`
			Skype    string `json:"skype"`
			Steam    string `json:"steam"`
			Twitter  string `json:"twitter"`
			Website  string `json:"website"`
			Yahoo    string `json:"yahoo"`
		} `json:"contacts"`
		Info struct {
			Board    string      `json:"board"`
			Closed   bool        `json:"closed"`
			ID       int         `json:"id"`
			Image    string      `json:"image"`
			Name     string      `json:"name"`
			Owner    interface{} `json:"owner"`
			Type     string      `json:"type"`
			Username string      `json:"username"`
			Website  string      `json:"website"`
		} `json:"info"`
		Personal struct {
			Biography string      `json:"biography"`
			Birthday  string      `json:"birthday"`
			Gender    bool        `json:"gender"`
			Gravatar  string      `json:"gravatar"`
			ID        int         `json:"id"`
			Interests interface{} `json:"interests"`
			Name      string      `json:"name"`
			Nation    string      `json:"nation"`
			Online    bool        `json:"online"`
			Quotes    []string    `json:"quotes"`
			Surname   string      `json:"surname"`
			Timezone  string      `json:"timezone"`
			Username  string      `json:"username"`
		} `json:"personal"`
	} `json:"data"`
	HumanMessage string `json:"humanMessage"`
	Message      string `json:"message"`
	Status       int    `json:"status"`
	Success      bool   `json:"success"`
}

type ProjectsIdFollowers struct {
	Data []struct {
		Contacts struct {
			Facebook string `json:"facebook"`
			Github   string `json:"github"`
			Jabber   string `json:"jabber"`
			Skype    string `json:"skype"`
			Steam    string `json:"steam"`
			Twitter  string `json:"twitter"`
			Website  string `json:"website"`
			Yahoo    string `json:"yahoo"`
		} `json:"contacts"`
		Info struct {
			Board    string      `json:"board"`
			Closed   bool        `json:"closed"`
			ID       int         `json:"id"`
			Image    string      `json:"image"`
			Name     string      `json:"name"`
			Owner    interface{} `json:"owner"`
			Type     string      `json:"type"`
			Username string      `json:"username"`
			Website  string      `json:"website"`
		} `json:"info"`
		Personal struct {
			Biography string   `json:"biography"`
			Birthday  string   `json:"birthday"`
			Gender    bool     `json:"gender"`
			Gravatar  string   `json:"gravatar"`
			ID        int      `json:"id"`
			Interests []string `json:"interests"`
			Name      string   `json:"name"`
			Nation    string   `json:"nation"`
			Online    bool     `json:"online"`
			Quotes    []string `json:"quotes"`
			Surname   string   `json:"surname"`
			Timezone  string   `json:"timezone"`
			Username  string   `json:"username"`
		} `json:"personal"`
	} `json:"data"`
	HumanMessage string `json:"humanMessage"`
	Message      string `json:"message"`
	Status       int    `json:"status"`
	Success      bool   `json:"success"`
}

type MeFriends struct {
	Data []struct {
		Contacts struct {
			Facebook string `json:"facebook"`
			Github   string `json:"github"`
			Jabber   string `json:"jabber"`
			Skype    string `json:"skype"`
			Steam    string `json:"steam"`
			Twitter  string `json:"twitter"`
			Website  string `json:"website"`
			Yahoo    string `json:"yahoo"`
		} `json:"contacts"`
		Info struct {
			Board    string      `json:"board"`
			Closed   bool        `json:"closed"`
			ID       int         `json:"id"`
			Image    string      `json:"image"`
			Name     string      `json:"name"`
			Owner    interface{} `json:"owner"`
			Type     string      `json:"type"`
			Username string      `json:"username"`
			Website  string      `json:"website"`
		} `json:"info"`
		Personal struct {
			Biography string      `json:"biography"`
			Birthday  string      `json:"birthday"`
			Gender    bool        `json:"gender"`
			Gravatar  string      `json:"gravatar"`
			ID        int         `json:"id"`
			Interests interface{} `json:"interests"`
			Name      string      `json:"name"`
			Nation    string      `json:"nation"`
			Online    bool        `json:"online"`
			Quotes    []string    `json:"quotes"`
			Surname   string      `json:"surname"`
			Timezone  string      `json:"timezone"`
			Username  string      `json:"username"`
		} `json:"personal"`
	} `json:"data"`
	HumanMessage string `json:"humanMessage"`
	Message      string `json:"message"`
	Status       int    `json:"status"`
	Success      bool   `json:"success"`
}

type MeFollowing struct {
	Data []struct {
		Contacts struct {
			Facebook string `json:"facebook"`
			Github   string `json:"github"`
			Jabber   string `json:"jabber"`
			Skype    string `json:"skype"`
			Steam    string `json:"steam"`
			Twitter  string `json:"twitter"`
			Website  string `json:"website"`
			Yahoo    string `json:"yahoo"`
		} `json:"contacts"`
		Info struct {
			Board    string      `json:"board"`
			Closed   bool        `json:"closed"`
			ID       int         `json:"id"`
			Image    string      `json:"image"`
			Name     string      `json:"name"`
			Owner    interface{} `json:"owner"`
			Type     string      `json:"type"`
			Username string      `json:"username"`
			Website  string      `json:"website"`
		} `json:"info"`
		Personal struct {
			Biography string      `json:"biography"`
			Birthday  string      `json:"birthday"`
			Gender    bool        `json:"gender"`
			Gravatar  string      `json:"gravatar"`
			ID        int         `json:"id"`
			Interests interface{} `json:"interests"`
			Name      string      `json:"name"`
			Nation    string      `json:"nation"`
			Online    bool        `json:"online"`
			Quotes    []string    `json:"quotes"`
			Surname   string      `json:"surname"`
			Timezone  string      `json:"timezone"`
			Username  string      `json:"username"`
		} `json:"personal"`
	} `json:"data"`
	HumanMessage string `json:"humanMessage"`
	Message      string `json:"message"`
	Status       int    `json:"status"`
	Success      bool   `json:"success"`
}

type UsersId struct {
	Data struct {
		Contacts struct {
			Facebook string `json:"facebook"`
			Github   string `json:"github"`
			Jabber   string `json:"jabber"`
			Skype    string `json:"skype"`
			Steam    string `json:"steam"`
			Twitter  string `json:"twitter"`
			Website  string `json:"website"`
			Yahoo    string `json:"yahoo"`
		} `json:"contacts"`
		Info struct {
			Board    string      `json:"board"`
			Closed   bool        `json:"closed"`
			ID       int         `json:"id"`
			Image    string      `json:"image"`
			Name     string      `json:"name"`
			Owner    interface{} `json:"owner"`
			Type     string      `json:"type"`
			Username string      `json:"username"`
			Website  string      `json:"website"`
		} `json:"info"`
		Personal struct {
			Biography string   `json:"biography"`
			Birthday  string   `json:"birthday"`
			Gender    bool     `json:"gender"`
			Gravatar  string   `json:"gravatar"`
			ID        int      `json:"id"`
			Interests []string `json:"interests"`
			Name      string   `json:"name"`
			Nation    string   `json:"nation"`
			Online    bool     `json:"online"`
			Quotes    []string `json:"quotes"`
			Surname   string   `json:"surname"`
			Timezone  string   `json:"timezone"`
			Username  string   `json:"username"`
		} `json:"personal"`
	} `json:"data"`
	HumanMessage string `json:"humanMessage"`
	Message      string `json:"message"`
	Status       int    `json:"status"`
	Success      bool   `json:"success"`
}

type MePmsOtherPmid struct {
	Data struct {
		CanDelete bool `json:"canDelete"`
		CanEdit   bool `json:"canEdit"`
		From      struct {
			Board    string      `json:"board"`
			Closed   bool        `json:"closed"`
			ID       int         `json:"id"`
			Image    string      `json:"image"`
			Name     string      `json:"name"`
			Owner    interface{} `json:"owner"`
			Type     string      `json:"type"`
			Username string      `json:"username"`
			Website  string      `json:"website"`
		} `json:"from"`
		Message string `json:"message"`
		Pmid    int    `json:"pmid"`
		Time    string `json:"time"`
		To      struct {
			Board    string      `json:"board"`
			Closed   bool        `json:"closed"`
			ID       int         `json:"id"`
			Image    string      `json:"image"`
			Name     string      `json:"name"`
			Owner    interface{} `json:"owner"`
			Type     string      `json:"type"`
			Username string      `json:"username"`
			Website  string      `json:"website"`
		} `json:"to"`
		ToRead bool `json:"toRead"`
	} `json:"data"`
	HumanMessage string `json:"humanMessage"`
	Message      string `json:"message"`
	Status       int    `json:"status"`
	Success      bool   `json:"success"`
}

type UsersIdPostsPidComments struct {
	Data []struct {
		CanDelete bool `json:"canDelete"`
		CanEdit   bool `json:"canEdit"`
		From      struct {
			Board    string      `json:"board"`
			Closed   bool        `json:"closed"`
			ID       int         `json:"id"`
			Image    string      `json:"image"`
			Name     string      `json:"name"`
			Owner    interface{} `json:"owner"`
			Type     string      `json:"type"`
			Username string      `json:"username"`
			Website  string      `json:"website"`
		} `json:"from"`
		Hcid    int    `json:"hcid"`
		Hpid    int    `json:"hpid"`
		Message string `json:"message"`
		Time    string `json:"time"`
		To      struct {
			Board    string      `json:"board"`
			Closed   bool        `json:"closed"`
			ID       int         `json:"id"`
			Image    string      `json:"image"`
			Name     string      `json:"name"`
			Owner    interface{} `json:"owner"`
			Type     string      `json:"type"`
			Username string      `json:"username"`
			Website  string      `json:"website"`
		} `json:"to"`
	} `json:"data"`
	HumanMessage string `json:"humanMessage"`
	Message      string `json:"message"`
	Status       int    `json:"status"`
	Success      bool   `json:"success"`
}

type MeFollowers struct {
	Data []struct {
		Contacts struct {
			Facebook string `json:"facebook"`
			Github   string `json:"github"`
			Jabber   string `json:"jabber"`
			Skype    string `json:"skype"`
			Steam    string `json:"steam"`
			Twitter  string `json:"twitter"`
			Website  string `json:"website"`
			Yahoo    string `json:"yahoo"`
		} `json:"contacts"`
		Info struct {
			Board    string      `json:"board"`
			Closed   bool        `json:"closed"`
			ID       int         `json:"id"`
			Image    string      `json:"image"`
			Name     string      `json:"name"`
			Owner    interface{} `json:"owner"`
			Type     string      `json:"type"`
			Username string      `json:"username"`
			Website  string      `json:"website"`
		} `json:"info"`
		Personal struct {
			Biography string      `json:"biography"`
			Birthday  string      `json:"birthday"`
			Gender    bool        `json:"gender"`
			Gravatar  string      `json:"gravatar"`
			ID        int         `json:"id"`
			Interests interface{} `json:"interests"`
			Name      string      `json:"name"`
			Nation    string      `json:"nation"`
			Online    bool        `json:"online"`
			Quotes    []string    `json:"quotes"`
			Surname   string      `json:"surname"`
			Timezone  string      `json:"timezone"`
			Username  string      `json:"username"`
		} `json:"personal"`
	} `json:"data"`
	HumanMessage string `json:"humanMessage"`
	Message      string `json:"message"`
	Status       int    `json:"status"`
	Success      bool   `json:"success"`
}

type MePostsPidCommentsCid struct {
	Data struct {
		CanDelete bool `json:"canDelete"`
		CanEdit   bool `json:"canEdit"`
		From      struct {
			Board    string      `json:"board"`
			Closed   bool        `json:"closed"`
			ID       int         `json:"id"`
			Image    string      `json:"image"`
			Name     string      `json:"name"`
			Owner    interface{} `json:"owner"`
			Type     string      `json:"type"`
			Username string      `json:"username"`
			Website  string      `json:"website"`
		} `json:"from"`
		Hcid    int    `json:"hcid"`
		Hpid    int    `json:"hpid"`
		Message string `json:"message"`
		Time    string `json:"time"`
		To      struct {
			Board    string      `json:"board"`
			Closed   bool        `json:"closed"`
			ID       int         `json:"id"`
			Image    string      `json:"image"`
			Name     string      `json:"name"`
			Owner    interface{} `json:"owner"`
			Type     string      `json:"type"`
			Username string      `json:"username"`
			Website  string      `json:"website"`
		} `json:"to"`
	} `json:"data"`
	HumanMessage string `json:"humanMessage"`
	Message      string `json:"message"`
	Status       int    `json:"status"`
	Success      bool   `json:"success"`
}

type MePostsPidComments struct {
	Data []struct {
		CanDelete bool `json:"canDelete"`
		CanEdit   bool `json:"canEdit"`
		From      struct {
			Board    string      `json:"board"`
			Closed   bool        `json:"closed"`
			ID       int         `json:"id"`
			Image    string      `json:"image"`
			Name     string      `json:"name"`
			Owner    interface{} `json:"owner"`
			Type     string      `json:"type"`
			Username string      `json:"username"`
			Website  string      `json:"website"`
		} `json:"from"`
		Hcid    int    `json:"hcid"`
		Hpid    int    `json:"hpid"`
		Message string `json:"message"`
		Time    string `json:"time"`
		To      struct {
			Board    string      `json:"board"`
			Closed   bool        `json:"closed"`
			ID       int         `json:"id"`
			Image    string      `json:"image"`
			Name     string      `json:"name"`
			Owner    interface{} `json:"owner"`
			Type     string      `json:"type"`
			Username string      `json:"username"`
			Website  string      `json:"website"`
		} `json:"to"`
	} `json:"data"`
	HumanMessage string `json:"humanMessage"`
	Message      string `json:"message"`
	Status       int    `json:"status"`
	Success      bool   `json:"success"`
}

type ProjectsIdPostsPidComments struct {
	Data []struct {
		CanDelete bool `json:"canDelete"`
		CanEdit   bool `json:"canEdit"`
		From      struct {
			Board    string      `json:"board"`
			Closed   bool        `json:"closed"`
			ID       int         `json:"id"`
			Image    string      `json:"image"`
			Name     string      `json:"name"`
			Owner    interface{} `json:"owner"`
			Type     string      `json:"type"`
			Username string      `json:"username"`
			Website  string      `json:"website"`
		} `json:"from"`
		Hcid    int    `json:"hcid"`
		Hpid    int    `json:"hpid"`
		Message string `json:"message"`
		Time    string `json:"time"`
		To      struct {
			Board  string `json:"board"`
			Closed bool   `json:"closed"`
			ID     int    `json:"id"`
			Image  string `json:"image"`
			Name   string `json:"name"`
			Owner  struct {
				Board    string      `json:"board"`
				Closed   bool        `json:"closed"`
				ID       int         `json:"id"`
				Image    string      `json:"image"`
				Name     string      `json:"name"`
				Owner    interface{} `json:"owner"`
				Type     string      `json:"type"`
				Username string      `json:"username"`
				Website  string      `json:"website"`
			} `json:"owner"`
			Type     string `json:"type"`
			Username string `json:"username"`
			Website  string `json:"website"`
		} `json:"to"`
	} `json:"data"`
	HumanMessage string `json:"humanMessage"`
	Message      string `json:"message"`
	Status       int    `json:"status"`
	Success      bool   `json:"success"`
}

type ProjectsIdPostsPidCommentsCid struct {
	Data struct {
		CanDelete bool `json:"canDelete"`
		CanEdit   bool `json:"canEdit"`
		From      struct {
			Board    string      `json:"board"`
			Closed   bool        `json:"closed"`
			ID       int         `json:"id"`
			Image    string      `json:"image"`
			Name     string      `json:"name"`
			Owner    interface{} `json:"owner"`
			Type     string      `json:"type"`
			Username string      `json:"username"`
			Website  string      `json:"website"`
		} `json:"from"`
		Hcid    int    `json:"hcid"`
		Hpid    int    `json:"hpid"`
		Message string `json:"message"`
		Time    string `json:"time"`
		To      struct {
			Board  string `json:"board"`
			Closed bool   `json:"closed"`
			ID     int    `json:"id"`
			Image  string `json:"image"`
			Name   string `json:"name"`
			Owner  struct {
				Board    string      `json:"board"`
				Closed   bool        `json:"closed"`
				ID       int         `json:"id"`
				Image    string      `json:"image"`
				Name     string      `json:"name"`
				Owner    interface{} `json:"owner"`
				Type     string      `json:"type"`
				Username string      `json:"username"`
				Website  string      `json:"website"`
			} `json:"owner"`
			Type     string `json:"type"`
			Username string `json:"username"`
			Website  string `json:"website"`
		} `json:"to"`
	} `json:"data"`
	HumanMessage string `json:"humanMessage"`
	Message      string `json:"message"`
	Status       int    `json:"status"`
	Success      bool   `json:"success"`
}

type UsersIdFollowers struct {
	Data []struct {
		Contacts struct {
			Facebook string `json:"facebook"`
			Github   string `json:"github"`
			Jabber   string `json:"jabber"`
			Skype    string `json:"skype"`
			Steam    string `json:"steam"`
			Twitter  string `json:"twitter"`
			Website  string `json:"website"`
			Yahoo    string `json:"yahoo"`
		} `json:"contacts"`
		Info struct {
			Board    string      `json:"board"`
			Closed   bool        `json:"closed"`
			ID       int         `json:"id"`
			Image    string      `json:"image"`
			Name     string      `json:"name"`
			Owner    interface{} `json:"owner"`
			Type     string      `json:"type"`
			Username string      `json:"username"`
			Website  string      `json:"website"`
		} `json:"info"`
		Personal struct {
			Biography string      `json:"biography"`
			Birthday  string      `json:"birthday"`
			Gender    bool        `json:"gender"`
			Gravatar  string      `json:"gravatar"`
			ID        int         `json:"id"`
			Interests interface{} `json:"interests"`
			Name      string      `json:"name"`
			Nation    string      `json:"nation"`
			Online    bool        `json:"online"`
			Quotes    []string    `json:"quotes"`
			Surname   string      `json:"surname"`
			Timezone  string      `json:"timezone"`
			Username  string      `json:"username"`
		} `json:"personal"`
	} `json:"data"`
	HumanMessage string `json:"humanMessage"`
	Message      string `json:"message"`
	Status       int    `json:"status"`
	Success      bool   `json:"success"`
}

type UsersIdFriends struct {
	Data []struct {
		Contacts struct {
			Facebook string `json:"facebook"`
			Github   string `json:"github"`
			Jabber   string `json:"jabber"`
			Skype    string `json:"skype"`
			Steam    string `json:"steam"`
			Twitter  string `json:"twitter"`
			Website  string `json:"website"`
			Yahoo    string `json:"yahoo"`
		} `json:"contacts"`
		Info struct {
			Board    string      `json:"board"`
			Closed   bool        `json:"closed"`
			ID       int         `json:"id"`
			Image    string      `json:"image"`
			Name     string      `json:"name"`
			Owner    interface{} `json:"owner"`
			Type     string      `json:"type"`
			Username string      `json:"username"`
			Website  string      `json:"website"`
		} `json:"info"`
		Personal struct {
			Biography string      `json:"biography"`
			Birthday  string      `json:"birthday"`
			Gender    bool        `json:"gender"`
			Gravatar  string      `json:"gravatar"`
			ID        int         `json:"id"`
			Interests interface{} `json:"interests"`
			Name      string      `json:"name"`
			Nation    string      `json:"nation"`
			Online    bool        `json:"online"`
			Quotes    []string    `json:"quotes"`
			Surname   string      `json:"surname"`
			Timezone  string      `json:"timezone"`
			Username  string      `json:"username"`
		} `json:"personal"`
	} `json:"data"`
	HumanMessage string `json:"humanMessage"`
	Message      string `json:"message"`
	Status       int    `json:"status"`
	Success      bool   `json:"success"`
}

type UsersIdPostsPid struct {
	Data struct {
		Bookmarkers int  `json:"bookmarkers"`
		CanBookmark bool `json:"canBookmark"`
		CanComment  bool `json:"canComment"`
		CanDelete   bool `json:"canDelete"`
		CanEdit     bool `json:"canEdit"`
		CanLurk     bool `json:"canLurk"`
		Closed      bool `json:"closed"`
		Comments    int  `json:"comments"`
		From        struct {
			Board    string      `json:"board"`
			Closed   bool        `json:"closed"`
			ID       int         `json:"id"`
			Image    string      `json:"image"`
			Name     string      `json:"name"`
			Owner    interface{} `json:"owner"`
			Type     string      `json:"type"`
			Username string      `json:"username"`
			Website  string      `json:"website"`
		} `json:"from"`
		Hpid      int    `json:"hpid"`
		Lang      string `json:"lang"`
		Lurkers   int    `json:"lurkers"`
		Message   string `json:"message"`
		News      bool   `json:"news"`
		Pid       int    `json:"pid"`
		Rate      int    `json:"rate"`
		Revisions int    `json:"revisions"`
		Time      string `json:"time"`
		Timestamp int    `json:"timestamp"`
		To        struct {
			Board    string      `json:"board"`
			Closed   bool        `json:"closed"`
			ID       int         `json:"id"`
			Image    string      `json:"image"`
			Name     string      `json:"name"`
			Owner    interface{} `json:"owner"`
			Type     string      `json:"type"`
			Username string      `json:"username"`
			Website  string      `json:"website"`
		} `json:"to"`
		Type string `json:"type"`
		URL  string `json:"url"`
	} `json:"data"`
	HumanMessage string `json:"humanMessage"`
	Message      string `json:"message"`
	Status       int    `json:"status"`
	Success      bool   `json:"success"`
}

type ProjectsIdPostsPid struct {
	Data struct {
		Bookmarkers int  `json:"bookmarkers"`
		CanBookmark bool `json:"canBookmark"`
		CanComment  bool `json:"canComment"`
		CanDelete   bool `json:"canDelete"`
		CanEdit     bool `json:"canEdit"`
		CanLurk     bool `json:"canLurk"`
		Closed      bool `json:"closed"`
		Comments    int  `json:"comments"`
		From        struct {
			Board    string      `json:"board"`
			Closed   bool        `json:"closed"`
			ID       int         `json:"id"`
			Image    string      `json:"image"`
			Name     string      `json:"name"`
			Owner    interface{} `json:"owner"`
			Type     string      `json:"type"`
			Username string      `json:"username"`
			Website  string      `json:"website"`
		} `json:"from"`
		Hpid      int    `json:"hpid"`
		Lang      string `json:"lang"`
		Lurkers   int    `json:"lurkers"`
		Message   string `json:"message"`
		News      bool   `json:"news"`
		Pid       int    `json:"pid"`
		Rate      int    `json:"rate"`
		Revisions int    `json:"revisions"`
		Time      string `json:"time"`
		Timestamp int    `json:"timestamp"`
		To        struct {
			Board  string `json:"board"`
			Closed bool   `json:"closed"`
			ID     int    `json:"id"`
			Image  string `json:"image"`
			Name   string `json:"name"`
			Owner  struct {
				Board    string      `json:"board"`
				Closed   bool        `json:"closed"`
				ID       int         `json:"id"`
				Image    string      `json:"image"`
				Name     string      `json:"name"`
				Owner    interface{} `json:"owner"`
				Type     string      `json:"type"`
				Username string      `json:"username"`
				Website  string      `json:"website"`
			} `json:"owner"`
			Type     string `json:"type"`
			Username string `json:"username"`
			Website  string `json:"website"`
		} `json:"to"`
		Type string `json:"type"`
		URL  string `json:"url"`
	} `json:"data"`
	HumanMessage string `json:"humanMessage"`
	Message      string `json:"message"`
	Status       int    `json:"status"`
	Success      bool   `json:"success"`
}

type MePms struct {
	Data []struct {
		From struct {
			Board    string      `json:"board"`
			Closed   bool        `json:"closed"`
			ID       int         `json:"id"`
			Image    string      `json:"image"`
			Name     string      `json:"name"`
			Owner    interface{} `json:"owner"`
			Type     string      `json:"type"`
			Username string      `json:"username"`
			Website  string      `json:"website"`
		} `json:"from"`
		LastMessage string `json:"lastMessage"`
		Time        string `json:"time"`
		To          struct {
			Board    string      `json:"board"`
			Closed   bool        `json:"closed"`
			ID       int         `json:"id"`
			Image    string      `json:"image"`
			Name     string      `json:"name"`
			Owner    interface{} `json:"owner"`
			Type     string      `json:"type"`
			Username string      `json:"username"`
			Website  string      `json:"website"`
		} `json:"to"`
		ToRead bool `json:"toRead"`
	} `json:"data"`
	HumanMessage string `json:"humanMessage"`
	Message      string `json:"message"`
	Status       int    `json:"status"`
	Success      bool   `json:"success"`
}

type MePmsOther struct {
	Data []struct {
		CanDelete bool `json:"canDelete"`
		CanEdit   bool `json:"canEdit"`
		From      struct {
			Board    string      `json:"board"`
			Closed   bool        `json:"closed"`
			ID       int         `json:"id"`
			Image    string      `json:"image"`
			Name     string      `json:"name"`
			Owner    interface{} `json:"owner"`
			Type     string      `json:"type"`
			Username string      `json:"username"`
			Website  string      `json:"website"`
		} `json:"from"`
		Message string `json:"message"`
		Pmid    int    `json:"pmid"`
		Time    string `json:"time"`
		To      struct {
			Board    string      `json:"board"`
			Closed   bool        `json:"closed"`
			ID       int         `json:"id"`
			Image    string      `json:"image"`
			Name     string      `json:"name"`
			Owner    interface{} `json:"owner"`
			Type     string      `json:"type"`
			Username string      `json:"username"`
			Website  string      `json:"website"`
		} `json:"to"`
		ToRead bool `json:"toRead"`
	} `json:"data"`
	HumanMessage string `json:"humanMessage"`
	Message      string `json:"message"`
	Status       int    `json:"status"`
	Success      bool   `json:"success"`
}

type ProjectsIdPosts struct {
	Data []struct {
		Bookmarkers int  `json:"bookmarkers"`
		CanBookmark bool `json:"canBookmark"`
		CanComment  bool `json:"canComment"`
		CanDelete   bool `json:"canDelete"`
		CanEdit     bool `json:"canEdit"`
		CanLurk     bool `json:"canLurk"`
		Closed      bool `json:"closed"`
		Comments    int  `json:"comments"`
		From        struct {
			Board    string      `json:"board"`
			Closed   bool        `json:"closed"`
			ID       int         `json:"id"`
			Image    string      `json:"image"`
			Name     string      `json:"name"`
			Owner    interface{} `json:"owner"`
			Type     string      `json:"type"`
			Username string      `json:"username"`
			Website  string      `json:"website"`
		} `json:"from"`
		Hpid      int    `json:"hpid"`
		Lang      string `json:"lang"`
		Lurkers   int    `json:"lurkers"`
		Message   string `json:"message"`
		News      bool   `json:"news"`
		Pid       int    `json:"pid"`
		Rate      int    `json:"rate"`
		Revisions int    `json:"revisions"`
		Time      string `json:"time"`
		Timestamp int    `json:"timestamp"`
		To        struct {
			Board  string `json:"board"`
			Closed bool   `json:"closed"`
			ID     int    `json:"id"`
			Image  string `json:"image"`
			Name   string `json:"name"`
			Owner  struct {
				Board    string      `json:"board"`
				Closed   bool        `json:"closed"`
				ID       int         `json:"id"`
				Image    string      `json:"image"`
				Name     string      `json:"name"`
				Owner    interface{} `json:"owner"`
				Type     string      `json:"type"`
				Username string      `json:"username"`
				Website  string      `json:"website"`
			} `json:"owner"`
			Type     string `json:"type"`
			Username string `json:"username"`
			Website  string `json:"website"`
		} `json:"to"`
		Type string `json:"type"`
		URL  string `json:"url"`
	} `json:"data"`
	HumanMessage string `json:"humanMessage"`
	Message      string `json:"message"`
	Status       int    `json:"status"`
	Success      bool   `json:"success"`
}

type MePosts struct {
	Data []struct {
		Bookmarkers int  `json:"bookmarkers"`
		CanBookmark bool `json:"canBookmark"`
		CanComment  bool `json:"canComment"`
		CanDelete   bool `json:"canDelete"`
		CanEdit     bool `json:"canEdit"`
		CanLurk     bool `json:"canLurk"`
		Closed      bool `json:"closed"`
		Comments    int  `json:"comments"`
		From        struct {
			Board    string      `json:"board"`
			Closed   bool        `json:"closed"`
			ID       int         `json:"id"`
			Image    string      `json:"image"`
			Name     string      `json:"name"`
			Owner    interface{} `json:"owner"`
			Type     string      `json:"type"`
			Username string      `json:"username"`
			Website  string      `json:"website"`
		} `json:"from"`
		Hpid      int    `json:"hpid"`
		Lang      string `json:"lang"`
		Lurkers   int    `json:"lurkers"`
		Message   string `json:"message"`
		News      bool   `json:"news"`
		Pid       int    `json:"pid"`
		Rate      int    `json:"rate"`
		Revisions int    `json:"revisions"`
		Time      string `json:"time"`
		Timestamp int    `json:"timestamp"`
		To        struct {
			Board    string      `json:"board"`
			Closed   bool        `json:"closed"`
			ID       int         `json:"id"`
			Image    string      `json:"image"`
			Name     string      `json:"name"`
			Owner    interface{} `json:"owner"`
			Type     string      `json:"type"`
			Username string      `json:"username"`
			Website  string      `json:"website"`
		} `json:"to"`
		Type string `json:"type"`
		URL  string `json:"url"`
	} `json:"data"`
	HumanMessage string `json:"humanMessage"`
	Message      string `json:"message"`
	Status       int    `json:"status"`
	Success      bool   `json:"success"`
}

type UsersIdPosts struct {
	Data []struct {
		Bookmarkers int  `json:"bookmarkers"`
		CanBookmark bool `json:"canBookmark"`
		CanComment  bool `json:"canComment"`
		CanDelete   bool `json:"canDelete"`
		CanEdit     bool `json:"canEdit"`
		CanLurk     bool `json:"canLurk"`
		Closed      bool `json:"closed"`
		Comments    int  `json:"comments"`
		From        struct {
			Board    string      `json:"board"`
			Closed   bool        `json:"closed"`
			ID       int         `json:"id"`
			Image    string      `json:"image"`
			Name     string      `json:"name"`
			Owner    interface{} `json:"owner"`
			Type     string      `json:"type"`
			Username string      `json:"username"`
			Website  string      `json:"website"`
		} `json:"from"`
		Hpid      int    `json:"hpid"`
		Lang      string `json:"lang"`
		Lurkers   int    `json:"lurkers"`
		Message   string `json:"message"`
		News      bool   `json:"news"`
		Pid       int    `json:"pid"`
		Rate      int    `json:"rate"`
		Revisions int    `json:"revisions"`
		Time      string `json:"time"`
		Timestamp int    `json:"timestamp"`
		To        struct {
			Board    string      `json:"board"`
			Closed   bool        `json:"closed"`
			ID       int         `json:"id"`
			Image    string      `json:"image"`
			Name     string      `json:"name"`
			Owner    interface{} `json:"owner"`
			Type     string      `json:"type"`
			Username string      `json:"username"`
			Website  string      `json:"website"`
		} `json:"to"`
		Type string `json:"type"`
		URL  string `json:"url"`
	} `json:"data"`
	HumanMessage string `json:"humanMessage"`
	Message      string `json:"message"`
	Status       int    `json:"status"`
	Success      bool   `json:"success"`
}

type MeHome struct {
	Data []struct {
		Bookmarkers int  `json:"bookmarkers"`
		CanBookmark bool `json:"canBookmark"`
		CanComment  bool `json:"canComment"`
		CanDelete   bool `json:"canDelete"`
		CanEdit     bool `json:"canEdit"`
		CanLurk     bool `json:"canLurk"`
		Closed      bool `json:"closed"`
		Comments    int  `json:"comments"`
		From        struct {
			Board    string      `json:"board"`
			Closed   bool        `json:"closed"`
			ID       int         `json:"id"`
			Image    string      `json:"image"`
			Name     string      `json:"name"`
			Owner    interface{} `json:"owner"`
			Type     string      `json:"type"`
			Username string      `json:"username"`
			Website  string      `json:"website"`
		} `json:"from"`
		Hpid      int    `json:"hpid"`
		Lang      string `json:"lang"`
		Lurkers   int    `json:"lurkers"`
		Message   string `json:"message"`
		News      bool   `json:"news"`
		Pid       int    `json:"pid"`
		Rate      int    `json:"rate"`
		Revisions int    `json:"revisions"`
		Time      string `json:"time"`
		Timestamp int    `json:"timestamp"`
		To        struct {
			Board    string      `json:"board"`
			Closed   bool        `json:"closed"`
			ID       int         `json:"id"`
			Image    string      `json:"image"`
			Name     string      `json:"name"`
			Owner    interface{} `json:"owner"`
			Type     string      `json:"type"`
			Username string      `json:"username"`
			Website  string      `json:"website"`
		} `json:"to"`
		Type string `json:"type"`
		URL  string `json:"url"`
	} `json:"data"`
	HumanMessage string `json:"humanMessage"`
	Message      string `json:"message"`
	Status       int    `json:"status"`
	Success      bool   `json:"success"`
}
