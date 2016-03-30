import React from 'react';
import Immutable from 'immutable';
import API from '../api.js';

export default class AddBook extends React.Component {
  constructor(props) {
    super(props)
    this.state = {
      atom: Immutable.fromJS({
        url: null
      })
    }
  }

  parseURL(url) {
    var p = document.createElement('a');
    p.href = url;
    return p
  }

  submitForm = (e) => {
    e.preventDefault()
    const url = this.parseURL(this.state.atom.get('url'));
    API.get(`/book?url=http://${url.hostname}${url.pathname}`)
      .then(res => console.log(res))
      .catch(err => console.log(err))
  }

  renderProxied = (res) => {

  }

  handleValChange = (e) => {
    this.setState({
      atom: this.state.atom.set(e.target.name, e.target.value)
    })
  }

  render() {
    return <form onSubmit={this.submitForm}>
      <div><input type="text" name="url" onChange={this.handleValChange}/></div>
      <input type="submit"/>
    </form>
  }
}
