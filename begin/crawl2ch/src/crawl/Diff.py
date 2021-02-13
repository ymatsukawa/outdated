#!/usr/bin/env python
# -*- coding: utf-8 -*-

import _os
import src.common.infoCrawl as info

class Diff:
	def __init__( self ):
		self.reLong = None
		self.reShort = None
		self.diff = None

	def get_diff( self ):
		return self.diff

	def first( self ):
		self.reLong = open( info.d['f_reLong'], 'rb' ).read().split('\n')[:]
		self.reShort = open( info.d['f_reShort'], 'rb' ).read().split('\n')[:]

	# compare old and new
	# if two have differentation return True, if not, return False
	def check( self ):
		LongNew = open( info.d['f_reLong'], 'rb' ).read().split('\n')[:]
		ShortNew = open( info.d['f_reShort'], 'rb' ).read().split('\n')[:]
		reD = set( LongNew[:] ) - set( self.reLong[:] )
		reI = set( ShortNew[:] ) - set( self.reShort[:] )
		L = map( lambda x: x, reD ) # List contains string of pattern matched 'regD'
		S = map( lambda x: x, reI ) # List contains string of pattern matched 'regI'
		if ( L != [] ) or ( S != [] ): # if any differentation exists
			self.diff = L + S
			self.reLong = LongNew[:]
			self.reShort = ShortNew[:]
			return True
		else:
			return False

