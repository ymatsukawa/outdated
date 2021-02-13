#!/usr/bin/env python
#-*- coding: utf-8 -*-

import _os
import src.common.infoCrawl as info

# file out DAT
def dat( src ):
	# write dat file
	print 'Writing DAT'
	f = open( info.d['f_dat'] , 'w' )
	f.write( src )
	f.flush()

# file out html header
# 'src' should be responce source
def header( src ):
	print 'Writing HTML header'
	f_header = open( info.d['f_header'], 'w' )
	s_header = str( src.code ) + " " + src.msg + "\n"
	for k, v in src.info().items():
		s_header += "%s: %s\n" % (k, v)
	f_header.write( s_header )

