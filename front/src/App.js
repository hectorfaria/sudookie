import React, { Component } from 'react';
import {connect, sendMsg } from './api/ws.js' 
import Header from './components/Header';
import ChatHistory from './components/ChatHistory';

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      chatHistory: []
    }
  }

  componentDidMount() {
    connect((msg) => {
      this.setState(prevState => ({
        chatHistory: [...this.state.chatHistory, msg]
      }))
    });
  }
  
  render() {
    return (
      <div className="App">
        <Header />
        <ChatHistory chatHistory={this.state.chatHistory} />
      </div>
    );
  }
}

export default App;
