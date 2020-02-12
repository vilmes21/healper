import React from "react";
import P3 from "./P3";

class Redirec extends React.Component {
  render(){
      const {to}=this.props;
      
    if (to==="/p3"){
      return <div>
        <P3 />
      </div>
    }
    
    return <div>
      {/* <NewDisease /> */}
      page 0
    </div>
  }
}

export default Redirec;