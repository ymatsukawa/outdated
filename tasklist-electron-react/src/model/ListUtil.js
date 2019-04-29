import { Record, List } from 'immutable';

const ListModel = new Record({
  list: null
});

class ListUtil extends ListModel {
  constructor(tasks) {
    super({
      list: List()
    });
  }

  pushItem(item) {
    return this.set('list', this.list.push(item));
  }

  removeAt(id) {
    return this.set('list', this.list.filter(t => t.id !== id));
  }
}

export default ListUtil;