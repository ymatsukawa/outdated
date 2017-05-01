import { createStore }                   from 'redux'
import helloWorldReducer, {initialState} from '../reducers/reducer'

/**
 * @SHOULD_MOD
 * preloadedState is temporary variable.
 * how avoid initialState.set
 */
const configureStore = (railsProps) => {
  var preloadedState = {}
  Object.keys(railsProps).forEach((key) => {
    preloadedState = initialState.set(key, railsProps[key])
  })
  return createStore(helloWorldReducer, preloadedState)
}

export default configureStore
