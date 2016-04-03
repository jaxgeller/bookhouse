import React from 'react';
import Immutable from 'immutable';
import API from '../api.js';

export default class AddBook extends React.Component {
  constructor(props) {
    super(props)
    this.state = {
      atom: Immutable.fromJS({
        url: null,
        books: [],
      }),
      loading: false
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
    this.setState({loading: true})
    API(`/book?url=http://${url.hostname}${url.pathname}`)
      .then(res => res.json())
      .then(res => {
        const books = this.state.atom.get('books')
        this.setState({
          atom: this.state.atom.set('books', this.state.atom.get('books').push(res)),
          loading: false
        })
      })
      .catch(err => console.log(err))
  }

  handleValChange = (e) => {
    this.setState({
      atom: this.state.atom.set(e.target.name, e.target.value)
    })
  }

  render() {
    return <div>
      <div>
        <form onSubmit={this.submitForm}>
          <div>
            <input style={{width:"100%", height: "50px"}} type="text" name="url" onChange={this.handleValChange}/>
          </div>
          <input type="submit"/>
        </form>
      </div>
      {this.state.loading ? <p>Loading</p> : null}
      <div className="book">
        {this.state.atom.get('books').map((item, index) => {
          return <div key={index}>
            <h3>{item.Title}</h3>
            <img src={item.Img.replace(/\"/g,'')} alt=""/>
          </div>;
        })}
      </div>
    </div>;
  }
}
