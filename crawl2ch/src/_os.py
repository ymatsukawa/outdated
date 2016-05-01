#!/usr/bin/env python
#-*- coding: utf-8 -*-

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

insert_pdir()


