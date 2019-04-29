import React, { Component } from 'react';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';

class TaskForm extends Component {
  constructor(props) {
    super(props);
    this.handleAddTask = props.handleAddTask;
  }

  render() {
    return (
      <div>
        <form onSubmit={ (e) => { this.handleAddTask(e) } } >
          <TextField name="title" className="input-text" type="text" placeholder="input task" />
          <Button type="submit" variant="contained" color="primary">
            Add Task
          </Button>
        </form>
      </div>
    );
  }
}

export default TaskForm;