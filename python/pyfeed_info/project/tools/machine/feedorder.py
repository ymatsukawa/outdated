from dateutil.parser import parse
import datetime
from project.tools.machine.mongodb import MongoDB
from project.tools.machine.regex import Regex
from project.settings import DATE_TYPES

class FeedOrder:
  _date_types = DATE_TYPES

  def __init__(self, origin_feed_set):
    self._feed_set = origin_feed_set
    self._mongodb  = MongoDB()
    self._regex    = Regex()

  def filter_by_keyword(self, key_item, keyword_set):
    '''
    Filters feeds by keyword.
    They are ordered by "keyword or search"
    After that, feed_set property is changed.
    :param key_item: where to search keyword by string. ex.) "title",
    :param keyword_set: set of keyword to filter by set. ex.) set(['cpu', 'HDD', 'RAM']), set(['cloud', 'business'])
    :return: self; instance
    '''
    # WANT - Data mining
    or_pattern = self._regex.or_pattern(keyword_set)
    self._feed_set = set(filter(lambda feed: or_pattern.match(feed[key_item]), self._feed_set))
    return self

  def filter_by_days(self, date_terms_from_now):
    '''
    Filters feeds by specific date terms from now.
    After filtered, feed_set property is changed.
    :param date_terms_from_now: how many limits by days. ex) 3 ... within 3 days, over 4 days are excluded.
    :return: self; instance
    '''
    now = datetime.datetime.now(datetime.timezone.utc)
    self._feed_set = set([feed for feed in self._feed_set if self._is_day_in_limit_days(feed, date_terms_from_now, now)])
    return self

  def exclude_cached(self, key_name):
    '''
    Filters feeds if data is cached.
    After filtered, feed_set property is changed.
    :param key_name: check where item(key) of feed is cached or not.
    :return: self; instance
    '''
    self._feed_set = set(filter(lambda feed: (not self._mongodb.is_saved({key_name : feed[key_name]})), self._feed_set))
    return self

  def _is_day_in_limit_days(self, feed, limit_days, datetime_now_utc):
    return True in set(map(lambda day_type:
      (day_type in feed) and
      ((datetime_now_utc - parse(feed[day_type])).days <= limit_days),
      self._date_types))

  @property
  def feed_set(self):
    return self._feed_set
