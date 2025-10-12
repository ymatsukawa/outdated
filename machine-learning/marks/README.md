# marks

mcp - markdown content parser

## Installation

```
git clone git@github.com/ymatsukawa/marks /path/to/marks
cd /path/to/marks

claude mcp add uv run /path/to/marks/src/server.py
```

## Prompt Examples

### full text search

**case-sensitive(default)**

```
Search text "TEST" from tests/test.md, using marks mcp
```

**case-insensitive**

```
Search text "TEST" from tests/test.md with case-insensitive, using marks mcp
```

### toc

```
Add Table of Contents at tests/test.md, using marks mcp
```

### content finder

```
Get content of heading "Caching" from tests/test.md, using marks mcp
```

## Tested on prompt layer

* mac OS X Sequoia
* Claude Code Sonnet(claude-sonnet-4-20250514)

## LICENSE

MIT
