import { connect }            from 'react-redux'
import HelloWorld             from '../components/HelloWorld'
import * as Actions           from '../actions/helloWorld'
import { bindActionCreators } from 'redux'

const mapStateToProps = (state) => ({
  helloModel: state
})

const mapDispatchToProps = (dispatch) => ({
  actions: bindActionCreators(Actions, dispatch)
})

export default connect(mapStateToProps, mapDispatchToProps)(HelloWorld)
