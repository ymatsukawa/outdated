import os
import unittest

loader = unittest.TestLoader()
unit_test_root_dir = os.path.dirname(os.path.abspath(__file__))
discovered_test = loader.discover(unit_test_root_dir)

runner = unittest.TextTestRunner()
runner.run(discovered_test)
