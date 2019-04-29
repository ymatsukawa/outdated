import React, { Component } from 'react';
import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';
import ListItemSecondaryAction from '@material-ui/core/ListItemSecondaryAction';
import IconButton from '@material-ui/core/IconButton';
import DeleteIcon from '@material-ui/icons/Delete';

class TaskList extends Component {
  constructor(props) {
    super(props);
    this.handleRemoveTask = props.handleRemoveTask;
  }

  render() {
    this.tasks = this.props.tasks;
    return (
      <div>
        <ul>
          {
            this.tasks.map((task, index) => {
              return (
                <ListItem key={index}>
                  <ListItemText primary={task.title} />
                  <ListItemSecondaryAction aria-label="Delete">
                    <IconButton aria-label="delete" onClick={(e) => { this.handleRemoveTask(e, task.id) } }>
                      <DeleteIcon  />
                    </IconButton>
                  </ListItemSecondaryAction>
                </ListItem>
              )
            })
          }
        </ul>
      </div>
    );
  }
}

export default TaskList;