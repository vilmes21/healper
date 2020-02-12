import React from "react";
import {connect} from 'react-redux'
import {redirect} from './action/page'

class Link2 extends React.Component {
    _click = () => {
        const {redirect, to} = this.props;
        redirect(to);
    }

    render() {
        const {children} = this.props;
        return <div onClick={this._click}>
            {children}
        </div>
    }
}

export default connect(null, {redirect})(Link2)