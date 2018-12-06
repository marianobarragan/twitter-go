package service

import (
	"github.com/marianobarragan/Twitter/src/domain"
)
import "time"
import "errors"

var tweets []*domain.Tweet
var idCounter int

func InitializeService(){
	tweets = make([]* domain.Tweet, 0) // Crea un slice vacío
}

func PublishTweet(tweet_p *domain.Tweet) (id int, error error){

	if len(tweet_p.User) == 0 {
		error = errors.New("user is required")
		return
	}
	if len(tweet_p.Text) == 0 {
		error = errors.New("text is required")
		return
	}
	if len(tweet_p.Text) > 140 {
		error =  errors.New("text is too long")
		return
	}

	t := time.Now()
	tweet_p.Date = &t
	id = idCounter
	idCounter ++
	tweet_p.Id = id
	tweets = append(tweets, tweet_p)
	return id, nil
}

func GetTweet(position int) *domain.Tweet{
	if position < len(tweets) {
		return tweets[position]
	}
	return nil
}

func GetTweets() []*domain.Tweet{
	return tweets
}

func GetTweetById(id int) *domain.Tweet{

	for _, tweet := range tweets {
		if tweet.Id == id {
			return tweet
		}
	}
	return nil
}

func CountTweetsByUser(user string) (count int)  {

	for _, tweet := range tweets {
		if tweet.User == user {
			count ++
		}
	}
	return
}



