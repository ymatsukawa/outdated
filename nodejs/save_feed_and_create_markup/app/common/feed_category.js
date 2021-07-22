FEED_CONST = {
    feeds: [
        {
            category: -1,
            name: "example.com",
            url: ["http://example.com/feed"],
            tag: "#example",
            markup_name: "example"
        },
        {
            category: 0,
            name: "devto-latest",
            url: [
                "https://dev.to/feed/",
                "https://dev.to/feed/tag/javascript"
            ],
            tag: "#devto",
            markup_name: "devto_latest"
        }
    ]
}

const getFeedCategory = (category_type) => {
    const foundObj = FEED_CONST.feeds.find((v) => v.name === category_type )
    if(!foundObj) {
        return -1
    }

    return foundObj.category
}

const getFeedLinks = (category_type) => {
    const foundObj = FEED_CONST.feeds.find((v) => v.name === category_type )
    if(!foundObj) {
        return []
    }

    return foundObj.url
}

const getFeedTag = (category_type) => {
    const foundObj = FEED_CONST.feeds.find((v) => v.name === category_type )
    if(!foundObj) {
        return "*"
    }

    return foundObj.tag
}

const getMarkupName = (category_type) => {
    const foundObj = FEED_CONST.feeds.find((v) => v.name === category_type )
    if(!foundObj) {
        return "sample"
    }

    return foundObj.markup_name
}

module.exports = {
    getFeedCategory,
    getFeedLinks,
    getFeedTag,
    getMarkupName
}
