#!/usr/bin/env python
#-*- coding: utf-8 -*-

import os.path as p
import _os

"""
return path joined with /crawl2ch/data
"""
def rep_on_datadir( path ):
	datadir = p.abspath( p.join( _os.gpdir, 'file' ) )
	return p.abspath( p.join(datadir, path) )

