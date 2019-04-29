const path = require('path');

/*
[settings convention]

io:
  .imagesDestDir
    Where to store png files. default is "{capture_transfer_dir}/dest_images"
    value should be absolute path and follows "{separator}dirname{separator}dirname...{separator}dirname"
    ex.) "/tmp/captures"
  .imagesTransferedDir
    Where to transfer files. default is "/tmp" (default is not expected windows env)
    value should be absolute path and follows "{separator}dirname{separator}dirname...{separator}dirname"
    ex.) "/path/to/myfolder/captures"
  .removePhotBeforeCapture
    Each time before take captures, remove files belows of .imagesDestDir or not.
    boolean true is yes, false is no.

authn: settings about authentication
  .authnId
    Unique id. ex.) "taskBoardService", "blahblahCloud". null or not string is not expected.
  .uri
    POST URI to authenticate.
  .postData
    POST form data to authenticate.
    ex.) { "userId": "username", "password": "mypassword", "url": urlencoded("/index") }
    note) latest version only supported content-type "x-www-form-urlencoded"
  .authenticatedStatusCode
    Status code when just authenticated and cookies are returned.
    ex.) 302

captures:
  .authnId
    If capture page is needed to authentication, link value of authnId at 'authn'.
    ex.) "taskBoardService"
  .uri
    Where to capture.
    ex.) "http://esample.com"
  .size
    Size to capture. WIDTHxHEIGHT
    ex.) "1920x1080"
  .selector
    To capture specific DOM, set this as CSS selectors.
    ex.) "div#foo div.topics"
*/
const settings = {
  io: {
    imagesDestDir: path.join(__dirname, '../dest_images'),
    imagesTransferedDir: '/tmp',
    removePhotosBeforeCapture: true
  },
  request: {
    headers: {
      'Accept-Language': 'ja'
    }
  },
  authn: [
    {
      authnId: "mytaskboard",
      uri: "https://exmaple.com",
      postData: { clientTimezoneOffset: "-540", url: encodeURIComponent("/dashboard"), username: "username", password: "password" },
      authenticatedStatusCode: 302
    }
  ],
  captures: [
    {
      "authnId": "mytaskboard",
      "uri": "https://example.com",
      "size": "1920x1080",
      "selector": "body"
    },
    {
      "authnId": null,
      "uri": "example.com",
      "size": "320x160",
      "selector": "body"
    }
  ]
};

module.exports = settings;
