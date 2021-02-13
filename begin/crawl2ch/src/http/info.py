#!/usr/bin/env python
#-*- coding: utf-8 -*-

import _os
import re
import src.common.infoCrawl as info

# Returns timestamp
def timestamp():
	timeStamp = open( info.d['f_header'], 'rb' ).read()
	timeStamp = re.search("date.*GMT", timeStamp).group(0).replace('date: ', '')
	return timestamp

