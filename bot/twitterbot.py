import tweepy
import time
from datetime import datetime
from tweepy import TweepError

api_key = 'pnDHE9x8v0UIUUann6IMQ5EHP'
api_secret_key = 'n2o1gLnW9k13GS8bXjwPWtBxw68R4SBP8LGD3qk8EApbeyW3Ry'
api_access_token = '914299970974490624-dOPMMTDeq9IzbeSydEtpUNWnba438ar'
api__secret_access_token = 'g52j8Mdr5oRq0E8t1KvSyIuHLBkl5R65Q3qklB1cMHHtu'


def getAPI():
    auth = tweepy.OAuthHandler(api_key, api_secret_key)
    auth.set_access_token(api_access_token, api__secret_access_token)
    api = tweepy.API(auth)
    try:
        api.verify_credentials()
    except:
        print('getAPI error')
    return api


def followerList(api):
    for follower in api.followers_ids():
        print("followerList:", follower)


def friendList(api):
    for friend in api.friends_ids():
        print("friendList:", friend)


def retweetofmeList(api):
    for tweet in api.retweets_of_me():
        print(tweet.text)


def sendMessage(api, sMsg):
    for follower in api.followers_ids():
        senddirectmessage(api, follower, sMsg)


def senddirectmessage(api, recepient_id, message):
    try:
        api.send_direct_message(recipient_id=recepient_id, text=message)
    except tweepy.TweepError as error:
        print(error.reason)


def hometimeline(api):
    htimeline = api.home_timeline()
    for tweet in htimeline:
        sMsg = '{0}/{1} - {2}'.format(tweet.id, tweet.user.name, tweet.text)
        print(sMsg)


def searchKeyword(api, skey):
    result = []
    searchresult = api.search(q=skey, count=1000)
    for val in searchresult:
        sMsg = '{0} - {1}'.format(val.id_str, val.text)
        # time.sleep(1)
        result.append(sMsg)
    # print(len(result))
    return result


def searchCursor(api, skey):
    result = []
    for tweet in tweepy.Cursor(api.search, q=(skey), since='2014-09-16', until='2019-12-31').items(1000):
        sMsg = '{0}/{1} - {2}'.format(tweet.created_at, tweet.user.location.encode('utf8'), tweet.text)
        # time.sleep(1)
        result.append(sMsg)
    # print(len(result))
    return result


def searchretweet(api, skey):
    try:
        for tweet in tweepy.Cursor(api.search, q=skey).items(10):
            print('\n@' + tweet.user.screen_name + ' - ' + tweet.text)
            tweet.retweet()
            time.sleep(10)
    except tweepy.TweepError as error:
        print(error.reason)


def checkfavorite(api, skey):
    try:
        for tweet in tweepy.Cursor(api.search, q=skey).items():
            print('\n@' + tweet.user.screen_name + ' - ' + tweet.text)
            # 좋아요 누르기
            tweet.favorite()
            # follow 하기
            if not tweet.user.following:
                tweet.user.follow()

    except tweepy.TweepError as error:
        print(error.reason)


def photopath(api, path):
    try:
        now = datetime.now()
        status = 'my photo:' + now.strftime("%Y-%m-%d %H:%M:%S")
        api.update_with_media(path, status=status)
    except tweepy.TweepError as error:
        print(error.reason)


if __name__ == "__main__":
    api = getAPI()
    user = api.me()

    print('user.friends_count:', user.friends_count)
    print('user.description:', user.description)
    print('user.location:', user.location)
    print('user.id:', user.id)
    print('user.id_str:', user.id_str)
    print('user.screen_name:', user.screen_name)
    print('user.name:', user.name)
    print('user.followers_count:', user.followers_count)
    print('user.statuses_count:', str(user.statuses_count))
    print('user.time_zone:', user.time_zone)

    # api.update_status(status="공부중")
    # senddirectmessage(api, user.id, '공부중')
    # api.update_status("@" + user.screen_name, in_reply_to_status_id=user.id)
    # hometimeline(api)
    # retweetofmeList(api)
    # followerList(api)
    # friendList(api)

    """
    results = searchCursor(api, '최신 영화')
    for result in results:
        print(result)
    """

    """
    results = searchKeyword(api, '최신 영화')
    for result in results:
        print(result)
    """

    # searchretweet(api, '최신 영화')
    # checkfavorite(api, '#python')
    # photopath(api, 'e:/a.jpg')
