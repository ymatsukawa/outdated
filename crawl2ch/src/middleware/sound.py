#!/usr/bin/env python
# -*- coding: utf-8 -*-

import winsound
import _os
import src.common.infoCrawl as info

class Sound:
	def __init__( self ):
		self.f_name = info.d['f_sound']
	
	# Plya Sound
	# ! Only on Windows !
	def play( self ):
		for i in range(2):
			winsound.PlaySound( self.f_name, winsound.SND_FILENAME )


