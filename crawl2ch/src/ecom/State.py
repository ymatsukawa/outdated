#!/usr/bin/env python
#-*- coding: utf-8 -*-

import Mail
import _os
import src.settings.user as User

class State:
	def __init__( self ):
		self.User = User.User()
		self.Mail = Mail.Mail()

	# Check password belongs ID with input.
	# If 'right user' set password and returns True, if not, returns False
	def check( self ):
		if self.User.check_passwd() == True:
			self.Mail.set_passwd( self.User.get_passwd() )
			return True
		else:
			return False

	# Make message
	# Send email
	def send_gmail( self, body ):
		self.Mail.convert_LtoS( body )
		self.Mail.def_create_msg()
		self.Mail.send_gmail()

