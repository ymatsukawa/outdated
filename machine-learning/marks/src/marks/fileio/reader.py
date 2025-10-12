# agents created
from pathlib import Path
from typing import Optional

class FileReaderError(Exception):
    pass


class Reader:
    def __init__(self, filepath: str) -> None:
        if not filepath or not filepath.strip():
            raise FileReaderError("File path cannot be empty or None")

        self._filepath = Path(filepath)
        self.content: Optional[str] = None

    @property
    def filepath(self) -> Path:
        return self._filepath

    def read(self) -> None:
        if not self._filepath.exists():
            raise FileReaderError(
                f"File not found: {self._filepath.absolute()}"
            )

        if not self._filepath.is_file():
            raise FileReaderError(
                f"Path is not a file: {self._filepath.absolute()}"
            )

        try:
            with self._filepath.open("r", encoding="utf-8") as file:
                self.content = file.read()
        except PermissionError as e:
            raise FileReaderError(
                f"Permission denied reading file: {self._filepath.absolute()}"
            ) from e
        except UnicodeDecodeError as e:
            raise FileReaderError(
                f"Unable to decode file as UTF-8: {self._filepath.absolute()}"
            ) from e
        except OSError as e:
            raise FileReaderError(
                f"Error reading file: {self._filepath.absolute()}"
            ) from e

    def __repr__(self) -> str:
        return f"Reader(filepath='{self._filepath}', content_loaded={self.content is not None})"
