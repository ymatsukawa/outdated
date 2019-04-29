const Pageres = require('pageres');
const fs = require('fs');
const glob = require('glob');
const path = require('path');

class Capture {
  constructor(captureSettings, ioSettings, requestSettings, authnCookies) {
    this.captureSettings = captureSettings;
    this.ioSettings = ioSettings;
    this.requestSettings = requestSettings;
    this.authnCookies = authnCookies;
  }

  /*
    captures where settings.captures' DOM
    when capture finished, callback() runs.
  */
  capture(captureFinishedCallback) {
    this._clear();
    var capturedTimes = 0;
    for(let capture of this.captureSettings) {
      var options = { delay: 3, selector: capture.selector, headers: this.requestSettings.headers };
      if(capture.authnId !== null) {
        options.cookies = this.authnCookies[capture.authnId];
      }
      const pageres = new Pageres(options);
      pageres.src(capture.uri, [capture.size])
                  .dest(this.ioSettings.imagesDestDir)
                  .run()
                  .then(() => {
                    if(++capturedTimes == this.captureSettings.length) {
                      captureFinishedCallback();
                    }
                  });
    }
  }

  _clear() {
    if(!this.ioSettings.removePhotosBeforeCapture) {
      return;
    }
    glob(this.ioSettings.imagesDestDir + path.sep + '**/*.png', (error, files) => {
      for(let file of files) {
        fs.unlinkSync(file);
      }
    });
  }
}

module.exports = Capture;
