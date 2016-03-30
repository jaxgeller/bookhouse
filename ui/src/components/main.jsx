import React from 'react'
import AddBook from './AddBook';

export default class Main extends React.Component {
  constructor(props) {
    super(props)

    this.state = {

    }
  }

  render() {
    return (
      <div>
        <h1>Add a Book</h1>
        <AddBook />
      </div>
    )
  }
}
