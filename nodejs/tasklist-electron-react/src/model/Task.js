import { Record } from 'immutable';

const TaskModelRecord = new Record({
  id: null,
  title: null
});

class Task extends TaskModelRecord {
  constructor(attributes) {
    super({
      id: Date.now(),
      title: attributes.title
    });
  }

  isValid(title) {
    if(this.title === "" || this.title.length > 255) {
      return false;
    }
    return true;
  }

  create(title) {
    return this.set('title', this.title);
  }
}

export default Task;