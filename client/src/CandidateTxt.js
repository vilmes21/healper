import React from "react";

class CandidateTxt extends React.Component {
    render() {
        const {objArr} = this.props;
        const arr = objArr.map(x => x.namezh);
        return <div>
            {arr.join("")}
        </div>
    }
}

export default CandidateTxt;