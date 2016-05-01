#!/usr/bin/env python
# -*- coding: utf-8 -*-

import ConfigParser
import getpass
import hashlib
import _os
import src.common.infoMail as infoML

class Config:
	def __init__( self ):
		self.conf = ConfigParser.SafeConfigParser()
		self.conf.read( infoML.dF['ini'] )
	
	def set( self, sect, item, val ):
		self.conf.set( sect, item, val )
	
	def update( self ):
		with open( infoML.dF['ini'], 'wb' ) as f_conf:
			self.conf.write( f_conf )


