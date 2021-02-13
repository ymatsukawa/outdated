const request = require('request');
class Authn {
  constructor(authnSettings, requestSettings) {
    this.authnSettings = authnSettings;
    this.requestSettings = requestSettings;
  }

  /*
    authenticate all of settings.authn > authnId
    when authned all, callback(cookies) is called.
    otherwise, error will be thrown.
  */
  authenticate(authnFinishedCallbacks) {
    // when no authn
    if(this.authnSettings.length == 0) {
      authnFinishedCallbacks({});
    }
    // any authn
    for(let authnSetting of this.authnSettings) {
      const requestOptions = {
        url: authnSetting.uri,
        headers: this.requestSettings.headers,
        form: authnSetting.postData,
        jar: true
      };
      var cookies = {};
      request.post(requestOptions, (error, response, body) => {
        if(response.statusCode == authnSetting.authenticatedStatusCode) {
          cookies[authnSetting.authnId] = response.headers['set-cookie'];
          if(Object.keys(cookies).length == this.authnSettings.length) {
            authnFinishedCallbacks(cookies);
          } else {
            // authn failed. this case is expected to below error case.
          }
        } else {
          console.log(cookies);
          throw new Error('authn failed. response status-code is, expected: ' + authnSetting.authenticatedStatusCode + ', got: ' + response.statusCode);
        }
      });
    }
  }
}

module.exports = Authn;
