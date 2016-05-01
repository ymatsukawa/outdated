#!/usr/bin/env python
# -*- coding: utf-8 -*-

import re
import sys
import _os
import src.common.infoCrawl as info
import Elem


class Seek:	
	def __init__( self ):
		self.lines = None
		self.rLong = Elem.Elem( info.d['regD'], info.d['f_reLong'] )
		self.rShort = Elem.Elem( info.d['regI'], info.d['f_reShort'] )

	def write( self ):
		self.read()
		self.w_rLong()
		self.w_rShort()

	"""
		Read only content from DAT
		Avoid deleted response
		Exit programm if full response happened
	"""
	def read( self ):
		self.lines = ''
		try:
			for line in open( info.d['f_dat'], 'rb' ):
				# Firstly check full response
				if line.find('Over 1000 Thread') != -1:
					raise AttributeError
				# convert deleted response to None-String
				line = line.replace( info.d['abone'], '' ).replace( info.d['pabone'], ' ' )
				# change newlines
				line = line.replace("\r\n", "\n")
				# get only content
				line = re.search('<>\s.*', line).group(0)
				# remove <>
				line = line.replace('<>', 'EOF')
				# adjust head and tail of sentence
				line = re.sub('^EOF\s', '', line)
				line = re.sub('\sEOF$', '\n', line )
				# replace '<br>' to Replacement-word
				line = line.replace(' <br> ', '\n' )
				self.lines += line
		except AttributeError:
			print 'happend FULL-RESPONSE'
			print 'Program will exit'
			sys.exit()

	# write out regD matched words
	def w_rLong( self ):
		self.rLong.gather( self.lines )
	
	# write out regI matched words
	def w_rShort( self ):
		self.rShort.gather( self.lines )


	
