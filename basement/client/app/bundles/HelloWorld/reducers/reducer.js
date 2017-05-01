import { handleActions } from 'redux-actions'
import Hello             from '../models/Hello'

export const initialState = new Hello()

const reducerMap = {
  UPDATE_INPUT(state, action) {
    return state.set(action.payload.id, action.payload.value)
  },
}

export default handleActions(reducerMap, initialState)
