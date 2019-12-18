import tweepy
import time
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
    except TweepError:
        print('senddirectmessage error')


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


if __name__ == "__main__":
    api = getAPI()
    user = api.me()

    """
    print(user.friends_count)
    print(user.description)
    print(user.location)
    print(user.id)
    print(user.id_str)
    print(user.screen_name)
    print(user.name)
    print(api.followers_ids())
    """

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

    results = searchKeyword(api, '최신 영화')
    for result in results:
        print(result)
