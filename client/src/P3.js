import React from "react";

class P3 extends React.Component {
    componentWillMount(){
        window.history.pushState('page2', 'Title', '/p3');
    }
    
  render(){
    return <div>
      page 3
    </div>
  }
}

export default P3;