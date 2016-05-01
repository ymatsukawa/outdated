#!/usr/bin/env python
# -*- coding: utf-8 -*-

import os
import Diff
import Crawl
import Seek

uos = os.name
if uos is 'nt':
	# windows NT
	import src.middleware.sound as Sound
else:
	# another OS
	# no implements
	pass

# check() should be invoked at least once
# At update(), if crawl is success( seeking words are grabbed ), returns True, if not, returns false
class State:
	def __init__( self ):
		self.Crawl = Crawl.Crawl()
		self.Seek = Seek.Seek()
		self.Diff = Diff.Diff()
		if uos is 'nt':
			self.Sound = Sound.Sound()
		self.updated_msg = None # List. Any don't have newline.
	def check( self ):
		self.Crawl.first()
		# Should seek first for over second times seeking
		self.Seek.write()
		self.Diff.first()
		self.Crawl.close()
	def update( self ):
		if self.Crawl.update() == False:
			print 'Not updated... Wait'
		else:	
			# before checking diff, should update .txt file
			self.Seek.write()
			print 'Dat has changed. Checking diff...'
			if self.Diff.check() == True:
				print '* Difference *'
				if uos is 'nt':
					self.Sound.play()
				self.updated_msg = self.Diff.get_diff()
				self.Crawl.close()
				return True
			else:
				print 'No Difference.'
				self.Crawl.close()
				return False

