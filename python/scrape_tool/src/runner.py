import os
from scrapy.crawler import CrawlerProcess
from scrapy.utils.project import get_project_settings
from scrape.common.const import Env
from scrape.scraper import Scraper
from iotool.pathaccess import PathAccess
from iotool.fileread import FileRead
from iotool.pathaccess import PathAccess
from dotenv import load_dotenv

class Runner:
    _env = None
    _scrapy_settings = None
    _user_settings = None

    def __init__(self, env=Env.DEV):
        load_dotenv()
        env = os.environ.get('SCRAPING_ENV')
        if env == "PROD":
            self._env = Env.PROD
        elif env == "DEV":
            self._env = Env.DEV
        elif env == "TEST":
            self._env = Env.TEST
        else:
            raise EnvironmentError("Value of SCRAPING_ENV is invalid. set PROD, DEV or TEST at .env")
    
    def initialize(self):
        setting_file = None
        if self._env == Env.PROD:
            setting_file = "settings.prod.json"
        elif self._env == Env.DEV:
            setting_file = "settings.dev.json"
        elif self._env == Env.TEST:
            setting_file = "settings.test.json"
        
        file_read = FileRead()
        self._user_settings = file_read.as_json(PathAccess.realpath(__file__, setting_file))

        self._scrapy_settings = get_project_settings()
        self._scrapy_settings.set('FEEDS', self._user_settings['feed'])
        self._scrapy_settings.set('DOWNLOAD_DELAY', self._user_settings['download']['delay_sec'])

    def run(self):
        process = CrawlerProcess(self._scrapy_settings)
        for target in self._user_settings["targets"]:
            scrape_args = {
                'name': target['name'],
                'url': target['url'],
                'search': target['search']
            }
            process.crawl(Scraper, scrape_args=scrape_args)

        process.start()