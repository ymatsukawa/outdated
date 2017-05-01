let filledAll = `@bot
create
title[test title]
projectId[1234]
priority[1]
issueTypeId[1]
`

let blank = ''
let onlyMention = '@bot'
let mentionAndOrder = `@bot
create`

let blankTitle = `@bot
create
title[]`

let noProjectPriorityIssuetype = `@bot
create
title[test title]`

let noPriorityIssuetype = `@bot
create
title[test title]
projectId[1234]`

let noIssuetype = `@bot
create
title[test title]
projectId[1234]
priority[1]`

export const errorMessages = {
  "blank":                      blank,
  "onlyMention":                onlyMention,
  "mentionAndOrder":            mentionAndOrder,
  "blankTitle":                 blankTitle,
  "noProjectPriorityIssuetype": noProjectPriorityIssuetype,
  "noPriorityIssuetype":        noPriorityIssuetype,
  "noIssuetype":                noIssuetype
}

export const successMessages = {
  "filledAll": filledAll  
}
