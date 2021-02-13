#!/usr/bin/env python
#-*- coding: utf-8 -*-

import os
import os.path as p
import sys


a = p.abspath
D = p.dirname


"""
Append path; parent directory at current directory.
"""
def insert_pdir():
	s = a( D( a( D( __file__ ) ) ) )
	sys.path.insert(0, s)
	return s

def insert_gpdir():
	# first directory is current directory
	# second directory is parent directory
	# third directory is grand parent directory
	s = a( D( a( D( a( D( __file__ ) ) ) ) ) )
	sys.path.insert(0, s)
	return s


insert_pdir()
gpdir = insert_gpdir()

