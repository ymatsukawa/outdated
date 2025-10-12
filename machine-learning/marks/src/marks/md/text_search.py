class TextSearch:
    def __init__(self, content: str):
        self._content = content
        self.result_full_text = None

    def full_text(self, keyword: str, case_sensitive: bool = False) -> list[dict]:
        """full text search"""
        lines = self._content.splitlines()
        results = []
        for i, line in enumerate(lines):
            target_txt = line if case_sensitive else line.lower()
            search_word = keyword if case_sensitive else keyword.lower()

            if search_word in target_txt:
              results.append({"line": i + 1, "text": line})

        self.result_full_text = results
