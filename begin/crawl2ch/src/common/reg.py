#!/usr/bin/env python
#-*- coding: utf-8 -*-

import re


def host( referer ):
	return re.search('\/\/[a-z0-9]+\.2ch\.net|\/\/.*\.bbspink\.com', referer ).group(0).replace('//', '')

def board( referer ):
	return re.sub( "/\d.*", '', re.search('cgi.*', referer ).group(0).replace('cgi/', '') )

def requestUrl( referer, board ):
	return re.sub( '/$', '.dat', referer.replace('test/read.cgi/' + board, board + '/dat' ) )

def datFile( requestUrl ):
	return re.search('\d\d+', requestUrl ).group(0)
