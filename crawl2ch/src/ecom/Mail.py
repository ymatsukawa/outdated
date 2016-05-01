#!/usr/bin/env python
#-*- coding: utf-8 -*-

import sys
import smtplib
from email.MIMEText import MIMEText
from email.Utils import formatdate
import _os
from src.common.infoMail import dML


class Mail:
	def __init__( self ):
		self.msg = None
		self.body = None
		self.From = None
		self.To = None
		self.passwd = None

	def set_passwd( self, p ):
		self.passwd = p
	
	# Convert 'string in List' to 'String with line'
	# the 'String with line' will be in self.body
	def convert_LtoS( self, body ):
		self.body = ''
		for line in body:
			self.body += line + "\n";

	# create message with MIME Text
	def def_create_msg( self ):
		self.msg = MIMEText( self.body )

		self.msg['Subject'] = 'TEST'
		self.msg['From'] = dML['from']
		self.msg['To'] = dML['to']
		self.msg['Date'] = formatdate()

	# send with gmail account
	# Before invocation of this function, ecom.User.User.check should be dealed at the point of self.passwd
	# Maybe mergine between Mail and User needed
	def send_gmail( self ):
		s = smtplib.SMTP('smtp.gmail.com', 587)
		s.ehlo()
		s.starttls()
		s.ehlo()
		s.login( dML['id'], self.passwd )
		s.sendmail( self.msg['From'], self.msg['To'], self.msg.as_string() )
		s.close()


