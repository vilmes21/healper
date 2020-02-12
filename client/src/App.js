import React from "react";
import NewDisease from "./NewDisease";
import Loading from "./Loading";
import config from "./helper/config";
import { connect } from 'react-redux'
import { redirect } from './action/page'
import Link2 from "./Link2";
import Home from "./Home";

class App extends React.Component {
  render(){
    const {redirect, page}=this.props;

    if (typeof(page) !== "string"){
      return <Loading />
    }

    const {urls}=config;

    switch (page) {
      case urls.newDisease:
        return <NewDisease />
      default:
        return <Home />
    }
  }
}

const mapState2Props = state => {
  return {
    page: state.page
  }
}

export default connect(
  mapState2Props,
  { redirect }
)(App)

//export default App;
