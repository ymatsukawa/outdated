import json

class FileRead:
    def __init__(self):
        pass

    def as_json(self, abs_path):
        return json.load(open(abs_path))
