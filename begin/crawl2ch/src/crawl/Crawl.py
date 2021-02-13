#!/usr/bin/env python
# -*- coding: utf-8 -*-

import _os
import src.common.infoCrawl as info
import src.middleware.gzip as gzip
import src.middleware.save as save
import src.http.check as check
import src.http.info as hinfo

# Think -- no procedure with code.
class Crawl:
	def __init__( self ):
		self.res  = None # responce data from Server
		self.dat  = None # dat file
		self.code = 404 # responce code from Server
		self.reqh = info.d['reqh_check'].copy() # request html header

	# First invocation
	#	Send request message
	# Write responce message
	# Write DAT file
	def first( self ):
		print 'First invocation for program'
		self.connect()
		self.write_header()
		self.decomp()
		self.write_dat()

	# Over second invocation
	# When page is modified, update and return True, if not, return False
	def update( self ):
		sign = self.is_modified()
		if sign == True:
			self.write_header()
			self.decomp()
			self.write_dat()
			return True
		else:
			print 'Wait more'
			return False

	#	get responce resouce and set it to self.res
	def connect( self ):
		self.res = check.url( self.reqh )
		if self.res is None:
			self.code = 304
		else:
			self.code = 200

	# file out header information
	def write_header( self ):
		save.header( self.res )

	# file out DAT
	def write_dat( self ):
		save.dat( self.dat )

	# Decompress 'gzipped' responce resource
	def decomp( self ):
		self.dat = gzip.decomp( self.res.read() )

	# Close connection
	def close( self ):
		if self.res != None:
			self.res.close()

	# Append 'If-Modified-Since' on Header
	# Get old timestamp
	# If dat is modified, return true, if not, return false.
	def is_modified( self ):
		print 'Checking modified part'
		self.reqh['If-Modified-Since'] = hinfo.timestamp()
		self.connect()
		if self.code == 200:
			print 'Modified before request'
			return True
		else: # Not Modified DAT
			print 'Not Modified...'
			return False

