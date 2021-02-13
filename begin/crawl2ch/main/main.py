#!/usr/bin/env python
#-*- coding: utf-8 -*-

import time
import sys
import _os
import src.crawl.State as State
import src.common.infoCrawl as info
import src.ecom.State as StateML
from datetime import datetime
TIME = info.d['looptime']
count = 0
def time_stamp(): tm = datetime.today(); print "%s:%s:%s" % (tm.hour, tm.minute, tm.second)


if TIME < 600:
	print 'Short term(lower is 600s). Set again'
	sys.exit()

time_stamp()
sCrawl = State.State()
""" sML = StateML.State() """ # ECOM --- To use function of email, remove bracket of comment 

print 'HOST: '    , info.d['host']
print 'URL: '     , info.d['reqUrl']
print 'Referer: ' , info.d['referer']
print '-' * 32

if __name__ == '__main__':
	""" sML.check() """ # ECOM
	sCrawl.check()
	print '-' * 32 
	sys.stdout.flush()
	time.sleep( TIME )

	while True:
		print '-' * 32
		time_stamp()
		if sCrawl.update() == True:
			""" sML.send_gmail( sCrawl.updated_msg ) """ # ECOM
		else:
			# updated but purpose word is not exists
			pass
		print '-' * 32
		sys.stdout.flush()
		count += 1
		if count >= info.d['freq']:
			print 'System is over. Program will exit'
			sys.exit()
		time.sleep( TIME )

