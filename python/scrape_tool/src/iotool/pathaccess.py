import os.path as p

class PathAccess:
    def __init__(self):
        pass
    
    @staticmethod
    def realpath(sender_filename, target_relative_path):
        current_dir = p.dirname(p.realpath(sender_filename))
        path = p.join(current_dir, target_relative_path)
        
        if not p.exists(path):
            raise Exception(path + " does not exist")
        
        return path