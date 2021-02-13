#!/usr/bin/env python
#-*- coding: utf-8 -*-

import zlib

# Decompress 'gzipped' src
# If object is responce resource, give value like 'res.read()' to parameter.
# Finally, returns dat
def decomp( buff ):
	# Raw data is gzipped, so it is needed to be decompressed
	print 'Decompressing DAT'
	dat = zlib.decompressobj(16 + zlib.MAX_WBITS ).decompress( buff ) # MUST -- zlib header setting to decompress gzip
	return dat

