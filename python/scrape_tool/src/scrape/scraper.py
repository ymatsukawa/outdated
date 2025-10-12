import scrapy

class Scraper(scrapy.Spider):
    name = None
    start_urls = ['http://example.com/setting_is_overwrited/by_scrape_args']
    desc = None
    search = None
    page_searched = None

    """
    args expcetation example:
        scrape_args = {
            name: 'example',
            url: 'http://example.com',
            search: [
                {
                    'desc': 'title',
                    'path': '/html/head/title'
                },
                {
                    'desc: 'each_p',
                    'path': '/html/body/div[1]/p',
                    'next_xpath_href': '/html/body/div[2]/a/@href',
                    'max_next_page': 2
                }
            ]
        }
    """
    def __init__(self, scrape_args):
        self.name = scrape_args['name']
        self.start_urls = [scrape_args['url']]
        self.search = scrape_args['search']
        self.page_crawled = 0

    def parse(self, response):
        for s in self.search:
            for res in response.xpath(s['xpath']):
                yield { s['desc'] : res.extract() }

                if 'next_xpath_href' in s:
                    self.page_crawled += 1
                    if self.page_crawled < s['max_crawl_page']:
                        next_page = response.xpath(s['next_xpath_href']).get()
                        if next_page is not None:
                            yield response.follow(next_page, callback=self.parse)
                        else:
                            self.page_searched = 0