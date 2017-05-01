import React      from 'react'
import { Record } from 'immutable'

const attributes = {
  name: "",
}

const HelloRecord = Record(attributes)

export default class Hello extends HelloRecord {
  displayName () {
    if(this.name) {
      return <div>Hello, {this.name}</div>
    }
    return null
  }
}
