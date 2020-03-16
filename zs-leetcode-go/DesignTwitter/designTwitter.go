package DesiginTwitter

// only support
type TweetNode struct {
	ID    int // tweet id
	Order int // global id or use time.Time to control
}

type Twitter struct {
	Followers map[int][]int // 1:2 1 follow 2
	Tweets    map[int][]*TweetNode
	GlobalID  int
	topTen    int
}


/** Initialize your data structure here. */
func Constructor() Twitter {
	return Twitter{
		Followers: map[int][]int{},
		Tweets:    map[int][]*TweetNode{},
		topTen:    10,
	}
}

/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int) {
	if _, ok := this.Tweets[userId]; !ok {
		this.Tweets[userId] = []*TweetNode{}
	}
	this.Tweets[userId] = append([]*TweetNode{{tweetId, this.GlobalID}}, this.Tweets[userId]...)
	this.GlobalID++
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	// first get an array include him and his followers
	followers := this.GetAllPost(userId)
	if followers == nil { return nil }
	return this.GetTopTenTweets(followers)
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int) {
	if followeeId == followerId { return } // not support self follow self
	users, ok := this.Followers[followerId]
	if !ok {
		this.Followers[followerId] = []int{}
	}
	if checkUserInFollowers(followeeId, users) { return }
	this.Followers[followerId] = append(this.Followers[followerId], followeeId)
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if users, ok := this.Followers[followerId]; ok {
		for index := range users {
			if users[index] == followeeId {
				this.Followers[followerId] = append(users[:index], users[index+1:]...)
				break
			}
		}
	}
}

func (self *Twitter) GetAllPost(userID int) [][]*TweetNode {
	allPosts := [][]*TweetNode{}
	if posts, ok := self.Tweets[userID]; ok {
		allPosts = append(allPosts, posts)
	}
	if followers, ok := self.Followers[userID]; ok {
		for _, item := range followers {
			if posts, ok := self.Tweets[item]; ok {
				allPosts = append(allPosts, posts)
			}
		}
	}
	return allPosts
}

type tmpStruct struct {
	Index int
	Length int
}
func (self *Twitter) GetTopTenTweets(posts [][]*TweetNode) []int {
	res := []int{}
	indexs := []tmpStruct{}
	for _, item := range posts {
		indexs = append(indexs, tmpStruct{0, len(item)})
	}
	if len(indexs) == 0 { return res}
	for len(res) < 10 {
		maxId, postID, needUpdateIndex := -1, -1, -1
		for index, item := range indexs {
			if item.Index < item.Length && posts[index][item.Index].Order > maxId {
				maxId = posts[index][item.Index].Order
				postID = posts[index][item.Index].ID
				needUpdateIndex = index
			}
		}
		if maxId == -1 { break}
		res = append(res, postID)
		indexs[needUpdateIndex].Index++
	}
	return res
}

func checkUserInFollowers(id int, ids []int) bool {
	for _, _id := range ids {
		if _id == id {
			return true
		}
	}
	return false
}

