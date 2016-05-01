#!/usr/bin/env python
# -*- coding: utf-8 -*-

import urllib2
import sys
import _os
import src.common.infoCrawl as info

# request check at the url
# If request passed(200), returns response source
# or None returns(304)
def url( reqh ):
	reqset = urllib2.Request( info.d['reqUrl'], None, reqh )
	try:	
		res = urllib2.urlopen( reqset )
	except urllib2.HTTPError as e:
		print 'The server couldn\'t fulfill the request.'
		print 'Error code: ', e.code
		return None
	except urllib2.URLError as e:
		print 'We failed to reach a server.'
		print 'Reason: ', e.reason
		print 'Program will exit'
		self.code = -1
		self.res.close()
		sys.exit() # INVOKE ERROR SIGNAL
	else:
		print 'Request passed'
		return res
		# everything is fine


