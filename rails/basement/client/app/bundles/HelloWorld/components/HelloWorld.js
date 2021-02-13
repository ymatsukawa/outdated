import React, { Component, PropTypes } from 'react'

export default class HelloWorld extends Component {
  updateInput(e) {
    const { actions } = this.props
    actions.updateInput(e.target)
  }

  render() {
    const { helloModel } = this.props

    return (
      <div>
        <h3>
          {helloModel.displayName()}
        </h3>
        <hr />
        <form >
          <label htmlFor="name">
            Say hello to:
          </label>
          <input
            id="name"
            type="text"
            onChange={this.updateInput.bind(this)}
          />
        </form>
      </div>
    )
  }
}

HelloWorld.propTypes = {
  updateName: PropTypes.func,
  name: PropTypes.string,
}
