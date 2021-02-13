import PreviewStrategy from './PreviewStrategy.js'

class ListPreviewStrategy extends PreviewStrategy {
  constructor(list) {
    super(list);
  }

  getPreviewText() {
    var result = '';
    this.stream.forEach((item, index) => {
      result += (item.title + "\n");
    });

    result = result.slice(0, -1); // chop last word
    return result;
  }
}

export default ListPreviewStrategy;
