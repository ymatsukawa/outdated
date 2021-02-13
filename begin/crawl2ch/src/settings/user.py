#!/usr/bin/env python
# -*- coding: utf-8 -*-

import getpass
import hashlib
import configMail as Config
import _os
import src.common.infoMail as infoML


class User:
	def __init__( self ):
		self.conf = Config.Config()
		self.passwd = None # passwd is raw string

	def get_passwd( self ):
		return self.passwd

	# Confirm default password
	# user-input -> confirm -> returns result
	# If input is right, returns true. If not, returns false
	def check_passwd( self ):
		print "Input Password belongs ID => %s:" % infoML.dML['id']
		pin = getpass.getpass()
		pkey = hashlib.sha1( pin ).hexdigest()
		if pkey == infoML.dML['key']:
			print 'Passed'
			self.passwd = pin
			return True
		else:
			print 'Not passed'
			return False
	


	# Remake password of ID
	# this key will be applied 'key' at 'ecom.ini'
	def change_passwd( self ):
		print 'You should pass confirmation of password before its change.'
		if self.check_passwd() == False:
			print 'Your input is wrong. Nothing has been changed.'
		else:
			print 'You passed. Input New password:'; pin = getpass.getpass()
			print 'Input again:'; again = getpass.getpass()
			if pin == again:
				print 'All right. Password has been changed'
				self.passwd = pin # passwd value is 'raw string'
				pkey = hashlib.sha1( pin ).hexdigest()
				# change ini.key and infoML.dML['key'] with 'hash-key'
				infoML.dML['key'] = pkey
				self.conf.set( 'user', 'key', pkey )
				self.conf.update()
			else:
				print '...Maybe type of confirmation is wrong from origin. Please do again.'

