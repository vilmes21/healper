import React from "react";
import Option1 from "./Option1";

class Options1 extends React.Component {
  render(){
      const {options, appendCandidate,cartPreBuilt}=this.props;
      if (options.length === 0){
          return <div>0 match</div>
      }

      return (
          <div>
              {
                  options.map(x => {
                      return <Option1 key={`${x.frontendType}-${x.id}`} option={x}
                      appendCandidate={appendCandidate}
                      cartPreBuilt={cartPreBuilt}
                      />
                  })
              }
          </div>
      )
    
  }
}

export default Options1;