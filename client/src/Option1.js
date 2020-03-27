import React from "react";
import config from "./helper/config";

const {organSymptom, patho, policy} = config.frontendType;
const preBuiltTypes = [organSymptom, patho, policy]

class Option1 extends React.Component {
    _click = () => {
        const {option, appendCandidate, cartPreBuilt} = this.props;
        const {frontendType} = option;
        
        if (preBuiltTypes.indexOf(frontendType) > -1) {
            cartPreBuilt(option)
        } else {
            appendCandidate(option);
        }
    }

    render() {
        const {option} = this.props;
        const {txt} = option;

        return (
            <div className="opt inlineBloc" onClick={this._click}>
                {txt}
            </div>
        )
    }
}

export default Option1;