import re

class Regex:
  _instance = None

  def __new__(cls):
    if cls._instance is None:
      cls._instance = super().__new__(cls)
    return cls._instance

  def __init__(self):
    None

  def or_pattern(self, words):
    '''
    can get or_pattern regex from words
    :param words: list of words to or search. ex.) ["hello", "world"]
    :return: when words is ["hello", "words"], returns re.compile("(.*)(hello|world)")
    '''
    pattern = '(.*)('
    for word in words:
      pattern += (word + '|')
    pattern = pattern[:-1]  # trim final char '|'
    pattern += ')'
    return re.compile(pattern)
