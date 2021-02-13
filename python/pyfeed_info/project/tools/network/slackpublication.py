import datetime
import json
import random
import requests
from project.settings import PUBLICATION, MAX_NUM_TO_PUB

class SlackPublication:
  def __init__(self):
    self._posted_data = None

  def post(self, raw_post_data):
    text = datetime.datetime.today().strftime("%Y/%m/%d") + "\n" + \
           self._format_data_to_post(raw_post_data)
    requests.post(
      url = PUBLICATION['url'],
      data = json.dumps({
        'text'       : text,
        'username'   : PUBLICATION['username'],
        'icon_emoji' : PUBLICATION['icon_emoji'],
        'link_names' : PUBLICATION['link_names'],
    }))

  def _format_data_to_post(self, raw_post_data):
    raw_data = list(raw_post_data)
    random.shuffle(raw_data)
    data = raw_data[0:MAX_NUM_TO_PUB]
    self._posted_data = data
    text = ''
    for d in data:
      text += ('*<' + d['link'] + '|' + d['title'] + '>*' + "\n") # "*<link|url>*" ... embedded bold link
      text += "\n"
    return text

  @property
  def posted_data(self):
    return self._posted_data
