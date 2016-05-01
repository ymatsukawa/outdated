#!/usr/bin/env python
#-*- coding: utf-8 -*-


import ConfigParser
import sys
import os.path as path
import infoFile as inF
import _os
import reg


d = {
'referer':      None , # referer; read.cgi
'reqUrl':       None , # request url; direct dat
'board':        None , # board name in url
'host':         None , # host name in url
'f_header':     None , # file name; response will be written
'f_dat':		    None , # file name; dat file will be written
'f_reLong':     None , # file name; matched words will be written
'f_reShort':    None , # file name; matched words will be written
'reqh_check':   None , # request header
'regD':         None , # regex; pattern
'regI':         None , # regex; pattern
'abone':        None , # deleted word at 2ch.net
'pabone':       None , # deleted word at bbspink.com
'looptime':     None , # interval time of running program running
'freq':         None , # frequency of program invocation
}

try:
	# Read init file
	datadir = path.abspath( path.join( _os.gpdir, 'data' ) )
	f_ini = path.abspath( path.join( datadir, 'crawl.ini' ) )
	conf = ConfigParser.SafeConfigParser()
	conf.read( f_ini )

# set around internet
	ref = conf.get('url', 'referer')
	d['referer'] = ref
	d['host'] = reg.host( ref )
	d['board'] = reg.board( ref )
	d['reqUrl'] = reg.requestUrl( ref, d['board'] )

# request header
	d['reqh_check'] = {
	'Host'            : d['host'],
	'User-Agent'      : 'Monazilla/1.0',
	'Connection'      : 'close',
	'Referer'         : d['referer'],
	'Accept-Encoding' : 'gzip',
	'Accept-Language' : 'ja',
	'Accept'          : '*/*'
	}

# information of file
	datname = reg.datFile( d['reqUrl'] ) + '.dat'
	d['f_dat'] = inF.rep_on_datadir( datname )
	d['f_header']  = inF.rep_on_datadir( 'header.txt' )
	d['f_reLong']  = inF.rep_on_datadir( 'long.txt' )
	d['f_reShort'] = inF.rep_on_datadir( 'short.txt' )
	d['f_sound']   = inF.rep_on_datadir( 'alart.wav' )

# user-defined regular expression
	d['regD'] = conf.get( 're', 'regD' )
	d['regI'] = unicode( conf.get( 're', 'regI' ), 'utf-8' ).encode('shift-jis')

# deleted word abone, ufoon
	d['abone']  = (u'あぼーん').encode('shift-jis')
	d['pabone'] = (u'うふ\u301cん').encode('shift-jis')

# frequency on program
	d['looptime'] = int( conf.get( 'time', 'interval' ) )
	d['freq'] = int( conf.get( 'time', 'freq' ) )
		
except Exception, e:
	print 'wrong at url parsing'
	print "http://HOST/test/read.cgi/BOARD/NUMBER/"
	sys.exit()

