import React from "react";
import {connect} from 'react-redux'
import {redirect} from './action/page'

class Loading extends React.Component {
    componentDidMount() {
        const f = () => {
            const d = window
                .location
                .href
                .split(`http://localhost:1234/`)[1];
            this
                .props
                .redirect(d);
        }

        f()
        // setTimeout(f, 1000)
    }

    render() {
        return <div>
            Loading...
        </div>
    }
}

export default connect(null, {redirect})(Loading)