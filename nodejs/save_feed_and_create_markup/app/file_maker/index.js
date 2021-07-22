const fs = require('fs')
const nodePandoc = require('node-pandoc');
const DBClient = require('../db/client')
const SQL = require('../db/sql')
const { FILE_CONTENT, INLINE_MARKUP, PANDOC_OPE } = require('./file_content')
const { getFeedCategory, getFeedTag, getMarkupName } = require('../common/feed_category')
const CHARSET = 'utf8'

const createMarkup = (category_type) => {
    const dbClient = new DBClient()
    fs.writeFileSync(PANDOC_OPE.src, INLINE_MARKUP, CHARSET)
    fs.appendFileSync(PANDOC_OPE.src, FILE_CONTENT.markdown.begin, CHARSET)
    fs.appendFileSync(PANDOC_OPE.src, FILE_CONTENT.markdown.page_header, CHARSET)

    var category = getFeedCategory(category_type)
    var sql = SQL.feed.select.periodic
    sql = sql.replace(/\$0/, category.toString())
    dbClient.select(sql)
        .then(records => {
            var body = ""
            if(records.length > 0) {
                var table_header = JSON.parse(JSON.stringify(FILE_CONTENT.markdown.table.header))
                table_header = table_header.replace(/\$0/, getFeedTag(category_type))
                fs.appendFileSync(PANDOC_OPE.src, table_header, CHARSET)
            
                records.forEach(record => {
                    var body = JSON.parse(JSON.stringify(FILE_CONTENT.markdown.table.body))
                    body = body.replace(/\$0/, record.title).replace(/\$1/, record.hyperlink)
                    fs.appendFileSync(PANDOC_OPE.src, body, CHARSET)
                })
            } else {
                fs.appendFileSync(PANDOC_OPE.src, FILE_CONTENT.markdown.excuseme, CHARSET)
            }
                
            pandoc_callback = (error, result) => {
                if(error) {
                    console.log("ERROR: pandoc got error")
                    console.log(error)
                }
        
                console.log(result)
            }
            
            fs.appendFileSync(PANDOC_OPE.src, FILE_CONTENT.markdown.end, CHARSET)

            const markupName = getMarkupName(category_type)
            var args = JSON.parse(JSON.stringify(PANDOC_OPE.args)) // deep copy because .push effects after execution
            args.push('-o', ['./created_files/', markupName, ".html"].join(""))
            nodePandoc(PANDOC_OPE.src, args, pandoc_callback)
        })
}

exports.createMarkup = createMarkup