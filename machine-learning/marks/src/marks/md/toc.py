import md_toc

class Toc:
    def __init__(self):
        self.tables = None
        self.tables_with_linenum = None

    def generate(self, filepath: str):
        self.tables = md_toc.api.build_toc(filepath)

    def generate_with_line_numbers(self, content: str, set_as: str = ""):
        tocs = []
        for i, line in enumerate(content.splitlines()):
            if line.lstrip().startswith("#"):
                tocs.append({"line": i + 1, "text": line})

        self.tables_with_linenum = tocs
