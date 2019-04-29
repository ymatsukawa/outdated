import { Record } from 'immutable';

const PreviewModel = new Record ({
  stream: null
});

class PreviewStrategy extends PreviewModel {
  constructor(stream) {
    super({
      stream: stream
    });
  }

  getPreviewText() {
     // set as abstract
  }
}

export default PreviewStrategy;