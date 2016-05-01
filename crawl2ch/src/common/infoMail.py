#!/usr/bin/env python
# -*- coding: utf-8 -*-

import ConfigParser
import re
import os
import _os

#dML
dML = {
	'id':   None ,
	'key':  None ,
	'from': None ,
	'to':   None 
}

dF = {
	'ini': None
}

try:
	datadir = os.path.abspath( os.path.join( _os.gpdir, 'data' ) )
	f_ini = os.path.abspath( os.path.join( datadir, 'ecom.ini' ) )
	dF['ini'] = f_ini
	conf = ConfigParser.SafeConfigParser()
	conf.read( dF['ini'] )

	dML['id']   = conf.get( 'user', 'id' )
	dML['key']  = conf.get( 'user', 'key' )
	dML['from'] = conf.get( 'send', 'from' )
	dML['to']   = conf.get( 'send', 'to' )

except Exception, e:
	print e

