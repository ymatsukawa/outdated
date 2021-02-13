# settings about RSS ====================

# URL to get rss feeds
RSS_URLS = [
 # "http://example.com/sample_rss_10.xml",
 # "http://example.com/sample_rss_20.xml",
 # "http://example.com/sample.rdf",
 # ...
]

# specify date type each rss feed item has
# ex.)
# rssX > dc:date 20XX-XX-XX
# and
# rssY > dc:published 20XX-0X-0X
# then DATE_TYPES = ['date', 'published']
DATE_TYPES = ['date', 'pubDate', 'lastBuildDate']

# Limit term from now to get rss
# ex.) when 3, rss will be gotten from 3 days before to now.
LIMIT_DATE_TERM = 3

# Keyword to aggregate from title in RSS-feed
# follow writing as ...> set(['key', 'word', ...])
# KEYWORD_SET = set(['python', 'Python', 'update', 'security', 'Security'])
KEYWORD_SET = set(['Python', 'python', 'update'])


# Settings about publication ====================
# Publication settings for slack
PUBLICATION = {
  'url'        : '', # webhook url
  'username'   :  'Py-News(beta)',
  'icon_emoji' : ':snake:',
  'link_names' : 0, # enable to mention(1) or not(0) when post data has mention
}

# max post numbers to publication
MAX_NUM_TO_PUB = 5


# Settings about MongoDB ====================
MONGODB = {
  'host'       : 'localhost',
  'port'       : 27017,
  'db'         : 'feed_info',
  'collection' : 'feed'
}
