import React from "react";

class Option1 extends React.Component {
  _click = () => {
    const {option,appendCandidate}=this.props;
    appendCandidate(option);
  }

  render(){
      const {option,appendCandidate}=this.props;
      const {id, name, namezh} = option;

      return (
          <div className="opt" onClick={this._click}>
              {namezh} ({name})
          </div>
      )
  }
}

export default Option1;