from enum import Enum, auto

class ContentFinderError(Enum):
    NOT_FOUND = auto()
    DUPLICATED = auto()

class ContentFinder:
    def __init__(self, content: str, toc_with_linenum: list[dict]):
        self._content = content
        self._toc_with_linenum = toc_with_linenum
        self._error = None

    def get_content_of_heading(self, heading: str) -> str:
        self._validate(heading)

        if self._error == ContentFinderError.NOT_FOUND:
            return f"heading is NOT FOUND: {heading}"
        elif self._error == ContentFinderError.DUPLICATED:
            return f"heading is DUPLICATED: {heading}"

        start_line = None
        end_line = None
        for toc in self._toc_with_linenum:
            if heading.lower() in toc["text"].lstrip("#").lstrip().lower():
                start_line = toc["line"]
                break

        for toc in self._toc_with_linenum:
            if toc["line"] > start_line:
                end_line = toc["line"] - 1
                break

        if end_line is None:
            end_line = len(self._content.splitlines())

        lines = self._content.splitlines()
        target_content = "\n".join(lines[start_line - 1:end_line])

        return target_content

    def _validate(self, heading: str) -> ContentFinderError:
        matches = [
            toc for toc in self._toc_with_linenum
            if heading.lower() in toc["text"].lstrip("#").lstrip().lower()
        ]

        if len(matches) == 0:
            self._error = ContentFinderError.NOT_FOUND
        elif len(matches) > 1:
            self._error = ContentFinderError.DUPLICATED
