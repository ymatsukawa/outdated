from pymongo import MongoClient
from project.settings import MONGODB

class MongoDB:
  _instance = None
  _db = None

  def __new__(cls):
    if cls._instance is None:
      cls._instance = super().__new__(cls)
      cls._client = MongoClient(MONGODB['host'], MONGODB['port'])
      cls._collection = cls._client[MONGODB['db']][MONGODB['collection']]
    return cls._instance

  def __init__(self):
    None

  def save(self, record, check_cache=True):
    if(not self.is_saved(record)):
      self._collection.insert(record)

  def is_saved(self, record):
    if(self.load(record) is None):
      return False
    else:
      return True

  def load(self, record):
    return self._collection.find_one(record)
