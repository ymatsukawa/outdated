#!/usr/bin/env python
# -*- coding: utf-8 -*-

import re

class Elem:
	def __init__( self, r, wf ):
		self.elem = None
		self.reg = re.compile( r )
		self.f_w = open( wf, 'wb' )

	# gather words grobaly that matches self.reg
	def gather( self, lineObj ):
		self.elem = []
		E = set( lineObj.split("\n") ) - set([''])
		lines = map( lambda x: x, E )
		for line in lines:
			s = None
			try:
				s = self.reg.search( line ).group(0)
			except AttributeError:
				pass # pass the line
			if s != None:
				self.elem.append( s )
		for line in self.elem:
			self.f_w.write( line + "\n" )
		self.f_w.flush()

