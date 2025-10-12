from fastmcp import FastMCP
from fileio.reader import Reader, FileReaderError
from md.toc import Toc
from md.text_search import TextSearch
from md.content_finder import ContentFinder

mcp = FastMCP(name="marks", version="0.0.1")

@mcp.tool
def read_markdown_all_content(filepath: str) -> str:
    """read markdown file all content"""
    try:
        reader = Reader(filepath)
        reader.read()
        return reader.content
    except FileReaderError as e:
        return f"Error reading file: {e}"

@mcp.tool
def full_text_search(content: str, keyword: str, case_sensitive: bool) -> str:
    """full text search. case_sensitivei should be True until prompt or instruction requires."""
    text_search = TextSearch(content=content)
    text_search.full_text(keyword=keyword, case_sensitive=case_sensitive)

    if len(text_search.result_full_text) == 0:
        return f"not found keyword - {keyword}"

    return f"searched and found - {keyword} \n{text_search.result_full_text}"

@mcp.tool
def create_toc(filepath: str) -> str:
    """create table of contents."""
    toc = Toc()
    toc.generate(filepath=filepath)
    return toc.tables

@mcp.tool
def create_toc_with_linenumbers(all_content: str) -> list[dict]:
    """create table of contents with line numbers"""
    toc = Toc()
    toc.generate_with_line_numbers(content=all_content)
    return toc.tables_with_linenum

@mcp.tool
def get_content_of_heading(all_content: str, toc_with_linenum: list[dict], heading: str) -> str:
    """
    get content of specific heading
    if heading has multiple matches, notify user that which heading to choose with showing line numbers
    """
    content_finder = ContentFinder(
        content=all_content,
        toc_with_linenum=toc_with_linenum
    )
    return content_finder.get_content_of_heading(heading=heading)

if __name__ == "__main__":
    mcp.run()
