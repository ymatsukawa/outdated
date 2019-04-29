import React, { Component } from 'react';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';


class TaskPreview extends Component {
  constructor(props) {
    super(props);
    this.handlePreviewTask = props.handlePreviewTask;
  }

  render() {
    this.previewText = this.props.previewText;
    return (
      <div>
        <TextField
          id="outlined-multiline-static"
          label=""
          multiline
          ows="4"
          value={this.previewText}
          defaultValue=""
          className="defautTextarea"
          margin="normal"
          variant="outlined"
        />
        <Button variant="contained" className="default-button" onClick={(e) => { this.handlePreviewTask(e) }}>
          Out Preview
        </Button>
      </div>
    )
  }
}

export default TaskPreview;