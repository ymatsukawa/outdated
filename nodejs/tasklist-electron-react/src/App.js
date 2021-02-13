import React, { Component } from 'react';
import TaskForm from './components/TaskForm.js'
import TaskList from './components/TaskList.js';
import TaskPreview from './components/TaskPreview.js';
import Task from './model/Task.js';
import ListUtil from './model/ListUtil.js';
import PreviewContext from './model/PreviewContext.js';
import Grid from '@material-ui/core/Grid';

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      taskList: new ListUtil(),
      previweText: ''
    };
    this.handleAddTask = this.handleAddTask.bind(this);
    this.handleRemoveTask = this.handleRemoveTask.bind(this);
    this.handlePreviewTask = this.handlePreviewTask.bind(this);
  }

  handleAddTask(e) {
    e.preventDefault();

    let inputTitle = e.target.title.value;
    const task = new Task( {title: inputTitle} );
    if(task.isValid())
    {
      const addedTaskList = this.state.taskList.pushItem(task.create());
      this.setState({ taskList: addedTaskList });
      e.target.title.value = '';
    }
  }

  handleRemoveTask(e, taskId) {
    e.preventDefault();
    const removedTaskList = this.state.taskList.removeAt(taskId);
    this.setState({ taskList: removedTaskList });
  }

  handlePreviewTask(e) {
    e.preventDefault();
    const preview = new PreviewContext('list', this.state.taskList.list);
    this.setState({ previewText: preview.getPreviewText() });
  }

  render() {
    return (
      <div className="App">
        <Grid container justify="center" alignItems="center">
          <Grid item>
            <TaskForm handleAddTask={this.handleAddTask} />
          </Grid>
          <Grid item>
            <TaskList tasks={this.state.taskList.list}
                      handleRemoveTask={this.handleRemoveTask}
            />
          </Grid>
          <Grid item>
            <TaskPreview previewText={this.state.previewText}
                        handlePreviewTask={this.handlePreviewTask} />
          </Grid>
        </Grid>
      </div>
    );
  }
}

export default App;
