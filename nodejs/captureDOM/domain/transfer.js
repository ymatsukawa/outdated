const fs = require('fs');
const glob = require('glob');
const path = require('path');

class Transfer {
  constructor(ioSettings) {
    this.ioSettings = ioSettings;
  }

  /*
    /path/to/imagesDestDir/*.png
    goes to
    /path/to/imagesTransferedDir/*.png
  */
  transferCaptures() {
    glob(this.ioSettings.imagesDestDir + path.sep + "**/*.png", (error, srcFiles) => {
      for(let srcFile of srcFiles) {
        fs.renameSync(srcFile, this.ioSettings.imagesTransferedDir + path.sep + srcFile.split(path.sep).pop());
      }
    });
  }
}

module.exports = Transfer;
