// about markdown.table.header > https://github.com/jgm/pandoc/issues/2574#issuecomment-162003916
const FILE_CONTENT = {
    markdown: {
        begin: "<main>\n",
        page_header: "### WEEKLY TOPICS\n",
        table: {
            header: "| $0 |\n| ------------------------------------------------ |\n",
            body: "| [$0]($1) |\n"
        },
        excuseme: "**NO post found in last week.**",
        end: "</main>"
    }
}



const INLINE_MARKUP = `
<link rel="stylesheet" href="https://unpkg.com/style.css@1.0.0">
<link rel="preconnect" href="https://fonts.gstatic.com">
<link href="https://fonts.googleapis.com/css2?family=Noto+Sans+JP:wght@100;300;400;500;700;900&display=swap" rel="stylesheet">
<style>
  body {
    font-family: 'Noto Sans JP', sans-serif;
  }
  table > tbody > tr > td {
      font-size: 24px;
      font-weight: 500px;
  }
  main {
    margin: 0 auto;
    max-width: 50em;
  }
</style>
`

const PANDOC_OPE = {
    src: ['/tmp/', Math.random().toString(16).substr(2, 8), '.md'].join(''),
    args: [
        '-f', 'markdown',
        '-t', 'html5'
    ]
}

module.exports = {
    FILE_CONTENT,
    INLINE_MARKUP,
    PANDOC_OPE
}