from logging import basicConfig, getLogger, INFO
from project.const import TITLE, LINK
from project.settings import RSS_URLS, KEYWORD_SET, LIMIT_DATE_TERM
from project.tools.machine.dataformatter import DataFormatter
from project.tools.machine.mongodb import MongoDB
from project.tools.machine.feedorder import FeedOrder
from project.tools.network.rssclient import RssClient
from project.tools.network.slackpublication import SlackPublication

basicConfig(level=INFO)
logger = getLogger(__name__)

class EntryPoint:
  def __init__(self):
    self._mongodb = MongoDB()

  def run_main(self):
    latest_feeds = RssClient(RSS_URLS).check_feeds().check_feeds().feeds
    if(self.no_feeds(latest_feeds)):
      logger.info('latest feeds not found, or network error happened')
      return

    ordered_feeds = FeedOrder(latest_feeds).filter_by_keyword(TITLE, KEYWORD_SET).filter_by_days(LIMIT_DATE_TERM).exclude_cached('link').feed_set
    if(self.no_feeds(ordered_feeds)):
      logger.info('after feed order, no feeds left')
      return

    data_formatter = DataFormatter(ordered_feeds)
    data_formatter.format_to_post()

    slack_publication = SlackPublication()
    slack_publication.post(data_formatter.formatted_data)

    posted_data = slack_publication.posted_data
    for data in posted_data:
      self._mongodb.save({LINK : data[LINK]})

    logger.info('process finished')

  def no_feeds(self, feeds):
    if(len(feeds) == 0):
      return True
    else:
      return False


#-------------------------

entry_point = EntryPoint()
entry_point.run_main()
