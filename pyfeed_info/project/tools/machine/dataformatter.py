from project.const import FEED_KEYNAME, TITLE, LINK
class DataFormatter:
  def __init__(self, feed_set):
    self._feed_set = feed_set
    self._formatted_data = None

  def format_to_post(self):
    '''
    Format feed set to "list of title and link dict"
    :return: no. after run, property "formatted_data" returns "list of title and link set"
    ex.)
    formatted_date
    [
      {
        "title" : "example",
        "link" : "http://example.com"
      },
      {
        "title" : "foo",
        "link" : "http://example.com/foo"
      },
      .
      .
      .
    ]
    '''
    result = []
    for feed in self._feed_set:
      value = {
        TITLE: feed[TITLE],
        LINK:  feed[LINK]
      }
      result.append(value)
    self._formatted_data = result

  @property
  def formatted_data(self):
      return self._formatted_data